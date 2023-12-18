package box

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wavynote/internal/gateway/http/handler/restapi"
	"github.com/wavynote/internal/platform/dbmsadapter/postgres"
	"github.com/wavynote/internal/wavynote"
)

type BoxHandler struct {
	dbInfo wavynote.DataBaseInfo
}

func NewBoxHandler(dbInfo wavynote.DataBaseInfo) *BoxHandler {
	h := &BoxHandler{
		dbInfo: dbInfo,
	}
	return h
}

// ShowConversation godoc
// @Summary      노트를 주고 받는 대화방 목록 조회(최대 3개)
// @Description  노트를 주고 받는 대화방 목록 조회(최대 3개)
// @Tags         Box 페이지
// @Security	 BasicAuth
// @Param        id   query     string  false  "user id"
// @Success      200  {object}  restapi.ConversationListResponse "조회된 대화방 목록 정보"
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /box/conversationlist [get]
func (h *BoxHandler) ShowConversation(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	userId := c.Query("id")
	fmt.Println("user_id:", userId)

	db := postgres.NewService(h.dbInfo.Host, h.dbInfo.Port, h.dbInfo.Login, h.dbInfo.Password, h.dbInfo.Database, h.dbInfo.SSLMode, h.dbInfo.AppName)
	err = db.Open()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	}

	defer func() {
		err := db.Close()
		if err != nil {
			//
		}
	}()

	query := fmt.Sprintf(`
		SELECT id, host_id, guest_id, create_at
		FROM public.conversation
		WHERE delete_at is NULL AND (host_id = '%s' OR guest_id = '%s')
	`, userId, userId)

	var conversationList []restapi.ConversationInfo
	rows, err := db.SelectQuery(query)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	} else {
		for i := 0; i < len(rows); i++ {
			var oppNicName string
			hostId := db.GetString(rows[i]["host_id"])
			if hostId == userId {
				oppNicName = db.GetString(rows[i]["guest_id"])
			} else {
				oppNicName = hostId
			}

			conversationId := db.GetUUID(rows[i]["id"])

			// note 테이블에서 동일한 conversation_id를 갖는 row의 개수 조회
			noteCount := 0
			innerQuery := fmt.Sprintf(`
				SELECT count(*) AS cnt
				FROM public.note
				WHERE conversation_id = '%s'
			`, conversationId)
			row, err := db.SelectQuery(innerQuery)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
					Code: http.StatusInternalServerError,
					Msg:  err.Error(),
				})
				return
			} else {
				for j := 0; j < len(row); j++ {
					noteCount = db.GetInteger(row[j]["cnt"])
				}
			}

			// 받은 노트(from_id가 사용자 id가 아닌 노트)중 isread값이 false인 row가 존재하는 경우에는 NewNote 필드를 true로 설정
			newNote := false
			innerQuery = fmt.Sprintf(`
				SELECT count(*) AS cnt
				FROM public.note
				WHERE conversation_id = '%s' AND from_id != '%s' AND isread = false
			`, conversationId, userId)
			row, err = db.SelectQuery(innerQuery)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
					Code: http.StatusInternalServerError,
					Msg:  err.Error(),
				})
				return
			} else {
				count := 0
				for j := 0; j < len(row); j++ {
					count = db.GetInteger(row[j]["cnt"])
				}

				if count != 0 {
					newNote = true
				}
			}

			conversationList = append(conversationList, restapi.ConversationInfo{
				ConverstaionId: conversationId,
				OppNickName:    oppNicName,
				NoteCount:      noteCount,
				NewNote:        newNote,
			})
		}
	}

	if len(conversationList) == 0 {
		c.IndentedJSON(http.StatusNotFound, restapi.Response404{
			Code: http.StatusNotFound,
			Msg:  "there is no conversation room",
		})
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		restapi.ConversationListResponse{
			Conversations: conversationList,
		},
	)
}

// ShowConversationNoteList godoc
// @Summary      특정 친구와 주고받은 노트 목록
// @Description  특정 친구와 주고받은 노트 목록
// @Tags         Box 페이지
// @Security	 BasicAuth
// @Param        id  query     string  false  "conversation id"
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /box/notelist [get]
func (h *BoxHandler) ShowConversationNoteList(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	conversationId := c.Query("id")
	fmt.Println("id:", conversationId)

	db := postgres.NewService(h.dbInfo.Host, h.dbInfo.Port, h.dbInfo.Login, h.dbInfo.Password, h.dbInfo.Database, h.dbInfo.SSLMode, h.dbInfo.AppName)
	err = db.Open()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	}

	defer func() {
		err := db.Close()
		if err != nil {
			//
		}
	}()

	query := fmt.Sprintf(`
		SELECT id, from_id, to_id, title, contents, send_at
		FROM public.note
		WHERE conversation_id = '%s'
		ORDER BY send_at DESC
	`, conversationId)

	var conversationNoteList []restapi.ConversationNoteInfo
	rows, err := db.SelectQuery(query)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	} else {
		for i := 0; i < len(rows); i++ {
			conversationNoteList = append(conversationNoteList, restapi.ConversationNoteInfo{
				NoteId:  db.GetUUID(rows[i]["id"]),
				FromId:  db.GetString(rows[i]["from_id"]),
				ToId:    db.GetString(rows[i]["to_id"]),
				Title:   db.GetString(rows[i]["title"]),
				Preview: db.GetString(rows[i]["contents"])[0:20], // preview는 본문의 시작부터 최대 20자까지 제공
				SendAt:  db.GetString(rows[i]["send_at"]),
			})
		}
	}

	if len(conversationNoteList) == 0 {
		c.IndentedJSON(http.StatusNotFound, restapi.Response404{
			Code: http.StatusNotFound,
			Msg:  "there is no note in target conversation room",
		})
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		restapi.ConverstaionNoteListResponse{
			ConversationNotes: conversationNoteList,
		},
	)
}

// ShowConversationNote godoc
// @Summary      특정 노트 조회
// @Description  특정 노트 조회
// @Tags         Box 페이지
// @Security	 BasicAuth
// @Param        cid  query     string  false  "conversation id"
// @Param        nid  query     string  false  "note id"
// @Param        uid  query     string  false  "user id"
// @Success      200  {object}  restapi.NoteInfo "조회한 노트 정보"
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /box/show [get]
func (h *BoxHandler) ShowConversationNote(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	conversationId := c.Query("cid")
	fmt.Println("user_id:", conversationId)

	noteId := c.Query("nid")
	fmt.Println("note_id:", noteId)

	userId := c.Query("uid")
	fmt.Println("user_id:", userId)

	db := postgres.NewService(h.dbInfo.Host, h.dbInfo.Port, h.dbInfo.Login, h.dbInfo.Password, h.dbInfo.Database, h.dbInfo.SSLMode, h.dbInfo.AppName)
	err = db.Open()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	}

	var tx *sql.Tx
	defer func() {
		err := db.RollbackTx(tx)
		if err != nil {
			// 이미 Commit이 수행된 경우에는 아래와 같은 에러메시지 확인 가능함
			//  - sql: transaction has already been committed or rolled back
		} else {
			// 롤백 성공
		}

		err = db.Close()
		if err != nil {
			//
		}
	}()

	// TRANSACTION BEGIN
	tx, err = db.BeginTx()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	}

	// 노트 정보 조회
	query := fmt.Sprintf(`
		SELECT id, folder_id, conversation_id, from_id, to_id, save_at, send_at, title, contents, keywords
		FROM public.note
		WHERE id = '%s' AND conversation_id = '%s'
	`, noteId, conversationId)

	var noteInfo restapi.NoteInfo
	rows, err := db.QueryTx(tx, query)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	} else {
		if len(rows) == 0 {
			c.IndentedJSON(http.StatusNotFound, restapi.Response404{
				Code: http.StatusNotFound,
				Msg:  "there is no target note",
			})
			return
		}

		for i := 0; i < len(rows); i++ {
			noteInfo = restapi.NoteInfo{
				NoteId:         db.GetUUID(rows[i]["id"]),
				FolderId:       db.GetUUID(rows[i]["folder_id"]),
				ConversationId: db.GetUUID(rows[i]["conversation_id"]),
				FromId:         db.GetString(rows[i]["from_id"]),
				ToId:           db.GetString(rows[i]["to_id"]),
				SaveAt:         db.GetString(rows[i]["save_at"]),
				SendAt:         db.GetString(rows[i]["send_at"]),
				Title:          db.GetString(rows[i]["title"]),
				Content:        db.GetString(rows[i]["contents"]),
				Keywords:       db.GetArray(rows[i]["keywords"]),
			}
		}
	}

	// isread 값 갱신
	//  - 상대방이 작성한 노트를 살펴보는 경우
	if userId != noteInfo.FromId {
		query = fmt.Sprintf(`
		UPDATE public.note SET isread = true
		WHERE id = '%s' AND conversation_id = '%s'
	`, noteId, conversationId)

		_, err = db.ExecTx(tx, query)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
				Code: http.StatusInternalServerError,
				Msg:  err.Error(),
			})
			return
		}
	}

	// TRANSACTION COMMIT
	err = db.CommitTx(tx)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		noteInfo,
	)
}

// DeleteConversation godoc
// @Summary      특정 대화방 삭제
// @Description  특정 대화방 삭제
// @Tags         Box 페이지
// @Security	 BasicAuth
// @Param        uid  query     string  false  "user id"
// @Param        cid  query     string  false  "conversation id"
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /box/conversation [delete]
func (h *BoxHandler) DeleteConversation(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	userId := c.Query("uid")
	fmt.Println("user_id:", userId)

	conversationId := c.Query("cid")
	fmt.Println("conversation_id:", conversationId)

	// TODO: 대화방이 삭제되는 경우에 해당 대화방에 남아있는 노트 정보를 서버단에서 어떻게 처리할 것 인가?
	//
	// 1. conversation 테이블에서 conversation_id에 해당하는 row의 delete_at 칼럼을 현재 시간으로 업데이트
	// 2. note 테이블의 conversation_id를 제거하고 folder_id를 d21c43a5-fa35-414f-92bb-b693b60aaee6로 업데이트
	//  - 노트 정보를 일정 기간 유지하기 위해 backup 폴더로 변경
	db := postgres.NewService(h.dbInfo.Host, h.dbInfo.Port, h.dbInfo.Login, h.dbInfo.Password, h.dbInfo.Database, h.dbInfo.SSLMode, h.dbInfo.AppName)
	err = db.Open()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	}

	var tx *sql.Tx
	defer func() {
		err := db.RollbackTx(tx)
		if err != nil {
			// 이미 Commit이 수행된 경우에는 아래와 같은 에러메시지 확인 가능함
			//  - sql: transaction has already been committed or rolled back
		} else {
			// 롤백 성공
		}

		err = db.Close()
		if err != nil {
			//
		}
	}()

	// TRANSACTION BEGIN
	tx, err = db.BeginTx()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	}

	query := fmt.Sprintf(`
		UPDATE public.conversation SET delete_at = '%s'
		WHERE id = '%s'
	`, time.Now().Format("2006-01-02 15:04:05"), conversationId)

	_, err = db.ExecTx(tx, query)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	}

	query = fmt.Sprintf(`
		UPDATE public.note SET conversation_id = NULL, folder_id = 'd21c43a5-fa35-414f-92bb-b693b60aaee6'
		WHERE conversation_id = '%s'
	`, conversationId)

	_, err = db.ExecTx(tx, query)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	}

	// TRANSACTION COMMIT
	err = db.CommitTx(tx)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		restapi.DefaultResponse{
			Result: "true",
			Msg:    "",
		},
	)
}
