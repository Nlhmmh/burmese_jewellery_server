// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package models

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for AccountAccountStatus.
const (
	AccountAccountStatusActive      AccountAccountStatus = "active"
	AccountAccountStatusDeactivated AccountAccountStatus = "deactivated"
	AccountAccountStatusLocked      AccountAccountStatus = "locked"
	AccountAccountStatusPending     AccountAccountStatus = "pending"
)

// Defines values for AccountLoginType.
const (
	Email    AccountLoginType = "email"
	Facebook AccountLoginType = "facebook"
	Google   AccountLoginType = "google"
	Phone    AccountLoginType = "phone"
)

// Defines values for AccountAdminAccountAdminRole.
const (
	AccountAdminAccountAdminRoleAdmin AccountAdminAccountAdminRole = "admin"
	AccountAdminAccountAdminRoleStaff AccountAdminAccountAdminRole = "staff"
)

// Defines values for AccountAdminAccountAdminStatus.
const (
	AccountAdminAccountAdminStatusActive      AccountAdminAccountAdminStatus = "active"
	AccountAdminAccountAdminStatusDeactivated AccountAdminAccountAdminStatus = "deactivated"
	AccountAdminAccountAdminStatusLocked      AccountAdminAccountAdminStatus = "locked"
)

// Defines values for AccountAdminPostParamAccountAdminRole.
const (
	AccountAdminPostParamAccountAdminRoleAdmin AccountAdminPostParamAccountAdminRole = "admin"
	AccountAdminPostParamAccountAdminRoleStaff AccountAdminPostParamAccountAdminRole = "staff"
)

// Defines values for AccountAdminPostParamAccountAdminStatus.
const (
	AccountAdminPostParamAccountAdminStatusActive      AccountAdminPostParamAccountAdminStatus = "active"
	AccountAdminPostParamAccountAdminStatusDeactivated AccountAdminPostParamAccountAdminStatus = "deactivated"
	AccountAdminPostParamAccountAdminStatusLocked      AccountAdminPostParamAccountAdminStatus = "locked"
)

// Defines values for AccountAdminPutParamAccountAdminRole.
const (
	Admin AccountAdminPutParamAccountAdminRole = "admin"
	Staff AccountAdminPutParamAccountAdminRole = "staff"
)

// Defines values for AccountAdminPutParamAccountAdminStatus.
const (
	AccountAdminPutParamAccountAdminStatusActive      AccountAdminPutParamAccountAdminStatus = "active"
	AccountAdminPutParamAccountAdminStatusDeactivated AccountAdminPutParamAccountAdminStatus = "deactivated"
	AccountAdminPutParamAccountAdminStatusLocked      AccountAdminPutParamAccountAdminStatus = "locked"
)

// Account defines model for Account.
type Account struct {
	AccountId     openapi_types.UUID   `json:"account_id"`
	AccountStatus AccountAccountStatus `json:"account_status"`
	CreatedAt     time.Time            `json:"created_at"`
	LoginId       *string              `json:"login_id,omitempty"`
	LoginType     AccountLoginType     `json:"login_type"`
	Mail          *openapi_types.Email `json:"mail,omitempty"`
	Phone         *string              `json:"phone,omitempty"`
	UpdatedAt     time.Time            `json:"updated_at"`
}

// AccountAccountStatus defines model for Account.AccountStatus.
type AccountAccountStatus string

// AccountLoginType defines model for Account.LoginType.
type AccountLoginType string

// AccountAdmin defines model for AccountAdmin.
type AccountAdmin struct {
	AccountAdminId     openapi_types.UUID             `json:"account_admin_id"`
	AccountAdminRole   AccountAdminAccountAdminRole   `json:"account_admin_role"`
	AccountAdminStatus AccountAdminAccountAdminStatus `json:"account_admin_status"`
	CreatedAt          time.Time                      `json:"created_at"`
	Mail               openapi_types.Email            `json:"mail"`
	UpdatedAt          time.Time                      `json:"updated_at"`
}

// AccountAdminAccountAdminRole defines model for AccountAdmin.AccountAdminRole.
type AccountAdminAccountAdminRole string

// AccountAdminAccountAdminStatus defines model for AccountAdmin.AccountAdminStatus.
type AccountAdminAccountAdminStatus string

// AccountAdminID defines model for AccountAdminID.
type AccountAdminID = openapi_types.UUID

// AccountAdminLoginParam defines model for AccountAdminLoginParam.
type AccountAdminLoginParam struct {
	Mail     openapi_types.Email `json:"mail"`
	Password string              `json:"password"`
}

// AccountAdminPostParam defines model for AccountAdminPostParam.
type AccountAdminPostParam struct {
	AccountAdminRole   AccountAdminPostParamAccountAdminRole   `json:"account_admin_role"`
	AccountAdminStatus AccountAdminPostParamAccountAdminStatus `json:"account_admin_status"`
	Mail               openapi_types.Email                     `json:"mail"`
	Password           string                                  `json:"password"`
}

// AccountAdminPostParamAccountAdminRole defines model for AccountAdminPostParam.AccountAdminRole.
type AccountAdminPostParamAccountAdminRole string

// AccountAdminPostParamAccountAdminStatus defines model for AccountAdminPostParam.AccountAdminStatus.
type AccountAdminPostParamAccountAdminStatus string

// AccountAdminPutParam defines model for AccountAdminPutParam.
type AccountAdminPutParam struct {
	AccountAdminRole   *AccountAdminPutParamAccountAdminRole   `json:"account_admin_role,omitempty"`
	AccountAdminStatus *AccountAdminPutParamAccountAdminStatus `json:"account_admin_status,omitempty"`
	Password           *string                                 `json:"password,omitempty"`
}

// AccountAdminPutParamAccountAdminRole defines model for AccountAdminPutParam.AccountAdminRole.
type AccountAdminPutParamAccountAdminRole string

// AccountAdminPutParamAccountAdminStatus defines model for AccountAdminPutParam.AccountAdminStatus.
type AccountAdminPutParamAccountAdminStatus string

// AuthGoogleCallbackResp defines model for AuthGoogleCallbackResp.
type AuthGoogleCallbackResp struct {
	Account      *Account `json:"account,omitempty"`
	IsRegistered bool     `json:"is_registered"`
}

// ErrMsg Error Message
type ErrMsg struct {
	// Message Error Message
	Message *string `json:"message,omitempty"`

	// Version Application Version
	Version *string `json:"version,omitempty"`
}

// Jewellery defines model for Jewellery.
type Jewellery struct {
	Description *string             `json:"description,omitempty"`
	ImageUrl    *string             `json:"image_url,omitempty"`
	JewelleryId *openapi_types.UUID `json:"jewellery_id,omitempty"`
	Name        *string             `json:"name,omitempty"`
	Price       *int                `json:"price,omitempty"`
	Quantity    *int64              `json:"quantity,omitempty"`
}

// GetApiAuthGoogleCallbackParams defines parameters for GetApiAuthGoogleCallback.
type GetApiAuthGoogleCallbackParams struct {
	// State OAuth2 state parameter
	State string `form:"state" json:"state"`

	// Code OAuth2 code parameter
	Code string `form:"code" json:"code"`
}

// PostApiAdminAccountAdminJSONRequestBody defines body for PostApiAdminAccountAdmin for application/json ContentType.
type PostApiAdminAccountAdminJSONRequestBody = AccountAdminPostParam

// PutApiAdminAccountAdminAccountAdminsIdJSONRequestBody defines body for PutApiAdminAccountAdminAccountAdminsId for application/json ContentType.
type PutApiAdminAccountAdminAccountAdminsIdJSONRequestBody = AccountAdminPutParam

// PostApiAdminLoginJSONRequestBody defines body for PostApiAdminLogin for application/json ContentType.
type PostApiAdminLoginJSONRequestBody = AccountAdminLoginParam

// GetApiJewelleryJSONRequestBody defines body for GetApiJewellery for application/json ContentType.
type GetApiJewelleryJSONRequestBody = Jewellery
