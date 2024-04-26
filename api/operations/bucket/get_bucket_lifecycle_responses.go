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

package bucket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// GetBucketLifecycleOKCode is the HTTP code returned for type GetBucketLifecycleOK
const GetBucketLifecycleOKCode int = 200

/*
GetBucketLifecycleOK A successful response.

swagger:response getBucketLifecycleOK
*/
type GetBucketLifecycleOK struct {

	/*
	  In: Body
	*/
	Payload *models.BucketLifecycleResponse `json:"body,omitempty"`
}

// NewGetBucketLifecycleOK creates GetBucketLifecycleOK with default headers values
func NewGetBucketLifecycleOK() *GetBucketLifecycleOK {

	return &GetBucketLifecycleOK{}
}

// WithPayload adds the payload to the get bucket lifecycle o k response
func (o *GetBucketLifecycleOK) WithPayload(payload *models.BucketLifecycleResponse) *GetBucketLifecycleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bucket lifecycle o k response
func (o *GetBucketLifecycleOK) SetPayload(payload *models.BucketLifecycleResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBucketLifecycleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetBucketLifecycleDefault Generic error response.

swagger:response getBucketLifecycleDefault
*/
type GetBucketLifecycleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetBucketLifecycleDefault creates GetBucketLifecycleDefault with default headers values
func NewGetBucketLifecycleDefault(code int) *GetBucketLifecycleDefault {
	if code <= 0 {
		code = 500
	}

	return &GetBucketLifecycleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get bucket lifecycle default response
func (o *GetBucketLifecycleDefault) WithStatusCode(code int) *GetBucketLifecycleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get bucket lifecycle default response
func (o *GetBucketLifecycleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get bucket lifecycle default response
func (o *GetBucketLifecycleDefault) WithPayload(payload *models.APIError) *GetBucketLifecycleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bucket lifecycle default response
func (o *GetBucketLifecycleDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBucketLifecycleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
