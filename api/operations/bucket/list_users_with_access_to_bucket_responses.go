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

	"github.com/luoling8192/console/models"
)

// ListUsersWithAccessToBucketOKCode is the HTTP code returned for type ListUsersWithAccessToBucketOK
const ListUsersWithAccessToBucketOKCode int = 200

/*
ListUsersWithAccessToBucketOK A successful response.

swagger:response listUsersWithAccessToBucketOK
*/
type ListUsersWithAccessToBucketOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewListUsersWithAccessToBucketOK creates ListUsersWithAccessToBucketOK with default headers values
func NewListUsersWithAccessToBucketOK() *ListUsersWithAccessToBucketOK {

	return &ListUsersWithAccessToBucketOK{}
}

// WithPayload adds the payload to the list users with access to bucket o k response
func (o *ListUsersWithAccessToBucketOK) WithPayload(payload []string) *ListUsersWithAccessToBucketOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list users with access to bucket o k response
func (o *ListUsersWithAccessToBucketOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUsersWithAccessToBucketOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
ListUsersWithAccessToBucketDefault Generic error response.

swagger:response listUsersWithAccessToBucketDefault
*/
type ListUsersWithAccessToBucketDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewListUsersWithAccessToBucketDefault creates ListUsersWithAccessToBucketDefault with default headers values
func NewListUsersWithAccessToBucketDefault(code int) *ListUsersWithAccessToBucketDefault {
	if code <= 0 {
		code = 500
	}

	return &ListUsersWithAccessToBucketDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list users with access to bucket default response
func (o *ListUsersWithAccessToBucketDefault) WithStatusCode(code int) *ListUsersWithAccessToBucketDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list users with access to bucket default response
func (o *ListUsersWithAccessToBucketDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list users with access to bucket default response
func (o *ListUsersWithAccessToBucketDefault) WithPayload(payload *models.APIError) *ListUsersWithAccessToBucketDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list users with access to bucket default response
func (o *ListUsersWithAccessToBucketDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUsersWithAccessToBucketDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
