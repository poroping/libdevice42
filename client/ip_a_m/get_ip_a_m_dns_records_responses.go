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

// GetIPAMDNSRecordsReader is a Reader for the GetIPAMDNSRecords structure.
type GetIPAMDNSRecordsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPAMDNSRecordsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIPAMDNSRecordsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIPAMDNSRecordsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIPAMDNSRecordsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIPAMDNSRecordsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIPAMDNSRecordsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetIPAMDNSRecordsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetIPAMDNSRecordsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIPAMDNSRecordsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIPAMDNSRecordsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIPAMDNSRecordsOK creates a GetIPAMDNSRecordsOK with default headers values
func NewGetIPAMDNSRecordsOK() *GetIPAMDNSRecordsOK {
	return &GetIPAMDNSRecordsOK{}
}

/*GetIPAMDNSRecordsOK handles this case with default header values.

The above command returns results like this:
*/
type GetIPAMDNSRecordsOK struct {
	Payload *GetIPAMDNSRecordsOKBody
}

func (o *GetIPAMDNSRecordsOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/dns/records/][%d] getIpAMDnsRecordsOK  %+v", 200, o.Payload)
}

func (o *GetIPAMDNSRecordsOK) GetPayload() *GetIPAMDNSRecordsOKBody {
	return o.Payload
}

func (o *GetIPAMDNSRecordsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetIPAMDNSRecordsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPAMDNSRecordsBadRequest creates a GetIPAMDNSRecordsBadRequest with default headers values
func NewGetIPAMDNSRecordsBadRequest() *GetIPAMDNSRecordsBadRequest {
	return &GetIPAMDNSRecordsBadRequest{}
}

/*GetIPAMDNSRecordsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetIPAMDNSRecordsBadRequest struct {
}

func (o *GetIPAMDNSRecordsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/dns/records/][%d] getIpAMDnsRecordsBadRequest ", 400)
}

func (o *GetIPAMDNSRecordsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMDNSRecordsUnauthorized creates a GetIPAMDNSRecordsUnauthorized with default headers values
func NewGetIPAMDNSRecordsUnauthorized() *GetIPAMDNSRecordsUnauthorized {
	return &GetIPAMDNSRecordsUnauthorized{}
}

/*GetIPAMDNSRecordsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetIPAMDNSRecordsUnauthorized struct {
}

func (o *GetIPAMDNSRecordsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/dns/records/][%d] getIpAMDnsRecordsUnauthorized ", 401)
}

func (o *GetIPAMDNSRecordsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMDNSRecordsForbidden creates a GetIPAMDNSRecordsForbidden with default headers values
func NewGetIPAMDNSRecordsForbidden() *GetIPAMDNSRecordsForbidden {
	return &GetIPAMDNSRecordsForbidden{}
}

/*GetIPAMDNSRecordsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetIPAMDNSRecordsForbidden struct {
}

func (o *GetIPAMDNSRecordsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/dns/records/][%d] getIpAMDnsRecordsForbidden ", 403)
}

func (o *GetIPAMDNSRecordsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMDNSRecordsNotFound creates a GetIPAMDNSRecordsNotFound with default headers values
func NewGetIPAMDNSRecordsNotFound() *GetIPAMDNSRecordsNotFound {
	return &GetIPAMDNSRecordsNotFound{}
}

/*GetIPAMDNSRecordsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetIPAMDNSRecordsNotFound struct {
}

func (o *GetIPAMDNSRecordsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/dns/records/][%d] getIpAMDnsRecordsNotFound ", 404)
}

func (o *GetIPAMDNSRecordsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMDNSRecordsMethodNotAllowed creates a GetIPAMDNSRecordsMethodNotAllowed with default headers values
func NewGetIPAMDNSRecordsMethodNotAllowed() *GetIPAMDNSRecordsMethodNotAllowed {
	return &GetIPAMDNSRecordsMethodNotAllowed{}
}

/*GetIPAMDNSRecordsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetIPAMDNSRecordsMethodNotAllowed struct {
}

func (o *GetIPAMDNSRecordsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/dns/records/][%d] getIpAMDnsRecordsMethodNotAllowed ", 405)
}

func (o *GetIPAMDNSRecordsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMDNSRecordsGone creates a GetIPAMDNSRecordsGone with default headers values
func NewGetIPAMDNSRecordsGone() *GetIPAMDNSRecordsGone {
	return &GetIPAMDNSRecordsGone{}
}

/*GetIPAMDNSRecordsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetIPAMDNSRecordsGone struct {
}

func (o *GetIPAMDNSRecordsGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/dns/records/][%d] getIpAMDnsRecordsGone ", 410)
}

func (o *GetIPAMDNSRecordsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMDNSRecordsInternalServerError creates a GetIPAMDNSRecordsInternalServerError with default headers values
func NewGetIPAMDNSRecordsInternalServerError() *GetIPAMDNSRecordsInternalServerError {
	return &GetIPAMDNSRecordsInternalServerError{}
}

/*GetIPAMDNSRecordsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetIPAMDNSRecordsInternalServerError struct {
}

func (o *GetIPAMDNSRecordsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/dns/records/][%d] getIpAMDnsRecordsInternalServerError ", 500)
}

func (o *GetIPAMDNSRecordsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMDNSRecordsServiceUnavailable creates a GetIPAMDNSRecordsServiceUnavailable with default headers values
func NewGetIPAMDNSRecordsServiceUnavailable() *GetIPAMDNSRecordsServiceUnavailable {
	return &GetIPAMDNSRecordsServiceUnavailable{}
}

/*GetIPAMDNSRecordsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetIPAMDNSRecordsServiceUnavailable struct {
}

func (o *GetIPAMDNSRecordsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/dns/records/][%d] getIpAMDnsRecordsServiceUnavailable ", 503)
}

func (o *GetIPAMDNSRecordsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetIPAMDNSRecordsOKBody get IP a m DNS records o k body
swagger:model GetIPAMDNSRecordsOKBody
*/
type GetIPAMDNSRecordsOKBody struct {

	// records
	Records []*models.IPMdnsRecords `json:"records"`

	// total count
	TotalCount interface{} `json:"total_count,omitempty"`
}

// Validate validates this get IP a m DNS records o k body
func (o *GetIPAMDNSRecordsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetIPAMDNSRecordsOKBody) validateRecords(formats strfmt.Registry) error {

	if swag.IsZero(o.Records) { // not required
		return nil
	}

	for i := 0; i < len(o.Records); i++ {
		if swag.IsZero(o.Records[i]) { // not required
			continue
		}

		if o.Records[i] != nil {
			if err := o.Records[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getIpAMDnsRecordsOK" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetIPAMDNSRecordsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIPAMDNSRecordsOKBody) UnmarshalBinary(b []byte) error {
	var res GetIPAMDNSRecordsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
