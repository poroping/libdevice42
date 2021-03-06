// Code generated by go-swagger; DO NOT EDIT.

package ip_a_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostIPAMDNSZonesReader is a Reader for the PostIPAMDNSZones structure.
type PostIPAMDNSZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIPAMDNSZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostIPAMDNSZonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostIPAMDNSZonesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostIPAMDNSZonesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostIPAMDNSZonesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostIPAMDNSZonesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostIPAMDNSZonesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostIPAMDNSZonesGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostIPAMDNSZonesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostIPAMDNSZonesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostIPAMDNSZonesOK creates a PostIPAMDNSZonesOK with default headers values
func NewPostIPAMDNSZonesOK() *PostIPAMDNSZonesOK {
	return &PostIPAMDNSZonesOK{}
}

/*PostIPAMDNSZonesOK handles this case with default header values.

The above command returns results like this:
*/
type PostIPAMDNSZonesOK struct {
	Payload *PostIPAMDNSZonesOKBody
}

func (o *PostIPAMDNSZonesOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/dns/zones/][%d] postIpAMDnsZonesOK  %+v", 200, o.Payload)
}

func (o *PostIPAMDNSZonesOK) GetPayload() *PostIPAMDNSZonesOKBody {
	return o.Payload
}

func (o *PostIPAMDNSZonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostIPAMDNSZonesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIPAMDNSZonesBadRequest creates a PostIPAMDNSZonesBadRequest with default headers values
func NewPostIPAMDNSZonesBadRequest() *PostIPAMDNSZonesBadRequest {
	return &PostIPAMDNSZonesBadRequest{}
}

/*PostIPAMDNSZonesBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostIPAMDNSZonesBadRequest struct {
}

func (o *PostIPAMDNSZonesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/dns/zones/][%d] postIpAMDnsZonesBadRequest ", 400)
}

func (o *PostIPAMDNSZonesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMDNSZonesUnauthorized creates a PostIPAMDNSZonesUnauthorized with default headers values
func NewPostIPAMDNSZonesUnauthorized() *PostIPAMDNSZonesUnauthorized {
	return &PostIPAMDNSZonesUnauthorized{}
}

/*PostIPAMDNSZonesUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostIPAMDNSZonesUnauthorized struct {
}

func (o *PostIPAMDNSZonesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/dns/zones/][%d] postIpAMDnsZonesUnauthorized ", 401)
}

func (o *PostIPAMDNSZonesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMDNSZonesForbidden creates a PostIPAMDNSZonesForbidden with default headers values
func NewPostIPAMDNSZonesForbidden() *PostIPAMDNSZonesForbidden {
	return &PostIPAMDNSZonesForbidden{}
}

/*PostIPAMDNSZonesForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostIPAMDNSZonesForbidden struct {
}

func (o *PostIPAMDNSZonesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/dns/zones/][%d] postIpAMDnsZonesForbidden ", 403)
}

func (o *PostIPAMDNSZonesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMDNSZonesNotFound creates a PostIPAMDNSZonesNotFound with default headers values
func NewPostIPAMDNSZonesNotFound() *PostIPAMDNSZonesNotFound {
	return &PostIPAMDNSZonesNotFound{}
}

/*PostIPAMDNSZonesNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostIPAMDNSZonesNotFound struct {
}

func (o *PostIPAMDNSZonesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/dns/zones/][%d] postIpAMDnsZonesNotFound ", 404)
}

func (o *PostIPAMDNSZonesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMDNSZonesMethodNotAllowed creates a PostIPAMDNSZonesMethodNotAllowed with default headers values
func NewPostIPAMDNSZonesMethodNotAllowed() *PostIPAMDNSZonesMethodNotAllowed {
	return &PostIPAMDNSZonesMethodNotAllowed{}
}

/*PostIPAMDNSZonesMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostIPAMDNSZonesMethodNotAllowed struct {
}

func (o *PostIPAMDNSZonesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/dns/zones/][%d] postIpAMDnsZonesMethodNotAllowed ", 405)
}

func (o *PostIPAMDNSZonesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMDNSZonesGone creates a PostIPAMDNSZonesGone with default headers values
func NewPostIPAMDNSZonesGone() *PostIPAMDNSZonesGone {
	return &PostIPAMDNSZonesGone{}
}

/*PostIPAMDNSZonesGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostIPAMDNSZonesGone struct {
}

func (o *PostIPAMDNSZonesGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/dns/zones/][%d] postIpAMDnsZonesGone ", 410)
}

func (o *PostIPAMDNSZonesGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMDNSZonesInternalServerError creates a PostIPAMDNSZonesInternalServerError with default headers values
func NewPostIPAMDNSZonesInternalServerError() *PostIPAMDNSZonesInternalServerError {
	return &PostIPAMDNSZonesInternalServerError{}
}

/*PostIPAMDNSZonesInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type PostIPAMDNSZonesInternalServerError struct {
}

func (o *PostIPAMDNSZonesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/dns/zones/][%d] postIpAMDnsZonesInternalServerError ", 500)
}

func (o *PostIPAMDNSZonesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMDNSZonesServiceUnavailable creates a PostIPAMDNSZonesServiceUnavailable with default headers values
func NewPostIPAMDNSZonesServiceUnavailable() *PostIPAMDNSZonesServiceUnavailable {
	return &PostIPAMDNSZonesServiceUnavailable{}
}

/*PostIPAMDNSZonesServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostIPAMDNSZonesServiceUnavailable struct {
}

func (o *PostIPAMDNSZonesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/dns/zones/][%d] postIpAMDnsZonesServiceUnavailable ", 503)
}

func (o *PostIPAMDNSZonesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostIPAMDNSZonesOKBody post IP a m DNS zones o k body
swagger:model PostIPAMDNSZonesOKBody
*/
type PostIPAMDNSZonesOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post IP a m DNS zones o k body
func (o *PostIPAMDNSZonesOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostIPAMDNSZonesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostIPAMDNSZonesOKBody) UnmarshalBinary(b []byte) error {
	var res PostIPAMDNSZonesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
