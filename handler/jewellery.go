package handler

import (
	"burmese_jewellery/ers"
	"burmese_jewellery/models"
	"burmese_jewellery/orm"
	"burmese_jewellery/query"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (h *Handler) GetApiJewellery(c *gin.Context, params models.GetApiJewelleryParams) {
	qList := []qm.QueryMod{
		qm.Offset(params.Offset),
		qm.Limit(params.Limit),
	}
	query.EqUUID(qList, params.Id, orm.JewelleryColumns.JewelleryID)
	query.EqUUID(qList, params.CategoryId, orm.JewelleryColumns.CategoryID)
	query.EqUUID(qList, params.GemId, orm.JewelleryColumns.GemID)
	query.EqUUID(qList, params.MaterialId, orm.JewelleryColumns.MaterialID)
	query.Like(qList, params.Name, orm.JewelleryColumns.Name)
	query.Eq(qList, params.IsPublished, orm.JewelleryColumns.IsPublished)

	list, err := orm.Jewelleries(qList...).AllG(c)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	resp, err := models.ConvListFromORM(list, models.ConvJewelleryFromORM)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetApiJewelleryJewelleryId(c *gin.Context, jewelleryId models.ID) {
	record, err := orm.FindJewelleryG(c, jewelleryId.String())
	if err != nil {
		ers.NotFound.New(err).Abort(c)
		return
	}

	resp, err := models.ConvJewelleryFromORM(record)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, resp)
}
