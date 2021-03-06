// Code generated by go-swagger; DO NOT EDIT.

package operating_systems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteDeviceOsReader is a Reader for the DeleteDeviceOs structure.
type DeleteDeviceOsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDeviceOsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDeviceOsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDeviceOsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteDeviceOsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDeviceOsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDeviceOsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeleteDeviceOsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewDeleteDeviceOsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDeviceOsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteDeviceOsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDeviceOsOK creates a DeleteDeviceOsOK with default headers values
func NewDeleteDeviceOsOK() *DeleteDeviceOsOK {
	return &DeleteDeviceOsOK{}
}

/*DeleteDeviceOsOK handles this case with default header values.

The above command returns results like this:
*/
type DeleteDeviceOsOK struct {
	Payload *DeleteDeviceOsOKBody
}

func (o *DeleteDeviceOsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device_os/{device_os_id}/][%d] deleteDeviceOsOK  %+v", 200, o.Payload)
}

func (o *DeleteDeviceOsOK) GetPayload() *DeleteDeviceOsOKBody {
	return o.Payload
}

func (o *DeleteDeviceOsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteDeviceOsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceOsBadRequest creates a DeleteDeviceOsBadRequest with default headers values
func NewDeleteDeviceOsBadRequest() *DeleteDeviceOsBadRequest {
	return &DeleteDeviceOsBadRequest{}
}

/*DeleteDeviceOsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type DeleteDeviceOsBadRequest struct {
}

func (o *DeleteDeviceOsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device_os/{device_os_id}/][%d] deleteDeviceOsBadRequest ", 400)
}

func (o *DeleteDeviceOsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceOsUnauthorized creates a DeleteDeviceOsUnauthorized with default headers values
func NewDeleteDeviceOsUnauthorized() *DeleteDeviceOsUnauthorized {
	return &DeleteDeviceOsUnauthorized{}
}

/*DeleteDeviceOsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type DeleteDeviceOsUnauthorized struct {
}

func (o *DeleteDeviceOsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device_os/{device_os_id}/][%d] deleteDeviceOsUnauthorized ", 401)
}

func (o *DeleteDeviceOsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceOsForbidden creates a DeleteDeviceOsForbidden with default headers values
func NewDeleteDeviceOsForbidden() *DeleteDeviceOsForbidden {
	return &DeleteDeviceOsForbidden{}
}

/*DeleteDeviceOsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type DeleteDeviceOsForbidden struct {
}

func (o *DeleteDeviceOsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device_os/{device_os_id}/][%d] deleteDeviceOsForbidden ", 403)
}

func (o *DeleteDeviceOsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceOsNotFound creates a DeleteDeviceOsNotFound with default headers values
func NewDeleteDeviceOsNotFound() *DeleteDeviceOsNotFound {
	return &DeleteDeviceOsNotFound{}
}

/*DeleteDeviceOsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type DeleteDeviceOsNotFound struct {
}

func (o *DeleteDeviceOsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device_os/{device_os_id}/][%d] deleteDeviceOsNotFound ", 404)
}

func (o *DeleteDeviceOsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceOsMethodNotAllowed creates a DeleteDeviceOsMethodNotAllowed with default headers values
func NewDeleteDeviceOsMethodNotAllowed() *DeleteDeviceOsMethodNotAllowed {
	return &DeleteDeviceOsMethodNotAllowed{}
}

/*DeleteDeviceOsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type DeleteDeviceOsMethodNotAllowed struct {
}

func (o *DeleteDeviceOsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device_os/{device_os_id}/][%d] deleteDeviceOsMethodNotAllowed ", 405)
}

func (o *DeleteDeviceOsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceOsGone creates a DeleteDeviceOsGone with default headers values
func NewDeleteDeviceOsGone() *DeleteDeviceOsGone {
	return &DeleteDeviceOsGone{}
}

/*DeleteDeviceOsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type DeleteDeviceOsGone struct {
}

func (o *DeleteDeviceOsGone) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device_os/{device_os_id}/][%d] deleteDeviceOsGone ", 410)
}

func (o *DeleteDeviceOsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceOsInternalServerError creates a DeleteDeviceOsInternalServerError with default headers values
func NewDeleteDeviceOsInternalServerError() *DeleteDeviceOsInternalServerError {
	return &DeleteDeviceOsInternalServerError{}
}

/*DeleteDeviceOsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type DeleteDeviceOsInternalServerError struct {
}

func (o *DeleteDeviceOsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device_os/{device_os_id}/][%d] deleteDeviceOsInternalServerError ", 500)
}

func (o *DeleteDeviceOsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceOsServiceUnavailable creates a DeleteDeviceOsServiceUnavailable with default headers values
func NewDeleteDeviceOsServiceUnavailable() *DeleteDeviceOsServiceUnavailable {
	return &DeleteDeviceOsServiceUnavailable{}
}

/*DeleteDeviceOsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type DeleteDeviceOsServiceUnavailable struct {
}

func (o *DeleteDeviceOsServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/device_os/{device_os_id}/][%d] deleteDeviceOsServiceUnavailable ", 503)
}

func (o *DeleteDeviceOsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeleteDeviceOsOKBody delete device os o k body
swagger:model DeleteDeviceOsOKBody
*/
type DeleteDeviceOsOKBody struct {

	// deleted
	Deleted interface{} `json:"deleted,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`
}

// Validate validates this delete device os o k body
func (o *DeleteDeviceOsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteDeviceOsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteDeviceOsOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteDeviceOsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
