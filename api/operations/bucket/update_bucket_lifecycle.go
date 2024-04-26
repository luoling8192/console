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
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/minio/console/models"
)

// UpdateBucketLifecycleHandlerFunc turns a function with the right signature into a update bucket lifecycle handler
type UpdateBucketLifecycleHandlerFunc func(UpdateBucketLifecycleParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateBucketLifecycleHandlerFunc) Handle(params UpdateBucketLifecycleParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// UpdateBucketLifecycleHandler interface for that can handle valid update bucket lifecycle params
type UpdateBucketLifecycleHandler interface {
	Handle(UpdateBucketLifecycleParams, *models.Principal) middleware.Responder
}

// NewUpdateBucketLifecycle creates a new http.Handler for the update bucket lifecycle operation
func NewUpdateBucketLifecycle(ctx *middleware.Context, handler UpdateBucketLifecycleHandler) *UpdateBucketLifecycle {
	return &UpdateBucketLifecycle{Context: ctx, Handler: handler}
}

/*
	UpdateBucketLifecycle swagger:route PUT /buckets/{bucket_name}/lifecycle/{lifecycle_id} Bucket updateBucketLifecycle

Update Lifecycle rule
*/
type UpdateBucketLifecycle struct {
	Context *middleware.Context
	Handler UpdateBucketLifecycleHandler
}

func (o *UpdateBucketLifecycle) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateBucketLifecycleParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
