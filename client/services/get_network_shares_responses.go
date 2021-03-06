// Code generated by go-swagger; DO NOT EDIT.

package services

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

// GetNetworkSharesReader is a Reader for the GetNetworkShares structure.
type GetNetworkSharesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSharesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSharesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNetworkSharesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetNetworkSharesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNetworkSharesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNetworkSharesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetNetworkSharesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetNetworkSharesGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNetworkSharesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetNetworkSharesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkSharesOK creates a GetNetworkSharesOK with default headers values
func NewGetNetworkSharesOK() *GetNetworkSharesOK {
	return &GetNetworkSharesOK{}
}

/*GetNetworkSharesOK handles this case with default header values.

The above command returns results like this:
*/
type GetNetworkSharesOK struct {
	Payload *GetNetworkSharesOKBody
}

func (o *GetNetworkSharesOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/network_shares/][%d] getNetworkSharesOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSharesOK) GetPayload() *GetNetworkSharesOKBody {
	return o.Payload
}

func (o *GetNetworkSharesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNetworkSharesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkSharesBadRequest creates a GetNetworkSharesBadRequest with default headers values
func NewGetNetworkSharesBadRequest() *GetNetworkSharesBadRequest {
	return &GetNetworkSharesBadRequest{}
}

/*GetNetworkSharesBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetNetworkSharesBadRequest struct {
}

func (o *GetNetworkSharesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/network_shares/][%d] getNetworkSharesBadRequest ", 400)
}

func (o *GetNetworkSharesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNetworkSharesUnauthorized creates a GetNetworkSharesUnauthorized with default headers values
func NewGetNetworkSharesUnauthorized() *GetNetworkSharesUnauthorized {
	return &GetNetworkSharesUnauthorized{}
}

/*GetNetworkSharesUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetNetworkSharesUnauthorized struct {
}

func (o *GetNetworkSharesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/network_shares/][%d] getNetworkSharesUnauthorized ", 401)
}

func (o *GetNetworkSharesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNetworkSharesForbidden creates a GetNetworkSharesForbidden with default headers values
func NewGetNetworkSharesForbidden() *GetNetworkSharesForbidden {
	return &GetNetworkSharesForbidden{}
}

/*GetNetworkSharesForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetNetworkSharesForbidden struct {
}

func (o *GetNetworkSharesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/network_shares/][%d] getNetworkSharesForbidden ", 403)
}

func (o *GetNetworkSharesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNetworkSharesNotFound creates a GetNetworkSharesNotFound with default headers values
func NewGetNetworkSharesNotFound() *GetNetworkSharesNotFound {
	return &GetNetworkSharesNotFound{}
}

/*GetNetworkSharesNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetNetworkSharesNotFound struct {
}

func (o *GetNetworkSharesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/network_shares/][%d] getNetworkSharesNotFound ", 404)
}

func (o *GetNetworkSharesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNetworkSharesMethodNotAllowed creates a GetNetworkSharesMethodNotAllowed with default headers values
func NewGetNetworkSharesMethodNotAllowed() *GetNetworkSharesMethodNotAllowed {
	return &GetNetworkSharesMethodNotAllowed{}
}

/*GetNetworkSharesMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetNetworkSharesMethodNotAllowed struct {
}

func (o *GetNetworkSharesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/network_shares/][%d] getNetworkSharesMethodNotAllowed ", 405)
}

func (o *GetNetworkSharesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNetworkSharesGone creates a GetNetworkSharesGone with default headers values
func NewGetNetworkSharesGone() *GetNetworkSharesGone {
	return &GetNetworkSharesGone{}
}

/*GetNetworkSharesGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetNetworkSharesGone struct {
}

func (o *GetNetworkSharesGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/network_shares/][%d] getNetworkSharesGone ", 410)
}

func (o *GetNetworkSharesGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNetworkSharesInternalServerError creates a GetNetworkSharesInternalServerError with default headers values
func NewGetNetworkSharesInternalServerError() *GetNetworkSharesInternalServerError {
	return &GetNetworkSharesInternalServerError{}
}

/*GetNetworkSharesInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type GetNetworkSharesInternalServerError struct {
}

func (o *GetNetworkSharesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/network_shares/][%d] getNetworkSharesInternalServerError ", 500)
}

func (o *GetNetworkSharesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNetworkSharesServiceUnavailable creates a GetNetworkSharesServiceUnavailable with default headers values
func NewGetNetworkSharesServiceUnavailable() *GetNetworkSharesServiceUnavailable {
	return &GetNetworkSharesServiceUnavailable{}
}

/*GetNetworkSharesServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetNetworkSharesServiceUnavailable struct {
}

func (o *GetNetworkSharesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/network_shares/][%d] getNetworkSharesServiceUnavailable ", 503)
}

func (o *GetNetworkSharesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetNetworkSharesOKBody get network shares o k body
swagger:model GetNetworkSharesOKBody
*/
type GetNetworkSharesOKBody struct {

	// networkshare details
	NetworkshareDetails []*models.NetworkshareDetails `json:"networkshare_details"`
}

// Validate validates this get network shares o k body
func (o *GetNetworkSharesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNetworkshareDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSharesOKBody) validateNetworkshareDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.NetworkshareDetails) { // not required
		return nil
	}

	for i := 0; i < len(o.NetworkshareDetails); i++ {
		if swag.IsZero(o.NetworkshareDetails[i]) { // not required
			continue
		}

		if o.NetworkshareDetails[i] != nil {
			if err := o.NetworkshareDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getNetworkSharesOK" + "." + "networkshare_details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSharesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSharesOKBody) UnmarshalBinary(b []byte) error {
	var res GetNetworkSharesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
