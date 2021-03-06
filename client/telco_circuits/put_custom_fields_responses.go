// Code generated by go-swagger; DO NOT EDIT.

package telco_circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PutCustomFieldsReader is a Reader for the PutCustomFields structure.
type PutCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomFieldsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutCustomFieldsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutCustomFieldsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutCustomFieldsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutCustomFieldsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPutCustomFieldsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPutCustomFieldsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutCustomFieldsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutCustomFieldsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCustomFieldsOK creates a PutCustomFieldsOK with default headers values
func NewPutCustomFieldsOK() *PutCustomFieldsOK {
	return &PutCustomFieldsOK{}
}

/*PutCustomFieldsOK handles this case with default header values.

The above command returns results like this:
*/
type PutCustomFieldsOK struct {
	Payload *PutCustomFieldsOKBody
}

func (o *PutCustomFieldsOK) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/circuit/][%d] putCustomFieldsOK  %+v", 200, o.Payload)
}

func (o *PutCustomFieldsOK) GetPayload() *PutCustomFieldsOKBody {
	return o.Payload
}

func (o *PutCustomFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutCustomFieldsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCustomFieldsBadRequest creates a PutCustomFieldsBadRequest with default headers values
func NewPutCustomFieldsBadRequest() *PutCustomFieldsBadRequest {
	return &PutCustomFieldsBadRequest{}
}

/*PutCustomFieldsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PutCustomFieldsBadRequest struct {
}

func (o *PutCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/circuit/][%d] putCustomFieldsBadRequest ", 400)
}

func (o *PutCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsUnauthorized creates a PutCustomFieldsUnauthorized with default headers values
func NewPutCustomFieldsUnauthorized() *PutCustomFieldsUnauthorized {
	return &PutCustomFieldsUnauthorized{}
}

/*PutCustomFieldsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PutCustomFieldsUnauthorized struct {
}

func (o *PutCustomFieldsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/circuit/][%d] putCustomFieldsUnauthorized ", 401)
}

func (o *PutCustomFieldsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsForbidden creates a PutCustomFieldsForbidden with default headers values
func NewPutCustomFieldsForbidden() *PutCustomFieldsForbidden {
	return &PutCustomFieldsForbidden{}
}

/*PutCustomFieldsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PutCustomFieldsForbidden struct {
}

func (o *PutCustomFieldsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/circuit/][%d] putCustomFieldsForbidden ", 403)
}

func (o *PutCustomFieldsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsNotFound creates a PutCustomFieldsNotFound with default headers values
func NewPutCustomFieldsNotFound() *PutCustomFieldsNotFound {
	return &PutCustomFieldsNotFound{}
}

/*PutCustomFieldsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PutCustomFieldsNotFound struct {
}

func (o *PutCustomFieldsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/circuit/][%d] putCustomFieldsNotFound ", 404)
}

func (o *PutCustomFieldsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsMethodNotAllowed creates a PutCustomFieldsMethodNotAllowed with default headers values
func NewPutCustomFieldsMethodNotAllowed() *PutCustomFieldsMethodNotAllowed {
	return &PutCustomFieldsMethodNotAllowed{}
}

/*PutCustomFieldsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PutCustomFieldsMethodNotAllowed struct {
}

func (o *PutCustomFieldsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/circuit/][%d] putCustomFieldsMethodNotAllowed ", 405)
}

func (o *PutCustomFieldsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsGone creates a PutCustomFieldsGone with default headers values
func NewPutCustomFieldsGone() *PutCustomFieldsGone {
	return &PutCustomFieldsGone{}
}

/*PutCustomFieldsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PutCustomFieldsGone struct {
}

func (o *PutCustomFieldsGone) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/circuit/][%d] putCustomFieldsGone ", 410)
}

func (o *PutCustomFieldsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsInternalServerError creates a PutCustomFieldsInternalServerError with default headers values
func NewPutCustomFieldsInternalServerError() *PutCustomFieldsInternalServerError {
	return &PutCustomFieldsInternalServerError{}
}

/*PutCustomFieldsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type PutCustomFieldsInternalServerError struct {
}

func (o *PutCustomFieldsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/circuit/][%d] putCustomFieldsInternalServerError ", 500)
}

func (o *PutCustomFieldsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldsServiceUnavailable creates a PutCustomFieldsServiceUnavailable with default headers values
func NewPutCustomFieldsServiceUnavailable() *PutCustomFieldsServiceUnavailable {
	return &PutCustomFieldsServiceUnavailable{}
}

/*PutCustomFieldsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PutCustomFieldsServiceUnavailable struct {
}

func (o *PutCustomFieldsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/circuit/][%d] putCustomFieldsServiceUnavailable ", 503)
}

func (o *PutCustomFieldsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PutCustomFieldsOKBody put custom fields o k body
swagger:model PutCustomFieldsOKBody
*/
type PutCustomFieldsOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this put custom fields o k body
func (o *PutCustomFieldsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutCustomFieldsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCustomFieldsOKBody) UnmarshalBinary(b []byte) error {
	var res PutCustomFieldsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
