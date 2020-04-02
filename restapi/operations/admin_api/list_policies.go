// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2020 MinIO, Inc.
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

package admin_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListPoliciesHandlerFunc turns a function with the right signature into a list policies handler
type ListPoliciesHandlerFunc func(ListPoliciesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListPoliciesHandlerFunc) Handle(params ListPoliciesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListPoliciesHandler interface for that can handle valid list policies params
type ListPoliciesHandler interface {
	Handle(ListPoliciesParams, interface{}) middleware.Responder
}

// NewListPolicies creates a new http.Handler for the list policies operation
func NewListPolicies(ctx *middleware.Context, handler ListPoliciesHandler) *ListPolicies {
	return &ListPolicies{Context: ctx, Handler: handler}
}

/*ListPolicies swagger:route GET /api/v1/policies AdminAPI listPolicies

List Policies

*/
type ListPolicies struct {
	Context *middleware.Context
	Handler ListPoliciesHandler
}

func (o *ListPolicies) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListPoliciesParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
