package handler

import (
	"burmese_jewellery/auth"
	"burmese_jewellery/ers"
	"burmese_jewellery/models"
	"burmese_jewellery/orm"
	"burmese_jewellery/tx"

	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// (POST /api/admin/login)
func (h *Handler) PostApiAdminLogin(c *gin.Context) {
	var req models.AccountAdminLoginParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	aa, err := orm.AccountAdmins(
		orm.AccountAdminWhere.Mail.EQ(string(req.Mail)),
	).OneG(c)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		ers.UserWithEmailNotExist.New(err).Abort(c)
		return
	}
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	if err := auth.CompareHashAndPassword(aa.Password, req.Password); err != nil {
		ers.PasswordWrong.New(err).Abort(c)
		return
	}

	token, err := auth.GenerateToken(aa.AccountAdminID, aa.AccountAdminRole)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, token)
}

// (GET /api/admin/account_admin)
func (h *Handler) GetApiAdminAccountAdmin(c *gin.Context) {
	list, err := orm.AccountAdmins().AllG(c)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	resp, err := models.ConvListFromORM(list, models.ConvAccountAdminFromORM)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// (GET /api/admin/account_admin/{account_admins_id})
func (h *Handler) GetApiAdminAccountAdminAccountAdminsId(c *gin.Context, accountAdminsId models.ID) {
	record, err := orm.FindAccountAdminG(c, accountAdminsId.String())
	if err != nil {
		ers.NotFound.New(err).Abort(c)
		return
	}

	resp, err := models.ConvAccountAdminFromORM(record)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// (POST /api/admin/account_admin)
func (h *Handler) PostApiAdminAccountAdmin(c *gin.Context) {
	var req models.AccountAdminPostParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	record := &orm.AccountAdmin{
		AccountAdminID:     uuid.New().String(),
		Mail:               string(req.Mail),
		Password:           req.Password,
		AccountAdminRole:   orm.AccountAdminRole(req.AccountAdminRole),
		AccountAdminStatus: orm.AccountAdminStatus(req.AccountAdminStatus),
	}
	if err := record.InsertG(c, boil.Infer()); err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.Status(http.StatusOK)
}

// (PUT /api/admin/account_admin/{account_admins_id})
func (h *Handler) PutApiAdminAccountAdminAccountAdminsId(c *gin.Context, accountAdminsId models.ID) {
	var req models.AccountAdminPutParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	record, err := orm.FindAccountAdminG(c, accountAdminsId.String())
	if err != nil {
		ers.NotFound.New(err).Abort(c)
		return
	}

	if v := req.Password; v != nil {
		newPW, err := auth.HashPassword(*v)
		if err != nil {
			ers.PasswordWrong.New(err).Abort(c)
			return
		}
		record.Password = newPW
	}
	if v := req.AccountAdminRole; v != nil {
		record.AccountAdminRole = orm.AccountAdminRole(*v)
	}
	if v := req.AccountAdminStatus; v != nil {
		record.AccountAdminStatus = orm.AccountAdminStatus(*v)
	}
	if _, err := record.UpdateG(c, boil.Infer()); err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.Status(http.StatusOK)
}

// (DELETE /api/admin/account_admin/{account_admins_id})
func (h *Handler) DeleteApiAdminAccountAdminAccountAdminsId(c *gin.Context, accountAdminsId models.ID) {
	if err := tx.Write(c, func(tx *sql.Tx) *ers.ErrResp {
		record, err := orm.FindAccountAdmin(c, tx, accountAdminsId.String())
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
