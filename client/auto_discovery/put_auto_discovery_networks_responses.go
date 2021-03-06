// Code generated by go-swagger; DO NOT EDIT.

package auto_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PutAutoDiscoveryNetworksReader is a Reader for the PutAutoDiscoveryNetworks structure.
type PutAutoDiscoveryNetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAutoDiscoveryNetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAutoDiscoveryNetworksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAutoDiscoveryNetworksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutAutoDiscoveryNetworksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutAutoDiscoveryNetworksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutAutoDiscoveryNetworksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPutAutoDiscoveryNetworksMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPutAutoDiscoveryNetworksGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutAutoDiscoveryNetworksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutAutoDiscoveryNetworksServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutAutoDiscoveryNetworksOK creates a PutAutoDiscoveryNetworksOK with default headers values
func NewPutAutoDiscoveryNetworksOK() *PutAutoDiscoveryNetworksOK {
	return &PutAutoDiscoveryNetworksOK{}
}

/*PutAutoDiscoveryNetworksOK handles this case with default header values.

The above command returns results like this:
*/
type PutAutoDiscoveryNetworksOK struct {
	Payload *PutAutoDiscoveryNetworksOKBody
}

func (o *PutAutoDiscoveryNetworksOK) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/auto_discovery/networks/][%d] putAutoDiscoveryNetworksOK  %+v", 200, o.Payload)
}

func (o *PutAutoDiscoveryNetworksOK) GetPayload() *PutAutoDiscoveryNetworksOKBody {
	return o.Payload
}

func (o *PutAutoDiscoveryNetworksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutAutoDiscoveryNetworksOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAutoDiscoveryNetworksBadRequest creates a PutAutoDiscoveryNetworksBadRequest with default headers values
func NewPutAutoDiscoveryNetworksBadRequest() *PutAutoDiscoveryNetworksBadRequest {
	return &PutAutoDiscoveryNetworksBadRequest{}
}

/*PutAutoDiscoveryNetworksBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PutAutoDiscoveryNetworksBadRequest struct {
}

func (o *PutAutoDiscoveryNetworksBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/auto_discovery/networks/][%d] putAutoDiscoveryNetworksBadRequest ", 400)
}

func (o *PutAutoDiscoveryNetworksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAutoDiscoveryNetworksUnauthorized creates a PutAutoDiscoveryNetworksUnauthorized with default headers values
func NewPutAutoDiscoveryNetworksUnauthorized() *PutAutoDiscoveryNetworksUnauthorized {
	return &PutAutoDiscoveryNetworksUnauthorized{}
}

/*PutAutoDiscoveryNetworksUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PutAutoDiscoveryNetworksUnauthorized struct {
}

func (o *PutAutoDiscoveryNetworksUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/auto_discovery/networks/][%d] putAutoDiscoveryNetworksUnauthorized ", 401)
}

func (o *PutAutoDiscoveryNetworksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAutoDiscoveryNetworksForbidden creates a PutAutoDiscoveryNetworksForbidden with default headers values
func NewPutAutoDiscoveryNetworksForbidden() *PutAutoDiscoveryNetworksForbidden {
	return &PutAutoDiscoveryNetworksForbidden{}
}

/*PutAutoDiscoveryNetworksForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PutAutoDiscoveryNetworksForbidden struct {
}

func (o *PutAutoDiscoveryNetworksForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/auto_discovery/networks/][%d] putAutoDiscoveryNetworksForbidden ", 403)
}

func (o *PutAutoDiscoveryNetworksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAutoDiscoveryNetworksNotFound creates a PutAutoDiscoveryNetworksNotFound with default headers values
func NewPutAutoDiscoveryNetworksNotFound() *PutAutoDiscoveryNetworksNotFound {
	return &PutAutoDiscoveryNetworksNotFound{}
}

/*PutAutoDiscoveryNetworksNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PutAutoDiscoveryNetworksNotFound struct {
}

func (o *PutAutoDiscoveryNetworksNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/auto_discovery/networks/][%d] putAutoDiscoveryNetworksNotFound ", 404)
}

func (o *PutAutoDiscoveryNetworksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAutoDiscoveryNetworksMethodNotAllowed creates a PutAutoDiscoveryNetworksMethodNotAllowed with default headers values
func NewPutAutoDiscoveryNetworksMethodNotAllowed() *PutAutoDiscoveryNetworksMethodNotAllowed {
	return &PutAutoDiscoveryNetworksMethodNotAllowed{}
}

/*PutAutoDiscoveryNetworksMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PutAutoDiscoveryNetworksMethodNotAllowed struct {
}

func (o *PutAutoDiscoveryNetworksMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/auto_discovery/networks/][%d] putAutoDiscoveryNetworksMethodNotAllowed ", 405)
}

func (o *PutAutoDiscoveryNetworksMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAutoDiscoveryNetworksGone creates a PutAutoDiscoveryNetworksGone with default headers values
func NewPutAutoDiscoveryNetworksGone() *PutAutoDiscoveryNetworksGone {
	return &PutAutoDiscoveryNetworksGone{}
}

/*PutAutoDiscoveryNetworksGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PutAutoDiscoveryNetworksGone struct {
}

func (o *PutAutoDiscoveryNetworksGone) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/auto_discovery/networks/][%d] putAutoDiscoveryNetworksGone ", 410)
}

func (o *PutAutoDiscoveryNetworksGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAutoDiscoveryNetworksInternalServerError creates a PutAutoDiscoveryNetworksInternalServerError with default headers values
func NewPutAutoDiscoveryNetworksInternalServerError() *PutAutoDiscoveryNetworksInternalServerError {
	return &PutAutoDiscoveryNetworksInternalServerError{}
}

/*PutAutoDiscoveryNetworksInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type PutAutoDiscoveryNetworksInternalServerError struct {
}

func (o *PutAutoDiscoveryNetworksInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/auto_discovery/networks/][%d] putAutoDiscoveryNetworksInternalServerError ", 500)
}

func (o *PutAutoDiscoveryNetworksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAutoDiscoveryNetworksServiceUnavailable creates a PutAutoDiscoveryNetworksServiceUnavailable with default headers values
func NewPutAutoDiscoveryNetworksServiceUnavailable() *PutAutoDiscoveryNetworksServiceUnavailable {
	return &PutAutoDiscoveryNetworksServiceUnavailable{}
}

/*PutAutoDiscoveryNetworksServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PutAutoDiscoveryNetworksServiceUnavailable struct {
}

func (o *PutAutoDiscoveryNetworksServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/auto_discovery/networks/][%d] putAutoDiscoveryNetworksServiceUnavailable ", 503)
}

func (o *PutAutoDiscoveryNetworksServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PutAutoDiscoveryNetworksOKBody put auto discovery networks o k body
swagger:model PutAutoDiscoveryNetworksOKBody
*/
type PutAutoDiscoveryNetworksOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this put auto discovery networks o k body
func (o *PutAutoDiscoveryNetworksOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutAutoDiscoveryNetworksOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutAutoDiscoveryNetworksOKBody) UnmarshalBinary(b []byte) error {
	var res PutAutoDiscoveryNetworksOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
