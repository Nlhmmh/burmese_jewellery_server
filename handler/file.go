package handler

import (
	"burmese_jewellery/env"
	"burmese_jewellery/ers"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func getFullFilePath(fileName string) string {
	return fmt.Sprintf("%s/%s", env.Get().FilePath, fileName)
}

// (POST /api/file)
func (h *Handler) PostApiFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	filePath := strconv.FormatInt(time.Now().Unix(), 10) + "-" + file.Filename
	if err := c.SaveUploadedFile(file, getFullFilePath(filePath)); err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, filePath)
}

// (GET /api/file/{file_name})
func (h *Handler) GetApiFileFileName(c *gin.Context, fileName string) {
	c.File(getFullFilePath(fileName))
}

// (DELETE /api/admin/file/{file_name})
func (h *Handler) DeleteApiAdminFileFileName(c *gin.Context, fileName string) {
	if err := os.Remove(getFullFilePath(fileName)); err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.Status(http.StatusOK)
}
