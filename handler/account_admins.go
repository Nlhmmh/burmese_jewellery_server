package handler

import (
	"burmese_jewellery/auth"
	"burmese_jewellery/ers"
	"burmese_jewellery/models"
	"burmese_jewellery/orm"

	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// (POST /api/admin/login)
func (h *Handler) PostApiAdminLogin(c *gin.Context) {
	var req models.AccountAdminLoginParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	aa, err := orm.AccountAdmins(
		qm.Where("mail=?", req.Mail),
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
	aaList, err := orm.AccountAdmins().AllG(c)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	resp, err := models.ConvAccountAdminListFromORM(aaList)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// (GET /api/admin/account_admin/{account_admins_id})
func (h *Handler) GetApiAdminAccountAdminAccountAdminsId(c *gin.Context, accountAdminsId models.ID) {
	aa, err := orm.AccountAdmins(
		qm.Where("account_admin_id=?", accountAdminsId),
	).OneG(c)
	if err != nil {
		ers.NotFound.New(err).Abort(c)
		return
	}

	resp := &models.AccountAdmin{}
	if err := resp.ConvFromORM(aa); err != nil {
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

	aa := &orm.AccountAdmin{
		AccountAdminID:     uuid.New().String(),
		Mail:               string(req.Mail),
		Password:           req.Password,
		AccountAdminRole:   orm.AccountAdminRole(req.AccountAdminRole),
		AccountAdminStatus: orm.AccountAdminStatus(req.AccountAdminStatus),
	}
	if err := aa.InsertG(c, boil.Infer()); err != nil {
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

	aa, err := orm.AccountAdmins(
		qm.Where("account_admin_id=?", accountAdminsId),
	).OneG(c)
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
		aa.Password = newPW
	}
	if v := req.AccountAdminRole; v != nil {
		aa.AccountAdminRole = orm.AccountAdminRole(*v)
	}
	if v := req.AccountAdminStatus; v != nil {
		aa.AccountAdminStatus = orm.AccountAdminStatus(*v)
	}
	if _, err := aa.UpdateG(c, boil.Infer()); err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.Status(http.StatusOK)
}

// (DELETE /api/admin/account_admin/{account_admins_id})
func (h *Handler) DeleteApiAdminAccountAdminAccountAdminsId(c *gin.Context, accountAdminsId models.ID) {
	aa, err := orm.AccountAdmins(
		qm.Where("account_admin_id=?", accountAdminsId),
	).OneG(c)
	if err != nil {
		ers.NotFound.New(err).Abort(c)
		return
	}

	if _, err := aa.DeleteG(c); err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.Status(http.StatusNoContent)
}
