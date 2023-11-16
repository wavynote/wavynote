package box

import "github.com/gin-gonic/gin"

type BoxHandler struct {
}

func NewBoxHandler() *BoxHandler {
	h := &BoxHandler{}
	return h
}

func (h *BoxHandler) DeleteConversation(c *gin.Context) {

}
