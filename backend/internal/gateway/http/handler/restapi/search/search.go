package search

import "github.com/gin-gonic/gin"

type SearchHandler struct {
}

func NewSearchHandler() *SearchHandler {
	h := &SearchHandler{}
	return h
}

func (h *SearchHandler) SearchFromTop(c *gin.Context) {

}
