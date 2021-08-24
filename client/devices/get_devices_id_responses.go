// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/poroping/libdevice42/models"
)

// GetDevicesIDReader is a Reader for the GetDevicesID structure.
type GetDevicesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetDevicesIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetDevicesIDGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetDevicesIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDevicesIDOK creates a GetDevicesIDOK with default headers values
func NewGetDevicesIDOK() *GetDevicesIDOK {
	return &GetDevicesIDOK{}
}

/*GetDevicesIDOK handles this case with default header values.

The above command returns results like this:
*/
type GetDevicesIDOK struct {
	Payload *models.DevicesAll
}

func (o *GetDevicesIDOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/id/{device-id}/][%d] getDevicesIdOK  %+v", 200, o.Payload)
}

func (o *GetDevicesIDOK) GetPayload() *models.DevicesAll {
	return o.Payload
}

func (o *GetDevicesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DevicesAll)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesIDBadRequest creates a GetDevicesIDBadRequest with default headers values
func NewGetDevicesIDBadRequest() *GetDevicesIDBadRequest {
	return &GetDevicesIDBadRequest{}
}

/*GetDevicesIDBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetDevicesIDBadRequest struct {
}

func (o *GetDevicesIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/id/{device-id}/][%d] getDevicesIdBadRequest ", 400)
}

func (o *GetDevicesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesIDUnauthorized creates a GetDevicesIDUnauthorized with default headers values
func NewGetDevicesIDUnauthorized() *GetDevicesIDUnauthorized {
	return &GetDevicesIDUnauthorized{}
}

/*GetDevicesIDUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetDevicesIDUnauthorized struct {
}

func (o *GetDevicesIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/id/{device-id}/][%d] getDevicesIdUnauthorized ", 401)
}

func (o *GetDevicesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesIDForbidden creates a GetDevicesIDForbidden with default headers values
func NewGetDevicesIDForbidden() *GetDevicesIDForbidden {
	return &GetDevicesIDForbidden{}
}

/*GetDevicesIDForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetDevicesIDForbidden struct {
}

func (o *GetDevicesIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/id/{device-id}/][%d] getDevicesIdForbidden ", 403)
}

func (o *GetDevicesIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesIDNotFound creates a GetDevicesIDNotFound with default headers values
func NewGetDevicesIDNotFound() *GetDevicesIDNotFound {
	return &GetDevicesIDNotFound{}
}

/*GetDevicesIDNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetDevicesIDNotFound struct {
}

func (o *GetDevicesIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/id/{device-id}/][%d] getDevicesIdNotFound ", 404)
}

func (o *GetDevicesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesIDMethodNotAllowed creates a GetDevicesIDMethodNotAllowed with default headers values
func NewGetDevicesIDMethodNotAllowed() *GetDevicesIDMethodNotAllowed {
	return &GetDevicesIDMethodNotAllowed{}
}

/*GetDevicesIDMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetDevicesIDMethodNotAllowed struct {
}

func (o *GetDevicesIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/id/{device-id}/][%d] getDevicesIdMethodNotAllowed ", 405)
}

func (o *GetDevicesIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesIDGone creates a GetDevicesIDGone with default headers values
func NewGetDevicesIDGone() *GetDevicesIDGone {
	return &GetDevicesIDGone{}
}

/*GetDevicesIDGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetDevicesIDGone struct {
}

func (o *GetDevicesIDGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/id/{device-id}/][%d] getDevicesIdGone ", 410)
}

func (o *GetDevicesIDGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesIDInternalServerError creates a GetDevicesIDInternalServerError with default headers values
func NewGetDevicesIDInternalServerError() *GetDevicesIDInternalServerError {
	return &GetDevicesIDInternalServerError{}
}

/*GetDevicesIDInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetDevicesIDInternalServerError struct {
}

func (o *GetDevicesIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/id/{device-id}/][%d] getDevicesIdInternalServerError ", 500)
}

func (o *GetDevicesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesIDServiceUnavailable creates a GetDevicesIDServiceUnavailable with default headers values
func NewGetDevicesIDServiceUnavailable() *GetDevicesIDServiceUnavailable {
	return &GetDevicesIDServiceUnavailable{}
}

/*GetDevicesIDServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetDevicesIDServiceUnavailable struct {
}

func (o *GetDevicesIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/id/{device-id}/][%d] getDevicesIdServiceUnavailable ", 503)
}

func (o *GetDevicesIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
