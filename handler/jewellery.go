package handler

import (
	"burmese_jewellery/ers"
	"burmese_jewellery/models"
	"burmese_jewellery/orm"
	"burmese_jewellery/query"
	"burmese_jewellery/tx"
	"database/sql"
	"fmt"

	"net/http"

	"github.com/aarondl/sqlboiler/v4/queries/qm"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetApiJewellery(c *gin.Context, params models.GetApiJewelleryParams) {
	var resp GetResp

	if err := tx.Write(c, func(tx *sql.Tx) *ers.ErrResp {
		qList := []qm.QueryMod{}
		qList = query.EqUUID(qList, params.Id, orm.JewelleryColumns.JewelleryID)
		qList = query.EqUUID(qList, params.CategoryId, orm.JewelleryColumns.CategoryID)
		qList = query.EqUUID(qList, params.GemId, orm.JewelleryColumns.GemID)
		qList = query.EqUUID(qList, params.MaterialId, orm.JewelleryColumns.MaterialID)
		qList = query.Like(qList, params.Name, orm.JewelleryColumns.Name)
		qList = query.Eq(qList, params.IsPublished, orm.JewelleryColumns.IsPublished)
		count, err := orm.Jewelleries(qList...).Count(c, tx)
		if err != nil {
			return ers.InternalServer.New(err)
		}
		resp.Count = count

		sort := models.Desc
		if v := params.Sort; v != nil {
			sort = *v
		}
		qList = append(qList, []qm.QueryMod{
			qm.Offset(params.Offset),
			qm.Limit(params.Limit),
			qm.OrderBy(fmt.Sprintf("%s %s", orm.AccountColumns.CreatedAt, string(sort))),
		}...)
		list, err := orm.Jewelleries(qList...).AllG(c)
		if err != nil {
			return ers.InternalServer.New(err)
		}

		data, err := models.ConvListFromORM(list, models.ConvJewelleryFromORM)
		if err != nil {
			return ers.InternalServer.New(err)
		}
		resp.Data = data

		return nil
	}); err != nil {
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
