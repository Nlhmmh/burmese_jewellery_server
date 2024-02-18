package handler

import (
	"burmese_jewellery/ers"
	"burmese_jewellery/models"
	"burmese_jewellery/orm"
	"burmese_jewellery/tx"
	"database/sql"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// (GET /api/admin/account)
func (h *Handler) GetApiAdminAccount(c *gin.Context) {
	aList, err := orm.Accounts().AllG(c)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	resp, err := models.ConvListFromORM(aList, models.ConvAccountFromORM)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// (GET /api/admin/account/{account_id})
func (h *Handler) GetApiAdminAccountAccountId(c *gin.Context, accountId models.ID) {
	record, err := orm.FindAccountG(c, accountId.String())
	if err != nil {
		ers.NotFound.New(err).Abort(c)
		return
	}

	resp, err := models.ConvAccountFromORM(record)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// (PUT /api/admin/account/{account_id})
func (h *Handler) PutApiAdminAccountAccountId(c *gin.Context, accountId models.ID) {
	var req models.AccountStatusPutParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	if err := tx.Write(c, func(tx *sql.Tx) *ers.ErrResp {
		record, err := orm.FindAccount(c, tx, accountId.String())
		if err != nil {
			return ers.InternalServer.New(err)
		}

		record.AccountStatus = orm.AccountStatus(req.AccountStatus)
		if _, err := record.Update(c, tx, boil.Infer()); err != nil {
			return ers.InternalServer.New(err)
		}

		return nil
	}); err != nil {
		return
	}

	c.Status(http.StatusOK)
}
