package profile

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
// @Success      200  {object}  restapi.LandingPageResonse "랜딩페이지 출력 시 필요한 정보"
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

	var reqInfo *restapi.SignInRequest

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

	defer func() {
		err := db.Close()
		if err != nil {
			//
		}
	}()

	// 비밀번호 유효성 검사
	query := fmt.Sprintf(`
		SELECT pwd
		FROM public.user
		WHERE id = '%s'
	`, reqInfo.Id)

	var password string
	rows, err := db.SelectQuery(query)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	} else {
		if len(rows) == 0 {
			c.IndentedJSON(http.StatusBadRequest, restapi.Response400{
				Code: http.StatusBadRequest,
				Msg:  "Invalid Id",
			})
			return
		}

		for i := 0; i < len(rows); i++ {
			password = db.GetString(rows[i]["pwd"])
		}
	}

	if !strings.EqualFold(password, reqInfo.Password) {
		c.IndentedJSON(http.StatusBadRequest, restapi.Response400{
			Code: http.StatusBadRequest,
			Msg:  "Invalid Password",
		})
		return
	}

	// 랜딩페이지 출력을 위한 디폴터 폴더 uuid 조회
	query = fmt.Sprintf(`
		SELECT id
		FROM public.folder
		WHERE user_id = '%s' AND isdefault = true
	`, reqInfo.Id)

	var folderId string
	rows, err = db.SelectQuery(query)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	} else {
		if len(rows) == 0 {
			c.IndentedJSON(http.StatusBadRequest, restapi.Response400{
				Code: http.StatusBadRequest,
				Msg:  "there is no default folder",
			})
			return
		}

		for i := 0; i < len(rows); i++ {
			folderId = db.GetUUID(rows[i]["id"])
		}
	}

	c.IndentedJSON(
		http.StatusOK,
		restapi.LandingPageResonse{
			FolderId: folderId,
			UserId:   reqInfo.Id,
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
	if userId == "" {
		c.IndentedJSON(http.StatusBadRequest, restapi.Response400{
			Code: http.StatusBadRequest,
			Msg:  "empty id",
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

	defer func() {
		err := db.Close()
		if err != nil {
			//
		}
	}()

	// ID 중복 체크
	query := fmt.Sprintf(`
		SELECT count(*) AS cnt
		FROM public.user
		WHERE id = '%s'
	`, userId)

	isValid := true
	rows, err := db.SelectQuery(query)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	} else {
		count := 0
		for i := 0; i < len(rows); i++ {
			count = db.GetInteger(rows[i]["cnt"])
		}

		if count != 0 {
			isValid = false
		}
	}

	var result restapi.DefaultResponse
	if isValid {
		result = restapi.DefaultResponse{
			Result: "true",
			Msg:    "",
		}
	} else {
		result = restapi.DefaultResponse{
			Result: "false",
			Msg:    "duplicated id",
		}
	}

	c.IndentedJSON(
		http.StatusOK,
		result,
	)
}

// SignUp godoc
// @Summary      회원 가입 요청
// @Description  회원 가입 요청
// @Tags         Profile 페이지
// @Param        body body      restapi.SignUpRequest  true  "회원 가입 요청 정보"
// @Security	 BasicAuth
// @Success      200  {object}  restapi.LandingPageResonse "랜딩페이지 출력 시 필요한 정보"
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

	var reqInfo *restapi.SignUpRequest

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

	query := fmt.Sprintf(`
		INSERT INTO public.user(id, pwd, nickname, create_at, keywords, emoji)
		VALUES('%s', '%s', '%s', '%s', ARRAY[%s]::uuid[], '%s')
	`, reqInfo.Id, reqInfo.Password, reqInfo.NickName, time.Now().Format("2006-01-02 15:04:05"), strings.Join(keywordList, ","), reqInfo.Emoji)

	_, err = db.ExecTx(tx, query)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, restapi.Response500{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	}

	// 디폴트 폴더(나의 노트) 생성
	uuid := uuid.New().String()
	query = fmt.Sprintf(`
		INSERT INTO public.folder(id, user_id, name, isdefault) 
		VALUES ('%s', '%s', '나의 노트', true)
	`, uuid, reqInfo.Id)

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
		restapi.LandingPageResonse{
			FolderId: uuid,
			UserId:   reqInfo.Id,
		},
	)
}
