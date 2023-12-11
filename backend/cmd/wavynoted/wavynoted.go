package main

import (
	"github.com/wavynote/internal/gateway/http"
	"github.com/wavynote/internal/wavynote"
)

func main() {
	httpInfo := wavynote.HTTPServerInfo{
		Ip:       "",
		Port:     16770,
		Cert:     "server.crt",
		Pkey:     "server.key",
		Rtimeout: 3600,
		Wtimeout: 3600,
	}

	dbInfo := wavynote.DataBaseInfo{
		Host:     "172.25.119.73",
		Port:     5432,
		Login:    "postgres",
		Password: "wavy20230914",
		Database: "wavynote",
		SSLMode:  "disable",
		AppName:  "wavynoted",
	}

	httpServer := http.NewHTTPServer(httpInfo, dbInfo)
	httpServer.StartServer()
}
