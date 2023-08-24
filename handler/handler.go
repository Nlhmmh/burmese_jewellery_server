package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

var _ ServerInterface = (*Handler)(nil)

func NewHandler() *Handler {
	return &Handler{}
}

func (*Handler) GetApiHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

func (*Handler) GetApiJewellery(c *gin.Context) {
	panic("unimplemented")
}
