package search

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
	"github.com/wavynote/internal/gateway/http/handler/restapi"
	"github.com/wavynote/internal/platform/dbmsadapter/postgres"
	"github.com/wavynote/internal/wavynote"
)

type SearchHandler struct {
	dbInfo wavynote.DataBaseInfo
}

func NewSearchHandler(dbInfo wavynote.DataBaseInfo) *SearchHandler {
	h := &SearchHandler{
		dbInfo: dbInfo,
	}
	return h
}

// SearchNoteFromTop godoc
// @Summary      전체 폴더를 대상으로 노트 내용 검색
// @Description  전체 폴더를 대상으로 노트 내용 검색
// @Tags         나의노트 페이지
// @Security	 BasicAuth
// @Param        id     query     string  false  "user id"
// @Param        query  query     string  false  "query for search"
// @Success      200  {object}  restapi.NoteListResponse "검색어가 포함된 노트 정보"
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /search/top [get]
func (h *SearchHandler) SearchNoteFromTop(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	userId := c.Query("id")
	fmt.Println("userId:", userId)

	queryInfo := c.Query("query")
	fmt.Println("query:", queryInfo)

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

	// 내가 작성한 모든 노트를 대상으로 검색
	query := fmt.Sprintf(`
		SELECT id, title, contents
		FROM public.note
		WHERE from_id = '%s'
		AND contents @@ to_tsquery('%s:*')
	`, userId, queryInfo)

	noteList := []restapi.NoteSimpleInfo{}
	rows, err := db.SelectQuery(query)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	} else {
		for i := 0; i < len(rows); i++ {
			noteList = append(noteList, restapi.NoteSimpleInfo{
				NoteId:  db.GetUUID(rows[i]["id"]),
				Title:   db.GetString(rows[i]["title"]),
				Preview: db.GetString(rows[i]["contents"])[0:20], // preview는 본문의 시작부터 최대 20자까지 제공
			})
		}
	}

	if len(noteList) == 0 {
		c.IndentedJSON(http.StatusNotFound, restapi.Response404{
			Code: http.StatusNotFound,
			Msg:  "there is no result",
		})
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		restapi.NoteListResponse{
			Notes: noteList,
		},
	)
}

// SearchNoteFromTargetFolder godoc
// @Summary      특정 폴더를 대상으로 노트 검색
// @Description  특정 폴더를 대상으로 노트 검색
// @Tags         나의노트 페이지
// @Security	 BasicAuth
// @Param        id  query     string  false  "target folder id"
// @Param        query  query     string  false  "query for search"
// @Success      200  {object}  restapi.NoteListResponse "검색어가 포함된 노트 정보"
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /search/folder [get]
func (h *SearchHandler) SearchNoteFromTargetFolder(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	folderId := c.Query("id")
	fmt.Println("folder_id:", folderId)
	queryInfo := c.Query("query")
	fmt.Println("query:", queryInfo)

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

	// 특정 폴더 내에 존재하는 노트를 대상으로 검색
	query := fmt.Sprintf(`
		SELECT id, title, contents
		FROM public.note
		WHERE folder_id = '%s'
		AND contents @@ to_tsquery('%s:*')
	`, folderId, queryInfo)

	noteList := []restapi.NoteSimpleInfo{}
	rows, err := db.SelectQuery(query)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	} else {
		for i := 0; i < len(rows); i++ {
			noteList = append(noteList, restapi.NoteSimpleInfo{
				NoteId:  db.GetUUID(rows[i]["id"]),
				Title:   db.GetString(rows[i]["title"]),
				Preview: db.GetString(rows[i]["contents"])[0:20], // preview는 본문의 시작부터 최대 20자까지 제공
			})
		}
	}

	if len(noteList) == 0 {
		c.IndentedJSON(http.StatusNotFound, restapi.Response404{
			Code: http.StatusNotFound,
			Msg:  "there is no result",
		})
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		restapi.NoteListResponse{
			Notes: noteList,
		},
	)
}

// SearchOpenNote godoc
// @Summary      오픈노트 대상으로 노트 내용 검색
// @Description  오픈노트 대상으로 노트 내용 검색
// @Tags         오픈노트 페이지
// @Security	 BasicAuth
// @Param        query     query     string  false  "query for search"
// @Success      200  {object}  restapi.OpenNoteListResponse "검색어가 포함된 오픈노트 정보"
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /search/opennote [get]
func (h *SearchHandler) SearchOpenNote(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	queryInfo := c.Query("query")
	fmt.Println("query:", queryInfo)

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

	// TODO: 해당 키워드가 포함된 오픈노트 내에서의 검색인가?
	//  - 그렇다면 인자로 선택된 키워드 정보를 받아야함
	//  - 아니라면 전체 오픈노트를 대상으로 검색을 수행하면됨(오픈노트의 조건은 note테이블의 to_id 값은 비어있고 conversation_id와 send_at 값은 존재하는 노트들이다.)
	query := fmt.Sprintf(`
		SELECT id, conversation_id, title, contents, send_at
		FROM public.note
		WHERE to_id is NULL
		AND send_at is not NULL
		AND contents @@ to_tsquery('%s:*')
	`, queryInfo)

	opennoteList := []restapi.OpenNoteSimpleInfo{}
	rows, err := db.SelectQuery(query)
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
			Msg:  "there is no result",
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
