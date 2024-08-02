package handler

import (
	"burmese_jewellery/ers"
	"burmese_jewellery/models"
	"burmese_jewellery/orm"
	"burmese_jewellery/orm_custom"
	"burmese_jewellery/query"
	"burmese_jewellery/tx"
	"database/sql"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// (GET /api/admin/account)
func (h *Handler) GetApiAdminAccount(c *gin.Context, params models.GetApiAdminAccountParams) {
	var resp GetResp

	if err := tx.Write(c, func(tx *sql.Tx) *ers.ErrResp {
		qList := []qm.QueryMod{}
		qList = query.EqUUID(qList, params.Id, orm.AccountColumns.AccountID)
		qList = query.Like(qList, params.FirstName, fmt.Sprintf("%s.%s", orm.TableNames.AccountProfiles, orm.AccountProfileColumns.FirstName))
		qList = query.Like(qList, params.LastName, fmt.Sprintf("%s.%s", orm.TableNames.AccountProfiles, orm.AccountProfileColumns.LastName))
		qList = query.Eq(qList, params.AccountStatus, orm.AccountColumns.AccountStatus)
		qList = append(qList, qm.InnerJoin(
			fmt.Sprintf(
				"%s ON %s.%s = %s.%s",
				orm.TableNames.AccountProfiles,
				orm.TableNames.AccountProfiles,
				orm.AccountProfileColumns.AccountID,
				orm.TableNames.Accounts,
				orm.AccountColumns.AccountID,
			),
		))
		count, err := orm.Accounts(qList...).Count(c, tx)
		if err != nil {
			return ers.InternalServer.New(err)
		}
		resp.Count = count

		sort := models.Desc
		if v := params.Sort; v != nil {
			sort = *v
		}
		qList = append(qList, []qm.QueryMod{
			qm.Select("accounts.*, account_profiles.*"),
			qm.From(orm.TableNames.Accounts),
			qm.Offset(params.Offset),
			qm.Limit(params.Limit),
			qm.OrderBy(fmt.Sprintf("%s.%s %s", orm.TableNames.Accounts, orm.AccountColumns.CreatedAt, string(sort))),
		}...)

		var list []*orm_custom.AccountWithProfile
		if err := orm.NewQuery(qList...).Bind(c, tx, &list); err != nil {
			return ers.InternalServer.New(err)
		}

		data, err := models.ConvListFromORM(list, models.ConvAccountWithProfileFromORM)
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
