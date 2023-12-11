// wavynote 모듈에 포함된 모든 패키지에서 공통적으로 사용되는 구조체를 정의
package wavynote

type HTTPServerInfo struct {
	Ip       string
	Port     int
	Cert     string
	Pkey     string
	Rtimeout int
	Wtimeout int
}

type DataBaseInfo struct {
	Host     string
	Port     int
	Login    string
	Password string
	Database string
	SSLMode  string
	AppName  string
}
