package restapi

const (
	METHOD_POST   = 0
	METHOD_GET    = 1
	METHOD_PUT    = 2
	METHOD_DELETE = 3

	RESTAPI_NAME     = "/wavynoteapi"
	RESTAPI_VERSION  = "v1.0"
	RESTAPI_BASEPATH = RESTAPI_NAME + "/" + RESTAPI_VERSION

	RESTAPI_SERVICENAME_BOX     = "/box"
	RESTAPI_SERVICENAME_KEYWORD = "/keyword"
	RESTAPI_SERVICENAME_WRITE   = "/write"
	RESTAPI_SERVICENAME_PROFILE = "/profile"
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
