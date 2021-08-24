// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/poroping/libdevice42/models"
)

// GetCustomersReader is a Reader for the GetCustomers structure.
type GetCustomersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCustomersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCustomersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCustomersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCustomersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetCustomersMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetCustomersGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCustomersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCustomersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCustomersOK creates a GetCustomersOK with default headers values
func NewGetCustomersOK() *GetCustomersOK {
	return &GetCustomersOK{}
}

/*GetCustomersOK handles this case with default header values.

The above command returns results like this:
*/
type GetCustomersOK struct {
	Payload *GetCustomersOKBody
}

func (o *GetCustomersOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/customers/][%d] getCustomersOK  %+v", 200, o.Payload)
}

func (o *GetCustomersOK) GetPayload() *GetCustomersOKBody {
	return o.Payload
}

func (o *GetCustomersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCustomersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersBadRequest creates a GetCustomersBadRequest with default headers values
func NewGetCustomersBadRequest() *GetCustomersBadRequest {
	return &GetCustomersBadRequest{}
}

/*GetCustomersBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetCustomersBadRequest struct {
}

func (o *GetCustomersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/customers/][%d] getCustomersBadRequest ", 400)
}

func (o *GetCustomersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomersUnauthorized creates a GetCustomersUnauthorized with default headers values
func NewGetCustomersUnauthorized() *GetCustomersUnauthorized {
	return &GetCustomersUnauthorized{}
}

/*GetCustomersUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetCustomersUnauthorized struct {
}

func (o *GetCustomersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/customers/][%d] getCustomersUnauthorized ", 401)
}

func (o *GetCustomersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomersForbidden creates a GetCustomersForbidden with default headers values
func NewGetCustomersForbidden() *GetCustomersForbidden {
	return &GetCustomersForbidden{}
}

/*GetCustomersForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetCustomersForbidden struct {
}

func (o *GetCustomersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/customers/][%d] getCustomersForbidden ", 403)
}

func (o *GetCustomersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomersNotFound creates a GetCustomersNotFound with default headers values
func NewGetCustomersNotFound() *GetCustomersNotFound {
	return &GetCustomersNotFound{}
}

/*GetCustomersNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetCustomersNotFound struct {
}

func (o *GetCustomersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/customers/][%d] getCustomersNotFound ", 404)
}

func (o *GetCustomersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomersMethodNotAllowed creates a GetCustomersMethodNotAllowed with default headers values
func NewGetCustomersMethodNotAllowed() *GetCustomersMethodNotAllowed {
	return &GetCustomersMethodNotAllowed{}
}

/*GetCustomersMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetCustomersMethodNotAllowed struct {
}

func (o *GetCustomersMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/customers/][%d] getCustomersMethodNotAllowed ", 405)
}

func (o *GetCustomersMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomersGone creates a GetCustomersGone with default headers values
func NewGetCustomersGone() *GetCustomersGone {
	return &GetCustomersGone{}
}

/*GetCustomersGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetCustomersGone struct {
}

func (o *GetCustomersGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/customers/][%d] getCustomersGone ", 410)
}

func (o *GetCustomersGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomersInternalServerError creates a GetCustomersInternalServerError with default headers values
func NewGetCustomersInternalServerError() *GetCustomersInternalServerError {
	return &GetCustomersInternalServerError{}
}

/*GetCustomersInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetCustomersInternalServerError struct {
}

func (o *GetCustomersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/customers/][%d] getCustomersInternalServerError ", 500)
}

func (o *GetCustomersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomersServiceUnavailable creates a GetCustomersServiceUnavailable with default headers values
func NewGetCustomersServiceUnavailable() *GetCustomersServiceUnavailable {
	return &GetCustomersServiceUnavailable{}
}

/*GetCustomersServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetCustomersServiceUnavailable struct {
}

func (o *GetCustomersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/customers/][%d] getCustomersServiceUnavailable ", 503)
}

func (o *GetCustomersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetCustomersOKBody get customers o k body
swagger:model GetCustomersOKBody
*/
type GetCustomersOKBody struct {

	// customers
	Customers []*models.Customers `json:"Customers"`
}

// Validate validates this get customers o k body
func (o *GetCustomersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCustomers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCustomersOKBody) validateCustomers(formats strfmt.Registry) error {

	if swag.IsZero(o.Customers) { // not required
		return nil
	}

	for i := 0; i < len(o.Customers); i++ {
		if swag.IsZero(o.Customers[i]) { // not required
			continue
		}

		if o.Customers[i] != nil {
			if err := o.Customers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getCustomersOK" + "." + "Customers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCustomersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCustomersOKBody) UnmarshalBinary(b []byte) error {
	var res GetCustomersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
