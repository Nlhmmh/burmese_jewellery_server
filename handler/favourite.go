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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (h *Handler) GetApiFavourite(c *gin.Context) {
	accID := auth.GetUserID(c)

	list, err := orm.AccountFavourites(
		orm.AccountFavouriteWhere.AccountID.EQ(accID),
		qm.OrderBy("created_at ASC"),
	).AllG(c)
	if err != nil {
		ers.NotFound.New(err).Abort(c)
		return
	}

	resp, err := models.ConvListFromORM(list, models.ConvAccountFavouriteFromORM)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) PostApiFavourite(c *gin.Context) {
	accID := auth.GetUserID(c)

	var req models.FavouritePostParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	if err := tx.Write(c, func(tx *sql.Tx) *ers.ErrResp {
		record := &orm.AccountFavourite{
			AccountID:   accID,
			JewelleryID: req.JewelleryId.String(),
		}
		if req.IsFavourite {
			if err := record.Upsert(c, tx, false, []string{orm.AccountFavouriteColumns.AccountID, orm.AccountFavouriteColumns.JewelleryID}, boil.Infer(), boil.Infer()); err != nil {
				return ers.InternalServer.New(err)
			}
		} else {
			if _, err := record.Delete(c, tx); err != nil {
				return ers.InternalServer.New(err)
			}
		}

		return nil
	}); err != nil {
		return
	}

	c.Status(http.StatusOK)
}
