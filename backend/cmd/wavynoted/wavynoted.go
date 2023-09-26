package main

import (
	"github.com/wavynote/internal/gateway/http"
)

func main() {
	httpServer := http.NewHTTPServer("", 6770, "server.crt", "server.key", 3600, 3600)
	httpServer.StartServer()
}
