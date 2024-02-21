package handler

import (
	"burmese_jewellery/auth"
	"burmese_jewellery/ers"
	"burmese_jewellery/handler/usecase"
	"burmese_jewellery/models"
	"burmese_jewellery/orm"
	"burmese_jewellery/tx"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (h *Handler) GetApiCart(c *gin.Context) {
	accID := auth.GetUserID(c)

	list, err := orm.AccountCartJewelleries(
		orm.AccountCartJewelleryWhere.AccountID.EQ(accID),
		qm.OrderBy("created_at ASC"),
	).AllG(c)
	if err != nil {
		ers.NotFound.New(err).Abort(c)
		return
	}

	resp, err := models.ConvListFromORM(list, models.ConvAccountCartJewelleryFromORM)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) PostApiCart(c *gin.Context) {
	accID := auth.GetUserID(c)

	var req models.CartPostParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	if err := tx.Write(c, func(tx *sql.Tx) *ers.ErrResp {
		if err := usecase.Upsert(
			func() (*orm.AccountCartJewellery, error) {
				return orm.FindAccountCartJewellery(c, tx, accID, req.JewelleryId.String())
			},
			func() error {
				record := &orm.AccountCartJewellery{
					AccountID:   accID,
					JewelleryID: req.JewelleryId.String(),
					Quantity:    req.Quantity,
				}
				return record.Insert(c, tx, boil.Infer())
			},
			func(record *orm.AccountCartJewellery) error {
				if req.Quantity == 0 {
					_, err := record.Delete(c, tx)
					return err
				}
				record.Quantity = req.Quantity
				_, err := record.Update(c, tx, boil.Infer())
				return err
			},
		); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return
	}

	c.Status(http.StatusOK)
}
