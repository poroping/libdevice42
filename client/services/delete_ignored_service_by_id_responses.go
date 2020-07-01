// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteIgnoredServiceByIDReader is a Reader for the DeleteIgnoredServiceByID structure.
type DeleteIgnoredServiceByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIgnoredServiceByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIgnoredServiceByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteIgnoredServiceByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteIgnoredServiceByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteIgnoredServiceByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteIgnoredServiceByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeleteIgnoredServiceByIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewDeleteIgnoredServiceByIDGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteIgnoredServiceByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteIgnoredServiceByIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteIgnoredServiceByIDOK creates a DeleteIgnoredServiceByIDOK with default headers values
func NewDeleteIgnoredServiceByIDOK() *DeleteIgnoredServiceByIDOK {
	return &DeleteIgnoredServiceByIDOK{}
}

/*DeleteIgnoredServiceByIDOK handles this case with default header values.

The above command returns results like this:
*/
type DeleteIgnoredServiceByIDOK struct {
	Payload *DeleteIgnoredServiceByIDOKBody
}

func (o *DeleteIgnoredServiceByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.0/ignored_service/{id}/][%d] deleteIgnoredServiceByIdOK  %+v", 200, o.Payload)
}

func (o *DeleteIgnoredServiceByIDOK) GetPayload() *DeleteIgnoredServiceByIDOKBody {
	return o.Payload
}

func (o *DeleteIgnoredServiceByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteIgnoredServiceByIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIgnoredServiceByIDBadRequest creates a DeleteIgnoredServiceByIDBadRequest with default headers values
func NewDeleteIgnoredServiceByIDBadRequest() *DeleteIgnoredServiceByIDBadRequest {
	return &DeleteIgnoredServiceByIDBadRequest{}
}

/*DeleteIgnoredServiceByIDBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type DeleteIgnoredServiceByIDBadRequest struct {
}

func (o *DeleteIgnoredServiceByIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/2.0/ignored_service/{id}/][%d] deleteIgnoredServiceByIdBadRequest ", 400)
}

func (o *DeleteIgnoredServiceByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIgnoredServiceByIDUnauthorized creates a DeleteIgnoredServiceByIDUnauthorized with default headers values
func NewDeleteIgnoredServiceByIDUnauthorized() *DeleteIgnoredServiceByIDUnauthorized {
	return &DeleteIgnoredServiceByIDUnauthorized{}
}

/*DeleteIgnoredServiceByIDUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type DeleteIgnoredServiceByIDUnauthorized struct {
}

func (o *DeleteIgnoredServiceByIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/2.0/ignored_service/{id}/][%d] deleteIgnoredServiceByIdUnauthorized ", 401)
}

func (o *DeleteIgnoredServiceByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIgnoredServiceByIDForbidden creates a DeleteIgnoredServiceByIDForbidden with default headers values
func NewDeleteIgnoredServiceByIDForbidden() *DeleteIgnoredServiceByIDForbidden {
	return &DeleteIgnoredServiceByIDForbidden{}
}

/*DeleteIgnoredServiceByIDForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type DeleteIgnoredServiceByIDForbidden struct {
}

func (o *DeleteIgnoredServiceByIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/2.0/ignored_service/{id}/][%d] deleteIgnoredServiceByIdForbidden ", 403)
}

func (o *DeleteIgnoredServiceByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIgnoredServiceByIDNotFound creates a DeleteIgnoredServiceByIDNotFound with default headers values
func NewDeleteIgnoredServiceByIDNotFound() *DeleteIgnoredServiceByIDNotFound {
	return &DeleteIgnoredServiceByIDNotFound{}
}

/*DeleteIgnoredServiceByIDNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type DeleteIgnoredServiceByIDNotFound struct {
}

func (o *DeleteIgnoredServiceByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/2.0/ignored_service/{id}/][%d] deleteIgnoredServiceByIdNotFound ", 404)
}

func (o *DeleteIgnoredServiceByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIgnoredServiceByIDMethodNotAllowed creates a DeleteIgnoredServiceByIDMethodNotAllowed with default headers values
func NewDeleteIgnoredServiceByIDMethodNotAllowed() *DeleteIgnoredServiceByIDMethodNotAllowed {
	return &DeleteIgnoredServiceByIDMethodNotAllowed{}
}

/*DeleteIgnoredServiceByIDMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type DeleteIgnoredServiceByIDMethodNotAllowed struct {
}

func (o *DeleteIgnoredServiceByIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /api/2.0/ignored_service/{id}/][%d] deleteIgnoredServiceByIdMethodNotAllowed ", 405)
}

func (o *DeleteIgnoredServiceByIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIgnoredServiceByIDGone creates a DeleteIgnoredServiceByIDGone with default headers values
func NewDeleteIgnoredServiceByIDGone() *DeleteIgnoredServiceByIDGone {
	return &DeleteIgnoredServiceByIDGone{}
}

/*DeleteIgnoredServiceByIDGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type DeleteIgnoredServiceByIDGone struct {
}

func (o *DeleteIgnoredServiceByIDGone) Error() string {
	return fmt.Sprintf("[DELETE /api/2.0/ignored_service/{id}/][%d] deleteIgnoredServiceByIdGone ", 410)
}

func (o *DeleteIgnoredServiceByIDGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIgnoredServiceByIDInternalServerError creates a DeleteIgnoredServiceByIDInternalServerError with default headers values
func NewDeleteIgnoredServiceByIDInternalServerError() *DeleteIgnoredServiceByIDInternalServerError {
	return &DeleteIgnoredServiceByIDInternalServerError{}
}

/*DeleteIgnoredServiceByIDInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type DeleteIgnoredServiceByIDInternalServerError struct {
}

func (o *DeleteIgnoredServiceByIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/2.0/ignored_service/{id}/][%d] deleteIgnoredServiceByIdInternalServerError ", 500)
}

func (o *DeleteIgnoredServiceByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIgnoredServiceByIDServiceUnavailable creates a DeleteIgnoredServiceByIDServiceUnavailable with default headers values
func NewDeleteIgnoredServiceByIDServiceUnavailable() *DeleteIgnoredServiceByIDServiceUnavailable {
	return &DeleteIgnoredServiceByIDServiceUnavailable{}
}

/*DeleteIgnoredServiceByIDServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type DeleteIgnoredServiceByIDServiceUnavailable struct {
}

func (o *DeleteIgnoredServiceByIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/2.0/ignored_service/{id}/][%d] deleteIgnoredServiceByIdServiceUnavailable ", 503)
}

func (o *DeleteIgnoredServiceByIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeleteIgnoredServiceByIDOKBody delete ignored service by ID o k body
swagger:model DeleteIgnoredServiceByIDOKBody
*/
type DeleteIgnoredServiceByIDOKBody struct {

	// deleted
	Deleted interface{} `json:"deleted,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`
}

// Validate validates this delete ignored service by ID o k body
func (o *DeleteIgnoredServiceByIDOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteIgnoredServiceByIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteIgnoredServiceByIDOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteIgnoredServiceByIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}