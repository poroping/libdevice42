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

// PostAutoDiscoveryPingsweepReader is a Reader for the PostAutoDiscoveryPingsweep structure.
type PostAutoDiscoveryPingsweepReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAutoDiscoveryPingsweepReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAutoDiscoveryPingsweepOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAutoDiscoveryPingsweepBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAutoDiscoveryPingsweepUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAutoDiscoveryPingsweepForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAutoDiscoveryPingsweepNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostAutoDiscoveryPingsweepMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostAutoDiscoveryPingsweepGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAutoDiscoveryPingsweepInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAutoDiscoveryPingsweepServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAutoDiscoveryPingsweepOK creates a PostAutoDiscoveryPingsweepOK with default headers values
func NewPostAutoDiscoveryPingsweepOK() *PostAutoDiscoveryPingsweepOK {
	return &PostAutoDiscoveryPingsweepOK{}
}

/*PostAutoDiscoveryPingsweepOK handles this case with default header values.

The above command returns results like this:
*/
type PostAutoDiscoveryPingsweepOK struct {
	Payload *PostAutoDiscoveryPingsweepOKBody
}

func (o *PostAutoDiscoveryPingsweepOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/pingsweep/][%d] postAutoDiscoveryPingsweepOK  %+v", 200, o.Payload)
}

func (o *PostAutoDiscoveryPingsweepOK) GetPayload() *PostAutoDiscoveryPingsweepOKBody {
	return o.Payload
}

func (o *PostAutoDiscoveryPingsweepOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostAutoDiscoveryPingsweepOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAutoDiscoveryPingsweepBadRequest creates a PostAutoDiscoveryPingsweepBadRequest with default headers values
func NewPostAutoDiscoveryPingsweepBadRequest() *PostAutoDiscoveryPingsweepBadRequest {
	return &PostAutoDiscoveryPingsweepBadRequest{}
}

/*PostAutoDiscoveryPingsweepBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostAutoDiscoveryPingsweepBadRequest struct {
}

func (o *PostAutoDiscoveryPingsweepBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/pingsweep/][%d] postAutoDiscoveryPingsweepBadRequest ", 400)
}

func (o *PostAutoDiscoveryPingsweepBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryPingsweepUnauthorized creates a PostAutoDiscoveryPingsweepUnauthorized with default headers values
func NewPostAutoDiscoveryPingsweepUnauthorized() *PostAutoDiscoveryPingsweepUnauthorized {
	return &PostAutoDiscoveryPingsweepUnauthorized{}
}

/*PostAutoDiscoveryPingsweepUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostAutoDiscoveryPingsweepUnauthorized struct {
}

func (o *PostAutoDiscoveryPingsweepUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/pingsweep/][%d] postAutoDiscoveryPingsweepUnauthorized ", 401)
}

func (o *PostAutoDiscoveryPingsweepUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryPingsweepForbidden creates a PostAutoDiscoveryPingsweepForbidden with default headers values
func NewPostAutoDiscoveryPingsweepForbidden() *PostAutoDiscoveryPingsweepForbidden {
	return &PostAutoDiscoveryPingsweepForbidden{}
}

/*PostAutoDiscoveryPingsweepForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostAutoDiscoveryPingsweepForbidden struct {
}

func (o *PostAutoDiscoveryPingsweepForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/pingsweep/][%d] postAutoDiscoveryPingsweepForbidden ", 403)
}

func (o *PostAutoDiscoveryPingsweepForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryPingsweepNotFound creates a PostAutoDiscoveryPingsweepNotFound with default headers values
func NewPostAutoDiscoveryPingsweepNotFound() *PostAutoDiscoveryPingsweepNotFound {
	return &PostAutoDiscoveryPingsweepNotFound{}
}

/*PostAutoDiscoveryPingsweepNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostAutoDiscoveryPingsweepNotFound struct {
}

func (o *PostAutoDiscoveryPingsweepNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/pingsweep/][%d] postAutoDiscoveryPingsweepNotFound ", 404)
}

func (o *PostAutoDiscoveryPingsweepNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryPingsweepMethodNotAllowed creates a PostAutoDiscoveryPingsweepMethodNotAllowed with default headers values
func NewPostAutoDiscoveryPingsweepMethodNotAllowed() *PostAutoDiscoveryPingsweepMethodNotAllowed {
	return &PostAutoDiscoveryPingsweepMethodNotAllowed{}
}

/*PostAutoDiscoveryPingsweepMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostAutoDiscoveryPingsweepMethodNotAllowed struct {
}

func (o *PostAutoDiscoveryPingsweepMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/pingsweep/][%d] postAutoDiscoveryPingsweepMethodNotAllowed ", 405)
}

func (o *PostAutoDiscoveryPingsweepMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryPingsweepGone creates a PostAutoDiscoveryPingsweepGone with default headers values
func NewPostAutoDiscoveryPingsweepGone() *PostAutoDiscoveryPingsweepGone {
	return &PostAutoDiscoveryPingsweepGone{}
}

/*PostAutoDiscoveryPingsweepGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostAutoDiscoveryPingsweepGone struct {
}

func (o *PostAutoDiscoveryPingsweepGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/pingsweep/][%d] postAutoDiscoveryPingsweepGone ", 410)
}

func (o *PostAutoDiscoveryPingsweepGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryPingsweepInternalServerError creates a PostAutoDiscoveryPingsweepInternalServerError with default headers values
func NewPostAutoDiscoveryPingsweepInternalServerError() *PostAutoDiscoveryPingsweepInternalServerError {
	return &PostAutoDiscoveryPingsweepInternalServerError{}
}

/*PostAutoDiscoveryPingsweepInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type PostAutoDiscoveryPingsweepInternalServerError struct {
}

func (o *PostAutoDiscoveryPingsweepInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/pingsweep/][%d] postAutoDiscoveryPingsweepInternalServerError ", 500)
}

func (o *PostAutoDiscoveryPingsweepInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryPingsweepServiceUnavailable creates a PostAutoDiscoveryPingsweepServiceUnavailable with default headers values
func NewPostAutoDiscoveryPingsweepServiceUnavailable() *PostAutoDiscoveryPingsweepServiceUnavailable {
	return &PostAutoDiscoveryPingsweepServiceUnavailable{}
}

/*PostAutoDiscoveryPingsweepServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostAutoDiscoveryPingsweepServiceUnavailable struct {
}

func (o *PostAutoDiscoveryPingsweepServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/pingsweep/][%d] postAutoDiscoveryPingsweepServiceUnavailable ", 503)
}

func (o *PostAutoDiscoveryPingsweepServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostAutoDiscoveryPingsweepOKBody post auto discovery pingsweep o k body
swagger:model PostAutoDiscoveryPingsweepOKBody
*/
type PostAutoDiscoveryPingsweepOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post auto discovery pingsweep o k body
func (o *PostAutoDiscoveryPingsweepOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAutoDiscoveryPingsweepOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAutoDiscoveryPingsweepOKBody) UnmarshalBinary(b []byte) error {
	var res PostAutoDiscoveryPingsweepOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
