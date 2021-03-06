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

// GetDevicesAssetReader is a Reader for the GetDevicesAsset structure.
type GetDevicesAssetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesAssetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesAssetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesAssetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesAssetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesAssetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesAssetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetDevicesAssetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetDevicesAssetGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesAssetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetDevicesAssetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDevicesAssetOK creates a GetDevicesAssetOK with default headers values
func NewGetDevicesAssetOK() *GetDevicesAssetOK {
	return &GetDevicesAssetOK{}
}

/*GetDevicesAssetOK handles this case with default header values.

The above command returns results like this:
*/
type GetDevicesAssetOK struct {
	Payload *models.DevicesCustomerID
}

func (o *GetDevicesAssetOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/asset/{device-asset}/][%d] getDevicesAssetOK  %+v", 200, o.Payload)
}

func (o *GetDevicesAssetOK) GetPayload() *models.DevicesCustomerID {
	return o.Payload
}

func (o *GetDevicesAssetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DevicesCustomerID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAssetBadRequest creates a GetDevicesAssetBadRequest with default headers values
func NewGetDevicesAssetBadRequest() *GetDevicesAssetBadRequest {
	return &GetDevicesAssetBadRequest{}
}

/*GetDevicesAssetBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetDevicesAssetBadRequest struct {
}

func (o *GetDevicesAssetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/asset/{device-asset}/][%d] getDevicesAssetBadRequest ", 400)
}

func (o *GetDevicesAssetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesAssetUnauthorized creates a GetDevicesAssetUnauthorized with default headers values
func NewGetDevicesAssetUnauthorized() *GetDevicesAssetUnauthorized {
	return &GetDevicesAssetUnauthorized{}
}

/*GetDevicesAssetUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetDevicesAssetUnauthorized struct {
}

func (o *GetDevicesAssetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/asset/{device-asset}/][%d] getDevicesAssetUnauthorized ", 401)
}

func (o *GetDevicesAssetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesAssetForbidden creates a GetDevicesAssetForbidden with default headers values
func NewGetDevicesAssetForbidden() *GetDevicesAssetForbidden {
	return &GetDevicesAssetForbidden{}
}

/*GetDevicesAssetForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetDevicesAssetForbidden struct {
}

func (o *GetDevicesAssetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/asset/{device-asset}/][%d] getDevicesAssetForbidden ", 403)
}

func (o *GetDevicesAssetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesAssetNotFound creates a GetDevicesAssetNotFound with default headers values
func NewGetDevicesAssetNotFound() *GetDevicesAssetNotFound {
	return &GetDevicesAssetNotFound{}
}

/*GetDevicesAssetNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetDevicesAssetNotFound struct {
}

func (o *GetDevicesAssetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/asset/{device-asset}/][%d] getDevicesAssetNotFound ", 404)
}

func (o *GetDevicesAssetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesAssetMethodNotAllowed creates a GetDevicesAssetMethodNotAllowed with default headers values
func NewGetDevicesAssetMethodNotAllowed() *GetDevicesAssetMethodNotAllowed {
	return &GetDevicesAssetMethodNotAllowed{}
}

/*GetDevicesAssetMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetDevicesAssetMethodNotAllowed struct {
}

func (o *GetDevicesAssetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/asset/{device-asset}/][%d] getDevicesAssetMethodNotAllowed ", 405)
}

func (o *GetDevicesAssetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesAssetGone creates a GetDevicesAssetGone with default headers values
func NewGetDevicesAssetGone() *GetDevicesAssetGone {
	return &GetDevicesAssetGone{}
}

/*GetDevicesAssetGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetDevicesAssetGone struct {
}

func (o *GetDevicesAssetGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/asset/{device-asset}/][%d] getDevicesAssetGone ", 410)
}

func (o *GetDevicesAssetGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesAssetInternalServerError creates a GetDevicesAssetInternalServerError with default headers values
func NewGetDevicesAssetInternalServerError() *GetDevicesAssetInternalServerError {
	return &GetDevicesAssetInternalServerError{}
}

/*GetDevicesAssetInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type GetDevicesAssetInternalServerError struct {
}

func (o *GetDevicesAssetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/asset/{device-asset}/][%d] getDevicesAssetInternalServerError ", 500)
}

func (o *GetDevicesAssetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDevicesAssetServiceUnavailable creates a GetDevicesAssetServiceUnavailable with default headers values
func NewGetDevicesAssetServiceUnavailable() *GetDevicesAssetServiceUnavailable {
	return &GetDevicesAssetServiceUnavailable{}
}

/*GetDevicesAssetServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetDevicesAssetServiceUnavailable struct {
}

func (o *GetDevicesAssetServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/devices/asset/{device-asset}/][%d] getDevicesAssetServiceUnavailable ", 503)
}

func (o *GetDevicesAssetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
