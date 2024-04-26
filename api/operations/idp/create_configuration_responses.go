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

	"github.com/luoling8192/console/models"
)

// CreateConfigurationCreatedCode is the HTTP code returned for type CreateConfigurationCreated
const CreateConfigurationCreatedCode int = 201

/*
CreateConfigurationCreated A successful response.

swagger:response createConfigurationCreated
*/
type CreateConfigurationCreated struct {

	/*
	  In: Body
	*/
	Payload *models.SetIDPResponse `json:"body,omitempty"`
}

// NewCreateConfigurationCreated creates CreateConfigurationCreated with default headers values
func NewCreateConfigurationCreated() *CreateConfigurationCreated {

	return &CreateConfigurationCreated{}
}

// WithPayload adds the payload to the create configuration created response
func (o *CreateConfigurationCreated) WithPayload(payload *models.SetIDPResponse) *CreateConfigurationCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create configuration created response
func (o *CreateConfigurationCreated) SetPayload(payload *models.SetIDPResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateConfigurationCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
CreateConfigurationDefault Generic error response.

swagger:response createConfigurationDefault
*/
type CreateConfigurationDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewCreateConfigurationDefault creates CreateConfigurationDefault with default headers values
func NewCreateConfigurationDefault(code int) *CreateConfigurationDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateConfigurationDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create configuration default response
func (o *CreateConfigurationDefault) WithStatusCode(code int) *CreateConfigurationDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create configuration default response
func (o *CreateConfigurationDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create configuration default response
func (o *CreateConfigurationDefault) WithPayload(payload *models.APIError) *CreateConfigurationDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create configuration default response
func (o *CreateConfigurationDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateConfigurationDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
