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

// SetPolicyNameRequest set policy name request
//
// swagger:model setPolicyNameRequest
type SetPolicyNameRequest struct {

	// entity name
	// Required: true
	EntityName *string `json:"entityName"`

	// entity type
	// Required: true
	EntityType *PolicyEntity `json:"entityType"`

	// name
	// Required: true
	Name []string `json:"name"`
}

// Validate validates this set policy name request
func (m *SetPolicyNameRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetPolicyNameRequest) validateEntityName(formats strfmt.Registry) error {

	if err := validate.Required("entityName", "body", m.EntityName); err != nil {
		return err
	}

	return nil
}

func (m *SetPolicyNameRequest) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entityType", "body", m.EntityType); err != nil {
		return err
	}

	if err := validate.Required("entityType", "body", m.EntityType); err != nil {
		return err
	}

	if m.EntityType != nil {
		if err := m.EntityType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityType")
			}
			return err
		}
	}

	return nil
}

func (m *SetPolicyNameRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this set policy name request based on the context it is used
func (m *SetPolicyNameRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntityType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetPolicyNameRequest) contextValidateEntityType(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityType != nil {

		if err := m.EntityType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SetPolicyNameRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetPolicyNameRequest) UnmarshalBinary(b []byte) error {
	var res SetPolicyNameRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
