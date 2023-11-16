package http

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wavynote/internal/gateway/http/handler/restapi"
	"github.com/wavynote/internal/gateway/http/handler/restapi/box"
	"github.com/wavynote/internal/gateway/http/handler/restapi/root"
	"github.com/wavynote/internal/gateway/http/handler/restapi/search"
	"github.com/wavynote/internal/gateway/http/handler/restapi/write"
)

type HTTPServer struct {
	ip       string
	port     int
	cert     string
	pkey     string
	rtimeout int
	wtimeout int
}

func NewHTTPServer(ip string, port int, cert string, pkey string, rtimeout int, wtimeout int) *HTTPServer {
	httpServer := &HTTPServer{
		ip:       ip,
		port:     port,
		cert:     cert,
		pkey:     pkey,
		rtimeout: rtimeout,
		wtimeout: wtimeout,
	}
	return httpServer
}

func (h *HTTPServer) StartServer() {
	connInfo := ":" + strconv.Itoa(h.port)

	// https://wiki.mozilla.org/Security/Server_Side_TLS
	//
	// not supported in go
	//  0x00,0x9E - DHE-RSA-AES128-GCM-SHA256
	//  0x00,0x9F - DHE-RSA-AES256-GCM-SHA384
	tlsCfg := &tls.Config{
		MinVersion:       tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256, tls.CurveP384},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_CHACHA20_POLY1305_SHA256,
		},
	}

	router := gin.Default()

	// NoRoute(404 Not Found)
	router.NoRoute(func(c *gin.Context) {
		c.IndentedJSON(
			http.StatusNotFound, restapi.Response404{
				Code: http.StatusNotFound,
				Msg:  fmt.Sprintf("%d %s no route", http.StatusNotFound, http.StatusText(http.StatusNotFound)),
			})
	})

	api := router.Group(restapi.RESTAPI_BASEPATH)
	{
		m := api.Group(restapi.RESTAPI_SERVICENAME_MAIN)
		{
			m.GET(
				restapi.LOCATION_FOR_MAIN_FOLDERLIST,
				root.NewRootHandler().GetFolderList,
			)
		}

		w := api.Group(restapi.RESTAPI_SERVICENAME_WRITE)
		{
			w.POST(
				restapi.LOCATION_FOR_WRITE_SAVE,
				write.NewWriteHandler().SaveNote,
			)
		}

		s := api.Group(restapi.RESTAPI_SERVICENAME_SEARCH)
		{
			s.GET(
				restapi.LOCATION_FOR_SEARCH_FROM_TOP,
				search.NewSearchHandler().SearchFromTop,
			)
		}

		b := api.Group(restapi.RESTAPI_SERVICENAME_BOX)
		{
			b.DELETE(
				restapi.LOCATION_FOR_BOX_CONVERSATION,
				box.NewBoxHandler().DeleteConversation,
			)
		}

		// p := api.Group(restapi.RESTAPI_SERVICENAME_PROFILE)
		// {

		// }
	}

	tlsSrv := &http.Server{
		Addr:         connInfo,
		TLSConfig:    tlsCfg,
		Handler:      router,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0), // disable HTTP2
		ReadTimeout:  time.Duration(h.rtimeout) * time.Second,
		WriteTimeout: time.Duration(h.wtimeout) * time.Second,
	}

	err := tlsSrv.ListenAndServeTLS(h.cert, h.pkey)
	if err != nil {
		return
	}
}
