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

// PostIPAMDNSRecordsReader is a Reader for the PostIPAMDNSRecords structure.
type PostIPAMDNSRecordsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIPAMDNSRecordsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostIPAMDNSRecordsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostIPAMDNSRecordsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostIPAMDNSRecordsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostIPAMDNSRecordsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostIPAMDNSRecordsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostIPAMDNSRecordsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostIPAMDNSRecordsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostIPAMDNSRecordsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostIPAMDNSRecordsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostIPAMDNSRecordsOK creates a PostIPAMDNSRecordsOK with default headers values
func NewPostIPAMDNSRecordsOK() *PostIPAMDNSRecordsOK {
	return &PostIPAMDNSRecordsOK{}
}

/*PostIPAMDNSRecordsOK handles this case with default header values.

The above command returns results like this:
*/
type PostIPAMDNSRecordsOK struct {
	Payload *PostIPAMDNSRecordsOKBody
}

func (o *PostIPAMDNSRecordsOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/dns/records/][%d] postIpAMDnsRecordsOK  %+v", 200, o.Payload)
}

func (o *PostIPAMDNSRecordsOK) GetPayload() *PostIPAMDNSRecordsOKBody {
	return o.Payload
}

func (o *PostIPAMDNSRecordsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostIPAMDNSRecordsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIPAMDNSRecordsBadRequest creates a PostIPAMDNSRecordsBadRequest with default headers values
func NewPostIPAMDNSRecordsBadRequest() *PostIPAMDNSRecordsBadRequest {
	return &PostIPAMDNSRecordsBadRequest{}
}

/*PostIPAMDNSRecordsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostIPAMDNSRecordsBadRequest struct {
}

func (o *PostIPAMDNSRecordsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/dns/records/][%d] postIpAMDnsRecordsBadRequest ", 400)
}

func (o *PostIPAMDNSRecordsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMDNSRecordsUnauthorized creates a PostIPAMDNSRecordsUnauthorized with default headers values
func NewPostIPAMDNSRecordsUnauthorized() *PostIPAMDNSRecordsUnauthorized {
	return &PostIPAMDNSRecordsUnauthorized{}
}

/*PostIPAMDNSRecordsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostIPAMDNSRecordsUnauthorized struct {
}

func (o *PostIPAMDNSRecordsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/dns/records/][%d] postIpAMDnsRecordsUnauthorized ", 401)
}

func (o *PostIPAMDNSRecordsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMDNSRecordsForbidden creates a PostIPAMDNSRecordsForbidden with default headers values
func NewPostIPAMDNSRecordsForbidden() *PostIPAMDNSRecordsForbidden {
	return &PostIPAMDNSRecordsForbidden{}
}

/*PostIPAMDNSRecordsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostIPAMDNSRecordsForbidden struct {
}

func (o *PostIPAMDNSRecordsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/dns/records/][%d] postIpAMDnsRecordsForbidden ", 403)
}

func (o *PostIPAMDNSRecordsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMDNSRecordsNotFound creates a PostIPAMDNSRecordsNotFound with default headers values
func NewPostIPAMDNSRecordsNotFound() *PostIPAMDNSRecordsNotFound {
	return &PostIPAMDNSRecordsNotFound{}
}

/*PostIPAMDNSRecordsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostIPAMDNSRecordsNotFound struct {
}

func (o *PostIPAMDNSRecordsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/dns/records/][%d] postIpAMDnsRecordsNotFound ", 404)
}

func (o *PostIPAMDNSRecordsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMDNSRecordsMethodNotAllowed creates a PostIPAMDNSRecordsMethodNotAllowed with default headers values
func NewPostIPAMDNSRecordsMethodNotAllowed() *PostIPAMDNSRecordsMethodNotAllowed {
	return &PostIPAMDNSRecordsMethodNotAllowed{}
}

/*PostIPAMDNSRecordsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostIPAMDNSRecordsMethodNotAllowed struct {
}

func (o *PostIPAMDNSRecordsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/dns/records/][%d] postIpAMDnsRecordsMethodNotAllowed ", 405)
}

func (o *PostIPAMDNSRecordsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMDNSRecordsGone creates a PostIPAMDNSRecordsGone with default headers values
func NewPostIPAMDNSRecordsGone() *PostIPAMDNSRecordsGone {
	return &PostIPAMDNSRecordsGone{}
}

/*PostIPAMDNSRecordsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostIPAMDNSRecordsGone struct {
}

func (o *PostIPAMDNSRecordsGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/dns/records/][%d] postIpAMDnsRecordsGone ", 410)
}

func (o *PostIPAMDNSRecordsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMDNSRecordsInternalServerError creates a PostIPAMDNSRecordsInternalServerError with default headers values
func NewPostIPAMDNSRecordsInternalServerError() *PostIPAMDNSRecordsInternalServerError {
	return &PostIPAMDNSRecordsInternalServerError{}
}

/*PostIPAMDNSRecordsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type PostIPAMDNSRecordsInternalServerError struct {
}

func (o *PostIPAMDNSRecordsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/dns/records/][%d] postIpAMDnsRecordsInternalServerError ", 500)
}

func (o *PostIPAMDNSRecordsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMDNSRecordsServiceUnavailable creates a PostIPAMDNSRecordsServiceUnavailable with default headers values
func NewPostIPAMDNSRecordsServiceUnavailable() *PostIPAMDNSRecordsServiceUnavailable {
	return &PostIPAMDNSRecordsServiceUnavailable{}
}

/*PostIPAMDNSRecordsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostIPAMDNSRecordsServiceUnavailable struct {
}

func (o *PostIPAMDNSRecordsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/dns/records/][%d] postIpAMDnsRecordsServiceUnavailable ", 503)
}

func (o *PostIPAMDNSRecordsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostIPAMDNSRecordsOKBody post IP a m DNS records o k body
swagger:model PostIPAMDNSRecordsOKBody
*/
type PostIPAMDNSRecordsOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post IP a m DNS records o k body
func (o *PostIPAMDNSRecordsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostIPAMDNSRecordsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostIPAMDNSRecordsOKBody) UnmarshalBinary(b []byte) error {
	var res PostIPAMDNSRecordsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
