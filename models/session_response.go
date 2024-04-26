// Code generated by go-swagger; DO NOT EDIT.

// This file is part of FST Console Server
// Copyright (c) 2023 FST, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SessionResponse session response
//
// swagger:model sessionResponse
type SessionResponse struct {

	// allow resources
	AllowResources []*PermissionResource `json:"allowResources"`

	// custom styles
	CustomStyles string `json:"customStyles,omitempty"`

	// distributed mode
	DistributedMode bool `json:"distributedMode,omitempty"`

	// env constants
	EnvConstants *EnvironmentConstants `json:"envConstants,omitempty"`

	// features
	Features []string `json:"features"`

	// operator
	Operator bool `json:"operator,omitempty"`

	// permissions
	Permissions map[string][]string `json:"permissions,omitempty"`

	// server end point
	ServerEndPoint string `json:"serverEndPoint,omitempty"`

	// status
	// Enum: [ok]
	Status string `json:"status,omitempty"`
}

// Validate validates this session response
func (m *SessionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvConstants(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SessionResponse) validateAllowResources(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowResources) { // not required
		return nil
	}

	for i := 0; i < len(m.AllowResources); i++ {
		if swag.IsZero(m.AllowResources[i]) { // not required
			continue
		}

		if m.AllowResources[i] != nil {
			if err := m.AllowResources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowResources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowResources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SessionResponse) validateEnvConstants(formats strfmt.Registry) error {
	if swag.IsZero(m.EnvConstants) { // not required
		return nil
	}

	if m.EnvConstants != nil {
		if err := m.EnvConstants.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("envConstants")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("envConstants")
			}
			return err
		}
	}

	return nil
}

var sessionResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sessionResponseTypeStatusPropEnum = append(sessionResponseTypeStatusPropEnum, v)
	}
}

const (

	// SessionResponseStatusOk captures enum value "ok"
	SessionResponseStatusOk string = "ok"
)

// prop value enum
func (m *SessionResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sessionResponseTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SessionResponse) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this session response based on the context it is used
func (m *SessionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllowResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnvConstants(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SessionResponse) contextValidateAllowResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AllowResources); i++ {

		if m.AllowResources[i] != nil {

			if swag.IsZero(m.AllowResources[i]) { // not required
				return nil
			}

			if err := m.AllowResources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowResources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowResources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SessionResponse) contextValidateEnvConstants(ctx context.Context, formats strfmt.Registry) error {

	if m.EnvConstants != nil {

		if swag.IsZero(m.EnvConstants) { // not required
			return nil
		}

		if err := m.EnvConstants.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("envConstants")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("envConstants")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SessionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SessionResponse) UnmarshalBinary(b []byte) error {
	var res SessionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
