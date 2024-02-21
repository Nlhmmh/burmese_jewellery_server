package handler

import (
	"burmese_jewellery/auth"
	"burmese_jewellery/ers"
	"burmese_jewellery/models"
	"burmese_jewellery/orm"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (h *Handler) GetApiProfile(c *gin.Context) {
	accountID := auth.GetUserID(c)

	record, err := orm.FindAccountProfileG(c, accountID)
	if err != nil {
		ers.NotFound.New(err).Abort(c)
		return
	}

	resp, err := models.ConvAccountProfileFromORM(record)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) PostApiProfile(c *gin.Context) {
	accountID := auth.GetUserID(c)

	var req models.AccountProfilePostParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	record := &orm.AccountProfile{
		AccountID: accountID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	if v := req.Birthday; v != nil {
		record.Birthday = null.TimeFrom(v.Time)
	}
	if v := req.Gender; v != nil {
		record.Gender = orm.NullGenderFrom(orm.Gender(*v))
	}
	if err := record.UpsertG(c, true, []string{orm.AccountProfileColumns.AccountID}, boil.Infer(), boil.Infer()); err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.Status(http.StatusOK)

}
