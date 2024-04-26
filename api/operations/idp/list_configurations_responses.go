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

package idp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// ListConfigurationsOKCode is the HTTP code returned for type ListConfigurationsOK
const ListConfigurationsOKCode int = 200

/*
ListConfigurationsOK A successful response.

swagger:response listConfigurationsOK
*/
type ListConfigurationsOK struct {

	/*
	  In: Body
	*/
	Payload *models.IdpListConfigurationsResponse `json:"body,omitempty"`
}

// NewListConfigurationsOK creates ListConfigurationsOK with default headers values
func NewListConfigurationsOK() *ListConfigurationsOK {

	return &ListConfigurationsOK{}
}

// WithPayload adds the payload to the list configurations o k response
func (o *ListConfigurationsOK) WithPayload(payload *models.IdpListConfigurationsResponse) *ListConfigurationsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list configurations o k response
func (o *ListConfigurationsOK) SetPayload(payload *models.IdpListConfigurationsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListConfigurationsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
ListConfigurationsDefault Generic error response.

swagger:response listConfigurationsDefault
*/
type ListConfigurationsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewListConfigurationsDefault creates ListConfigurationsDefault with default headers values
func NewListConfigurationsDefault(code int) *ListConfigurationsDefault {
	if code <= 0 {
		code = 500
	}

	return &ListConfigurationsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list configurations default response
func (o *ListConfigurationsDefault) WithStatusCode(code int) *ListConfigurationsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list configurations default response
func (o *ListConfigurationsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list configurations default response
func (o *ListConfigurationsDefault) WithPayload(payload *models.APIError) *ListConfigurationsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list configurations default response
func (o *ListConfigurationsDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListConfigurationsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
