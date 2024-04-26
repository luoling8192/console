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

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/luoling8192/console/models"
)

// CheckUserServiceAccountsOKCode is the HTTP code returned for type CheckUserServiceAccountsOK
const CheckUserServiceAccountsOKCode int = 200

/*
CheckUserServiceAccountsOK A successful response.

swagger:response checkUserServiceAccountsOK
*/
type CheckUserServiceAccountsOK struct {

	/*
	  In: Body
	*/
	Payload *models.UserServiceAccountSummary `json:"body,omitempty"`
}

// NewCheckUserServiceAccountsOK creates CheckUserServiceAccountsOK with default headers values
func NewCheckUserServiceAccountsOK() *CheckUserServiceAccountsOK {

	return &CheckUserServiceAccountsOK{}
}

// WithPayload adds the payload to the check user service accounts o k response
func (o *CheckUserServiceAccountsOK) WithPayload(payload *models.UserServiceAccountSummary) *CheckUserServiceAccountsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check user service accounts o k response
func (o *CheckUserServiceAccountsOK) SetPayload(payload *models.UserServiceAccountSummary) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckUserServiceAccountsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
CheckUserServiceAccountsDefault Generic error response.

swagger:response checkUserServiceAccountsDefault
*/
type CheckUserServiceAccountsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewCheckUserServiceAccountsDefault creates CheckUserServiceAccountsDefault with default headers values
func NewCheckUserServiceAccountsDefault(code int) *CheckUserServiceAccountsDefault {
	if code <= 0 {
		code = 500
	}

	return &CheckUserServiceAccountsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the check user service accounts default response
func (o *CheckUserServiceAccountsDefault) WithStatusCode(code int) *CheckUserServiceAccountsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the check user service accounts default response
func (o *CheckUserServiceAccountsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the check user service accounts default response
func (o *CheckUserServiceAccountsDefault) WithPayload(payload *models.APIError) *CheckUserServiceAccountsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check user service accounts default response
func (o *CheckUserServiceAccountsDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckUserServiceAccountsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
