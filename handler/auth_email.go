package handler

import (
	"burmese_jewellery/auth"
	"burmese_jewellery/ers"
	"burmese_jewellery/handler/usecase"
	"burmese_jewellery/mail"
	"burmese_jewellery/models"
	"burmese_jewellery/orm"
	"burmese_jewellery/tx"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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
		).One(c, tx)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return ers.InternalServer.New(err)
		}
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			a = &orm.Account{
				AccountID:     uuid.New().String(),
				LoginType:     orm.LoginTypeEmail,
				Mail:          null.StringFrom(string(req.Mail)),
				AccountStatus: orm.AccountStatusPending,
			}
			if err := a.Insert(c, tx, boil.Infer()); err != nil {
				return ers.InternalServer.New(err)
			}
		}

		apExists, err := orm.AccountProfileExists(c, tx, a.AccountID)
		if err != nil {
			return ers.InternalServer.New(err)
		}
		isRegistered = apExists

		if !isRegistered {
			if err := sendOTP(c, tx, a.AccountID, string(req.Mail)); err != nil {
				return err
			}
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

func (*Handler) PostApiAuthEmailOtp(c *gin.Context) {
	var req models.AuthEmailOTPParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	a, err := orm.Accounts(
		orm.AccountWhere.LoginType.EQ(orm.LoginTypeEmail),
		orm.AccountWhere.Mail.EQ(null.StringFrom(string(req.Mail))),
		qm.Load(orm.AccountRels.AccountOtp),
	).OneG(c)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		ers.UserWithEmailNotExist.Abort(c)
		return
	}
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	if a.R.AccountOtp == nil {
		ers.OTPNotExist.Abort(c)
		return
	}

	if usecase.IsOTPExpired(a.R.AccountOtp.UpdatedAt) {
		ers.OTPExpired.Abort(c)
		return
	}

	if a.R.AccountOtp.Otp != req.Otp {
		ers.OTPWrong.Abort(c)
		return
	}

	c.Status(http.StatusOK)
}

func (*Handler) PostApiAuthEmailOtpResend(c *gin.Context) {
	var req models.AuthEmailOTPResendParam
	if err := c.ShouldBindJSON(&req); err != nil {
		ers.BadRequest.New(err).Abort(c)
		return
	}

	if err := tx.Write(c, func(tx *sql.Tx) *ers.ErrResp {
		a, err := orm.Accounts(
			orm.AccountWhere.LoginType.EQ(orm.LoginTypeEmail),
			orm.AccountWhere.Mail.EQ(null.StringFrom(string(req.Mail))),
		).One(c, tx)
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			return ers.UserWithEmailNotExist
		}
		if err != nil {
			return ers.InternalServer.New(err)
		}

		if err := sendOTP(c, tx, a.AccountID, string(req.Mail)); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return
	}

	c.Status(http.StatusOK)
}

func sendOTP(c *gin.Context, tx boil.ContextExecutor, accID string, accMail string) *ers.ErrResp {
	// generate otp
	otp := usecase.GenOTP()

	// store otp to match in account table
	aOTP := &orm.AccountOtp{
		AccountID: accID,
		Otp:       otp,
		UpdatedAt: time.Now(),
	}
	if err := aOTP.Upsert(c, tx, true, []string{orm.AccountOtpColumns.AccountID}, boil.Infer(), boil.Infer()); err != nil {
		return ers.InternalServer.New(err)
	}

	// send otp to user email
	if err := mail.NewMail().SendMail(accMail, aOTP.Otp); err != nil {
		return ers.InternalServer.New(err)
	}

	return nil
}
