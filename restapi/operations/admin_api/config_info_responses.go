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
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/m3/mcs/models"
)

// ConfigInfoOKCode is the HTTP code returned for type ConfigInfoOK
const ConfigInfoOKCode int = 200

/*ConfigInfoOK A successful response.

swagger:response configInfoOK
*/
type ConfigInfoOK struct {

	/*
	  In: Body
	*/
	Payload *models.Configuration `json:"body,omitempty"`
}

// NewConfigInfoOK creates ConfigInfoOK with default headers values
func NewConfigInfoOK() *ConfigInfoOK {

	return &ConfigInfoOK{}
}

// WithPayload adds the payload to the config info o k response
func (o *ConfigInfoOK) WithPayload(payload *models.Configuration) *ConfigInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the config info o k response
func (o *ConfigInfoOK) SetPayload(payload *models.Configuration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConfigInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ConfigInfoDefault Generic error response.

swagger:response configInfoDefault
*/
type ConfigInfoDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewConfigInfoDefault creates ConfigInfoDefault with default headers values
func NewConfigInfoDefault(code int) *ConfigInfoDefault {
	if code <= 0 {
		code = 500
	}

	return &ConfigInfoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the config info default response
func (o *ConfigInfoDefault) WithStatusCode(code int) *ConfigInfoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the config info default response
func (o *ConfigInfoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the config info default response
func (o *ConfigInfoDefault) WithPayload(payload *models.Error) *ConfigInfoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the config info default response
func (o *ConfigInfoDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConfigInfoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
