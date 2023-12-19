package write

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wavynote/internal/gateway/http/handler/restapi"
	"github.com/wavynote/internal/platform/dbmsadapter/postgres"
	"github.com/wavynote/internal/wavynote"
)

type WriteHandler struct {
	dbInfo wavynote.DataBaseInfo
}

func NewWriteHandler(dbInfo wavynote.DataBaseInfo) *WriteHandler {
	h := &WriteHandler{
		dbInfo: dbInfo,
	}
	return h
}

// SaveNote godoc
// @Summary      내가 쓴 노트 저장
// @Description  내가 쓴 노트 저장
// @Tags         나의노트 페이지
// @Security	 BasicAuth
// @Param        body body      restapi.SaveNoteRequest  true  "저장할 노트 정보"
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /write/save [post]
func (h *WriteHandler) SaveNote(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	var reqInfo *restapi.SaveNoteRequest

	err = c.BindJSON(&reqInfo)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, restapi.Response400{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

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

	var keywordList []string
	for _, keyword := range reqInfo.Keywords {
		keywordList = append(keywordList, fmt.Sprintf("'%s'", keyword))
	}
	fmt.Println(keywordList)

	query := fmt.Sprintf(`
		INSERT INTO public.note(id, folder_id, from_id, save_at, title, contents, keywords)
		VALUES(uuid_generate_v4(), '%s', '%s', '%s', '%s', '%s', ARRAY[%s]::uuid[])
	`, reqInfo.FolderId, reqInfo.FromId, time.Now().Format("2006-01-02 15:04:05"), reqInfo.Title, reqInfo.Content, strings.Join(keywordList, ","))

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

	// TODO: 성공에 대한 응답값 강화
	c.IndentedJSON(
		http.StatusOK,
		restapi.DefaultResponse{
			Result: "true",
			Msg:    "",
		},
	)
}

// UpdateNote godoc
// @Summary      내가 쓴 노트 갱신
// @Description  내가 쓴 노트 갱신
// @Tags         나의노트 페이지
// @Security	 BasicAuth
// @Param        body body      restapi.UpdateNoteRequest  true  "갱신할 노트 정보"
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /write/save [put]
func (h *WriteHandler) UpdateNote(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	var reqInfo *restapi.UpdateNoteRequest

	err = c.BindJSON(&reqInfo)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, restapi.Response400{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

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

	var keywordList []string
	for _, keyword := range reqInfo.Keywords {
		keywordList = append(keywordList, fmt.Sprintf("'%s'", keyword))
	}

	// 수정될 수 있는 항목
	//  - 제목, 본문, 키워드
	var targets []string
	if reqInfo.Title != "" {
		targets = append(targets, fmt.Sprintf("title = '%s'", reqInfo.Title))
	}

	if reqInfo.Content != "" {
		targets = append(targets, fmt.Sprintf("contents = '%s'", reqInfo.Content))
	}

	if len(keywordList) != 0 {
		targets = append(targets, fmt.Sprintf("keywords = ARRAY[%s]::uuid[]", strings.Join(keywordList, ", ")))
	}

	if len(targets) == 0 {
		c.IndentedJSON(http.StatusBadRequest, restapi.Response400{
			Code: http.StatusBadRequest,
			Msg:  "there is no target to update",
		})
		return
	}

	query := fmt.Sprintf(`
		UPDATE public.note SET %s, save_at = '%s'
		WHERE id = '%s' AND from_id = '%s'
	`, strings.Join(targets, ","), time.Now().Format("2006-01-02 15:04:05"), reqInfo.NoteId, reqInfo.FromId)

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

	// TODO: 성공에 대한 응답값 강화
	c.IndentedJSON(
		http.StatusOK,
		restapi.DefaultResponse{
			Result: "true",
			Msg:    "",
		},
	)

}

// SendNote godoc
// @Summary      내가 쓴 노트를 특정 대상에게 보내기
// @Description  내가 쓴 노트를 특정 대상에게 보내기
// @Tags         나의노트 페이지
// @Security	 BasicAuth
// @Param        body body      restapi.SendNoteRequest  true  "특정 대상에게 보낼 노트 정보"
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /write/send [post]
func (h *WriteHandler) SendNote(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	var reqInfo *restapi.SendNoteRequest

	err = c.BindJSON(&reqInfo)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, restapi.Response400{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

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
		UPDATE public.note SET conversation_id = '%s', to_id = '%s', send_at = '%s'
		WHERE id = '%s' AND from_id = '%s'
	`, reqInfo.ConversationId, reqInfo.ToId, time.Now().Format("2006-01-02 15:04:05"), reqInfo.NoteId, reqInfo.FromId)

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

// ShareToOpenNote godoc
// @Summary      내가 쓴 노트를 오픈 노트에 공유하기
// @Description  내가 쓴 노트를 오픈 노트에 공유하기
// @Tags         나의노트 페이지
// @Security	 BasicAuth
// @Param        body body      restapi.ShareNoteRequest  true  "오픈 노트에 공유를 위한 정보"
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /write/opennote [post]
func (h *WriteHandler) ShareToOpenNote(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	var reqInfo *restapi.ShareNoteRequest

	err = c.BindJSON(&reqInfo)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, restapi.Response400{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

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

	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// 대화방 생성
	uuid := uuid.New().String()
	query := fmt.Sprintf(`
		INSERT INTO public.conversation(id, host_id, create_at)
		VALUES('%s', '%s', '%s')
	`, uuid, reqInfo.HostId, currentTime)

	_, err = db.ExecTx(tx, query)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	}

	// 공유할 노트의 conversation_id 및 send_at 값 갱신
	//  - 오픈노트에 공유하는 시점을 send_at에 기록함(to_id 칼럼 값이 비어있고 send_at에 값이 존재하는 경우는 오픈노트에 공유된 노트임)
	query = fmt.Sprintf(`
		UPDATE public.note SET conversation_id = '%s', send_at = '%s'
		WHERE id = '%s' AND from_id = '%s'
	`, uuid, currentTime, reqInfo.NoteId, reqInfo.HostId)

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

// SendNoteToRandomUser godoc
// @Summary      내가 쓴 노트를 랜덤 매칭을 통해 (비슷한 관심 주제를 갖는)임의의 대상에게 보내기
// @Description  내가 쓴 노트를 랜덤 매칭을 통해 (비슷한 관심 주제를 갖는)임의의 대상에게 보내기
// @Tags         나의노트 페이지
// @Security	 BasicAuth
// @Param        body body      restapi.RandomMatchRequest  true  "랜덤 매칭을 통해 보낼 노트 정보"
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /write/random [post]
func (h *WriteHandler) SendNoteToRandomUser(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	// TODO: 구현
	//  - conversation 테이블의 guest_id 칼럼을 갱신해야함

	c.IndentedJSON(
		http.StatusOK,
		restapi.DefaultResponse{
			Result: "true",
			Msg:    "[구현중] 내가 쓴 노트를 랜덤 매칭을 통해 (비슷한 관심 주제를 갖는)임의의 대상의 노트로 보내기",
		},
	)
}

// ShowNote godoc
// @Summary      내가 쓴 노트 조회
// @Description  내가 쓴 노트 조회
// @Tags         나의노트 페이지
// @Security	 BasicAuth
// @Param        uid  query     string  false  "user id"
// @Param        nid  query     string  false  "note id"
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /write/show [get]
func (h *WriteHandler) ShowNote(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	userId := c.Query("uid")
	fmt.Println("id:", userId)

	noteId := c.Query("nid")
	fmt.Println("id:", noteId)

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
		SELECT id, folder_id, conversation_id, from_id, to_id, save_at, send_at, title, contents, keywords
		FROM public.note
		WHERE id = '%s' AND from_id = '%s'
	`, noteId, userId)

	var noteInfo restapi.NoteInfo
	rows, err := db.SelectQuery(query)
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

	c.IndentedJSON(
		http.StatusOK,
		noteInfo,
	)
}
