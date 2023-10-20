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
func (h *Handler) GetApiAdminAccountAdmin(c *gin.Context) {}

// (GET /api/admin/account_admin/{account_admins_id})
func (h *Handler) GetApiAdminAccountAdminAccountAdminsId(c *gin.Context, accountAdminsId models.AccountAdminID) {
}

// (POST /api/admin/account_admin)
func (h *Handler) PostApiAdminAccountAdmin(c *gin.Context) {}

// (PUT /api/admin/account_admin/{account_admins_id})
func (h *Handler) PutApiAdminAccountAdminAccountAdminsId(c *gin.Context, accountAdminsId models.AccountAdminID) {
}

// (DELETE /api/admin/account_admin/{account_admins_id})
func (h *Handler) DeleteApiAdminAccountAdminAccountAdminsId(c *gin.Context, accountAdminsId models.AccountAdminID) {
}
