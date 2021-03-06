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

// PostAutoDiscoveryVserverReader is a Reader for the PostAutoDiscoveryVserver structure.
type PostAutoDiscoveryVserverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAutoDiscoveryVserverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAutoDiscoveryVserverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAutoDiscoveryVserverBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAutoDiscoveryVserverUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAutoDiscoveryVserverForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAutoDiscoveryVserverNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostAutoDiscoveryVserverMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostAutoDiscoveryVserverGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAutoDiscoveryVserverInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAutoDiscoveryVserverServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAutoDiscoveryVserverOK creates a PostAutoDiscoveryVserverOK with default headers values
func NewPostAutoDiscoveryVserverOK() *PostAutoDiscoveryVserverOK {
	return &PostAutoDiscoveryVserverOK{}
}

/*PostAutoDiscoveryVserverOK handles this case with default header values.

The above command returns results like this:
*/
type PostAutoDiscoveryVserverOK struct {
	Payload *PostAutoDiscoveryVserverOKBody
}

func (o *PostAutoDiscoveryVserverOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/vserver/][%d] postAutoDiscoveryVserverOK  %+v", 200, o.Payload)
}

func (o *PostAutoDiscoveryVserverOK) GetPayload() *PostAutoDiscoveryVserverOKBody {
	return o.Payload
}

func (o *PostAutoDiscoveryVserverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostAutoDiscoveryVserverOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAutoDiscoveryVserverBadRequest creates a PostAutoDiscoveryVserverBadRequest with default headers values
func NewPostAutoDiscoveryVserverBadRequest() *PostAutoDiscoveryVserverBadRequest {
	return &PostAutoDiscoveryVserverBadRequest{}
}

/*PostAutoDiscoveryVserverBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostAutoDiscoveryVserverBadRequest struct {
}

func (o *PostAutoDiscoveryVserverBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/vserver/][%d] postAutoDiscoveryVserverBadRequest ", 400)
}

func (o *PostAutoDiscoveryVserverBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryVserverUnauthorized creates a PostAutoDiscoveryVserverUnauthorized with default headers values
func NewPostAutoDiscoveryVserverUnauthorized() *PostAutoDiscoveryVserverUnauthorized {
	return &PostAutoDiscoveryVserverUnauthorized{}
}

/*PostAutoDiscoveryVserverUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostAutoDiscoveryVserverUnauthorized struct {
}

func (o *PostAutoDiscoveryVserverUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/vserver/][%d] postAutoDiscoveryVserverUnauthorized ", 401)
}

func (o *PostAutoDiscoveryVserverUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryVserverForbidden creates a PostAutoDiscoveryVserverForbidden with default headers values
func NewPostAutoDiscoveryVserverForbidden() *PostAutoDiscoveryVserverForbidden {
	return &PostAutoDiscoveryVserverForbidden{}
}

/*PostAutoDiscoveryVserverForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostAutoDiscoveryVserverForbidden struct {
}

func (o *PostAutoDiscoveryVserverForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/vserver/][%d] postAutoDiscoveryVserverForbidden ", 403)
}

func (o *PostAutoDiscoveryVserverForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryVserverNotFound creates a PostAutoDiscoveryVserverNotFound with default headers values
func NewPostAutoDiscoveryVserverNotFound() *PostAutoDiscoveryVserverNotFound {
	return &PostAutoDiscoveryVserverNotFound{}
}

/*PostAutoDiscoveryVserverNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostAutoDiscoveryVserverNotFound struct {
}

func (o *PostAutoDiscoveryVserverNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/vserver/][%d] postAutoDiscoveryVserverNotFound ", 404)
}

func (o *PostAutoDiscoveryVserverNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryVserverMethodNotAllowed creates a PostAutoDiscoveryVserverMethodNotAllowed with default headers values
func NewPostAutoDiscoveryVserverMethodNotAllowed() *PostAutoDiscoveryVserverMethodNotAllowed {
	return &PostAutoDiscoveryVserverMethodNotAllowed{}
}

/*PostAutoDiscoveryVserverMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostAutoDiscoveryVserverMethodNotAllowed struct {
}

func (o *PostAutoDiscoveryVserverMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/vserver/][%d] postAutoDiscoveryVserverMethodNotAllowed ", 405)
}

func (o *PostAutoDiscoveryVserverMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryVserverGone creates a PostAutoDiscoveryVserverGone with default headers values
func NewPostAutoDiscoveryVserverGone() *PostAutoDiscoveryVserverGone {
	return &PostAutoDiscoveryVserverGone{}
}

/*PostAutoDiscoveryVserverGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostAutoDiscoveryVserverGone struct {
}

func (o *PostAutoDiscoveryVserverGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/vserver/][%d] postAutoDiscoveryVserverGone ", 410)
}

func (o *PostAutoDiscoveryVserverGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryVserverInternalServerError creates a PostAutoDiscoveryVserverInternalServerError with default headers values
func NewPostAutoDiscoveryVserverInternalServerError() *PostAutoDiscoveryVserverInternalServerError {
	return &PostAutoDiscoveryVserverInternalServerError{}
}

/*PostAutoDiscoveryVserverInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type PostAutoDiscoveryVserverInternalServerError struct {
}

func (o *PostAutoDiscoveryVserverInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/vserver/][%d] postAutoDiscoveryVserverInternalServerError ", 500)
}

func (o *PostAutoDiscoveryVserverInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAutoDiscoveryVserverServiceUnavailable creates a PostAutoDiscoveryVserverServiceUnavailable with default headers values
func NewPostAutoDiscoveryVserverServiceUnavailable() *PostAutoDiscoveryVserverServiceUnavailable {
	return &PostAutoDiscoveryVserverServiceUnavailable{}
}

/*PostAutoDiscoveryVserverServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostAutoDiscoveryVserverServiceUnavailable struct {
}

func (o *PostAutoDiscoveryVserverServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/auto_discovery/vserver/][%d] postAutoDiscoveryVserverServiceUnavailable ", 503)
}

func (o *PostAutoDiscoveryVserverServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostAutoDiscoveryVserverOKBody post auto discovery vserver o k body
swagger:model PostAutoDiscoveryVserverOKBody
*/
type PostAutoDiscoveryVserverOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post auto discovery vserver o k body
func (o *PostAutoDiscoveryVserverOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAutoDiscoveryVserverOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAutoDiscoveryVserverOKBody) UnmarshalBinary(b []byte) error {
	var res PostAutoDiscoveryVserverOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
