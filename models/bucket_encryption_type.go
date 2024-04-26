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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// BucketEncryptionType bucket encryption type
//
// swagger:model bucketEncryptionType
type BucketEncryptionType string

func NewBucketEncryptionType(value BucketEncryptionType) *BucketEncryptionType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated BucketEncryptionType.
func (m BucketEncryptionType) Pointer() *BucketEncryptionType {
	return &m
}

const (

	// BucketEncryptionTypeSseDashS3 captures enum value "sse-s3"
	BucketEncryptionTypeSseDashS3 BucketEncryptionType = "sse-s3"

	// BucketEncryptionTypeSseDashKms captures enum value "sse-kms"
	BucketEncryptionTypeSseDashKms BucketEncryptionType = "sse-kms"
)

// for schema
var bucketEncryptionTypeEnum []interface{}

func init() {
	var res []BucketEncryptionType
	if err := json.Unmarshal([]byte(`["sse-s3","sse-kms"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bucketEncryptionTypeEnum = append(bucketEncryptionTypeEnum, v)
	}
}

func (m BucketEncryptionType) validateBucketEncryptionTypeEnum(path, location string, value BucketEncryptionType) error {
	if err := validate.EnumCase(path, location, value, bucketEncryptionTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this bucket encryption type
func (m BucketEncryptionType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBucketEncryptionTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this bucket encryption type based on context it is used
func (m BucketEncryptionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
