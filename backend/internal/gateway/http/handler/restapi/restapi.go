package restapi

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	METHOD_POST   = 0
	METHOD_GET    = 1
	METHOD_PUT    = 2
	METHOD_DELETE = 3

	RESTAPI_NAME     = "/wavynote"
	RESTAPI_VERSION  = "v1.0"
	RESTAPI_BASEPATH = RESTAPI_NAME + "/" + RESTAPI_VERSION

	RESTAPI_SERVICENAME_MAIN    = "/main"
	RESTAPI_SERVICENAME_SEARCH  = "/search"
	RESTAPI_SERVICENAME_WRITE   = "/write"
	RESTAPI_SERVICENAME_BOX     = "/box"
	RESTAPI_SERVICENAME_PROFILE = "/profile"

	BASIC_AUTH_USER     = "wavynote"
	BASIC_AUTH_PASSWORD = "wavy20230914"

	LOCATION_FOR_MAIN_FOLDERLIST = "/folderlist"
	LOCATION_FOR_MAIN_NOTELIST   = "/notelist"
	LOCATION_FOR_MAIN_FOLDER     = "/folder"
	LOCATION_FOR_MAIN_NOTE       = "/note"

	LOCATION_FOR_SEARCH_FROM_TOP    = "/top"
	LOCATION_FOR_SEARCH_FROM_FOLDER = "/folder"

	LOCATION_FOR_WRITE_SAVE     = "/save"
	LOCATION_FOR_WRITE_SEND     = "/send"
	LOCATION_FOR_WRITE_OPENNOTE = "/opennote"
	LOCATION_FOR_WRITE_RANDOM   = "/random"
	LOCATION_FOR_WRITE_SHOW     = "/show"

	LOCATION_FOR_BOX_CONVERSATION_LIST = "/conversationlist"
	LOCATION_FOR_BOX_CONVERSATION      = "/conversation"
	LOCATION_FOR_BOX_NOTELIST          = "/notelist"
	LOCATION_FOR_BOX_SHOW              = "/show"
)

type Response400 struct {
	Code int    `json:"code" example:"400"`
	Msg  string `json:"msg" example:"Bad request"`
}

type Response401 struct {
	Code int    `json:"code" example:"401"`
	Msg  string `json:"msg" example:"Fail to parse Authorization header"`
}

type Response403 struct {
	Code int    `json:"code" example:"403"`
	Msg  string `json:"msg" example:"you can't access this target resource when using swagger API test"`
}

type Response404 struct {
	Code int    `json:"code" example:"404"`
	Msg  string `json:"msg" example:"Not exist the target resource in our server"`
}

type Response500 struct {
	Code int    `json:"code" example:"500"`
	Msg  string `json:"msg" example:"Internal server error"`
}

func BasicAuth(c *gin.Context) {
	user, pwd, ok := c.Request.BasicAuth()
	if !ok {
		c.Abort()
		c.Header("WWW-Authenticate", `Basic realm="webkeeper"`)
		c.IndentedJSON(http.StatusUnauthorized, Response401{
			Code: http.StatusUnauthorized,
			Msg:  fmt.Sprintf("%d %s fail to parse Authorization header", http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)),
		})
		return
	}

	if user != BASIC_AUTH_USER {
		c.Abort()
		c.Header("WWW-Authenticate", `Basic realm="webkeeper"`)
		c.IndentedJSON(http.StatusUnauthorized, Response401{
			Code: http.StatusUnauthorized,
			Msg:  fmt.Sprintf("%d %s user incorrect", http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)),
		})
		return
	}

	if pwd != BASIC_AUTH_PASSWORD {
		c.Abort()
		c.Header("WWW-Authenticate", `Basic realm="webkeeper"`)
		c.IndentedJSON(http.StatusUnauthorized, Response401{
			Code: http.StatusUnauthorized,
			Msg:  fmt.Sprintf("%d %s password incorrect", http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)),
		})
		return
	}
}
