package write

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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
// @Tags         Write 페이지
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

	c.IndentedJSON(
		http.StatusOK,
		restapi.DefaultResponse{
			Result: "true",
			Msg:    "내가 쓴 노트 저장",
		},
	)
}

// SendNote godoc
// @Summary      내가 쓴 노트를 특정 대상에게 보내기
// @Description  내가 쓴 노트를 특정 대상에게 보내기
// @Tags         Write 페이지
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

	c.IndentedJSON(
		http.StatusOK,
		restapi.DefaultResponse{
			Result: "true",
			Msg:    "내가 쓴 노트를 특정 대상에게 보내기",
		},
	)
}

// ShareToOpenNote godoc
// @Summary      내가 쓴 노트를 오픈 노트에 공유하기
// @Description  내가 쓴 노트를 오픈 노트에 공유하기
// @Tags         Write 페이지
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

	c.IndentedJSON(
		http.StatusOK,
		restapi.DefaultResponse{
			Result: "true",
			Msg:    "내가 쓴 노트를 오픈 노트에 공유하기",
		},
	)
}

// SendNoteToRandomUser godoc
// @Summary      내가 쓴 노트를 랜덤 매칭을 통해 (비슷한 관심 주제를 갖는)임의의 대상의 노트로 보내기
// @Description  내가 쓴 노트를 랜덤 매칭을 통해 (비슷한 관심 주제를 갖는)임의의 대상의 노트로 보내기
// @Tags         Write 페이지
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

	c.IndentedJSON(
		http.StatusOK,
		restapi.DefaultResponse{
			Result: "true",
			Msg:    "내가 쓴 노트를 랜덤 매칭을 통해 (비슷한 관심 주제를 갖는)임의의 대상의 노트로 보내기",
		},
	)
}

// ShowNote godoc
// @Summary      내가 쓴 노트 조회
// @Description  내가 쓴 노트 조회
// @Tags         Write 페이지
// @Security	 BasicAuth
// @Param        id  query     string  false  "note id"
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

	noteId := c.Query("id")
	fmt.Println("id:", noteId)

	c.IndentedJSON(
		http.StatusOK,
		restapi.NoteInfo{
			NoteId:   "09d05df1-2958-4a3d-b910-3b4fb079327b",
			FolderId: "980e71ba-0395-49aa-833e-3ebc76b3ec88",
			FromId:   "wavynoteadmin@gmail.com",
			SaveAt:   "2023-11-01 21:00:00",
			Title:    "나의첫번째노트",
			Content:  "나의 첫 웨이비노트 본문 내용입니다.",
			Keywords: []string{"일상∙생각", "마음챙김"},
		},
	)
}
