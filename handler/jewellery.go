package handler

import (
	"burmese_jewellery/ers"
	"burmese_jewellery/models"
	"burmese_jewellery/orm"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (h *Handler) GetApiJewellery(c *gin.Context) {
	list, err := orm.Jewelleries().AllG(c)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	resp, err := models.ConvJewelleryListFromORM(list)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) PostApiJewellery(c *gin.Context) {
	var req models.JewelleryPostParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	model := &orm.Jewellery{
		JewelleryID: uuid.New().String(),
	}
	if err := model.InsertG(c, boil.Infer()); err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.Status(http.StatusOK)
}
