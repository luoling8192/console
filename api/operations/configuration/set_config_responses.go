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

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// SetConfigOKCode is the HTTP code returned for type SetConfigOK
const SetConfigOKCode int = 200

/*
SetConfigOK A successful response.

swagger:response setConfigOK
*/
type SetConfigOK struct {

	/*
	  In: Body
	*/
	Payload *models.SetConfigResponse `json:"body,omitempty"`
}

// NewSetConfigOK creates SetConfigOK with default headers values
func NewSetConfigOK() *SetConfigOK {

	return &SetConfigOK{}
}

// WithPayload adds the payload to the set config o k response
func (o *SetConfigOK) WithPayload(payload *models.SetConfigResponse) *SetConfigOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set config o k response
func (o *SetConfigOK) SetPayload(payload *models.SetConfigResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetConfigOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
SetConfigDefault Generic error response.

swagger:response setConfigDefault
*/
type SetConfigDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewSetConfigDefault creates SetConfigDefault with default headers values
func NewSetConfigDefault(code int) *SetConfigDefault {
	if code <= 0 {
		code = 500
	}

	return &SetConfigDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the set config default response
func (o *SetConfigDefault) WithStatusCode(code int) *SetConfigDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the set config default response
func (o *SetConfigDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the set config default response
func (o *SetConfigDefault) WithPayload(payload *models.APIError) *SetConfigDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set config default response
func (o *SetConfigDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetConfigDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
