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

// GetBucketReplicationRuleOKCode is the HTTP code returned for type GetBucketReplicationRuleOK
const GetBucketReplicationRuleOKCode int = 200

/*
GetBucketReplicationRuleOK A successful response.

swagger:response getBucketReplicationRuleOK
*/
type GetBucketReplicationRuleOK struct {

	/*
	  In: Body
	*/
	Payload *models.BucketReplicationRule `json:"body,omitempty"`
}

// NewGetBucketReplicationRuleOK creates GetBucketReplicationRuleOK with default headers values
func NewGetBucketReplicationRuleOK() *GetBucketReplicationRuleOK {

	return &GetBucketReplicationRuleOK{}
}

// WithPayload adds the payload to the get bucket replication rule o k response
func (o *GetBucketReplicationRuleOK) WithPayload(payload *models.BucketReplicationRule) *GetBucketReplicationRuleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bucket replication rule o k response
func (o *GetBucketReplicationRuleOK) SetPayload(payload *models.BucketReplicationRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBucketReplicationRuleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetBucketReplicationRuleDefault Generic error response.

swagger:response getBucketReplicationRuleDefault
*/
type GetBucketReplicationRuleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetBucketReplicationRuleDefault creates GetBucketReplicationRuleDefault with default headers values
func NewGetBucketReplicationRuleDefault(code int) *GetBucketReplicationRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &GetBucketReplicationRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get bucket replication rule default response
func (o *GetBucketReplicationRuleDefault) WithStatusCode(code int) *GetBucketReplicationRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get bucket replication rule default response
func (o *GetBucketReplicationRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get bucket replication rule default response
func (o *GetBucketReplicationRuleDefault) WithPayload(payload *models.APIError) *GetBucketReplicationRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bucket replication rule default response
func (o *GetBucketReplicationRuleDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBucketReplicationRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
