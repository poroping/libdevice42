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

// DeleteCustomersReader is a Reader for the DeleteCustomers structure.
type DeleteCustomersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCustomersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCustomersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteCustomersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCustomersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCustomersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeleteCustomersMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewDeleteCustomersGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCustomersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteCustomersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCustomersOK creates a DeleteCustomersOK with default headers values
func NewDeleteCustomersOK() *DeleteCustomersOK {
	return &DeleteCustomersOK{}
}

/*DeleteCustomersOK handles this case with default header values.

The above command returns results like this:
*/
type DeleteCustomersOK struct {
	Payload *DeleteCustomersOKBody
}

func (o *DeleteCustomersOK) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/customers/{id}/][%d] deleteCustomersOK  %+v", 200, o.Payload)
}

func (o *DeleteCustomersOK) GetPayload() *DeleteCustomersOKBody {
	return o.Payload
}

func (o *DeleteCustomersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteCustomersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomersBadRequest creates a DeleteCustomersBadRequest with default headers values
func NewDeleteCustomersBadRequest() *DeleteCustomersBadRequest {
	return &DeleteCustomersBadRequest{}
}

/*DeleteCustomersBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type DeleteCustomersBadRequest struct {
}

func (o *DeleteCustomersBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/customers/{id}/][%d] deleteCustomersBadRequest ", 400)
}

func (o *DeleteCustomersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomersUnauthorized creates a DeleteCustomersUnauthorized with default headers values
func NewDeleteCustomersUnauthorized() *DeleteCustomersUnauthorized {
	return &DeleteCustomersUnauthorized{}
}

/*DeleteCustomersUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type DeleteCustomersUnauthorized struct {
}

func (o *DeleteCustomersUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/customers/{id}/][%d] deleteCustomersUnauthorized ", 401)
}

func (o *DeleteCustomersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomersForbidden creates a DeleteCustomersForbidden with default headers values
func NewDeleteCustomersForbidden() *DeleteCustomersForbidden {
	return &DeleteCustomersForbidden{}
}

/*DeleteCustomersForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type DeleteCustomersForbidden struct {
}

func (o *DeleteCustomersForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/customers/{id}/][%d] deleteCustomersForbidden ", 403)
}

func (o *DeleteCustomersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomersNotFound creates a DeleteCustomersNotFound with default headers values
func NewDeleteCustomersNotFound() *DeleteCustomersNotFound {
	return &DeleteCustomersNotFound{}
}

/*DeleteCustomersNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type DeleteCustomersNotFound struct {
}

func (o *DeleteCustomersNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/customers/{id}/][%d] deleteCustomersNotFound ", 404)
}

func (o *DeleteCustomersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomersMethodNotAllowed creates a DeleteCustomersMethodNotAllowed with default headers values
func NewDeleteCustomersMethodNotAllowed() *DeleteCustomersMethodNotAllowed {
	return &DeleteCustomersMethodNotAllowed{}
}

/*DeleteCustomersMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type DeleteCustomersMethodNotAllowed struct {
}

func (o *DeleteCustomersMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/customers/{id}/][%d] deleteCustomersMethodNotAllowed ", 405)
}

func (o *DeleteCustomersMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomersGone creates a DeleteCustomersGone with default headers values
func NewDeleteCustomersGone() *DeleteCustomersGone {
	return &DeleteCustomersGone{}
}

/*DeleteCustomersGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type DeleteCustomersGone struct {
}

func (o *DeleteCustomersGone) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/customers/{id}/][%d] deleteCustomersGone ", 410)
}

func (o *DeleteCustomersGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomersInternalServerError creates a DeleteCustomersInternalServerError with default headers values
func NewDeleteCustomersInternalServerError() *DeleteCustomersInternalServerError {
	return &DeleteCustomersInternalServerError{}
}

/*DeleteCustomersInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type DeleteCustomersInternalServerError struct {
}

func (o *DeleteCustomersInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/customers/{id}/][%d] deleteCustomersInternalServerError ", 500)
}

func (o *DeleteCustomersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomersServiceUnavailable creates a DeleteCustomersServiceUnavailable with default headers values
func NewDeleteCustomersServiceUnavailable() *DeleteCustomersServiceUnavailable {
	return &DeleteCustomersServiceUnavailable{}
}

/*DeleteCustomersServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type DeleteCustomersServiceUnavailable struct {
}

func (o *DeleteCustomersServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/customers/{id}/][%d] deleteCustomersServiceUnavailable ", 503)
}

func (o *DeleteCustomersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeleteCustomersOKBody delete customers o k body
swagger:model DeleteCustomersOKBody
*/
type DeleteCustomersOKBody struct {

	// deleted
	Deleted interface{} `json:"deleted,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`
}

// Validate validates this delete customers o k body
func (o *DeleteCustomersOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCustomersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCustomersOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteCustomersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
