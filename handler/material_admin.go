package handler

import (
	"burmese_jewellery/ers"
	"burmese_jewellery/models"
	"burmese_jewellery/orm"
	"burmese_jewellery/tx"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (h *Handler) PostApiAdminMaterial(c *gin.Context) {
	var req models.MaterialPostParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	record := &orm.Material{
		MaterialID: uuid.New().String(),
		Name:       req.Name,
	}
	if err := record.InsertG(c, boil.Infer()); err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) PutApiAdminMaterialMaterialId(c *gin.Context, materialId models.ID) {
	var req models.MaterialPutParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	if err := tx.Write(c, func(tx *sql.Tx) *ers.ErrResp {
		record, err := orm.FindMaterial(c, tx, materialId.String())
		if err != nil {
			return ers.NotFound.New(err)
		}

		record.Name = req.Name
		if _, err := record.Update(c, tx, boil.Infer()); err != nil {
			return ers.InternalServer.New(err)
		}

		return nil
	}); err != nil {
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteApiAdminMaterialMaterialId(c *gin.Context, materialId models.ID) {
	if err := tx.Write(c, func(tx *sql.Tx) *ers.ErrResp {
		record, err := orm.FindMaterial(c, tx, materialId.String())
		if err != nil {
			return ers.NotFound.New(err)
		}

		if _, err := record.Delete(c, tx); err != nil {
			return ers.InternalServer.New(err)
		}

		return nil
	}); err != nil {
		return
	}

	c.Status(http.StatusNoContent)
}
