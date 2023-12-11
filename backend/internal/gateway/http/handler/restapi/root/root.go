package root

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
	"github.com/wavynote/internal/gateway/http/handler/restapi"
	"github.com/wavynote/internal/platform/dbmsadapter/postgres"
	"github.com/wavynote/internal/wavynote"
)

type RootHandler struct {
	dbInfo wavynote.DataBaseInfo
}

func NewRootHandler(dbInfo wavynote.DataBaseInfo) *RootHandler {
	h := &RootHandler{
		dbInfo: dbInfo,
	}
	return h
}

// GetFolderList godoc
// @Summary      존재하는 모든 폴더 목록 조회
// @Description  존재하는 모든 폴더 목록 조회
// @Tags         Main 페이지
// @Security	 BasicAuth
// @Param        id   query     string  false  "user id"
// @Success      200  {object}  restapi.FolderListResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /main/folderlist [get]
func (h *RootHandler) GetFolderList(c *gin.Context) {
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

	folderList := []restapi.FolderSimpleInfo{}
	folderList = append(folderList, restapi.FolderSimpleInfo{
		FolderId:   "a3106a0c-5ce7-40f6-81f4-ff9b8ebb240b",
		FolderName: "생각정리",
		NoteCount:  5,
	})
	folderList = append(folderList, restapi.FolderSimpleInfo{
		FolderId:   "980e71ba-0395-49aa-833e-3ebc76b3ec88",
		FolderName: "나의웨이비노트",
		NoteCount:  3,
	})

	c.IndentedJSON(
		http.StatusOK,
		restapi.FolderListResponse{
			Folders: folderList,
		},
	)
}

// GetNoteList godoc
// @Summary      특정 폴더에 존재하는 모든 노트 조회
// @Description  특정 폴더에 존재하는 모든 노트 조회
// @Tags         Main 페이지
// @Security	 BasicAuth
// @Param        uid  query     string  false  "user id"
// @Param        fid  query     string  false  "folder id"
// @Success      200  {object}  restapi.NoteListResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /main/notelist [get]
func (h *RootHandler) GetNoteList(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	userId := c.Query("uid")
	fmt.Println("user_id:", userId)

	folderId := c.Query("fid")
	fmt.Println("folder_id:", folderId)

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

// ChangeFolderName godoc
// @Summary      특정 폴더 이름 변경
// @Description  특정 폴더 이름 변경
// @Tags         Main 페이지
// @Param        body body      restapi.ChangeFolderNameRequest  true  "변경하고자 하는 폴더 정보"
// @Security	 BasicAuth
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /main/folder [post]
func (h *RootHandler) ChangeFolderName(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	c.IndentedJSON(
		http.StatusOK,
		restapi.DefaultResponse{
			Result: "true",
			Msg:    "특정 폴더 이름 변경",
		},
	)
}

// RemoveFolder godoc
// @Summary      특정 폴더 삭제
// @Description  특정 폴더 삭제
// @Tags         Main 페이지
// @Param        body body      restapi.RemoveFolderRequest  true  "삭제할 폴다 정보"
// @Security	 BasicAuth
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /main/folder [delete]
func (h *RootHandler) RemoveFolder(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	c.IndentedJSON(
		http.StatusOK,
		restapi.DefaultResponse{
			Result: "true",
			Msg:    "특정 폴더 삭제",
		},
	)
}

// RemoveNote godoc
// @Summary      내가 쓴 특정 노트 삭제
// @Description  내가 쓴 특정 노트 삭제
// @Tags         Main 페이지
// @Param        body body      restapi.RemoveNoteRequest  true  "삭제할 노트 정보"
// @Security	 BasicAuth
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /main/note [delete]
func (h *RootHandler) RemoveNote(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	c.IndentedJSON(
		http.StatusOK,
		restapi.DefaultResponse{
			Result: "true",
			Msg:    "내가 쓴 특정 노트 삭제",
		},
	)
}
