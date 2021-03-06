// Code generated by go-swagger; DO NOT EDIT.

package software

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostUpdateSoftwareLicensesReader is a Reader for the PostUpdateSoftwareLicenses structure.
type PostUpdateSoftwareLicensesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUpdateSoftwareLicensesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUpdateSoftwareLicensesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUpdateSoftwareLicensesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUpdateSoftwareLicensesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUpdateSoftwareLicensesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUpdateSoftwareLicensesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostUpdateSoftwareLicensesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostUpdateSoftwareLicensesGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUpdateSoftwareLicensesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUpdateSoftwareLicensesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostUpdateSoftwareLicensesOK creates a PostUpdateSoftwareLicensesOK with default headers values
func NewPostUpdateSoftwareLicensesOK() *PostUpdateSoftwareLicensesOK {
	return &PostUpdateSoftwareLicensesOK{}
}

/*PostUpdateSoftwareLicensesOK handles this case with default header values.

The above command returns results like this:
*/
type PostUpdateSoftwareLicensesOK struct {
	Payload *PostUpdateSoftwareLicensesOKBody
}

func (o *PostUpdateSoftwareLicensesOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/software/license_keys/][%d] postUpdateSoftwareLicensesOK  %+v", 200, o.Payload)
}

func (o *PostUpdateSoftwareLicensesOK) GetPayload() *PostUpdateSoftwareLicensesOKBody {
	return o.Payload
}

func (o *PostUpdateSoftwareLicensesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostUpdateSoftwareLicensesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUpdateSoftwareLicensesBadRequest creates a PostUpdateSoftwareLicensesBadRequest with default headers values
func NewPostUpdateSoftwareLicensesBadRequest() *PostUpdateSoftwareLicensesBadRequest {
	return &PostUpdateSoftwareLicensesBadRequest{}
}

/*PostUpdateSoftwareLicensesBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostUpdateSoftwareLicensesBadRequest struct {
}

func (o *PostUpdateSoftwareLicensesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/software/license_keys/][%d] postUpdateSoftwareLicensesBadRequest ", 400)
}

func (o *PostUpdateSoftwareLicensesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdateSoftwareLicensesUnauthorized creates a PostUpdateSoftwareLicensesUnauthorized with default headers values
func NewPostUpdateSoftwareLicensesUnauthorized() *PostUpdateSoftwareLicensesUnauthorized {
	return &PostUpdateSoftwareLicensesUnauthorized{}
}

/*PostUpdateSoftwareLicensesUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostUpdateSoftwareLicensesUnauthorized struct {
}

func (o *PostUpdateSoftwareLicensesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/software/license_keys/][%d] postUpdateSoftwareLicensesUnauthorized ", 401)
}

func (o *PostUpdateSoftwareLicensesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdateSoftwareLicensesForbidden creates a PostUpdateSoftwareLicensesForbidden with default headers values
func NewPostUpdateSoftwareLicensesForbidden() *PostUpdateSoftwareLicensesForbidden {
	return &PostUpdateSoftwareLicensesForbidden{}
}

/*PostUpdateSoftwareLicensesForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostUpdateSoftwareLicensesForbidden struct {
}

func (o *PostUpdateSoftwareLicensesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/software/license_keys/][%d] postUpdateSoftwareLicensesForbidden ", 403)
}

func (o *PostUpdateSoftwareLicensesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdateSoftwareLicensesNotFound creates a PostUpdateSoftwareLicensesNotFound with default headers values
func NewPostUpdateSoftwareLicensesNotFound() *PostUpdateSoftwareLicensesNotFound {
	return &PostUpdateSoftwareLicensesNotFound{}
}

/*PostUpdateSoftwareLicensesNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostUpdateSoftwareLicensesNotFound struct {
}

func (o *PostUpdateSoftwareLicensesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/software/license_keys/][%d] postUpdateSoftwareLicensesNotFound ", 404)
}

func (o *PostUpdateSoftwareLicensesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdateSoftwareLicensesMethodNotAllowed creates a PostUpdateSoftwareLicensesMethodNotAllowed with default headers values
func NewPostUpdateSoftwareLicensesMethodNotAllowed() *PostUpdateSoftwareLicensesMethodNotAllowed {
	return &PostUpdateSoftwareLicensesMethodNotAllowed{}
}

/*PostUpdateSoftwareLicensesMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostUpdateSoftwareLicensesMethodNotAllowed struct {
}

func (o *PostUpdateSoftwareLicensesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/software/license_keys/][%d] postUpdateSoftwareLicensesMethodNotAllowed ", 405)
}

func (o *PostUpdateSoftwareLicensesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdateSoftwareLicensesGone creates a PostUpdateSoftwareLicensesGone with default headers values
func NewPostUpdateSoftwareLicensesGone() *PostUpdateSoftwareLicensesGone {
	return &PostUpdateSoftwareLicensesGone{}
}

/*PostUpdateSoftwareLicensesGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostUpdateSoftwareLicensesGone struct {
}

func (o *PostUpdateSoftwareLicensesGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/software/license_keys/][%d] postUpdateSoftwareLicensesGone ", 410)
}

func (o *PostUpdateSoftwareLicensesGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdateSoftwareLicensesInternalServerError creates a PostUpdateSoftwareLicensesInternalServerError with default headers values
func NewPostUpdateSoftwareLicensesInternalServerError() *PostUpdateSoftwareLicensesInternalServerError {
	return &PostUpdateSoftwareLicensesInternalServerError{}
}

/*PostUpdateSoftwareLicensesInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type PostUpdateSoftwareLicensesInternalServerError struct {
}

func (o *PostUpdateSoftwareLicensesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/software/license_keys/][%d] postUpdateSoftwareLicensesInternalServerError ", 500)
}

func (o *PostUpdateSoftwareLicensesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdateSoftwareLicensesServiceUnavailable creates a PostUpdateSoftwareLicensesServiceUnavailable with default headers values
func NewPostUpdateSoftwareLicensesServiceUnavailable() *PostUpdateSoftwareLicensesServiceUnavailable {
	return &PostUpdateSoftwareLicensesServiceUnavailable{}
}

/*PostUpdateSoftwareLicensesServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostUpdateSoftwareLicensesServiceUnavailable struct {
}

func (o *PostUpdateSoftwareLicensesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/software/license_keys/][%d] postUpdateSoftwareLicensesServiceUnavailable ", 503)
}

func (o *PostUpdateSoftwareLicensesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostUpdateSoftwareLicensesOKBody post update software licenses o k body
swagger:model PostUpdateSoftwareLicensesOKBody
*/
type PostUpdateSoftwareLicensesOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post update software licenses o k body
func (o *PostUpdateSoftwareLicensesOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostUpdateSoftwareLicensesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostUpdateSoftwareLicensesOKBody) UnmarshalBinary(b []byte) error {
	var res PostUpdateSoftwareLicensesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
