package write

import "github.com/gin-gonic/gin"

type WriteHandler struct {
}

func NewWriteHandler() *WriteHandler {
	h := &WriteHandler{}
	return h
}

func (h *WriteHandler) SaveNote(c *gin.Context) {

}
