package profile

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
	"github.com/wavynote/internal/gateway/http/handler/restapi"
	"github.com/wavynote/internal/wavynote"
)

type ProfileHandler struct {
	dbInfo wavynote.DataBaseInfo
}

func NewProfileHandler(dbInfo wavynote.DataBaseInfo) *ProfileHandler {
	h := &ProfileHandler{
		dbInfo: dbInfo,
	}
	return h
}

// SignIn godoc
// @Summary      로그인 요청
// @Description  로그인 요청
// @Tags         Profile 페이지
// @Param        body body      restapi.SignInRequest  true  "로그인 요청"
// @Security	 BasicAuth
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /profile/signin [post]
func (h *ProfileHandler) SignIn(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	c.IndentedJSON(
		http.StatusOK,
		restapi.DefaultResponse{
			Result: "true",
			Msg:    "로그인 요청",
		},
	)
}

// CheckDuplicateID godoc
// @Summary      회원 가입 시 ID 중복 체크
// @Description  회원 가입 시 ID 중복 체크
// @Tags         Profile 페이지
// @Security	 BasicAuth
// @Param        id   query     string  false  "user id"
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /profile/duplicate [get]
func (h *ProfileHandler) CheckDuplicateID(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	userId := c.Query("id")
	fmt.Println("user_id:", userId)

	c.IndentedJSON(
		http.StatusOK,
		restapi.DefaultResponse{
			Result: "true",
			Msg:    "회원 가입 시 ID 중복 체크",
		},
	)
}

// SignUp godoc
// @Summary      회원 가입 요청
// @Description  회원 가입 요청
// @Tags         Profile 페이지
// @Param        body body      restapi.SignUpRequest  true  "회원 가입 요청"
// @Security	 BasicAuth
// @Success      200  {object}  restapi.DefaultResponse ""
// @Failure      400  {object}  restapi.Response400 "요청에 포함된 파라미터 값이 잘못된 경우입니다"
// @Failure		 401  {object}  restapi.Response401 "인증에 실패한 경우이며, 실패 사유가 전달됩니다"
// @Failure      404  {object}  restapi.Response404 "요청한 리소스가 서버에 존재하지 않는 경우입니다"
// @Failure      500  {object}  restapi.Response500 "요청을 처리하는 과정에서 서버에 문제가 발생한 경우입니다"
// @Router       /profile/signup [post]
func (h *ProfileHandler) SignUp(c *gin.Context) {
	dmp, err := httputil.DumpRequest(c.Request, true)
	if err == nil {
		fmt.Printf("dump request:\n%s\n", string(dmp))
	}

	c.IndentedJSON(
		http.StatusOK,
		restapi.DefaultResponse{
			Result: "true",
			Msg:    "회원 가입 요청",
		},
	)
}
