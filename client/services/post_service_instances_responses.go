// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostServiceInstancesReader is a Reader for the PostServiceInstances structure.
type PostServiceInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostServiceInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostServiceInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostServiceInstancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostServiceInstancesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostServiceInstancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostServiceInstancesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostServiceInstancesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostServiceInstancesGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostServiceInstancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostServiceInstancesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostServiceInstancesOK creates a PostServiceInstancesOK with default headers values
func NewPostServiceInstancesOK() *PostServiceInstancesOK {
	return &PostServiceInstancesOK{}
}

/*PostServiceInstancesOK handles this case with default header values.

The above command returns results like this:
*/
type PostServiceInstancesOK struct {
	Payload *PostServiceInstancesOKBody
}

func (o *PostServiceInstancesOK) Error() string {
	return fmt.Sprintf("[POST /api/2.0/service_instances/][%d] postServiceInstancesOK  %+v", 200, o.Payload)
}

func (o *PostServiceInstancesOK) GetPayload() *PostServiceInstancesOKBody {
	return o.Payload
}

func (o *PostServiceInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostServiceInstancesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostServiceInstancesBadRequest creates a PostServiceInstancesBadRequest with default headers values
func NewPostServiceInstancesBadRequest() *PostServiceInstancesBadRequest {
	return &PostServiceInstancesBadRequest{}
}

/*PostServiceInstancesBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostServiceInstancesBadRequest struct {
}

func (o *PostServiceInstancesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/2.0/service_instances/][%d] postServiceInstancesBadRequest ", 400)
}

func (o *PostServiceInstancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostServiceInstancesUnauthorized creates a PostServiceInstancesUnauthorized with default headers values
func NewPostServiceInstancesUnauthorized() *PostServiceInstancesUnauthorized {
	return &PostServiceInstancesUnauthorized{}
}

/*PostServiceInstancesUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostServiceInstancesUnauthorized struct {
}

func (o *PostServiceInstancesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/2.0/service_instances/][%d] postServiceInstancesUnauthorized ", 401)
}

func (o *PostServiceInstancesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostServiceInstancesForbidden creates a PostServiceInstancesForbidden with default headers values
func NewPostServiceInstancesForbidden() *PostServiceInstancesForbidden {
	return &PostServiceInstancesForbidden{}
}

/*PostServiceInstancesForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostServiceInstancesForbidden struct {
}

func (o *PostServiceInstancesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/2.0/service_instances/][%d] postServiceInstancesForbidden ", 403)
}

func (o *PostServiceInstancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostServiceInstancesNotFound creates a PostServiceInstancesNotFound with default headers values
func NewPostServiceInstancesNotFound() *PostServiceInstancesNotFound {
	return &PostServiceInstancesNotFound{}
}

/*PostServiceInstancesNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostServiceInstancesNotFound struct {
}

func (o *PostServiceInstancesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/2.0/service_instances/][%d] postServiceInstancesNotFound ", 404)
}

func (o *PostServiceInstancesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostServiceInstancesMethodNotAllowed creates a PostServiceInstancesMethodNotAllowed with default headers values
func NewPostServiceInstancesMethodNotAllowed() *PostServiceInstancesMethodNotAllowed {
	return &PostServiceInstancesMethodNotAllowed{}
}

/*PostServiceInstancesMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostServiceInstancesMethodNotAllowed struct {
}

func (o *PostServiceInstancesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/2.0/service_instances/][%d] postServiceInstancesMethodNotAllowed ", 405)
}

func (o *PostServiceInstancesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostServiceInstancesGone creates a PostServiceInstancesGone with default headers values
func NewPostServiceInstancesGone() *PostServiceInstancesGone {
	return &PostServiceInstancesGone{}
}

/*PostServiceInstancesGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostServiceInstancesGone struct {
}

func (o *PostServiceInstancesGone) Error() string {
	return fmt.Sprintf("[POST /api/2.0/service_instances/][%d] postServiceInstancesGone ", 410)
}

func (o *PostServiceInstancesGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostServiceInstancesInternalServerError creates a PostServiceInstancesInternalServerError with default headers values
func NewPostServiceInstancesInternalServerError() *PostServiceInstancesInternalServerError {
	return &PostServiceInstancesInternalServerError{}
}

/*PostServiceInstancesInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type PostServiceInstancesInternalServerError struct {
}

func (o *PostServiceInstancesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/2.0/service_instances/][%d] postServiceInstancesInternalServerError ", 500)
}

func (o *PostServiceInstancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostServiceInstancesServiceUnavailable creates a PostServiceInstancesServiceUnavailable with default headers values
func NewPostServiceInstancesServiceUnavailable() *PostServiceInstancesServiceUnavailable {
	return &PostServiceInstancesServiceUnavailable{}
}

/*PostServiceInstancesServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostServiceInstancesServiceUnavailable struct {
}

func (o *PostServiceInstancesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/2.0/service_instances/][%d] postServiceInstancesServiceUnavailable ", 503)
}

func (o *PostServiceInstancesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostServiceInstancesOKBody post service instances o k body
swagger:model PostServiceInstancesOKBody
*/
type PostServiceInstancesOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post service instances o k body
func (o *PostServiceInstancesOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostServiceInstancesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostServiceInstancesOKBody) UnmarshalBinary(b []byte) error {
	var res PostServiceInstancesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
