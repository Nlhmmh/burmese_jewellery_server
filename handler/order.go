package handler

import (
	"burmese_jewellery/auth"
	"burmese_jewellery/ers"
	"burmese_jewellery/models"
	"burmese_jewellery/orm"
	"burmese_jewellery/tx"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (h *Handler) GetApiOrder(c *gin.Context) {
	accID := auth.GetUserID(c)

	list, err := orm.AccountOrders(
		orm.AccountOrderWhere.AccountID.EQ(accID),
		qm.OrderBy("created_at ASC"),
		qm.Load(orm.AccountOrderRels.OrderAccountOrderJewelleries),
		qm.Load(orm.AccountOrderRels.OrderAccountOrderAddress),
	).AllG(c)
	if err != nil {
		ers.NotFound.New(err).Abort(c)
		return
	}

	resp, err := models.ConvListFromORM(list, models.ConvAccountOrderFromORM)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetApiOrderAccountOrderId(c *gin.Context, accountOrderId models.ID) {
	accID := auth.GetUserID(c)

	record, err := orm.AccountOrders(
		orm.AccountOrderWhere.OrderID.EQ(accountOrderId.String()),
		orm.AccountOrderWhere.AccountID.EQ(accID),
		qm.Load(orm.AccountOrderRels.OrderAccountOrderJewelleries),
		qm.Load(orm.AccountOrderRels.OrderAccountOrderAddress),
	).OneG(c)
	if err != nil {
		ers.NotFound.New(err).Abort(c)
		return
	}

	model, err := models.ConvAccountOrderFromORM(record)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, model)
}

func (h *Handler) PostApiOrder(c *gin.Context) {
	accID := auth.GetUserID(c)

	var req models.OrderPostParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	if err := tx.Write(c, func(tx *sql.Tx) *ers.ErrResp {
		// List CartJewelleries
		list, err := orm.AccountCartJewelleries(
			orm.AccountCartJewelleryWhere.AccountID.EQ(accID),
			qm.OrderBy("created_at ASC"),
			qm.Load(orm.AccountCartJewelleryRels.Jewellery),
		).All(c, tx)
		if err != nil {
			return ers.NotFound.New(err)
		}

		// If cart is empty, return error
		if len(list) == 0 {
			return ers.CartEmpty
		}

		// Create Order
		ao := &orm.AccountOrder{
			OrderID:     uuid.New().String(),
			AccountID:   accID,
			OrderStatus: orm.OrderStatusProceeding,
		}
		if err := ao.Insert(c, tx, boil.Infer()); err != nil {
			return ers.InternalServer.New(err)
		}

		// Create OrderJewelleries
		for _, v := range list {
			aoj := &orm.AccountOrderJewellery{
				OrderID:     ao.OrderID,
				JewelleryID: v.JewelleryID,
				Quantity:    v.Quantity,
				Price:       v.R.Jewellery.Price,
			}
			if err := aoj.Insert(c, tx, boil.Infer()); err != nil {
				return ers.InternalServer.New(err)
			}

			if _, err := v.Delete(c, tx); err != nil {
				return ers.InternalServer.New(err)
			}
		}

		// Create OrderAddress
		aoa := &orm.AccountOrderAddress{
			OrderID:     ao.OrderID,
			CountryCode: req.Address.CountryCode,
			PostCode:    null.StringFromPtr(req.Address.PostCode),
			State:       null.StringFromPtr(req.Address.State),
			City:        req.Address.City,
			Address:     req.Address.Address,
			Remarks:     null.StringFromPtr(req.Address.Remarks),
		}
		if err := aoa.Insert(c, tx, boil.Infer()); err != nil {
			return ers.InternalServer.New(err)
		}

		return nil
	}); err != nil {
		return
	}

	c.Status(http.StatusOK)
}
