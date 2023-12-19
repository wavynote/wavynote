package opennote

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wavynote/internal/gateway/http/handler/restapi"
	"github.com/wavynote/internal/platform/dbmsadapter/postgres"
	"github.com/wavynote/internal/wavynote"
)

type OpenNoteHandler struct {
	dbInfo wavynote.DataBaseInfo
}

func NewOpenNotehHandler(dbInfo wavynote.DataBaseInfo) *OpenNoteHandler {
	h := &OpenNoteHandler{
		dbInfo: dbInfo,
	}
	return h
}

// GetOpenNoteList godoc
// @Summary      사용자 키워드 기반으로 오픈노트에 공유된 노트 목록 조회
// @Description  사용자 키워드 기반으로 오픈노트에 공유된 노트 목록 조회
// @Tags         오픈노트 페이지
// @Security	 BasicAuth
// @Param        uid     query     string  false  "user id"
// @Success      200  {object}  restapi.OpenNoteListResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /opennote/list [get]
func (h *OpenNoteHandler) GetOpenNoteList(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	userId := c.Query("uid")
	fmt.Println("userId:", userId)

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

	// 사용자가 설정한 키워드 조회
	query := fmt.Sprintf(`
		SELECT keywords
		FROM public.user
		WHERE id = '%s'
	`, userId)

	var keywords []string
	rows, err := db.SelectQuery(query)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	} else {
		for i := 0; i < len(rows); i++ {
			keywords = db.GetArray(rows[i]["keywords"])
		}
	}

	if len(keywords) == 0 {
		c.IndentedJSON(http.StatusNotFound, restapi.Response404{
			Code: http.StatusNotFound,
			Msg:  "there is no keywords in target user information",
		})
		return
	}

	var where []string
	for _, keyword := range keywords {
		where = append(where, fmt.Sprintf("'%s' = ANY(keywords)", keyword))
	}

	// 해당 키워드가 포함된 오픈노트 조회
	//  - 오픈노트의 조건은 note테이블의 to_id 값은 비어있고 conversation_id와 send_at 값은 존재하는 노트들이다.
	query = fmt.Sprintf(`
		SELECT id, conversation_id, title, contents, send_at
		FROM public.note
		WHERE %s
		AND to_id is NULL
		AND send_at is not NULL
	`, strings.Join(where, " OR "))
	fmt.Println(strings.Join(where, " OR "))

	opennoteList := []restapi.OpenNoteSimpleInfo{}
	rows, err = db.SelectQuery(query)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	} else {
		for i := 0; i < len(rows); i++ {
			opennoteList = append(opennoteList, restapi.OpenNoteSimpleInfo{
				NoteId:         db.GetUUID(rows[i]["id"]),
				ConversationId: db.GetUUID(rows[i]["conversation_id"]),
				Title:          db.GetString(rows[i]["title"]),
				Preview:        db.GetString(rows[i]["contents"])[0:20], // preview는 본문의 시작부터 최대 20자까지 제공
				PostAt:         db.GetString(rows[i]["send_at"]),
			})
		}
	}

	if len(opennoteList) == 0 {
		c.IndentedJSON(http.StatusNotFound, restapi.Response404{
			Code: http.StatusNotFound,
			Msg:  "there is no opennote",
		})
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		restapi.OpenNoteListResponse{
			Notes: opennoteList,
		},
	)
}

// ShowOpenNote godoc
// @Summary      특정 오픈노트 조회
// @Description  특정 오픈노트 조회
// @Tags         오픈노트 페이지
// @Security	 BasicAuth
// @Param        nid     query     string  false  "note id"
// @Param        cid     query     string  false  "conversation id"
// @Success      200  {object}  restapi.NoteInfo ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /opennote/show [get]
func (h *OpenNoteHandler) ShowOpenNote(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	noteId := c.Query("nid")
	fmt.Println("noteId:", noteId)

	conversationId := c.Query("cid")
	fmt.Println("conversationId:", conversationId)

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
		WHERE id = '%s' AND conversation_id = '%s'
	`, noteId, conversationId)

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
