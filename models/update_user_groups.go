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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateUserGroups update user groups
//
// swagger:model updateUserGroups
type UpdateUserGroups struct {

	// groups
	// Required: true
	Groups []string `json:"groups"`
}

// Validate validates this update user groups
func (m *UpdateUserGroups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateUserGroups) validateGroups(formats strfmt.Registry) error {

	if err := validate.Required("groups", "body", m.Groups); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update user groups based on context it is used
func (m *UpdateUserGroups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateUserGroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateUserGroups) UnmarshalBinary(b []byte) error {
	var res UpdateUserGroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
