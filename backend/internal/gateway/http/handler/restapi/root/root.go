package root

import "github.com/gin-gonic/gin"

type RootHandler struct {
}

func NewRootHandler() *RootHandler {
	h := &RootHandler{}
	return h
}

func (h *RootHandler) GetFolderList(c *gin.Context) {

}
