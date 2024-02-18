package handler

import (
	"burmese_jewellery/ers"
	"burmese_jewellery/models"
	"burmese_jewellery/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetApiCategory(c *gin.Context) {
	list, err := orm.Categories().AllG(c)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	resp, err := models.ConvListFromORM(list, models.ConvCategoryFromORM)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, resp)
}
