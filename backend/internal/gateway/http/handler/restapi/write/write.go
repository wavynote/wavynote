package write

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
	"github.com/wavynote/internal/gateway/http/handler/restapi"
)

type WriteHandler struct {
}

func NewWriteHandler() *WriteHandler {
	h := &WriteHandler{}
	return h
}

// SaveNote godoc
// @Summary      내가 쓴 노트 저장
// @Description  내가 쓴 노트 저장
// @Tags         Write 페이지
// @Security	 BasicAuth
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /write/note [post]
func (h *WriteHandler) SaveNote(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
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
