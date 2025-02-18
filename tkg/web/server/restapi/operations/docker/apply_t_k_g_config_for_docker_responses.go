// Code generated by go-swagger; DO NOT EDIT.

package docker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// ApplyTKGConfigForDockerOKCode is the HTTP code returned for type ApplyTKGConfigForDockerOK
const ApplyTKGConfigForDockerOKCode int = 200

/*ApplyTKGConfigForDockerOK Apply change to TKG configuration successfully

swagger:response applyTKGConfigForDockerOK
*/
type ApplyTKGConfigForDockerOK struct {

	/*
	  In: Body
	*/
	Payload *models.ConfigFileInfo `json:"body,omitempty"`
}

// NewApplyTKGConfigForDockerOK creates ApplyTKGConfigForDockerOK with default headers values
func NewApplyTKGConfigForDockerOK() *ApplyTKGConfigForDockerOK {

	return &ApplyTKGConfigForDockerOK{}
}

// WithPayload adds the payload to the apply t k g config for docker o k response
func (o *ApplyTKGConfigForDockerOK) WithPayload(payload *models.ConfigFileInfo) *ApplyTKGConfigForDockerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the apply t k g config for docker o k response
func (o *ApplyTKGConfigForDockerOK) SetPayload(payload *models.ConfigFileInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ApplyTKGConfigForDockerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ApplyTKGConfigForDockerBadRequestCode is the HTTP code returned for type ApplyTKGConfigForDockerBadRequest
const ApplyTKGConfigForDockerBadRequestCode int = 400

/*ApplyTKGConfigForDockerBadRequest Bad request

swagger:response applyTKGConfigForDockerBadRequest
*/
type ApplyTKGConfigForDockerBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewApplyTKGConfigForDockerBadRequest creates ApplyTKGConfigForDockerBadRequest with default headers values
func NewApplyTKGConfigForDockerBadRequest() *ApplyTKGConfigForDockerBadRequest {

	return &ApplyTKGConfigForDockerBadRequest{}
}

// WithPayload adds the payload to the apply t k g config for docker bad request response
func (o *ApplyTKGConfigForDockerBadRequest) WithPayload(payload *models.Error) *ApplyTKGConfigForDockerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the apply t k g config for docker bad request response
func (o *ApplyTKGConfigForDockerBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ApplyTKGConfigForDockerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ApplyTKGConfigForDockerInternalServerErrorCode is the HTTP code returned for type ApplyTKGConfigForDockerInternalServerError
const ApplyTKGConfigForDockerInternalServerErrorCode int = 500

/*ApplyTKGConfigForDockerInternalServerError Internal server error

swagger:response applyTKGConfigForDockerInternalServerError
*/
type ApplyTKGConfigForDockerInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewApplyTKGConfigForDockerInternalServerError creates ApplyTKGConfigForDockerInternalServerError with default headers values
func NewApplyTKGConfigForDockerInternalServerError() *ApplyTKGConfigForDockerInternalServerError {

	return &ApplyTKGConfigForDockerInternalServerError{}
}

// WithPayload adds the payload to the apply t k g config for docker internal server error response
func (o *ApplyTKGConfigForDockerInternalServerError) WithPayload(payload *models.Error) *ApplyTKGConfigForDockerInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the apply t k g config for docker internal server error response
func (o *ApplyTKGConfigForDockerInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ApplyTKGConfigForDockerInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
