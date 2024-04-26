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

package release

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// ListReleasesOKCode is the HTTP code returned for type ListReleasesOK
const ListReleasesOKCode int = 200

/*
ListReleasesOK A successful response.

swagger:response listReleasesOK
*/
type ListReleasesOK struct {

	/*
	  In: Body
	*/
	Payload *models.ReleaseListResponse `json:"body,omitempty"`
}

// NewListReleasesOK creates ListReleasesOK with default headers values
func NewListReleasesOK() *ListReleasesOK {

	return &ListReleasesOK{}
}

// WithPayload adds the payload to the list releases o k response
func (o *ListReleasesOK) WithPayload(payload *models.ReleaseListResponse) *ListReleasesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list releases o k response
func (o *ListReleasesOK) SetPayload(payload *models.ReleaseListResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListReleasesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
ListReleasesDefault Generic error response.

swagger:response listReleasesDefault
*/
type ListReleasesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewListReleasesDefault creates ListReleasesDefault with default headers values
func NewListReleasesDefault(code int) *ListReleasesDefault {
	if code <= 0 {
		code = 500
	}

	return &ListReleasesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list releases default response
func (o *ListReleasesDefault) WithStatusCode(code int) *ListReleasesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list releases default response
func (o *ListReleasesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list releases default response
func (o *ListReleasesDefault) WithPayload(payload *models.APIError) *ListReleasesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list releases default response
func (o *ListReleasesDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListReleasesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
