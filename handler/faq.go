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

func (h *Handler) GetApiFaq(c *gin.Context, params models.GetApiFaqParams) {
	qList := []qm.QueryMod{
		qm.Offset(params.Offset),
		qm.Limit(params.Limit),
	}
	qList = query.EqUUID(qList, params.Id, orm.FaqColumns.FaqID)
	qList = query.Like(qList, params.Question, orm.FaqColumns.Question)
	qList = query.Like(qList, params.Answer, orm.FaqColumns.Answer)
	qList = query.Eq(qList, params.IsActive, orm.FaqColumns.IsActive)

	list, err := orm.Faqs(qList...).AllG(c)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	resp, err := models.ConvListFromORM(list, models.ConvFAQFromORM)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, resp)
}
