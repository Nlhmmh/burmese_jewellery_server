package handler

import (
	"burmese_jewellery/ers"
	"burmese_jewellery/models"
	"burmese_jewellery/orm"
	"burmese_jewellery/query"
	"burmese_jewellery/tx"
	"database/sql"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

	resp, err := models.ConvJewelleryListFromORM(list)
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

	resp := &models.Jewellery{}
	if err := resp.ConvFromORM(record); err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) PostApiAdminJewellery(c *gin.Context) {
	var req models.JewelleryPostParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	record := &orm.Jewellery{
		JewelleryID: uuid.New().String(),
		CategoryID:  req.CategoryId.String(),
		GemID:       req.GemId.String(),
		MaterialID:  req.MaterialId.String(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		ImageURL:    req.ImageUrl,
		IsPublished: req.IsPublished,
	}
	if err := record.InsertG(c, boil.Infer()); err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) PutApiAdminJewelleryJewelleryId(c *gin.Context, jewelleryId models.ID) {
	var req models.JewelleryPutParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	if err := tx.Write(c, func(tx *sql.Tx) *ers.ErrResp {
		record, err := orm.FindJewellery(c, tx, jewelleryId.String())
		if err != nil {
			return ers.NotFound.New(err)
		}

		record.CategoryID = req.CategoryId.String()
		record.GemID = req.GemId.String()
		record.MaterialID = req.MaterialId.String()
		record.Name = req.Name
		record.Description = req.Description
		record.Price = req.Price
		record.ImageURL = req.ImageUrl
		record.IsPublished = req.IsPublished
		if _, err := record.Update(c, tx, boil.Infer()); err != nil {
			return ers.InternalServer.New(err)
		}

		return nil
	}); err != nil {
		return
	}

	c.Status(http.StatusOK)
}
