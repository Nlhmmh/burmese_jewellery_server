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

// Health Check
// (GET /api/health_check)
func (*Handler) GetApiHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
