// Code generated by go-swagger; DO NOT EDIT.

package racks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PutCustomFieldsRackReader is a Reader for the PutCustomFieldsRack structure.
type PutCustomFieldsRackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomFieldsRackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomFieldsRackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutCustomFieldsRackBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutCustomFieldsRackUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutCustomFieldsRackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutCustomFieldsRackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPutCustomFieldsRackMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPutCustomFieldsRackGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutCustomFieldsRackInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutCustomFieldsRackServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCustomFieldsRackOK creates a PutCustomFieldsRackOK with default headers values
func NewPutCustomFieldsRackOK() *PutCustomFieldsRackOK {
	return &PutCustomFieldsRackOK{}
}

/*PutCustomFieldsRackOK handles this case with default header values.

The above command returns results like this:
*/
type PutCustomFieldsRackOK struct {
	Payload *PutCustomFieldsRackOKBody
}

func (o *PutCustomFieldsRackOK) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/rack/][%d] putCustomFieldsRackOK  %+v", 200, o.Payload)
}

func (o *PutCustomFieldsRackOK) GetPayload() *PutCustomFieldsRackOKBody {
	return o.Payload
}

func (o *PutCustomFieldsRackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutCustomFieldsRackOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCustomFieldsRackBadRequest creates a PutCustomFieldsRackBadRequest with default headers values
func NewPutCustomFieldsRackBadRequest() *PutCustomFieldsRackBadRequest {
	return &PutCustomFieldsRackBadRequest{}
}

/*PutCustomFieldsRackBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PutCustomFieldsRackBadRequest struct {
}

func (o *PutCustomFieldsRackBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/rack/][%d] putCustomFieldsRackBadRequest ", 400)
}

func (o *PutCustomFieldsRackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsRackUnauthorized creates a PutCustomFieldsRackUnauthorized with default headers values
func NewPutCustomFieldsRackUnauthorized() *PutCustomFieldsRackUnauthorized {
	return &PutCustomFieldsRackUnauthorized{}
}

/*PutCustomFieldsRackUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PutCustomFieldsRackUnauthorized struct {
}

func (o *PutCustomFieldsRackUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/rack/][%d] putCustomFieldsRackUnauthorized ", 401)
}

func (o *PutCustomFieldsRackUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsRackForbidden creates a PutCustomFieldsRackForbidden with default headers values
func NewPutCustomFieldsRackForbidden() *PutCustomFieldsRackForbidden {
	return &PutCustomFieldsRackForbidden{}
}

/*PutCustomFieldsRackForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PutCustomFieldsRackForbidden struct {
}

func (o *PutCustomFieldsRackForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/rack/][%d] putCustomFieldsRackForbidden ", 403)
}

func (o *PutCustomFieldsRackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsRackNotFound creates a PutCustomFieldsRackNotFound with default headers values
func NewPutCustomFieldsRackNotFound() *PutCustomFieldsRackNotFound {
	return &PutCustomFieldsRackNotFound{}
}

/*PutCustomFieldsRackNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PutCustomFieldsRackNotFound struct {
}

func (o *PutCustomFieldsRackNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/rack/][%d] putCustomFieldsRackNotFound ", 404)
}

func (o *PutCustomFieldsRackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsRackMethodNotAllowed creates a PutCustomFieldsRackMethodNotAllowed with default headers values
func NewPutCustomFieldsRackMethodNotAllowed() *PutCustomFieldsRackMethodNotAllowed {
	return &PutCustomFieldsRackMethodNotAllowed{}
}

/*PutCustomFieldsRackMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PutCustomFieldsRackMethodNotAllowed struct {
}

func (o *PutCustomFieldsRackMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/rack/][%d] putCustomFieldsRackMethodNotAllowed ", 405)
}

func (o *PutCustomFieldsRackMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsRackGone creates a PutCustomFieldsRackGone with default headers values
func NewPutCustomFieldsRackGone() *PutCustomFieldsRackGone {
	return &PutCustomFieldsRackGone{}
}

/*PutCustomFieldsRackGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PutCustomFieldsRackGone struct {
}

func (o *PutCustomFieldsRackGone) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/rack/][%d] putCustomFieldsRackGone ", 410)
}

func (o *PutCustomFieldsRackGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsRackInternalServerError creates a PutCustomFieldsRackInternalServerError with default headers values
func NewPutCustomFieldsRackInternalServerError() *PutCustomFieldsRackInternalServerError {
	return &PutCustomFieldsRackInternalServerError{}
}

/*PutCustomFieldsRackInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type PutCustomFieldsRackInternalServerError struct {
}

func (o *PutCustomFieldsRackInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/rack/][%d] putCustomFieldsRackInternalServerError ", 500)
}

func (o *PutCustomFieldsRackInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsRackServiceUnavailable creates a PutCustomFieldsRackServiceUnavailable with default headers values
func NewPutCustomFieldsRackServiceUnavailable() *PutCustomFieldsRackServiceUnavailable {
	return &PutCustomFieldsRackServiceUnavailable{}
}

/*PutCustomFieldsRackServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PutCustomFieldsRackServiceUnavailable struct {
}

func (o *PutCustomFieldsRackServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/rack/][%d] putCustomFieldsRackServiceUnavailable ", 503)
}

func (o *PutCustomFieldsRackServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PutCustomFieldsRackOKBody put custom fields rack o k body
swagger:model PutCustomFieldsRackOKBody
*/
type PutCustomFieldsRackOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this put custom fields rack o k body
func (o *PutCustomFieldsRackOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutCustomFieldsRackOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCustomFieldsRackOKBody) UnmarshalBinary(b []byte) error {
	var res PutCustomFieldsRackOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
