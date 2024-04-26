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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MultiBucketResponseState multi bucket response state
//
// swagger:model multiBucketResponseState
type MultiBucketResponseState struct {

	// replication state
	ReplicationState []*MultiBucketResponseItem `json:"replicationState"`
}

// Validate validates this multi bucket response state
func (m *MultiBucketResponseState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReplicationState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MultiBucketResponseState) validateReplicationState(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplicationState) { // not required
		return nil
	}

	for i := 0; i < len(m.ReplicationState); i++ {
		if swag.IsZero(m.ReplicationState[i]) { // not required
			continue
		}

		if m.ReplicationState[i] != nil {
			if err := m.ReplicationState[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("replicationState" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("replicationState" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this multi bucket response state based on the context it is used
func (m *MultiBucketResponseState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReplicationState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MultiBucketResponseState) contextValidateReplicationState(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReplicationState); i++ {

		if m.ReplicationState[i] != nil {

			if swag.IsZero(m.ReplicationState[i]) { // not required
				return nil
			}

			if err := m.ReplicationState[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("replicationState" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("replicationState" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MultiBucketResponseState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MultiBucketResponseState) UnmarshalBinary(b []byte) error {
	var res MultiBucketResponseState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
