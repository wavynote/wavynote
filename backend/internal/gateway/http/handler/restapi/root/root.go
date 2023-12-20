package root

import (
	"database/sql"
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
// @Tags         나의노트 페이지
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

	query := fmt.Sprintf(`
		SELECT id, name
		FROM public.folder
		WHERE user_id = '%s'
		ORDER BY name
	`, userId)

	var folderList []restapi.FolderSimpleInfo
	rows, err := db.SelectQuery(query)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	} else {
		for i := 0; i < len(rows); i++ {
			folderId := db.GetUUID(rows[i]["id"])
			// note 테이블에서 동일한 folder_id를 갖는 row의 개수 조회
			noteCount := 0
			innerQuery := fmt.Sprintf(`
				SELECT count(*) AS cnt
				FROM public.note
				WHERE folder_id = '%s'
			`, folderId)
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

			folderList = append(folderList, restapi.FolderSimpleInfo{
				FolderId:   folderId,
				FolderName: db.GetString(rows[i]["name"]),
				NoteCount:  noteCount,
			})
		}
	}

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
// @Tags         나의노트 페이지
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
		SELECT id, title, contents
		FROM public.note
		WHERE from_id = '%s' AND folder_id = '%s'
		ORDER BY save_at
	`, userId, folderId)

	var noteList []restapi.NoteSimpleInfo
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
			Msg:  "there is no note in target folder",
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

// ChangeNoteFolder godoc
// @Summary      노트가 저장되는 폴더 변경
// @Description  노트가 저장되는 폴더 변경
// @Tags         나의노트 페이지
// @Param        body body      restapi.ChangeNoteFolderRequest  true  "노트를 저장할 (변경할)폴더 정보"
// @Security	 BasicAuth
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /main/notefolder [put]
func (h *RootHandler) ChangeNoteFolder(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	var reqInfo *restapi.ChangeNoteFolderRequest

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

	for _, noteId := range reqInfo.Notes {
		query := fmt.Sprintf(`
			UPDATE public.note SET folder_id = '%s'
			WHERE id = '%s' AND from_id = '%s'
		`, reqInfo.FolderId, noteId, reqInfo.UserId)

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
		restapi.DefaultResponse{
			Result: "true",
			Msg:    "",
		},
	)
}

// AddFolder godoc
// @Summary      폴더 추가
// @Description  폴더 추가
// @Tags         나의노트 페이지
// @Param        body body      restapi.AddFolderRequest  true  "추가할 폴더 정보"
// @Security	 BasicAuth
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /main/folder [post]
func (h *RootHandler) AddFolder(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	var reqInfo *restapi.AddFolderRequest

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
		INSERT INTO public.folder(id, user_id, name) 
		VALUES (uuid_generate_v4(), '%s', '%s')
	`, reqInfo.UserId, reqInfo.FolderName)

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

// ChangeFolderName godoc
// @Summary      특정 폴더 이름 변경
// @Description  특정 폴더 이름 변경
// @Tags         나의노트 페이지
// @Param        body body      restapi.ChangeFolderNameRequest  true  "변경하고자 하는 폴더 정보"
// @Security	 BasicAuth
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /main/folder [put]
func (h *RootHandler) ChangeFolderName(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	var reqInfo *restapi.ChangeFolderNameRequest

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
		UPDATE public.folder SET name = '%s'
		WHERE id = '%s' AND user_id = '%s'
	`, reqInfo.FolderName, reqInfo.FolderId, reqInfo.UserId)

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

// RemoveFolder godoc
// @Summary      특정 폴더 삭제
// @Description  특정 폴더 삭제
// @Tags         나의노트 페이지
// @Param        body body      restapi.RemoveFolderRequest  true  "삭제할 폴더 정보"
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

	var reqInfo *restapi.RemoveFolderRequest

	err = c.BindJSON(&reqInfo)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, restapi.Response400{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	// TODO: 폴더 내에 존재하던 노트들은 어떻게 처리 할 것인가?
	//  - 폴더 테이블에서 관리하는 정보가 데이터로서의 가치가 그렇게 크진않을거라 판단하여 row 자체를 삭제하도록 함
	//  - 단, note 테이블에 포함된 folder_id 칼럼의 경우에는 백업 폴더의 uuid(d21c43a5-fa35-414f-92bb-b693b60aaee6)로 업데이트함

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

	for _, folder := range reqInfo.RemoveFolders {
		query := fmt.Sprintf(`
			DELETE FROM public.folder
			WHERE id = '%s' AND user_id = '%s'
		`, folder.FolderId, reqInfo.UserId)

		_, err = db.ExecTx(tx, query)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
				Code: http.StatusInternalServerError,
				Msg:  err.Error(),
			})
			return
		}

		query = fmt.Sprintf(`
			UPDATE public.note SET folder_id = 'd21c43a5-fa35-414f-92bb-b693b60aaee6'
			WHERE folder_id = '%s' AND from_id = '%s'
		`, folder.FolderId, reqInfo.UserId)

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
		restapi.DefaultResponse{
			Result: "true",
			Msg:    "",
		},
	)
}

// RemoveNote godoc
// @Summary      내가 쓴 특정 노트 삭제
// @Description  내가 쓴 특정 노트 삭제
// @Tags         나의노트 페이지
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

	// 데이터 삭제하지말고 folder만 backup으로 변경해두자
	// TODO: backup 폴더에 존재하는 노트 복구방법은? 영구삭제 시점은?

	var reqInfo *restapi.RemoveNoteRequest

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

	// backup 폴더의 id는 d21c43a5-fa35-414f-92bb-b693b60aaee6로 고정임
	for _, noteInfo := range reqInfo.RemoveNotes {
		query := fmt.Sprintf(`
			UPDATE public.note SET folder_id = 'd21c43a5-fa35-414f-92bb-b693b60aaee6'
			WHERE id = '%s' AND folder_id = '%s' AND from_id = '%s'
		`, noteInfo.NoteId, noteInfo.FolderId, reqInfo.UserId)

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
		restapi.DefaultResponse{
			Result: "true",
			Msg:    "",
		},
	)
}
