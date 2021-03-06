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

// PostAutoDiscoveryIpmiReader is a Reader for the PostAutoDiscoveryIpmi structure.
type PostAutoDiscoveryIpmiReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAutoDiscoveryIpmiReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAutoDiscoveryIpmiOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAutoDiscoveryIpmiBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAutoDiscoveryIpmiUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAutoDiscoveryIpmiForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAutoDiscoveryIpmiNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostAutoDiscoveryIpmiMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostAutoDiscoveryIpmiGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAutoDiscoveryIpmiInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAutoDiscoveryIpmiServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAutoDiscoveryIpmiOK creates a PostAutoDiscoveryIpmiOK with default headers values
func NewPostAutoDiscoveryIpmiOK() *PostAutoDiscoveryIpmiOK {
	return &PostAutoDiscoveryIpmiOK{}
}

/*PostAutoDiscoveryIpmiOK handles this case with default header values.

The above command returns results like this:
*/
type PostAutoDiscoveryIpmiOK struct {
	Payload *PostAutoDiscoveryIpmiOKBody
}

func (o *PostAutoDiscoveryIpmiOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/ipmi/][%d] postAutoDiscoveryIpmiOK  %+v", 200, o.Payload)
}

func (o *PostAutoDiscoveryIpmiOK) GetPayload() *PostAutoDiscoveryIpmiOKBody {
	return o.Payload
}

func (o *PostAutoDiscoveryIpmiOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostAutoDiscoveryIpmiOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAutoDiscoveryIpmiBadRequest creates a PostAutoDiscoveryIpmiBadRequest with default headers values
func NewPostAutoDiscoveryIpmiBadRequest() *PostAutoDiscoveryIpmiBadRequest {
	return &PostAutoDiscoveryIpmiBadRequest{}
}

/*PostAutoDiscoveryIpmiBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostAutoDiscoveryIpmiBadRequest struct {
}

func (o *PostAutoDiscoveryIpmiBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/ipmi/][%d] postAutoDiscoveryIpmiBadRequest ", 400)
}

func (o *PostAutoDiscoveryIpmiBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryIpmiUnauthorized creates a PostAutoDiscoveryIpmiUnauthorized with default headers values
func NewPostAutoDiscoveryIpmiUnauthorized() *PostAutoDiscoveryIpmiUnauthorized {
	return &PostAutoDiscoveryIpmiUnauthorized{}
}

/*PostAutoDiscoveryIpmiUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostAutoDiscoveryIpmiUnauthorized struct {
}

func (o *PostAutoDiscoveryIpmiUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/ipmi/][%d] postAutoDiscoveryIpmiUnauthorized ", 401)
}

func (o *PostAutoDiscoveryIpmiUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryIpmiForbidden creates a PostAutoDiscoveryIpmiForbidden with default headers values
func NewPostAutoDiscoveryIpmiForbidden() *PostAutoDiscoveryIpmiForbidden {
	return &PostAutoDiscoveryIpmiForbidden{}
}

/*PostAutoDiscoveryIpmiForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostAutoDiscoveryIpmiForbidden struct {
}

func (o *PostAutoDiscoveryIpmiForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/ipmi/][%d] postAutoDiscoveryIpmiForbidden ", 403)
}

func (o *PostAutoDiscoveryIpmiForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryIpmiNotFound creates a PostAutoDiscoveryIpmiNotFound with default headers values
func NewPostAutoDiscoveryIpmiNotFound() *PostAutoDiscoveryIpmiNotFound {
	return &PostAutoDiscoveryIpmiNotFound{}
}

/*PostAutoDiscoveryIpmiNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostAutoDiscoveryIpmiNotFound struct {
}

func (o *PostAutoDiscoveryIpmiNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/ipmi/][%d] postAutoDiscoveryIpmiNotFound ", 404)
}

func (o *PostAutoDiscoveryIpmiNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryIpmiMethodNotAllowed creates a PostAutoDiscoveryIpmiMethodNotAllowed with default headers values
func NewPostAutoDiscoveryIpmiMethodNotAllowed() *PostAutoDiscoveryIpmiMethodNotAllowed {
	return &PostAutoDiscoveryIpmiMethodNotAllowed{}
}

/*PostAutoDiscoveryIpmiMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostAutoDiscoveryIpmiMethodNotAllowed struct {
}

func (o *PostAutoDiscoveryIpmiMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/ipmi/][%d] postAutoDiscoveryIpmiMethodNotAllowed ", 405)
}

func (o *PostAutoDiscoveryIpmiMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryIpmiGone creates a PostAutoDiscoveryIpmiGone with default headers values
func NewPostAutoDiscoveryIpmiGone() *PostAutoDiscoveryIpmiGone {
	return &PostAutoDiscoveryIpmiGone{}
}

/*PostAutoDiscoveryIpmiGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostAutoDiscoveryIpmiGone struct {
}

func (o *PostAutoDiscoveryIpmiGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/ipmi/][%d] postAutoDiscoveryIpmiGone ", 410)
}

func (o *PostAutoDiscoveryIpmiGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryIpmiInternalServerError creates a PostAutoDiscoveryIpmiInternalServerError with default headers values
func NewPostAutoDiscoveryIpmiInternalServerError() *PostAutoDiscoveryIpmiInternalServerError {
	return &PostAutoDiscoveryIpmiInternalServerError{}
}

/*PostAutoDiscoveryIpmiInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type PostAutoDiscoveryIpmiInternalServerError struct {
}

func (o *PostAutoDiscoveryIpmiInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/ipmi/][%d] postAutoDiscoveryIpmiInternalServerError ", 500)
}

func (o *PostAutoDiscoveryIpmiInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryIpmiServiceUnavailable creates a PostAutoDiscoveryIpmiServiceUnavailable with default headers values
func NewPostAutoDiscoveryIpmiServiceUnavailable() *PostAutoDiscoveryIpmiServiceUnavailable {
	return &PostAutoDiscoveryIpmiServiceUnavailable{}
}

/*PostAutoDiscoveryIpmiServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostAutoDiscoveryIpmiServiceUnavailable struct {
}

func (o *PostAutoDiscoveryIpmiServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/ipmi/][%d] postAutoDiscoveryIpmiServiceUnavailable ", 503)
}

func (o *PostAutoDiscoveryIpmiServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostAutoDiscoveryIpmiOKBody post auto discovery ipmi o k body
swagger:model PostAutoDiscoveryIpmiOKBody
*/
type PostAutoDiscoveryIpmiOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post auto discovery ipmi o k body
func (o *PostAutoDiscoveryIpmiOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAutoDiscoveryIpmiOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAutoDiscoveryIpmiOKBody) UnmarshalBinary(b []byte) error {
	var res PostAutoDiscoveryIpmiOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
