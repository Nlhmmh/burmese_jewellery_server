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

func (h *Handler) PostApiAdminCategory(c *gin.Context) {
	var req models.CategoryPostParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	record := &orm.Category{
		CategoryID:  uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
		ImageURL:    req.ImageUrl,
	}
	if err := record.InsertG(c, boil.Infer()); err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) PutApiAdminCategoryCategoryId(c *gin.Context, categoryId models.ID) {
	var req models.CategoryPutParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	if err := tx.Write(c, func(tx *sql.Tx) *ers.ErrResp {
		record, err := orm.FindCategory(c, tx, categoryId.String())
		if err != nil {
			return ers.NotFound.New(err)
		}

		record.Name = req.Name
		record.Description = req.Description
		record.ImageURL = req.ImageUrl
		if _, err := record.Update(c, tx, boil.Infer()); err != nil {
			return ers.InternalServer.New(err)
		}

		return nil
	}); err != nil {
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteApiAdminCategoryCategoryId(c *gin.Context, categoryId models.ID) {
	if err := tx.Write(c, func(tx *sql.Tx) *ers.ErrResp {
		record, err := orm.FindCategory(c, tx, categoryId.String())
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
