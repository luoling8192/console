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

package subnet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/luoling8192/console/models"
)

// SubnetRegTokenOKCode is the HTTP code returned for type SubnetRegTokenOK
const SubnetRegTokenOKCode int = 200

/*
SubnetRegTokenOK A successful response.

swagger:response subnetRegTokenOK
*/
type SubnetRegTokenOK struct {

	/*
	  In: Body
	*/
	Payload *models.SubnetRegTokenResponse `json:"body,omitempty"`
}

// NewSubnetRegTokenOK creates SubnetRegTokenOK with default headers values
func NewSubnetRegTokenOK() *SubnetRegTokenOK {

	return &SubnetRegTokenOK{}
}

// WithPayload adds the payload to the subnet reg token o k response
func (o *SubnetRegTokenOK) WithPayload(payload *models.SubnetRegTokenResponse) *SubnetRegTokenOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the subnet reg token o k response
func (o *SubnetRegTokenOK) SetPayload(payload *models.SubnetRegTokenResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SubnetRegTokenOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
SubnetRegTokenDefault Generic error response.

swagger:response subnetRegTokenDefault
*/
type SubnetRegTokenDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewSubnetRegTokenDefault creates SubnetRegTokenDefault with default headers values
func NewSubnetRegTokenDefault(code int) *SubnetRegTokenDefault {
	if code <= 0 {
		code = 500
	}

	return &SubnetRegTokenDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the subnet reg token default response
func (o *SubnetRegTokenDefault) WithStatusCode(code int) *SubnetRegTokenDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the subnet reg token default response
func (o *SubnetRegTokenDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the subnet reg token default response
func (o *SubnetRegTokenDefault) WithPayload(payload *models.APIError) *SubnetRegTokenDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the subnet reg token default response
func (o *SubnetRegTokenDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SubnetRegTokenDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
