// Code generated by go-swagger; DO NOT EDIT.

package iam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"iam/models"
)

// PostIamServiceTokenOKCode is the HTTP code returned for type PostIamServiceTokenOK
const PostIamServiceTokenOKCode int = 200

/*PostIamServiceTokenOK Success

swagger:response postIamServiceTokenOK
*/
type PostIamServiceTokenOK struct {

	/*
	  In: Body
	*/
	Payload *models.IamTokenResponse `json:"body,omitempty"`
}

// NewPostIamServiceTokenOK creates PostIamServiceTokenOK with default headers values
func NewPostIamServiceTokenOK() *PostIamServiceTokenOK {

	return &PostIamServiceTokenOK{}
}

// WithPayload adds the payload to the post iam service token o k response
func (o *PostIamServiceTokenOK) WithPayload(payload *models.IamTokenResponse) *PostIamServiceTokenOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post iam service token o k response
func (o *PostIamServiceTokenOK) SetPayload(payload *models.IamTokenResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostIamServiceTokenOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostIamServiceTokenDefault Internal Error

swagger:response postIamServiceTokenDefault
*/
type PostIamServiceTokenDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Status `json:"body,omitempty"`
}

// NewPostIamServiceTokenDefault creates PostIamServiceTokenDefault with default headers values
func NewPostIamServiceTokenDefault(code int) *PostIamServiceTokenDefault {
	if code <= 0 {
		code = 500
	}

	return &PostIamServiceTokenDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post iam service token default response
func (o *PostIamServiceTokenDefault) WithStatusCode(code int) *PostIamServiceTokenDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post iam service token default response
func (o *PostIamServiceTokenDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post iam service token default response
func (o *PostIamServiceTokenDefault) WithPayload(payload *models.Status) *PostIamServiceTokenDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post iam service token default response
func (o *PostIamServiceTokenDefault) SetPayload(payload *models.Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostIamServiceTokenDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}