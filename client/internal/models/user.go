// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// User User 实体
//
// swagger:model User
type User struct {

	// User ID
	ID int64 `json:"id,omitempty"`

	// 用户名
	Name string `json:"name,omitempty"`

	// 状态
	// active StatusActive
	// inactive StatusInactive
	// Enum: [active inactive]
	Status string `json:"status,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var userTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userTypeStatusPropEnum = append(userTypeStatusPropEnum, v)
	}
}

const (

	// UserStatusActive captures enum value "active"
	UserStatusActive string = "active"

	// UserStatusInactive captures enum value "inactive"
	UserStatusInactive string = "inactive"
)

// prop value enum
func (m *User) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *User) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user based on context it is used
func (m *User) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
