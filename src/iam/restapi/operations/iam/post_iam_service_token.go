// Code generated by go-swagger; DO NOT EDIT.

package iam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostIamServiceTokenHandlerFunc turns a function with the right signature into a post iam service token handler
type PostIamServiceTokenHandlerFunc func(PostIamServiceTokenParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostIamServiceTokenHandlerFunc) Handle(params PostIamServiceTokenParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostIamServiceTokenHandler interface for that can handle valid post iam service token params
type PostIamServiceTokenHandler interface {
	Handle(PostIamServiceTokenParams, interface{}) middleware.Responder
}

// NewPostIamServiceToken creates a new http.Handler for the post iam service token operation
func NewPostIamServiceToken(ctx *middleware.Context, handler PostIamServiceTokenHandler) *PostIamServiceToken {
	return &PostIamServiceToken{Context: ctx, Handler: handler}
}

/*PostIamServiceToken swagger:route POST /iam/service_token iam postIamServiceToken

Token endpoint for service to service authentication

Token endpoint for service to service authentication

*/
type PostIamServiceToken struct {
	Context *middleware.Context
	Handler PostIamServiceTokenHandler
}

func (o *PostIamServiceToken) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostIamServiceTokenParams()

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
