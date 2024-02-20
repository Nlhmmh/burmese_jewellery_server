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

func (h *Handler) PostApiAdminFaq(c *gin.Context) {
	var req models.FAQPostParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	record := &orm.Faq{
		FaqID:    uuid.New().String(),
		Question: req.Question,
		Answer:   req.Answer,
		IsActive: req.IsActive,
	}
	if err := record.InsertG(c, boil.Infer()); err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) PutApiAdminFaqFaqId(c *gin.Context, faqId models.ID) {
	var req models.FAQPutParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	if err := tx.Write(c, func(tx *sql.Tx) *ers.ErrResp {
		record, err := orm.FindFaq(c, tx, faqId.String())
		if err != nil {
			return ers.NotFound.New(err)
		}

		record.Question = req.Question
		record.Answer = req.Answer
		record.IsActive = req.IsActive
		if _, err := record.Update(c, tx, boil.Infer()); err != nil {
			return ers.InternalServer.New(err)
		}

		return nil
	}); err != nil {
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteApiAdminFaqFaqId(c *gin.Context, faqId models.ID) {
	if err := tx.Write(c, func(tx *sql.Tx) *ers.ErrResp {
		record, err := orm.FindFaq(c, tx, faqId.String())
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
