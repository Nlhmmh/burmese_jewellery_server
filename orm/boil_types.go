// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package orm

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/null/v8/convert"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/strmangle"
)

// M type is for providing columns and column values to UpdateAll.
type M map[string]interface{}

// ErrSyncFail occurs during insert when the record could not be retrieved in
// order to populate default value information. This usually happens when LastInsertId
// fails or there was a primary key configuration that was not resolvable.
var ErrSyncFail = errors.New("orm: failed to synchronize data after insert")

type insertCache struct {
	query        string
	retQuery     string
	valueMapping []uint64
	retMapping   []uint64
}

type updateCache struct {
	query        string
	valueMapping []uint64
}

func makeCacheKey(cols boil.Columns, nzDefaults []string) string {
	buf := strmangle.GetBuffer()

	buf.WriteString(strconv.Itoa(cols.Kind))
	for _, w := range cols.Cols {
		buf.WriteString(w)
	}

	if len(nzDefaults) != 0 {
		buf.WriteByte('.')
	}
	for _, nz := range nzDefaults {
		buf.WriteString(nz)
	}

	str := buf.String()
	strmangle.PutBuffer(buf)
	return str
}

type AccountAdminRole string

// Enum values for AccountAdminRole
const (
	AccountAdminRoleAdmin AccountAdminRole = "admin"
	AccountAdminRoleStaff AccountAdminRole = "staff"
)

func AllAccountAdminRole() []AccountAdminRole {
	return []AccountAdminRole{
		AccountAdminRoleAdmin,
		AccountAdminRoleStaff,
	}
}

func (e AccountAdminRole) IsValid() error {
	switch e {
	case AccountAdminRoleAdmin, AccountAdminRoleStaff:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e AccountAdminRole) String() string {
	return string(e)
}

type AccountAdminStatus string

// Enum values for AccountAdminStatus
const (
	AccountAdminStatusActive      AccountAdminStatus = "active"
	AccountAdminStatusLocked      AccountAdminStatus = "locked"
	AccountAdminStatusDeactivated AccountAdminStatus = "deactivated"
)

func AllAccountAdminStatus() []AccountAdminStatus {
	return []AccountAdminStatus{
		AccountAdminStatusActive,
		AccountAdminStatusLocked,
		AccountAdminStatusDeactivated,
	}
}

func (e AccountAdminStatus) IsValid() error {
	switch e {
	case AccountAdminStatusActive, AccountAdminStatusLocked, AccountAdminStatusDeactivated:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e AccountAdminStatus) String() string {
	return string(e)
}

type OrderStatus string

// Enum values for OrderStatus
const (
	OrderStatusProceeding OrderStatus = "proceeding"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusReturned   OrderStatus = "returned"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

func AllOrderStatus() []OrderStatus {
	return []OrderStatus{
		OrderStatusProceeding,
		OrderStatusShipped,
		OrderStatusDelivered,
		OrderStatusReturned,
		OrderStatusCancelled,
	}
}

func (e OrderStatus) IsValid() error {
	switch e {
	case OrderStatusProceeding, OrderStatusShipped, OrderStatusDelivered, OrderStatusReturned, OrderStatusCancelled:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e OrderStatus) String() string {
	return string(e)
}

type Gender string

// Enum values for Gender
const (
	GenderMale        Gender = "male"
	GenderFemale      Gender = "female"
	GenderUnspecified Gender = "unspecified"
)

func AllGender() []Gender {
	return []Gender{
		GenderMale,
		GenderFemale,
		GenderUnspecified,
	}
}

func (e Gender) IsValid() error {
	switch e {
	case GenderMale, GenderFemale, GenderUnspecified:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e Gender) String() string {
	return string(e)
}

// NullGender is a nullable Gender enum type. It supports SQL and JSON serialization.
type NullGender struct {
	Val   Gender
	Valid bool
}

// NullGenderFrom creates a new Gender that will never be blank.
func NullGenderFrom(v Gender) NullGender {
	return NewNullGender(v, true)
}

// NullGenderFromPtr creates a new NullGender that be null if s is nil.
func NullGenderFromPtr(v *Gender) NullGender {
	if v == nil {
		return NewNullGender("", false)
	}
	return NewNullGender(*v, true)
}

// NewNullGender creates a new NullGender
func NewNullGender(v Gender, valid bool) NullGender {
	return NullGender{
		Val:   v,
		Valid: valid,
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NullGender) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, null.NullBytes) {
		e.Val = ""
		e.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &e.Val); err != nil {
		return err
	}

	e.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
func (e NullGender) MarshalJSON() ([]byte, error) {
	if !e.Valid {
		return null.NullBytes, nil
	}
	return json.Marshal(e.Val)
}

// MarshalText implements encoding.TextMarshaler.
func (e NullGender) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	return []byte(e.Val), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *NullGender) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		e.Valid = false
		return nil
	}

	e.Val = Gender(text)
	e.Valid = true
	return nil
}

// SetValid changes this NullGender value and also sets it to be non-null.
func (e *NullGender) SetValid(v Gender) {
	e.Val = v
	e.Valid = true
}

// Ptr returns a pointer to this NullGender value, or a nil pointer if this NullGender is null.
func (e NullGender) Ptr() *Gender {
	if !e.Valid {
		return nil
	}
	return &e.Val
}

// IsZero returns true for null types.
func (e NullGender) IsZero() bool {
	return !e.Valid
}

// Scan implements the Scanner interface.
func (e *NullGender) Scan(value interface{}) error {
	if value == nil {
		e.Val, e.Valid = "", false
		return nil
	}
	e.Valid = true
	return convert.ConvertAssign((*string)(&e.Val), value)
}

// Value implements the driver Valuer interface.
func (e NullGender) Value() (driver.Value, error) {
	if !e.Valid {
		return nil, nil
	}
	return string(e.Val), nil
}

type LoginType string

// Enum values for LoginType
const (
	LoginTypeEmail    LoginType = "email"
	LoginTypePhone    LoginType = "phone"
	LoginTypeGoogle   LoginType = "google"
	LoginTypeFacebook LoginType = "facebook"
)

func AllLoginType() []LoginType {
	return []LoginType{
		LoginTypeEmail,
		LoginTypePhone,
		LoginTypeGoogle,
		LoginTypeFacebook,
	}
}

func (e LoginType) IsValid() error {
	switch e {
	case LoginTypeEmail, LoginTypePhone, LoginTypeGoogle, LoginTypeFacebook:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e LoginType) String() string {
	return string(e)
}

type AccountStatus string

// Enum values for AccountStatus
const (
	AccountStatusPending     AccountStatus = "pending"
	AccountStatusActive      AccountStatus = "active"
	AccountStatusLocked      AccountStatus = "locked"
	AccountStatusDeactivated AccountStatus = "deactivated"
)

func AllAccountStatus() []AccountStatus {
	return []AccountStatus{
		AccountStatusPending,
		AccountStatusActive,
		AccountStatusLocked,
		AccountStatusDeactivated,
	}
}

func (e AccountStatus) IsValid() error {
	switch e {
	case AccountStatusPending, AccountStatusActive, AccountStatusLocked, AccountStatusDeactivated:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e AccountStatus) String() string {
	return string(e)
}
