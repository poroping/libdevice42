// Code generated by go-swagger; DO NOT EDIT.

package parts_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostPartmodelsReader is a Reader for the PostPartmodels structure.
type PostPartmodelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPartmodelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostPartmodelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPartmodelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostPartmodelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostPartmodelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostPartmodelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostPartmodelsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostPartmodelsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostPartmodelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostPartmodelsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPartmodelsOK creates a PostPartmodelsOK with default headers values
func NewPostPartmodelsOK() *PostPartmodelsOK {
	return &PostPartmodelsOK{}
}

/*PostPartmodelsOK handles this case with default header values.

The above command returns results like this:
*/
type PostPartmodelsOK struct {
	Payload *PostPartmodelsOKBody
}

func (o *PostPartmodelsOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/partmodels/][%d] postPartmodelsOK  %+v", 200, o.Payload)
}

func (o *PostPartmodelsOK) GetPayload() *PostPartmodelsOKBody {
	return o.Payload
}

func (o *PostPartmodelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostPartmodelsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPartmodelsBadRequest creates a PostPartmodelsBadRequest with default headers values
func NewPostPartmodelsBadRequest() *PostPartmodelsBadRequest {
	return &PostPartmodelsBadRequest{}
}

/*PostPartmodelsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostPartmodelsBadRequest struct {
}

func (o *PostPartmodelsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/partmodels/][%d] postPartmodelsBadRequest ", 400)
}

func (o *PostPartmodelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPartmodelsUnauthorized creates a PostPartmodelsUnauthorized with default headers values
func NewPostPartmodelsUnauthorized() *PostPartmodelsUnauthorized {
	return &PostPartmodelsUnauthorized{}
}

/*PostPartmodelsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostPartmodelsUnauthorized struct {
}

func (o *PostPartmodelsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/partmodels/][%d] postPartmodelsUnauthorized ", 401)
}

func (o *PostPartmodelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPartmodelsForbidden creates a PostPartmodelsForbidden with default headers values
func NewPostPartmodelsForbidden() *PostPartmodelsForbidden {
	return &PostPartmodelsForbidden{}
}

/*PostPartmodelsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostPartmodelsForbidden struct {
}

func (o *PostPartmodelsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/partmodels/][%d] postPartmodelsForbidden ", 403)
}

func (o *PostPartmodelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPartmodelsNotFound creates a PostPartmodelsNotFound with default headers values
func NewPostPartmodelsNotFound() *PostPartmodelsNotFound {
	return &PostPartmodelsNotFound{}
}

/*PostPartmodelsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostPartmodelsNotFound struct {
}

func (o *PostPartmodelsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/partmodels/][%d] postPartmodelsNotFound ", 404)
}

func (o *PostPartmodelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPartmodelsMethodNotAllowed creates a PostPartmodelsMethodNotAllowed with default headers values
func NewPostPartmodelsMethodNotAllowed() *PostPartmodelsMethodNotAllowed {
	return &PostPartmodelsMethodNotAllowed{}
}

/*PostPartmodelsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostPartmodelsMethodNotAllowed struct {
}

func (o *PostPartmodelsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/partmodels/][%d] postPartmodelsMethodNotAllowed ", 405)
}

func (o *PostPartmodelsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPartmodelsGone creates a PostPartmodelsGone with default headers values
func NewPostPartmodelsGone() *PostPartmodelsGone {
	return &PostPartmodelsGone{}
}

/*PostPartmodelsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostPartmodelsGone struct {
}

func (o *PostPartmodelsGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/partmodels/][%d] postPartmodelsGone ", 410)
}

func (o *PostPartmodelsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPartmodelsInternalServerError creates a PostPartmodelsInternalServerError with default headers values
func NewPostPartmodelsInternalServerError() *PostPartmodelsInternalServerError {
	return &PostPartmodelsInternalServerError{}
}

/*PostPartmodelsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type PostPartmodelsInternalServerError struct {
}

func (o *PostPartmodelsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/partmodels/][%d] postPartmodelsInternalServerError ", 500)
}

func (o *PostPartmodelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPartmodelsServiceUnavailable creates a PostPartmodelsServiceUnavailable with default headers values
func NewPostPartmodelsServiceUnavailable() *PostPartmodelsServiceUnavailable {
	return &PostPartmodelsServiceUnavailable{}
}

/*PostPartmodelsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostPartmodelsServiceUnavailable struct {
}

func (o *PostPartmodelsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/partmodels/][%d] postPartmodelsServiceUnavailable ", 503)
}

func (o *PostPartmodelsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostPartmodelsOKBody post partmodels o k body
swagger:model PostPartmodelsOKBody
*/
type PostPartmodelsOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post partmodels o k body
func (o *PostPartmodelsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostPartmodelsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostPartmodelsOKBody) UnmarshalBinary(b []byte) error {
	var res PostPartmodelsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
