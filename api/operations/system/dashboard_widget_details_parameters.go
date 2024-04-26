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

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDashboardWidgetDetailsParams creates a new DashboardWidgetDetailsParams object
//
// There are no default values defined in the spec.
func NewDashboardWidgetDetailsParams() DashboardWidgetDetailsParams {

	return DashboardWidgetDetailsParams{}
}

// DashboardWidgetDetailsParams contains all the bound params for the dashboard widget details operation
// typically these are obtained from a http.Request
//
// swagger:parameters DashboardWidgetDetails
type DashboardWidgetDetailsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	End *int64
	/*
	  In: query
	*/
	Start *int64
	/*
	  In: query
	*/
	Step *int32
	/*
	  Required: true
	  In: path
	*/
	WidgetID int32
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDashboardWidgetDetailsParams() beforehand.
func (o *DashboardWidgetDetailsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qEnd, qhkEnd, _ := qs.GetOK("end")
	if err := o.bindEnd(qEnd, qhkEnd, route.Formats); err != nil {
		res = append(res, err)
	}

	qStart, qhkStart, _ := qs.GetOK("start")
	if err := o.bindStart(qStart, qhkStart, route.Formats); err != nil {
		res = append(res, err)
	}

	qStep, qhkStep, _ := qs.GetOK("step")
	if err := o.bindStep(qStep, qhkStep, route.Formats); err != nil {
		res = append(res, err)
	}

	rWidgetID, rhkWidgetID, _ := route.Params.GetOK("widgetId")
	if err := o.bindWidgetID(rWidgetID, rhkWidgetID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindEnd binds and validates parameter End from query.
func (o *DashboardWidgetDetailsParams) bindEnd(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("end", "query", "int64", raw)
	}
	o.End = &value

	return nil
}

// bindStart binds and validates parameter Start from query.
func (o *DashboardWidgetDetailsParams) bindStart(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("start", "query", "int64", raw)
	}
	o.Start = &value

	return nil
}

// bindStep binds and validates parameter Step from query.
func (o *DashboardWidgetDetailsParams) bindStep(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("step", "query", "int32", raw)
	}
	o.Step = &value

	return nil
}

// bindWidgetID binds and validates parameter WidgetID from path.
func (o *DashboardWidgetDetailsParams) bindWidgetID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("widgetId", "path", "int32", raw)
	}
	o.WidgetID = value

	return nil
}
