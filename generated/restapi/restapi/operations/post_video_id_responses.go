// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dkgv/dislikes/generated/restapi/models"
)

// PostVideoIDOKCode is the HTTP code returned for type PostVideoIDOK
const PostVideoIDOKCode int = 200

/*PostVideoIDOK Success

swagger:response postVideoIdOK
*/
type PostVideoIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.VideoResponse `json:"body,omitempty"`
}

// NewPostVideoIDOK creates PostVideoIDOK with default headers values
func NewPostVideoIDOK() *PostVideoIDOK {

	return &PostVideoIDOK{}
}

// WithPayload adds the payload to the post video Id o k response
func (o *PostVideoIDOK) WithPayload(payload *models.VideoResponse) *PostVideoIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post video Id o k response
func (o *PostVideoIDOK) SetPayload(payload *models.VideoResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostVideoIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostVideoIDBadRequestCode is the HTTP code returned for type PostVideoIDBadRequest
const PostVideoIDBadRequestCode int = 400

/*PostVideoIDBadRequest Bad Request

swagger:response postVideoIdBadRequest
*/
type PostVideoIDBadRequest struct {
}

// NewPostVideoIDBadRequest creates PostVideoIDBadRequest with default headers values
func NewPostVideoIDBadRequest() *PostVideoIDBadRequest {

	return &PostVideoIDBadRequest{}
}

// WriteResponse to the client
func (o *PostVideoIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
