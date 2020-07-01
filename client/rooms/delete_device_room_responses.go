// Code generated by go-swagger; DO NOT EDIT.

package rooms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteDeviceRoomReader is a Reader for the DeleteDeviceRoom structure.
type DeleteDeviceRoomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDeviceRoomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDeviceRoomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDeviceRoomBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteDeviceRoomUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDeviceRoomForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDeviceRoomNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeleteDeviceRoomMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewDeleteDeviceRoomGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDeviceRoomInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteDeviceRoomServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDeviceRoomOK creates a DeleteDeviceRoomOK with default headers values
func NewDeleteDeviceRoomOK() *DeleteDeviceRoomOK {
	return &DeleteDeviceRoomOK{}
}

/*DeleteDeviceRoomOK handles this case with default header values.

The above command returns results like this:
*/
type DeleteDeviceRoomOK struct {
	Payload *DeleteDeviceRoomOKBody
}

func (o *DeleteDeviceRoomOK) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device/room/{ID}/][%d] deleteDeviceRoomOK  %+v", 200, o.Payload)
}

func (o *DeleteDeviceRoomOK) GetPayload() *DeleteDeviceRoomOKBody {
	return o.Payload
}

func (o *DeleteDeviceRoomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteDeviceRoomOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceRoomBadRequest creates a DeleteDeviceRoomBadRequest with default headers values
func NewDeleteDeviceRoomBadRequest() *DeleteDeviceRoomBadRequest {
	return &DeleteDeviceRoomBadRequest{}
}

/*DeleteDeviceRoomBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type DeleteDeviceRoomBadRequest struct {
}

func (o *DeleteDeviceRoomBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device/room/{ID}/][%d] deleteDeviceRoomBadRequest ", 400)
}

func (o *DeleteDeviceRoomBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceRoomUnauthorized creates a DeleteDeviceRoomUnauthorized with default headers values
func NewDeleteDeviceRoomUnauthorized() *DeleteDeviceRoomUnauthorized {
	return &DeleteDeviceRoomUnauthorized{}
}

/*DeleteDeviceRoomUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type DeleteDeviceRoomUnauthorized struct {
}

func (o *DeleteDeviceRoomUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device/room/{ID}/][%d] deleteDeviceRoomUnauthorized ", 401)
}

func (o *DeleteDeviceRoomUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceRoomForbidden creates a DeleteDeviceRoomForbidden with default headers values
func NewDeleteDeviceRoomForbidden() *DeleteDeviceRoomForbidden {
	return &DeleteDeviceRoomForbidden{}
}

/*DeleteDeviceRoomForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type DeleteDeviceRoomForbidden struct {
}

func (o *DeleteDeviceRoomForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device/room/{ID}/][%d] deleteDeviceRoomForbidden ", 403)
}

func (o *DeleteDeviceRoomForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceRoomNotFound creates a DeleteDeviceRoomNotFound with default headers values
func NewDeleteDeviceRoomNotFound() *DeleteDeviceRoomNotFound {
	return &DeleteDeviceRoomNotFound{}
}

/*DeleteDeviceRoomNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type DeleteDeviceRoomNotFound struct {
}

func (o *DeleteDeviceRoomNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device/room/{ID}/][%d] deleteDeviceRoomNotFound ", 404)
}

func (o *DeleteDeviceRoomNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceRoomMethodNotAllowed creates a DeleteDeviceRoomMethodNotAllowed with default headers values
func NewDeleteDeviceRoomMethodNotAllowed() *DeleteDeviceRoomMethodNotAllowed {
	return &DeleteDeviceRoomMethodNotAllowed{}
}

/*DeleteDeviceRoomMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type DeleteDeviceRoomMethodNotAllowed struct {
}

func (o *DeleteDeviceRoomMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device/room/{ID}/][%d] deleteDeviceRoomMethodNotAllowed ", 405)
}

func (o *DeleteDeviceRoomMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceRoomGone creates a DeleteDeviceRoomGone with default headers values
func NewDeleteDeviceRoomGone() *DeleteDeviceRoomGone {
	return &DeleteDeviceRoomGone{}
}

/*DeleteDeviceRoomGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type DeleteDeviceRoomGone struct {
}

func (o *DeleteDeviceRoomGone) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device/room/{ID}/][%d] deleteDeviceRoomGone ", 410)
}

func (o *DeleteDeviceRoomGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceRoomInternalServerError creates a DeleteDeviceRoomInternalServerError with default headers values
func NewDeleteDeviceRoomInternalServerError() *DeleteDeviceRoomInternalServerError {
	return &DeleteDeviceRoomInternalServerError{}
}

/*DeleteDeviceRoomInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type DeleteDeviceRoomInternalServerError struct {
}

func (o *DeleteDeviceRoomInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device/room/{ID}/][%d] deleteDeviceRoomInternalServerError ", 500)
}

func (o *DeleteDeviceRoomInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceRoomServiceUnavailable creates a DeleteDeviceRoomServiceUnavailable with default headers values
func NewDeleteDeviceRoomServiceUnavailable() *DeleteDeviceRoomServiceUnavailable {
	return &DeleteDeviceRoomServiceUnavailable{}
}

/*DeleteDeviceRoomServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type DeleteDeviceRoomServiceUnavailable struct {
}

func (o *DeleteDeviceRoomServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device/room/{ID}/][%d] deleteDeviceRoomServiceUnavailable ", 503)
}

func (o *DeleteDeviceRoomServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeleteDeviceRoomOKBody delete device room o k body
swagger:model DeleteDeviceRoomOKBody
*/
type DeleteDeviceRoomOKBody struct {

	// deleted
	Deleted interface{} `json:"deleted,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`
}

// Validate validates this delete device room o k body
func (o *DeleteDeviceRoomOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteDeviceRoomOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteDeviceRoomOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteDeviceRoomOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}