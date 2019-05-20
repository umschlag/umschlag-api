// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/umschlag/umschlag-api/pkg/api/v1/models"
)

// DelteTeamFromUserOKCode is the HTTP code returned for type DelteTeamFromUserOK
const DelteTeamFromUserOKCode int = 200

/*DelteTeamFromUserOK Plain success message

swagger:response delteTeamFromUserOK
*/
type DelteTeamFromUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDelteTeamFromUserOK creates DelteTeamFromUserOK with default headers values
func NewDelteTeamFromUserOK() *DelteTeamFromUserOK {

	return &DelteTeamFromUserOK{}
}

// WithPayload adds the payload to the delte team from user o k response
func (o *DelteTeamFromUserOK) WithPayload(payload *models.GeneralError) *DelteTeamFromUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delte team from user o k response
func (o *DelteTeamFromUserOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DelteTeamFromUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DelteTeamFromUserForbiddenCode is the HTTP code returned for type DelteTeamFromUserForbidden
const DelteTeamFromUserForbiddenCode int = 403

/*DelteTeamFromUserForbidden User is not authorized

swagger:response delteTeamFromUserForbidden
*/
type DelteTeamFromUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDelteTeamFromUserForbidden creates DelteTeamFromUserForbidden with default headers values
func NewDelteTeamFromUserForbidden() *DelteTeamFromUserForbidden {

	return &DelteTeamFromUserForbidden{}
}

// WithPayload adds the payload to the delte team from user forbidden response
func (o *DelteTeamFromUserForbidden) WithPayload(payload *models.GeneralError) *DelteTeamFromUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delte team from user forbidden response
func (o *DelteTeamFromUserForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DelteTeamFromUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DelteTeamFromUserPreconditionFailedCode is the HTTP code returned for type DelteTeamFromUserPreconditionFailed
const DelteTeamFromUserPreconditionFailedCode int = 412

/*DelteTeamFromUserPreconditionFailed Failed to parse request body

swagger:response delteTeamFromUserPreconditionFailed
*/
type DelteTeamFromUserPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDelteTeamFromUserPreconditionFailed creates DelteTeamFromUserPreconditionFailed with default headers values
func NewDelteTeamFromUserPreconditionFailed() *DelteTeamFromUserPreconditionFailed {

	return &DelteTeamFromUserPreconditionFailed{}
}

// WithPayload adds the payload to the delte team from user precondition failed response
func (o *DelteTeamFromUserPreconditionFailed) WithPayload(payload *models.GeneralError) *DelteTeamFromUserPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delte team from user precondition failed response
func (o *DelteTeamFromUserPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DelteTeamFromUserPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DelteTeamFromUserUnprocessableEntityCode is the HTTP code returned for type DelteTeamFromUserUnprocessableEntity
const DelteTeamFromUserUnprocessableEntityCode int = 422

/*DelteTeamFromUserUnprocessableEntity User is not assigned

swagger:response delteTeamFromUserUnprocessableEntity
*/
type DelteTeamFromUserUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDelteTeamFromUserUnprocessableEntity creates DelteTeamFromUserUnprocessableEntity with default headers values
func NewDelteTeamFromUserUnprocessableEntity() *DelteTeamFromUserUnprocessableEntity {

	return &DelteTeamFromUserUnprocessableEntity{}
}

// WithPayload adds the payload to the delte team from user unprocessable entity response
func (o *DelteTeamFromUserUnprocessableEntity) WithPayload(payload *models.GeneralError) *DelteTeamFromUserUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delte team from user unprocessable entity response
func (o *DelteTeamFromUserUnprocessableEntity) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DelteTeamFromUserUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DelteTeamFromUserDefault Some error unrelated to the handler

swagger:response delteTeamFromUserDefault
*/
type DelteTeamFromUserDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDelteTeamFromUserDefault creates DelteTeamFromUserDefault with default headers values
func NewDelteTeamFromUserDefault(code int) *DelteTeamFromUserDefault {
	if code <= 0 {
		code = 500
	}

	return &DelteTeamFromUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delte team from user default response
func (o *DelteTeamFromUserDefault) WithStatusCode(code int) *DelteTeamFromUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delte team from user default response
func (o *DelteTeamFromUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delte team from user default response
func (o *DelteTeamFromUserDefault) WithPayload(payload *models.GeneralError) *DelteTeamFromUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delte team from user default response
func (o *DelteTeamFromUserDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DelteTeamFromUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
