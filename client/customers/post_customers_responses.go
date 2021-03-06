// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostCustomersReader is a Reader for the PostCustomers structure.
type PostCustomersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCustomersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostCustomersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostCustomersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostCustomersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostCustomersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostCustomersMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostCustomersGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostCustomersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostCustomersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCustomersOK creates a PostCustomersOK with default headers values
func NewPostCustomersOK() *PostCustomersOK {
	return &PostCustomersOK{}
}

/*PostCustomersOK handles this case with default header values.

The above command returns results like this:
*/
type PostCustomersOK struct {
	Payload *PostCustomersOKBody
}

func (o *PostCustomersOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/customers/][%d] postCustomersOK  %+v", 200, o.Payload)
}

func (o *PostCustomersOK) GetPayload() *PostCustomersOKBody {
	return o.Payload
}

func (o *PostCustomersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostCustomersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCustomersBadRequest creates a PostCustomersBadRequest with default headers values
func NewPostCustomersBadRequest() *PostCustomersBadRequest {
	return &PostCustomersBadRequest{}
}

/*PostCustomersBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostCustomersBadRequest struct {
}

func (o *PostCustomersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/customers/][%d] postCustomersBadRequest ", 400)
}

func (o *PostCustomersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomersUnauthorized creates a PostCustomersUnauthorized with default headers values
func NewPostCustomersUnauthorized() *PostCustomersUnauthorized {
	return &PostCustomersUnauthorized{}
}

/*PostCustomersUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostCustomersUnauthorized struct {
}

func (o *PostCustomersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/customers/][%d] postCustomersUnauthorized ", 401)
}

func (o *PostCustomersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomersForbidden creates a PostCustomersForbidden with default headers values
func NewPostCustomersForbidden() *PostCustomersForbidden {
	return &PostCustomersForbidden{}
}

/*PostCustomersForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostCustomersForbidden struct {
}

func (o *PostCustomersForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/customers/][%d] postCustomersForbidden ", 403)
}

func (o *PostCustomersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomersNotFound creates a PostCustomersNotFound with default headers values
func NewPostCustomersNotFound() *PostCustomersNotFound {
	return &PostCustomersNotFound{}
}

/*PostCustomersNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostCustomersNotFound struct {
}

func (o *PostCustomersNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/customers/][%d] postCustomersNotFound ", 404)
}

func (o *PostCustomersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomersMethodNotAllowed creates a PostCustomersMethodNotAllowed with default headers values
func NewPostCustomersMethodNotAllowed() *PostCustomersMethodNotAllowed {
	return &PostCustomersMethodNotAllowed{}
}

/*PostCustomersMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostCustomersMethodNotAllowed struct {
}

func (o *PostCustomersMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/customers/][%d] postCustomersMethodNotAllowed ", 405)
}

func (o *PostCustomersMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomersGone creates a PostCustomersGone with default headers values
func NewPostCustomersGone() *PostCustomersGone {
	return &PostCustomersGone{}
}

/*PostCustomersGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostCustomersGone struct {
}

func (o *PostCustomersGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/customers/][%d] postCustomersGone ", 410)
}

func (o *PostCustomersGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomersInternalServerError creates a PostCustomersInternalServerError with default headers values
func NewPostCustomersInternalServerError() *PostCustomersInternalServerError {
	return &PostCustomersInternalServerError{}
}

/*PostCustomersInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type PostCustomersInternalServerError struct {
}

func (o *PostCustomersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/customers/][%d] postCustomersInternalServerError ", 500)
}

func (o *PostCustomersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomersServiceUnavailable creates a PostCustomersServiceUnavailable with default headers values
func NewPostCustomersServiceUnavailable() *PostCustomersServiceUnavailable {
	return &PostCustomersServiceUnavailable{}
}

/*PostCustomersServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostCustomersServiceUnavailable struct {
}

func (o *PostCustomersServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/customers/][%d] postCustomersServiceUnavailable ", 503)
}

func (o *PostCustomersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostCustomersOKBody post customers o k body
swagger:model PostCustomersOKBody
*/
type PostCustomersOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post customers o k body
func (o *PostCustomersOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCustomersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCustomersOKBody) UnmarshalBinary(b []byte) error {
	var res PostCustomersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
