package search

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
	"github.com/wavynote/internal/gateway/http/handler/restapi"
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

	query := c.Query("query")
	fmt.Println("query:", query)

	// TODO: 구현

	noteList := []restapi.NoteSimpleInfo{}
	noteList = append(noteList, restapi.NoteSimpleInfo{
		NoteId:  "a3106a0c-5ce7-40f6-81f4-ff9b8ebb240b",
		Title:   "나의첫번째노트",
		Preview: "나의 첫 웨이비노트 본문 내용입니다.",
	})
	noteList = append(noteList, restapi.NoteSimpleInfo{
		NoteId:  "1a092b35-dc9e-472e-be39-7391ca176040",
		Title:   "나의두번째노트",
		Preview: "나의 두번째 웨이비노트 본문 내용입니다.",
	})

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
	query := c.Query("query")
	fmt.Println("folder_id:", folderId)
	fmt.Println("query:", query)

	// TODO: 구현

	noteList := []restapi.NoteSimpleInfo{}
	noteList = append(noteList, restapi.NoteSimpleInfo{
		NoteId:  "a3106a0c-5ce7-40f6-81f4-ff9b8ebb240b",
		Title:   "나의첫번째노트",
		Preview: "나의 첫 웨이비노트 본문 내용입니다.",
	})
	noteList = append(noteList, restapi.NoteSimpleInfo{
		NoteId:  "1a092b35-dc9e-472e-be39-7391ca176040",
		Title:   "나의두번째노트",
		Preview: "나의 두번째 웨이비노트 본문 내용입니다.",
	})

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

	query := c.Query("query")
	fmt.Println("query:", query)
}
