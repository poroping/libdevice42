// Code generated by go-swagger; DO NOT EDIT.

package parts_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PutCustomFieldPartReader is a Reader for the PutCustomFieldPart structure.
type PutCustomFieldPartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomFieldPartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomFieldPartOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutCustomFieldPartBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutCustomFieldPartUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutCustomFieldPartForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutCustomFieldPartNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPutCustomFieldPartMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPutCustomFieldPartGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutCustomFieldPartInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutCustomFieldPartServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCustomFieldPartOK creates a PutCustomFieldPartOK with default headers values
func NewPutCustomFieldPartOK() *PutCustomFieldPartOK {
	return &PutCustomFieldPartOK{}
}

/*PutCustomFieldPartOK handles this case with default header values.

The above command returns results like this:
*/
type PutCustomFieldPartOK struct {
	Payload *PutCustomFieldPartOKBody
}

func (o *PutCustomFieldPartOK) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/part/][%d] putCustomFieldPartOK  %+v", 200, o.Payload)
}

func (o *PutCustomFieldPartOK) GetPayload() *PutCustomFieldPartOKBody {
	return o.Payload
}

func (o *PutCustomFieldPartOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutCustomFieldPartOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCustomFieldPartBadRequest creates a PutCustomFieldPartBadRequest with default headers values
func NewPutCustomFieldPartBadRequest() *PutCustomFieldPartBadRequest {
	return &PutCustomFieldPartBadRequest{}
}

/*PutCustomFieldPartBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PutCustomFieldPartBadRequest struct {
}

func (o *PutCustomFieldPartBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/part/][%d] putCustomFieldPartBadRequest ", 400)
}

func (o *PutCustomFieldPartBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldPartUnauthorized creates a PutCustomFieldPartUnauthorized with default headers values
func NewPutCustomFieldPartUnauthorized() *PutCustomFieldPartUnauthorized {
	return &PutCustomFieldPartUnauthorized{}
}

/*PutCustomFieldPartUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PutCustomFieldPartUnauthorized struct {
}

func (o *PutCustomFieldPartUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/part/][%d] putCustomFieldPartUnauthorized ", 401)
}

func (o *PutCustomFieldPartUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldPartForbidden creates a PutCustomFieldPartForbidden with default headers values
func NewPutCustomFieldPartForbidden() *PutCustomFieldPartForbidden {
	return &PutCustomFieldPartForbidden{}
}

/*PutCustomFieldPartForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PutCustomFieldPartForbidden struct {
}

func (o *PutCustomFieldPartForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/part/][%d] putCustomFieldPartForbidden ", 403)
}

func (o *PutCustomFieldPartForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldPartNotFound creates a PutCustomFieldPartNotFound with default headers values
func NewPutCustomFieldPartNotFound() *PutCustomFieldPartNotFound {
	return &PutCustomFieldPartNotFound{}
}

/*PutCustomFieldPartNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PutCustomFieldPartNotFound struct {
}

func (o *PutCustomFieldPartNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/part/][%d] putCustomFieldPartNotFound ", 404)
}

func (o *PutCustomFieldPartNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldPartMethodNotAllowed creates a PutCustomFieldPartMethodNotAllowed with default headers values
func NewPutCustomFieldPartMethodNotAllowed() *PutCustomFieldPartMethodNotAllowed {
	return &PutCustomFieldPartMethodNotAllowed{}
}

/*PutCustomFieldPartMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PutCustomFieldPartMethodNotAllowed struct {
}

func (o *PutCustomFieldPartMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/part/][%d] putCustomFieldPartMethodNotAllowed ", 405)
}

func (o *PutCustomFieldPartMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldPartGone creates a PutCustomFieldPartGone with default headers values
func NewPutCustomFieldPartGone() *PutCustomFieldPartGone {
	return &PutCustomFieldPartGone{}
}

/*PutCustomFieldPartGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PutCustomFieldPartGone struct {
}

func (o *PutCustomFieldPartGone) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/part/][%d] putCustomFieldPartGone ", 410)
}

func (o *PutCustomFieldPartGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldPartInternalServerError creates a PutCustomFieldPartInternalServerError with default headers values
func NewPutCustomFieldPartInternalServerError() *PutCustomFieldPartInternalServerError {
	return &PutCustomFieldPartInternalServerError{}
}

/*PutCustomFieldPartInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type PutCustomFieldPartInternalServerError struct {
}

func (o *PutCustomFieldPartInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/part/][%d] putCustomFieldPartInternalServerError ", 500)
}

func (o *PutCustomFieldPartInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomFieldPartServiceUnavailable creates a PutCustomFieldPartServiceUnavailable with default headers values
func NewPutCustomFieldPartServiceUnavailable() *PutCustomFieldPartServiceUnavailable {
	return &PutCustomFieldPartServiceUnavailable{}
}

/*PutCustomFieldPartServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PutCustomFieldPartServiceUnavailable struct {
}

func (o *PutCustomFieldPartServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/custom_fields/part/][%d] putCustomFieldPartServiceUnavailable ", 503)
}

func (o *PutCustomFieldPartServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PutCustomFieldPartOKBody put custom field part o k body
swagger:model PutCustomFieldPartOKBody
*/
type PutCustomFieldPartOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this put custom field part o k body
func (o *PutCustomFieldPartOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutCustomFieldPartOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCustomFieldPartOKBody) UnmarshalBinary(b []byte) error {
	var res PutCustomFieldPartOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}