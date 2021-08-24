// Code generated by go-swagger; DO NOT EDIT.

package ip_a_m

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

// GetIPAMsubnetsReader is a Reader for the GetIPAMsubnets structure.
type GetIPAMsubnetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPAMsubnetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIPAMsubnetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIPAMsubnetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIPAMsubnetsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIPAMsubnetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIPAMsubnetsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetIPAMsubnetsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetIPAMsubnetsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIPAMsubnetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIPAMsubnetsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIPAMsubnetsOK creates a GetIPAMsubnetsOK with default headers values
func NewGetIPAMsubnetsOK() *GetIPAMsubnetsOK {
	return &GetIPAMsubnetsOK{}
}

/*GetIPAMsubnetsOK handles this case with default header values.

The above command returns results like this:
*/
type GetIPAMsubnetsOK struct {
	Payload *GetIPAMsubnetsOKBody
}

func (o *GetIPAMsubnetsOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/subnets/][%d] getIpAMsubnetsOK  %+v", 200, o.Payload)
}

func (o *GetIPAMsubnetsOK) GetPayload() *GetIPAMsubnetsOKBody {
	return o.Payload
}

func (o *GetIPAMsubnetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetIPAMsubnetsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPAMsubnetsBadRequest creates a GetIPAMsubnetsBadRequest with default headers values
func NewGetIPAMsubnetsBadRequest() *GetIPAMsubnetsBadRequest {
	return &GetIPAMsubnetsBadRequest{}
}

/*GetIPAMsubnetsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetIPAMsubnetsBadRequest struct {
}

func (o *GetIPAMsubnetsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/subnets/][%d] getIpAMsubnetsBadRequest ", 400)
}

func (o *GetIPAMsubnetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMsubnetsUnauthorized creates a GetIPAMsubnetsUnauthorized with default headers values
func NewGetIPAMsubnetsUnauthorized() *GetIPAMsubnetsUnauthorized {
	return &GetIPAMsubnetsUnauthorized{}
}

/*GetIPAMsubnetsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetIPAMsubnetsUnauthorized struct {
}

func (o *GetIPAMsubnetsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/subnets/][%d] getIpAMsubnetsUnauthorized ", 401)
}

func (o *GetIPAMsubnetsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMsubnetsForbidden creates a GetIPAMsubnetsForbidden with default headers values
func NewGetIPAMsubnetsForbidden() *GetIPAMsubnetsForbidden {
	return &GetIPAMsubnetsForbidden{}
}

/*GetIPAMsubnetsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetIPAMsubnetsForbidden struct {
}

func (o *GetIPAMsubnetsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/subnets/][%d] getIpAMsubnetsForbidden ", 403)
}

func (o *GetIPAMsubnetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMsubnetsNotFound creates a GetIPAMsubnetsNotFound with default headers values
func NewGetIPAMsubnetsNotFound() *GetIPAMsubnetsNotFound {
	return &GetIPAMsubnetsNotFound{}
}

/*GetIPAMsubnetsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetIPAMsubnetsNotFound struct {
}

func (o *GetIPAMsubnetsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/subnets/][%d] getIpAMsubnetsNotFound ", 404)
}

func (o *GetIPAMsubnetsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMsubnetsMethodNotAllowed creates a GetIPAMsubnetsMethodNotAllowed with default headers values
func NewGetIPAMsubnetsMethodNotAllowed() *GetIPAMsubnetsMethodNotAllowed {
	return &GetIPAMsubnetsMethodNotAllowed{}
}

/*GetIPAMsubnetsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetIPAMsubnetsMethodNotAllowed struct {
}

func (o *GetIPAMsubnetsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/subnets/][%d] getIpAMsubnetsMethodNotAllowed ", 405)
}

func (o *GetIPAMsubnetsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMsubnetsGone creates a GetIPAMsubnetsGone with default headers values
func NewGetIPAMsubnetsGone() *GetIPAMsubnetsGone {
	return &GetIPAMsubnetsGone{}
}

/*GetIPAMsubnetsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetIPAMsubnetsGone struct {
}

func (o *GetIPAMsubnetsGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/subnets/][%d] getIpAMsubnetsGone ", 410)
}

func (o *GetIPAMsubnetsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMsubnetsInternalServerError creates a GetIPAMsubnetsInternalServerError with default headers values
func NewGetIPAMsubnetsInternalServerError() *GetIPAMsubnetsInternalServerError {
	return &GetIPAMsubnetsInternalServerError{}
}

/*GetIPAMsubnetsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetIPAMsubnetsInternalServerError struct {
}

func (o *GetIPAMsubnetsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/subnets/][%d] getIpAMsubnetsInternalServerError ", 500)
}

func (o *GetIPAMsubnetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMsubnetsServiceUnavailable creates a GetIPAMsubnetsServiceUnavailable with default headers values
func NewGetIPAMsubnetsServiceUnavailable() *GetIPAMsubnetsServiceUnavailable {
	return &GetIPAMsubnetsServiceUnavailable{}
}

/*GetIPAMsubnetsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetIPAMsubnetsServiceUnavailable struct {
}

func (o *GetIPAMsubnetsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/subnets/][%d] getIpAMsubnetsServiceUnavailable ", 503)
}

func (o *GetIPAMsubnetsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetIPAMsubnetsOKBody get IP a msubnets o k body
swagger:model GetIPAMsubnetsOKBody
*/
type GetIPAMsubnetsOKBody struct {

	// limit
	Limit interface{} `json:"limit,omitempty"`

	// offset
	Offset interface{} `json:"offset,omitempty"`

	// subnets
	Subnets []*models.IPAMsubnets `json:"subnets"`

	// total count
	TotalCount interface{} `json:"total_count,omitempty"`
}

// Validate validates this get IP a msubnets o k body
func (o *GetIPAMsubnetsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSubnets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetIPAMsubnetsOKBody) validateSubnets(formats strfmt.Registry) error {

	if swag.IsZero(o.Subnets) { // not required
		return nil
	}

	for i := 0; i < len(o.Subnets); i++ {
		if swag.IsZero(o.Subnets[i]) { // not required
			continue
		}

		if o.Subnets[i] != nil {
			if err := o.Subnets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getIpAMsubnetsOK" + "." + "subnets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetIPAMsubnetsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIPAMsubnetsOKBody) UnmarshalBinary(b []byte) error {
	var res GetIPAMsubnetsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
