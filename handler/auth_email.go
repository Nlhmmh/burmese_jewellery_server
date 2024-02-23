package handler

import (
	"burmese_jewellery/auth"
	"burmese_jewellery/ers"
	"burmese_jewellery/models"
	"burmese_jewellery/orm"
	"burmese_jewellery/tx"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (*Handler) PostApiAuthEmailRegister(c *gin.Context) {
	var req models.AuthEmailParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	var isRegistered bool

	if err := tx.Write(c, func(tx *sql.Tx) *ers.ErrResp {
		a, err := orm.Accounts(
			orm.AccountWhere.LoginType.EQ(orm.LoginTypeEmail),
			orm.AccountWhere.Mail.EQ(null.StringFrom(string(req.Mail))),
		).OneG(c)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return ers.InternalServer.New(err)
		}
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			a = &orm.Account{
				AccountID:     uuid.New().String(),
				LoginType:     orm.LoginTypeGoogle,
				Mail:          null.StringFrom(string(req.Mail)),
				AccountStatus: orm.AccountStatusActive,
			}
			if err := a.InsertG(c, boil.Infer()); err != nil {
				return ers.InternalServer.New(err)
			}
		}

		apExists, err := orm.AccountProfileExists(c, tx, a.AccountID)
		if err != nil {
			return ers.InternalServer.New(err)
		}
		isRegistered = apExists

		if !isRegistered {
			// TODO: generate otp
			// TODO: store otp to match in account table
			// TODO: send otp to user email
			fmt.Println("todo")
		}

		return nil
	}); err != nil {
		return
	}

	c.JSON(http.StatusOK, &models.AuthResp{
		IsRegistered: isRegistered,
	})
}

func (*Handler) PostApiAuthEmailLogin(c *gin.Context) {
	var req models.AuthEmailParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	a, err := orm.Accounts(
		orm.AccountWhere.LoginType.EQ(orm.LoginTypeEmail),
		orm.AccountWhere.Mail.EQ(null.StringFrom(string(req.Mail))),
	).OneG(c)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		ers.UserWithEmailNotExist.New(err).Abort(c)
		return
	}
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	apExists, err := orm.AccountProfileExistsG(c, a.AccountID)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}
	if !apExists {
		ers.UserWithEmailNotExist.Abort(c)
		return
	}

	if v := a.Password; !v.Valid {
		ers.InternalServer.New(errors.New("account password is null")).Abort(c)
		return
	}

	if err := auth.CompareHashAndPassword(a.Password.String, req.Password); err != nil {
		ers.PasswordWrong.New(err).Abort(c)
		return
	}

	token, err := auth.GenerateToken(a.AccountID, auth.RoleUser)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, token)
}
