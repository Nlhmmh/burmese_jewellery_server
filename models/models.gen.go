// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package models

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for AccountLoginType.
const (
	Email    AccountLoginType = "email"
	Facebook AccountLoginType = "facebook"
	Google   AccountLoginType = "google"
	Phone    AccountLoginType = "phone"
)

// Defines values for AccountAdminRole.
const (
	Admin AccountAdminRole = "admin"
	Staff AccountAdminRole = "staff"
)

// Defines values for AccountAdminStatus.
const (
	AccountAdminStatusActive      AccountAdminStatus = "active"
	AccountAdminStatusDeactivated AccountAdminStatus = "deactivated"
	AccountAdminStatusLocked      AccountAdminStatus = "locked"
)

// Defines values for AccountStatus.
const (
	AccountStatusActive      AccountStatus = "active"
	AccountStatusDeactivated AccountStatus = "deactivated"
	AccountStatusLocked      AccountStatus = "locked"
	AccountStatusPending     AccountStatus = "pending"
)

// Defines values for Gender.
const (
	Female      Gender = "female"
	Male        Gender = "male"
	Unspecified Gender = "unspecified"
)

// Defines values for OrderStatus.
const (
	Cancelled  OrderStatus = "cancelled"
	Delivered  OrderStatus = "delivered"
	Proceeding OrderStatus = "proceeding"
	Returned   OrderStatus = "returned"
	Shipped    OrderStatus = "shipped"
)

// Account defines model for Account.
type Account struct {
	AccountId     openapi_types.UUID   `json:"account_id"`
	AccountStatus AccountStatus        `json:"account_status"`
	CreatedAt     time.Time            `json:"created_at"`
	LoginId       *string              `json:"login_id,omitempty"`
	LoginType     AccountLoginType     `json:"login_type"`
	Mail          *openapi_types.Email `json:"mail,omitempty"`
	Phone         *string              `json:"phone,omitempty"`
	UpdatedAt     time.Time            `json:"updated_at"`
}

// AccountLoginType defines model for Account.LoginType.
type AccountLoginType string

// AccountAdmin defines model for AccountAdmin.
type AccountAdmin struct {
	AccountAdminId     openapi_types.UUID  `json:"account_admin_id"`
	AccountAdminRole   AccountAdminRole    `json:"account_admin_role"`
	AccountAdminStatus AccountAdminStatus  `json:"account_admin_status"`
	CreatedAt          time.Time           `json:"created_at"`
	Mail               openapi_types.Email `json:"mail"`
	UpdatedAt          time.Time           `json:"updated_at"`
}

// AccountAdminLoginParam defines model for AccountAdminLoginParam.
type AccountAdminLoginParam struct {
	Mail     openapi_types.Email `json:"mail"`
	Password string              `json:"password"`
}

// AccountAdminPostParam defines model for AccountAdminPostParam.
type AccountAdminPostParam struct {
	AccountAdminRole   AccountAdminRole    `json:"account_admin_role"`
	AccountAdminStatus AccountAdminStatus  `json:"account_admin_status"`
	Mail               openapi_types.Email `json:"mail"`
	Password           string              `json:"password"`
}

// AccountAdminPutParam defines model for AccountAdminPutParam.
type AccountAdminPutParam struct {
	AccountAdminRole   *AccountAdminRole   `json:"account_admin_role,omitempty"`
	AccountAdminStatus *AccountAdminStatus `json:"account_admin_status,omitempty"`
	Password           *string             `json:"password,omitempty"`
}

// AccountAdminRole defines model for AccountAdminRole.
type AccountAdminRole string

// AccountAdminStatus defines model for AccountAdminStatus.
type AccountAdminStatus string

// AccountCartJewellery defines model for AccountCartJewellery.
type AccountCartJewellery struct {
	AccountId   openapi_types.UUID `json:"account_id"`
	CreatedAt   time.Time          `json:"created_at"`
	JewelleryId openapi_types.UUID `json:"jewellery_id"`
	Quantity    int                `json:"quantity"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

// AccountFavourite defines model for AccountFavourite.
type AccountFavourite struct {
	AccountId   openapi_types.UUID `json:"account_id"`
	CreatedAt   time.Time          `json:"created_at"`
	JewelleryId openapi_types.UUID `json:"jewellery_id"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

// AccountOrder defines model for AccountOrder.
type AccountOrder struct {
	AccountId               openapi_types.UUID      `json:"account_id"`
	AccountOrderAddress     AccountOrderAddress     `json:"account_order_address"`
	AccountOrderJewelleries []AccountOrderJewellery `json:"account_order_jewelleries"`
	CreatedAt               time.Time               `json:"created_at"`
	OrderId                 openapi_types.UUID      `json:"order_id"`
	OrderStatus             OrderStatus             `json:"order_status"`
	UpdatedAt               time.Time               `json:"updated_at"`
}

// AccountOrderAddress defines model for AccountOrderAddress.
type AccountOrderAddress struct {
	Address     string  `json:"address"`
	City        string  `json:"city"`
	CountryCode string  `json:"country_code"`
	PostCode    *string `json:"post_code,omitempty"`
	Remarks     *string `json:"remarks,omitempty"`
	State       *string `json:"state,omitempty"`
}

// AccountOrderJewellery defines model for AccountOrderJewellery.
type AccountOrderJewellery struct {
	CreatedAt   time.Time          `json:"created_at"`
	JewelleryId openapi_types.UUID `json:"jewellery_id"`
	OrderId     openapi_types.UUID `json:"order_id"`
	Price       int                `json:"price"`
	Quantity    int                `json:"quantity"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

// AccountProfile defines model for AccountProfile.
type AccountProfile struct {
	AccountId openapi_types.UUID `json:"account_id"`
	Birthday  openapi_types.Date `json:"birthday"`
	CreatedAt time.Time          `json:"created_at"`
	FirstName string             `json:"first_name"`
	Gender    Gender             `json:"gender"`
	LastName  string             `json:"last_name"`
	UpdatedAt time.Time          `json:"updated_at"`
}

// AccountProfilePostParam defines model for AccountProfilePostParam.
type AccountProfilePostParam struct {
	Birthday  *openapi_types.Date `json:"birthday,omitempty"`
	FirstName string              `json:"first_name"`
	Gender    *Gender             `json:"gender,omitempty"`
	LastName  string              `json:"last_name"`
}

// AccountStatus defines model for AccountStatus.
type AccountStatus string

// AccountStatusPutParam defines model for AccountStatusPutParam.
type AccountStatusPutParam struct {
	AccountStatus AccountStatus `json:"account_status"`
}

// AccountWithProfile defines model for AccountWithProfile.
type AccountWithProfile struct {
	Account        Account        `json:"account"`
	AccountProfile AccountProfile `json:"account_profile"`
}

// AuthEmailParam defines model for AuthEmailParam.
type AuthEmailParam struct {
	Mail     openapi_types.Email `json:"mail"`
	Password string              `json:"password"`
}

// AuthResp defines model for AuthResp.
type AuthResp struct {
	IsRegistered bool `json:"is_registered"`
}

// CartPostParam defines model for CartPostParam.
type CartPostParam struct {
	JewelleryId openapi_types.UUID `json:"jewellery_id"`
	Quantity    int                `json:"quantity"`
}

// Category defines model for Category.
type Category struct {
	CategoryId  openapi_types.UUID `json:"category_id"`
	CreatedAt   time.Time          `json:"created_at"`
	Description string             `json:"description"`
	ImageUrl    string             `json:"image_url"`
	Name        string             `json:"name"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

// CategoryPostParam defines model for CategoryPostParam.
type CategoryPostParam struct {
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	Name        string `json:"name"`
}

// CategoryPutParam defines model for CategoryPutParam.
type CategoryPutParam struct {
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	Name        string `json:"name"`
}

// ErrMsg Error Message
type ErrMsg struct {
	// Message Error message
	Message *string `json:"message,omitempty"`

	// Version Application Version
	Version *string `json:"version,omitempty"`
}

// FAQ defines model for FAQ.
type FAQ struct {
	Answer    string             `json:"answer"`
	CreatedAt time.Time          `json:"created_at"`
	FaqId     openapi_types.UUID `json:"faq_id"`
	IsActive  bool               `json:"is_active"`
	Question  string             `json:"question"`
	UpdatedAt time.Time          `json:"updated_at"`
}

// FAQPostParam defines model for FAQPostParam.
type FAQPostParam struct {
	Answer   string `json:"answer"`
	IsActive bool   `json:"is_active"`
	Question string `json:"question"`
}

// FAQPutParam defines model for FAQPutParam.
type FAQPutParam struct {
	Answer   string `json:"answer"`
	IsActive bool   `json:"is_active"`
	Question string `json:"question"`
}

// FavouritePostParam defines model for FavouritePostParam.
type FavouritePostParam struct {
	IsFavourite bool               `json:"is_favourite"`
	JewelleryId openapi_types.UUID `json:"jewellery_id"`
}

// Gem defines model for Gem.
type Gem struct {
	CreatedAt time.Time          `json:"created_at"`
	GemId     openapi_types.UUID `json:"gem_id"`
	Name      string             `json:"name"`
	UpdatedAt time.Time          `json:"updated_at"`
}

// GemPostParam defines model for GemPostParam.
type GemPostParam struct {
	Name string `json:"name"`
}

// GemPutParam defines model for GemPutParam.
type GemPutParam struct {
	Name string `json:"name"`
}

// Gender defines model for Gender.
type Gender string

// ID defines model for ID.
type ID = openapi_types.UUID

// Jewellery defines model for Jewellery.
type Jewellery struct {
	CategoryId  openapi_types.UUID `json:"category_id"`
	CreatedAt   time.Time          `json:"created_at"`
	Description string             `json:"description"`
	GemId       openapi_types.UUID `json:"gem_id"`
	ImageUrl    string             `json:"image_url"`
	IsPublished bool               `json:"is_published"`
	JewelleryId openapi_types.UUID `json:"jewellery_id"`
	MaterialId  openapi_types.UUID `json:"material_id"`
	Name        string             `json:"name"`
	Price       int                `json:"price"`
	Quantity    int                `json:"quantity"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

// JewelleryPostParam defines model for JewelleryPostParam.
type JewelleryPostParam struct {
	CategoryId  openapi_types.UUID `json:"category_id"`
	Description string             `json:"description"`
	GemId       openapi_types.UUID `json:"gem_id"`
	ImageUrl    string             `json:"image_url"`
	IsPublished bool               `json:"is_published"`
	MaterialId  openapi_types.UUID `json:"material_id"`
	Name        string             `json:"name"`
	Price       int                `json:"price"`
	Quantity    int                `json:"quantity"`
}

// JewelleryPutParam defines model for JewelleryPutParam.
type JewelleryPutParam struct {
	CategoryId  openapi_types.UUID `json:"category_id"`
	Description string             `json:"description"`
	GemId       openapi_types.UUID `json:"gem_id"`
	ImageUrl    string             `json:"image_url"`
	IsPublished bool               `json:"is_published"`
	MaterialId  openapi_types.UUID `json:"material_id"`
	Name        string             `json:"name"`
	Price       int                `json:"price"`
	Quantity    int                `json:"quantity"`
}

// Material defines model for Material.
type Material struct {
	CreatedAt  time.Time          `json:"created_at"`
	MaterialId openapi_types.UUID `json:"material_id"`
	Name       string             `json:"name"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

// MaterialPostParam defines model for MaterialPostParam.
type MaterialPostParam struct {
	Name string `json:"name"`
}

// MaterialPutParam defines model for MaterialPutParam.
type MaterialPutParam struct {
	Name string `json:"name"`
}

// OrderPostParam defines model for OrderPostParam.
type OrderPostParam struct {
	Address AccountOrderAddress `json:"address"`
}

// OrderStatus defines model for OrderStatus.
type OrderStatus string

// GetApiAdminAccountParams defines parameters for GetApiAdminAccount.
type GetApiAdminAccountParams struct {
	Offset        int            `form:"offset" json:"offset"`
	Limit         int            `form:"limit" json:"limit"`
	Id            *ID            `form:"id,omitempty" json:"id,omitempty"`
	FirstName     *string        `form:"first_name,omitempty" json:"first_name,omitempty"`
	LastName      *string        `form:"last_name,omitempty" json:"last_name,omitempty"`
	AccountStatus *AccountStatus `form:"account_status,omitempty" json:"account_status,omitempty"`
}

// GetApiAdminAccountAdminParams defines parameters for GetApiAdminAccountAdmin.
type GetApiAdminAccountAdminParams struct {
	Offset             int                 `form:"offset" json:"offset"`
	Limit              int                 `form:"limit" json:"limit"`
	Id                 *ID                 `form:"id,omitempty" json:"id,omitempty"`
	Mail               *string             `form:"mail,omitempty" json:"mail,omitempty"`
	AccountAdminRole   *AccountAdminRole   `form:"account_admin_role,omitempty" json:"account_admin_role,omitempty"`
	AccountAdminStatus *AccountAdminStatus `form:"account_admin_status,omitempty" json:"account_admin_status,omitempty"`
}

// GetApiAuthGoogleCallbackParams defines parameters for GetApiAuthGoogleCallback.
type GetApiAuthGoogleCallbackParams struct {
	// State OAuth2 state parameter
	State string `form:"state" json:"state"`

	// Code OAuth2 code parameter
	Code string `form:"code" json:"code"`
}

// GetApiFaqParams defines parameters for GetApiFaq.
type GetApiFaqParams struct {
	Offset   int     `form:"offset" json:"offset"`
	Limit    int     `form:"limit" json:"limit"`
	Id       *ID     `form:"id,omitempty" json:"id,omitempty"`
	Question *string `form:"question,omitempty" json:"question,omitempty"`
	Answer   *string `form:"answer,omitempty" json:"answer,omitempty"`
	IsActive *bool   `form:"is_active,omitempty" json:"is_active,omitempty"`
}

// GetApiJewelleryParams defines parameters for GetApiJewellery.
type GetApiJewelleryParams struct {
	Offset      int     `form:"offset" json:"offset"`
	Limit       int     `form:"limit" json:"limit"`
	Id          *ID     `form:"id,omitempty" json:"id,omitempty"`
	CategoryId  *ID     `form:"category_id,omitempty" json:"category_id,omitempty"`
	GemId       *ID     `form:"gem_id,omitempty" json:"gem_id,omitempty"`
	MaterialId  *ID     `form:"material_id,omitempty" json:"material_id,omitempty"`
	Name        *string `form:"name,omitempty" json:"name,omitempty"`
	IsPublished *bool   `form:"is_published,omitempty" json:"is_published,omitempty"`
}

// PutApiAdminAccountAccountIdJSONRequestBody defines body for PutApiAdminAccountAccountId for application/json ContentType.
type PutApiAdminAccountAccountIdJSONRequestBody = AccountStatusPutParam

// PostApiAdminAccountAdminJSONRequestBody defines body for PostApiAdminAccountAdmin for application/json ContentType.
type PostApiAdminAccountAdminJSONRequestBody = AccountAdminPostParam

// PutApiAdminAccountAdminAccountAdminsIdJSONRequestBody defines body for PutApiAdminAccountAdminAccountAdminsId for application/json ContentType.
type PutApiAdminAccountAdminAccountAdminsIdJSONRequestBody = AccountAdminPutParam

// PostApiAdminCategoryJSONRequestBody defines body for PostApiAdminCategory for application/json ContentType.
type PostApiAdminCategoryJSONRequestBody = CategoryPostParam

// PutApiAdminCategoryCategoryIdJSONRequestBody defines body for PutApiAdminCategoryCategoryId for application/json ContentType.
type PutApiAdminCategoryCategoryIdJSONRequestBody = CategoryPutParam

// PostApiAdminFaqJSONRequestBody defines body for PostApiAdminFaq for application/json ContentType.
type PostApiAdminFaqJSONRequestBody = FAQPostParam

// PutApiAdminFaqFaqIdJSONRequestBody defines body for PutApiAdminFaqFaqId for application/json ContentType.
type PutApiAdminFaqFaqIdJSONRequestBody = FAQPutParam

// PostApiAdminGemJSONRequestBody defines body for PostApiAdminGem for application/json ContentType.
type PostApiAdminGemJSONRequestBody = GemPostParam

// PutApiAdminGemGemIdJSONRequestBody defines body for PutApiAdminGemGemId for application/json ContentType.
type PutApiAdminGemGemIdJSONRequestBody = GemPutParam

// PostApiAdminJewelleryJSONRequestBody defines body for PostApiAdminJewellery for application/json ContentType.
type PostApiAdminJewelleryJSONRequestBody = JewelleryPostParam

// PutApiAdminJewelleryJewelleryIdJSONRequestBody defines body for PutApiAdminJewelleryJewelleryId for application/json ContentType.
type PutApiAdminJewelleryJewelleryIdJSONRequestBody = JewelleryPutParam

// PostApiAdminLoginJSONRequestBody defines body for PostApiAdminLogin for application/json ContentType.
type PostApiAdminLoginJSONRequestBody = AccountAdminLoginParam

// PostApiAdminMaterialJSONRequestBody defines body for PostApiAdminMaterial for application/json ContentType.
type PostApiAdminMaterialJSONRequestBody = MaterialPostParam

// PutApiAdminMaterialMaterialIdJSONRequestBody defines body for PutApiAdminMaterialMaterialId for application/json ContentType.
type PutApiAdminMaterialMaterialIdJSONRequestBody = MaterialPutParam

// PostApiAuthEmailLoginJSONRequestBody defines body for PostApiAuthEmailLogin for application/json ContentType.
type PostApiAuthEmailLoginJSONRequestBody = AuthEmailParam

// PostApiAuthEmailRegisterJSONRequestBody defines body for PostApiAuthEmailRegister for application/json ContentType.
type PostApiAuthEmailRegisterJSONRequestBody = AuthEmailParam

// PostApiCartJSONRequestBody defines body for PostApiCart for application/json ContentType.
type PostApiCartJSONRequestBody = CartPostParam

// PostApiFavouriteJSONRequestBody defines body for PostApiFavourite for application/json ContentType.
type PostApiFavouriteJSONRequestBody = FavouritePostParam

// PostApiOrderJSONRequestBody defines body for PostApiOrder for application/json ContentType.
type PostApiOrderJSONRequestBody = OrderPostParam

// PostApiProfileJSONRequestBody defines body for PostApiProfile for application/json ContentType.
type PostApiProfileJSONRequestBody = AccountProfilePostParam
