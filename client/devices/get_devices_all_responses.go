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

// GetDevicesAllReader is a Reader for the GetDevicesAll structure.
type GetDevicesAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesAllBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesAllUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesAllForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesAllNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetDevicesAllMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetDevicesAllGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesAllInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetDevicesAllServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDevicesAllOK creates a GetDevicesAllOK with default headers values
func NewGetDevicesAllOK() *GetDevicesAllOK {
	return &GetDevicesAllOK{}
}

/*GetDevicesAllOK handles this case with default header values.

The above command returns results like this:
*/
type GetDevicesAllOK struct {
	Payload *models.DevicesAll
}

func (o *GetDevicesAllOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/all/][%d] getDevicesAllOK  %+v", 200, o.Payload)
}

func (o *GetDevicesAllOK) GetPayload() *models.DevicesAll {
	return o.Payload
}

func (o *GetDevicesAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DevicesAll)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAllBadRequest creates a GetDevicesAllBadRequest with default headers values
func NewGetDevicesAllBadRequest() *GetDevicesAllBadRequest {
	return &GetDevicesAllBadRequest{}
}

/*GetDevicesAllBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetDevicesAllBadRequest struct {
}

func (o *GetDevicesAllBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/all/][%d] getDevicesAllBadRequest ", 400)
}

func (o *GetDevicesAllBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesAllUnauthorized creates a GetDevicesAllUnauthorized with default headers values
func NewGetDevicesAllUnauthorized() *GetDevicesAllUnauthorized {
	return &GetDevicesAllUnauthorized{}
}

/*GetDevicesAllUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetDevicesAllUnauthorized struct {
}

func (o *GetDevicesAllUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/all/][%d] getDevicesAllUnauthorized ", 401)
}

func (o *GetDevicesAllUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesAllForbidden creates a GetDevicesAllForbidden with default headers values
func NewGetDevicesAllForbidden() *GetDevicesAllForbidden {
	return &GetDevicesAllForbidden{}
}

/*GetDevicesAllForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetDevicesAllForbidden struct {
}

func (o *GetDevicesAllForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/all/][%d] getDevicesAllForbidden ", 403)
}

func (o *GetDevicesAllForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesAllNotFound creates a GetDevicesAllNotFound with default headers values
func NewGetDevicesAllNotFound() *GetDevicesAllNotFound {
	return &GetDevicesAllNotFound{}
}

/*GetDevicesAllNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetDevicesAllNotFound struct {
}

func (o *GetDevicesAllNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/all/][%d] getDevicesAllNotFound ", 404)
}

func (o *GetDevicesAllNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesAllMethodNotAllowed creates a GetDevicesAllMethodNotAllowed with default headers values
func NewGetDevicesAllMethodNotAllowed() *GetDevicesAllMethodNotAllowed {
	return &GetDevicesAllMethodNotAllowed{}
}

/*GetDevicesAllMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetDevicesAllMethodNotAllowed struct {
}

func (o *GetDevicesAllMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/all/][%d] getDevicesAllMethodNotAllowed ", 405)
}

func (o *GetDevicesAllMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesAllGone creates a GetDevicesAllGone with default headers values
func NewGetDevicesAllGone() *GetDevicesAllGone {
	return &GetDevicesAllGone{}
}

/*GetDevicesAllGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetDevicesAllGone struct {
}

func (o *GetDevicesAllGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/all/][%d] getDevicesAllGone ", 410)
}

func (o *GetDevicesAllGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesAllInternalServerError creates a GetDevicesAllInternalServerError with default headers values
func NewGetDevicesAllInternalServerError() *GetDevicesAllInternalServerError {
	return &GetDevicesAllInternalServerError{}
}

/*GetDevicesAllInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type GetDevicesAllInternalServerError struct {
}

func (o *GetDevicesAllInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/all/][%d] getDevicesAllInternalServerError ", 500)
}

func (o *GetDevicesAllInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesAllServiceUnavailable creates a GetDevicesAllServiceUnavailable with default headers values
func NewGetDevicesAllServiceUnavailable() *GetDevicesAllServiceUnavailable {
	return &GetDevicesAllServiceUnavailable{}
}

/*GetDevicesAllServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetDevicesAllServiceUnavailable struct {
}

func (o *GetDevicesAllServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/all/][%d] getDevicesAllServiceUnavailable ", 503)
}

func (o *GetDevicesAllServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
