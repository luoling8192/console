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

package object

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// PutObjectTagsOKCode is the HTTP code returned for type PutObjectTagsOK
const PutObjectTagsOKCode int = 200

/*
PutObjectTagsOK A successful response.

swagger:response putObjectTagsOK
*/
type PutObjectTagsOK struct {
}

// NewPutObjectTagsOK creates PutObjectTagsOK with default headers values
func NewPutObjectTagsOK() *PutObjectTagsOK {

	return &PutObjectTagsOK{}
}

// WriteResponse to the client
func (o *PutObjectTagsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*
PutObjectTagsDefault Generic error response.

swagger:response putObjectTagsDefault
*/
type PutObjectTagsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPutObjectTagsDefault creates PutObjectTagsDefault with default headers values
func NewPutObjectTagsDefault(code int) *PutObjectTagsDefault {
	if code <= 0 {
		code = 500
	}

	return &PutObjectTagsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put object tags default response
func (o *PutObjectTagsDefault) WithStatusCode(code int) *PutObjectTagsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put object tags default response
func (o *PutObjectTagsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put object tags default response
func (o *PutObjectTagsDefault) WithPayload(payload *models.APIError) *PutObjectTagsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put object tags default response
func (o *PutObjectTagsDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutObjectTagsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
