// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2022 MinIO, Inc.
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
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// ListPoliciesWithBucketOKCode is the HTTP code returned for type ListPoliciesWithBucketOK
const ListPoliciesWithBucketOKCode int = 200

/*ListPoliciesWithBucketOK A successful response.

swagger:response listPoliciesWithBucketOK
*/
type ListPoliciesWithBucketOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListPoliciesResponse `json:"body,omitempty"`
}

// NewListPoliciesWithBucketOK creates ListPoliciesWithBucketOK with default headers values
func NewListPoliciesWithBucketOK() *ListPoliciesWithBucketOK {

	return &ListPoliciesWithBucketOK{}
}

// WithPayload adds the payload to the list policies with bucket o k response
func (o *ListPoliciesWithBucketOK) WithPayload(payload *models.ListPoliciesResponse) *ListPoliciesWithBucketOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list policies with bucket o k response
func (o *ListPoliciesWithBucketOK) SetPayload(payload *models.ListPoliciesResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPoliciesWithBucketOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ListPoliciesWithBucketDefault Generic error response.

swagger:response listPoliciesWithBucketDefault
*/
type ListPoliciesWithBucketDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListPoliciesWithBucketDefault creates ListPoliciesWithBucketDefault with default headers values
func NewListPoliciesWithBucketDefault(code int) *ListPoliciesWithBucketDefault {
	if code <= 0 {
		code = 500
	}

	return &ListPoliciesWithBucketDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list policies with bucket default response
func (o *ListPoliciesWithBucketDefault) WithStatusCode(code int) *ListPoliciesWithBucketDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list policies with bucket default response
func (o *ListPoliciesWithBucketDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list policies with bucket default response
func (o *ListPoliciesWithBucketDefault) WithPayload(payload *models.Error) *ListPoliciesWithBucketDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list policies with bucket default response
func (o *ListPoliciesWithBucketDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPoliciesWithBucketDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
