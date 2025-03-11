// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2023 MinIO, Inc.
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

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/minio/console/api/operations/auth"
	"github.com/minio/console/api/operations/bucket"
	"github.com/minio/console/api/operations/license"
	"github.com/minio/console/api/operations/object"
	"github.com/minio/console/api/operations/public"
	"github.com/minio/console/api/operations/system"
	"github.com/minio/console/api/operations/user"
	"github.com/minio/console/models"
)

// NewConsoleAPI creates a new Console instance
func NewConsoleAPI(spec *loads.Document) *ConsoleAPI {
	return &ConsoleAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,

		BinProducer:  runtime.ByteStreamProducer(),
		JSONProducer: runtime.JSONProducer(),

		SystemAdminInfoHandler: system.AdminInfoHandlerFunc(func(params system.AdminInfoParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation system.AdminInfo has not yet been implemented")
		}),
		BucketBucketInfoHandler: bucket.BucketInfoHandlerFunc(func(params bucket.BucketInfoParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation bucket.BucketInfo has not yet been implemented")
		}),
		ObjectDeleteMultipleObjectsHandler: object.DeleteMultipleObjectsHandlerFunc(func(params object.DeleteMultipleObjectsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation object.DeleteMultipleObjects has not yet been implemented")
		}),
		ObjectDeleteObjectHandler: object.DeleteObjectHandlerFunc(func(params object.DeleteObjectParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation object.DeleteObject has not yet been implemented")
		}),
		ObjectDownloadObjectHandler: object.DownloadObjectHandlerFunc(func(params object.DownloadObjectParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation object.DownloadObject has not yet been implemented")
		}),
		ObjectDownloadMultipleObjectsHandler: object.DownloadMultipleObjectsHandlerFunc(func(params object.DownloadMultipleObjectsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation object.DownloadMultipleObjects has not yet been implemented")
		}),
		PublicDownloadSharedObjectHandler: public.DownloadSharedObjectHandlerFunc(func(params public.DownloadSharedObjectParams) middleware.Responder {
			return middleware.NotImplemented("operation public.DownloadSharedObject has not yet been implemented")
		}),
		BucketGetBucketQuotaHandler: bucket.GetBucketQuotaHandlerFunc(func(params bucket.GetBucketQuotaParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation bucket.GetBucketQuota has not yet been implemented")
		}),
		BucketGetBucketRewindHandler: bucket.GetBucketRewindHandlerFunc(func(params bucket.GetBucketRewindParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation bucket.GetBucketRewind has not yet been implemented")
		}),
		BucketGetBucketVersioningHandler: bucket.GetBucketVersioningHandlerFunc(func(params bucket.GetBucketVersioningParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation bucket.GetBucketVersioning has not yet been implemented")
		}),
		BucketGetMaxShareLinkExpHandler: bucket.GetMaxShareLinkExpHandlerFunc(func(params bucket.GetMaxShareLinkExpParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation bucket.GetMaxShareLinkExp has not yet been implemented")
		}),
		ObjectGetObjectMetadataHandler: object.GetObjectMetadataHandlerFunc(func(params object.GetObjectMetadataParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation object.GetObjectMetadata has not yet been implemented")
		}),
		LicenseLicenseAcknowledgeHandler: license.LicenseAcknowledgeHandlerFunc(func(params license.LicenseAcknowledgeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation license.LicenseAcknowledge has not yet been implemented")
		}),
		BucketListBucketsHandler: bucket.ListBucketsHandlerFunc(func(params bucket.ListBucketsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation bucket.ListBuckets has not yet been implemented")
		}),
		ObjectListObjectsHandler: object.ListObjectsHandlerFunc(func(params object.ListObjectsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation object.ListObjects has not yet been implemented")
		}),
		UserListUsersHandler: user.ListUsersHandlerFunc(func(params user.ListUsersParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.ListUsers has not yet been implemented")
		}),
		AuthLoginHandler: auth.LoginHandlerFunc(func(params auth.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.Login has not yet been implemented")
		}),
		AuthLoginDetailHandler: auth.LoginDetailHandlerFunc(func(params auth.LoginDetailParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.LoginDetail has not yet been implemented")
		}),
		AuthLoginOauth2AuthHandler: auth.LoginOauth2AuthHandlerFunc(func(params auth.LoginOauth2AuthParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.LoginOauth2Auth has not yet been implemented")
		}),
		AuthLogoutHandler: auth.LogoutHandlerFunc(func(params auth.LogoutParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation auth.Logout has not yet been implemented")
		}),
		BucketMakeBucketHandler: bucket.MakeBucketHandlerFunc(func(params bucket.MakeBucketParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation bucket.MakeBucket has not yet been implemented")
		}),
		ObjectPostBucketsBucketNameObjectsUploadHandler: object.PostBucketsBucketNameObjectsUploadHandlerFunc(func(params object.PostBucketsBucketNameObjectsUploadParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation object.PostBucketsBucketNameObjectsUpload has not yet been implemented")
		}),
		BucketPutBucketTagsHandler: bucket.PutBucketTagsHandlerFunc(func(params bucket.PutBucketTagsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation bucket.PutBucketTags has not yet been implemented")
		}),
		ObjectPutObjectRestoreHandler: object.PutObjectRestoreHandlerFunc(func(params object.PutObjectRestoreParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation object.PutObjectRestore has not yet been implemented")
		}),
		ObjectPutObjectTagsHandler: object.PutObjectTagsHandlerFunc(func(params object.PutObjectTagsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation object.PutObjectTags has not yet been implemented")
		}),
		AuthSessionCheckHandler: auth.SessionCheckHandlerFunc(func(params auth.SessionCheckParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation auth.SessionCheck has not yet been implemented")
		}),
		BucketSetBucketVersioningHandler: bucket.SetBucketVersioningHandlerFunc(func(params bucket.SetBucketVersioningParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation bucket.SetBucketVersioning has not yet been implemented")
		}),
		ObjectShareObjectHandler: object.ShareObjectHandlerFunc(func(params object.ShareObjectParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation object.ShareObject has not yet been implemented")
		}),

		// Applies when the "X-Anonymous" header is set
		AnonymousAuth: func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (anonymous) X-Anonymous from header param [X-Anonymous] has not yet been implemented")
		},
		KeyAuth: func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (key) has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*ConsoleAPI the console API */
type ConsoleAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for the following mime types:
	//   - multipart/form-data
	MultipartformConsumer runtime.Consumer

	// BinProducer registers a producer for the following mime types:
	//   - application/octet-stream
	BinProducer runtime.Producer
	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// AnonymousAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key X-Anonymous provided in the header
	AnonymousAuth func(string) (*models.Principal, error)

	// KeyAuth registers a function that takes an access token and a collection of required scopes and returns a principal
	// it performs authentication based on an oauth2 bearer token provided in the request
	KeyAuth func(string, []string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// SystemAdminInfoHandler sets the operation handler for the admin info operation
	SystemAdminInfoHandler system.AdminInfoHandler
	// BucketBucketInfoHandler sets the operation handler for the bucket info operation
	BucketBucketInfoHandler bucket.BucketInfoHandler
	// ObjectDeleteMultipleObjectsHandler sets the operation handler for the delete multiple objects operation
	ObjectDeleteMultipleObjectsHandler object.DeleteMultipleObjectsHandler
	// ObjectDeleteObjectHandler sets the operation handler for the delete object operation
	ObjectDeleteObjectHandler object.DeleteObjectHandler
	// ObjectDownloadObjectHandler sets the operation handler for the download object operation
	ObjectDownloadObjectHandler object.DownloadObjectHandler
	// ObjectDownloadMultipleObjectsHandler sets the operation handler for the download multiple objects operation
	ObjectDownloadMultipleObjectsHandler object.DownloadMultipleObjectsHandler
	// PublicDownloadSharedObjectHandler sets the operation handler for the download shared object operation
	PublicDownloadSharedObjectHandler public.DownloadSharedObjectHandler
	// BucketGetBucketQuotaHandler sets the operation handler for the get bucket quota operation
	BucketGetBucketQuotaHandler bucket.GetBucketQuotaHandler
	// BucketGetBucketRewindHandler sets the operation handler for the get bucket rewind operation
	BucketGetBucketRewindHandler bucket.GetBucketRewindHandler
	// BucketGetBucketVersioningHandler sets the operation handler for the get bucket versioning operation
	BucketGetBucketVersioningHandler bucket.GetBucketVersioningHandler
	// BucketGetMaxShareLinkExpHandler sets the operation handler for the get max share link exp operation
	BucketGetMaxShareLinkExpHandler bucket.GetMaxShareLinkExpHandler
	// ObjectGetObjectMetadataHandler sets the operation handler for the get object metadata operation
	ObjectGetObjectMetadataHandler object.GetObjectMetadataHandler
	// LicenseLicenseAcknowledgeHandler sets the operation handler for the license acknowledge operation
	LicenseLicenseAcknowledgeHandler license.LicenseAcknowledgeHandler
	// BucketListBucketsHandler sets the operation handler for the list buckets operation
	BucketListBucketsHandler bucket.ListBucketsHandler
	// ObjectListObjectsHandler sets the operation handler for the list objects operation
	ObjectListObjectsHandler object.ListObjectsHandler
	// UserListUsersHandler sets the operation handler for the list users operation
	UserListUsersHandler user.ListUsersHandler
	// AuthLoginHandler sets the operation handler for the login operation
	AuthLoginHandler auth.LoginHandler
	// AuthLoginDetailHandler sets the operation handler for the login detail operation
	AuthLoginDetailHandler auth.LoginDetailHandler
	// AuthLoginOauth2AuthHandler sets the operation handler for the login oauth2 auth operation
	AuthLoginOauth2AuthHandler auth.LoginOauth2AuthHandler
	// AuthLogoutHandler sets the operation handler for the logout operation
	AuthLogoutHandler auth.LogoutHandler
	// BucketMakeBucketHandler sets the operation handler for the make bucket operation
	BucketMakeBucketHandler bucket.MakeBucketHandler
	// ObjectPostBucketsBucketNameObjectsUploadHandler sets the operation handler for the post buckets bucket name objects upload operation
	ObjectPostBucketsBucketNameObjectsUploadHandler object.PostBucketsBucketNameObjectsUploadHandler
	// BucketPutBucketTagsHandler sets the operation handler for the put bucket tags operation
	BucketPutBucketTagsHandler bucket.PutBucketTagsHandler
	// ObjectPutObjectRestoreHandler sets the operation handler for the put object restore operation
	ObjectPutObjectRestoreHandler object.PutObjectRestoreHandler
	// ObjectPutObjectTagsHandler sets the operation handler for the put object tags operation
	ObjectPutObjectTagsHandler object.PutObjectTagsHandler
	// AuthSessionCheckHandler sets the operation handler for the session check operation
	AuthSessionCheckHandler auth.SessionCheckHandler
	// BucketSetBucketVersioningHandler sets the operation handler for the set bucket versioning operation
	BucketSetBucketVersioningHandler bucket.SetBucketVersioningHandler
	// ObjectShareObjectHandler sets the operation handler for the share object operation
	ObjectShareObjectHandler object.ShareObjectHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *ConsoleAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *ConsoleAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *ConsoleAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ConsoleAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ConsoleAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ConsoleAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ConsoleAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ConsoleAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ConsoleAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ConsoleAPI
func (o *ConsoleAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.BinProducer == nil {
		unregistered = append(unregistered, "BinProducer")
	}
	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.AnonymousAuth == nil {
		unregistered = append(unregistered, "XAnonymousAuth")
	}
	if o.KeyAuth == nil {
		unregistered = append(unregistered, "KeyAuth")
	}

	if o.SystemAdminInfoHandler == nil {
		unregistered = append(unregistered, "system.AdminInfoHandler")
	}
	if o.BucketBucketInfoHandler == nil {
		unregistered = append(unregistered, "bucket.BucketInfoHandler")
	}
	if o.ObjectDeleteMultipleObjectsHandler == nil {
		unregistered = append(unregistered, "object.DeleteMultipleObjectsHandler")
	}
	if o.ObjectDeleteObjectHandler == nil {
		unregistered = append(unregistered, "object.DeleteObjectHandler")
	}
	if o.ObjectDownloadObjectHandler == nil {
		unregistered = append(unregistered, "object.DownloadObjectHandler")
	}
	if o.ObjectDownloadMultipleObjectsHandler == nil {
		unregistered = append(unregistered, "object.DownloadMultipleObjectsHandler")
	}
	if o.PublicDownloadSharedObjectHandler == nil {
		unregistered = append(unregistered, "public.DownloadSharedObjectHandler")
	}
	if o.BucketGetBucketQuotaHandler == nil {
		unregistered = append(unregistered, "bucket.GetBucketQuotaHandler")
	}
	if o.BucketGetBucketRewindHandler == nil {
		unregistered = append(unregistered, "bucket.GetBucketRewindHandler")
	}
	if o.BucketGetBucketVersioningHandler == nil {
		unregistered = append(unregistered, "bucket.GetBucketVersioningHandler")
	}
	if o.BucketGetMaxShareLinkExpHandler == nil {
		unregistered = append(unregistered, "bucket.GetMaxShareLinkExpHandler")
	}
	if o.ObjectGetObjectMetadataHandler == nil {
		unregistered = append(unregistered, "object.GetObjectMetadataHandler")
	}
	if o.LicenseLicenseAcknowledgeHandler == nil {
		unregistered = append(unregistered, "license.LicenseAcknowledgeHandler")
	}
	if o.BucketListBucketsHandler == nil {
		unregistered = append(unregistered, "bucket.ListBucketsHandler")
	}
	if o.ObjectListObjectsHandler == nil {
		unregistered = append(unregistered, "object.ListObjectsHandler")
	}
	if o.UserListUsersHandler == nil {
		unregistered = append(unregistered, "user.ListUsersHandler")
	}
	if o.AuthLoginHandler == nil {
		unregistered = append(unregistered, "auth.LoginHandler")
	}
	if o.AuthLoginDetailHandler == nil {
		unregistered = append(unregistered, "auth.LoginDetailHandler")
	}
	if o.AuthLoginOauth2AuthHandler == nil {
		unregistered = append(unregistered, "auth.LoginOauth2AuthHandler")
	}
	if o.AuthLogoutHandler == nil {
		unregistered = append(unregistered, "auth.LogoutHandler")
	}
	if o.BucketMakeBucketHandler == nil {
		unregistered = append(unregistered, "bucket.MakeBucketHandler")
	}
	if o.ObjectPostBucketsBucketNameObjectsUploadHandler == nil {
		unregistered = append(unregistered, "object.PostBucketsBucketNameObjectsUploadHandler")
	}
	if o.BucketPutBucketTagsHandler == nil {
		unregistered = append(unregistered, "bucket.PutBucketTagsHandler")
	}
	if o.ObjectPutObjectRestoreHandler == nil {
		unregistered = append(unregistered, "object.PutObjectRestoreHandler")
	}
	if o.ObjectPutObjectTagsHandler == nil {
		unregistered = append(unregistered, "object.PutObjectTagsHandler")
	}
	if o.AuthSessionCheckHandler == nil {
		unregistered = append(unregistered, "auth.SessionCheckHandler")
	}
	if o.BucketSetBucketVersioningHandler == nil {
		unregistered = append(unregistered, "bucket.SetBucketVersioningHandler")
	}
	if o.ObjectShareObjectHandler == nil {
		unregistered = append(unregistered, "object.ShareObjectHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ConsoleAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ConsoleAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "anonymous":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.AnonymousAuth(token)
			})

		case "key":
			result[name] = o.BearerAuthenticator(name, func(token string, scopes []string) (interface{}, error) {
				return o.KeyAuth(token, scopes)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *ConsoleAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *ConsoleAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *ConsoleAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/octet-stream":
			result["application/octet-stream"] = o.BinProducer
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ConsoleAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the console API
func (o *ConsoleAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ConsoleAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/admin/info"] = system.NewAdminInfo(o.context, o.SystemAdminInfoHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/buckets/{name}"] = bucket.NewBucketInfo(o.context, o.BucketBucketInfoHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/buckets/{bucket_name}/delete-objects"] = object.NewDeleteMultipleObjects(o.context, o.ObjectDeleteMultipleObjectsHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/buckets/{bucket_name}/objects"] = object.NewDeleteObject(o.context, o.ObjectDeleteObjectHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/buckets/{bucket_name}/objects/download"] = object.NewDownloadObject(o.context, o.ObjectDownloadObjectHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/buckets/{bucket_name}/objects/download-multiple"] = object.NewDownloadMultipleObjects(o.context, o.ObjectDownloadMultipleObjectsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/download-shared-object/{url}"] = public.NewDownloadSharedObject(o.context, o.PublicDownloadSharedObjectHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/buckets/{name}/quota"] = bucket.NewGetBucketQuota(o.context, o.BucketGetBucketQuotaHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/buckets/{bucket_name}/rewind/{date}"] = bucket.NewGetBucketRewind(o.context, o.BucketGetBucketRewindHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/buckets/{bucket_name}/versioning"] = bucket.NewGetBucketVersioning(o.context, o.BucketGetBucketVersioningHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/buckets/max-share-exp"] = bucket.NewGetMaxShareLinkExp(o.context, o.BucketGetMaxShareLinkExpHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/buckets/{bucket_name}/objects/metadata"] = object.NewGetObjectMetadata(o.context, o.ObjectGetObjectMetadataHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/license/acknowledge"] = license.NewLicenseAcknowledge(o.context, o.LicenseLicenseAcknowledgeHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/buckets"] = bucket.NewListBuckets(o.context, o.BucketListBucketsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/buckets/{bucket_name}/objects"] = object.NewListObjects(o.context, o.ObjectListObjectsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users"] = user.NewListUsers(o.context, o.UserListUsersHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/login"] = auth.NewLogin(o.context, o.AuthLoginHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/login"] = auth.NewLoginDetail(o.context, o.AuthLoginDetailHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/login/oauth2/auth"] = auth.NewLoginOauth2Auth(o.context, o.AuthLoginOauth2AuthHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/logout"] = auth.NewLogout(o.context, o.AuthLogoutHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/buckets"] = bucket.NewMakeBucket(o.context, o.BucketMakeBucketHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/buckets/{bucket_name}/objects/upload"] = object.NewPostBucketsBucketNameObjectsUpload(o.context, o.ObjectPostBucketsBucketNameObjectsUploadHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/buckets/{bucket_name}/tags"] = bucket.NewPutBucketTags(o.context, o.BucketPutBucketTagsHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/buckets/{bucket_name}/objects/restore"] = object.NewPutObjectRestore(o.context, o.ObjectPutObjectRestoreHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/buckets/{bucket_name}/objects/tags"] = object.NewPutObjectTags(o.context, o.ObjectPutObjectTagsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/session"] = auth.NewSessionCheck(o.context, o.AuthSessionCheckHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/buckets/{bucket_name}/versioning"] = bucket.NewSetBucketVersioning(o.context, o.BucketSetBucketVersioningHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/buckets/{bucket_name}/objects/share"] = object.NewShareObject(o.context, o.ObjectShareObjectHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ConsoleAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *ConsoleAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *ConsoleAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *ConsoleAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *ConsoleAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}
