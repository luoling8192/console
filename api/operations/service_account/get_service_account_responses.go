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

package service_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/luoling8192/console/models"
)

// GetServiceAccountOKCode is the HTTP code returned for type GetServiceAccountOK
const GetServiceAccountOKCode int = 200

/*
GetServiceAccountOK A successful response.

swagger:response getServiceAccountOK
*/
type GetServiceAccountOK struct {

	/*
	  In: Body
	*/
	Payload *models.ServiceAccount `json:"body,omitempty"`
}

// NewGetServiceAccountOK creates GetServiceAccountOK with default headers values
func NewGetServiceAccountOK() *GetServiceAccountOK {

	return &GetServiceAccountOK{}
}

// WithPayload adds the payload to the get service account o k response
func (o *GetServiceAccountOK) WithPayload(payload *models.ServiceAccount) *GetServiceAccountOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service account o k response
func (o *GetServiceAccountOK) SetPayload(payload *models.ServiceAccount) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceAccountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetServiceAccountDefault Generic error response.

swagger:response getServiceAccountDefault
*/
type GetServiceAccountDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetServiceAccountDefault creates GetServiceAccountDefault with default headers values
func NewGetServiceAccountDefault(code int) *GetServiceAccountDefault {
	if code <= 0 {
		code = 500
	}

	return &GetServiceAccountDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get service account default response
func (o *GetServiceAccountDefault) WithStatusCode(code int) *GetServiceAccountDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get service account default response
func (o *GetServiceAccountDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get service account default response
func (o *GetServiceAccountDefault) WithPayload(payload *models.APIError) *GetServiceAccountDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service account default response
func (o *GetServiceAccountDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceAccountDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
