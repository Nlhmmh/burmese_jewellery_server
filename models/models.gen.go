// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.4 DO NOT EDIT.
package models

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// ErrorMessage Error Message
type ErrorMessage struct {
	// Message Error MEssage
	Message *string `json:"message,omitempty"`

	// Version Application Version
	Version *string `json:"version,omitempty"`
}

// Jewellery defines model for Jewellery.
type Jewellery struct {
	Description *string `db:"description" json:"description,omitempty"`
	ImageUrl    *string `db:"image_url" json:"image_url,omitempty"`

	// JewelleryId Jewellery ID
	JewelleryId *JewelleryID `json:"jewellery_id,omitempty"`
	Name        *string      `db:"name" json:"name,omitempty"`
	Price       *int         `db:"price" json:"price,omitempty"`
	Quantity    *int64       `db:"quantity" json:"quantity,omitempty"`
}

// JewelleryID Jewellery ID
type JewelleryID = openapi_types.UUID

// JewelleryIDParam Jewellery ID
type JewelleryIDParam = JewelleryID

// GetApiAuthGoogleCallbackParams defines parameters for GetApiAuthGoogleCallback.
type GetApiAuthGoogleCallbackParams struct {
	// State OAuth2 state parameter
	State string `form:"state" json:"state"`

	// Code OAuth2 code parameter
	Code string `form:"code" json:"code"`
}

// GetApiJewelleryJSONRequestBody defines body for GetApiJewellery for application/json ContentType.
type GetApiJewelleryJSONRequestBody = Jewellery