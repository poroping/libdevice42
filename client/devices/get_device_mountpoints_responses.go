// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetDeviceMountpointsReader is a Reader for the GetDeviceMountpoints structure.
type GetDeviceMountpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceMountpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceMountpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDeviceMountpointsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDeviceMountpointsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDeviceMountpointsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeviceMountpointsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetDeviceMountpointsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetDeviceMountpointsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDeviceMountpointsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetDeviceMountpointsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDeviceMountpointsOK creates a GetDeviceMountpointsOK with default headers values
func NewGetDeviceMountpointsOK() *GetDeviceMountpointsOK {
	return &GetDeviceMountpointsOK{}
}

/*GetDeviceMountpointsOK handles this case with default header values.

The above command returns results like this:
*/
type GetDeviceMountpointsOK struct {
	Payload *GetDeviceMountpointsOKBody
}

func (o *GetDeviceMountpointsOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/device/mountpoints/][%d] getDeviceMountpointsOK  %+v", 200, o.Payload)
}

func (o *GetDeviceMountpointsOK) GetPayload() *GetDeviceMountpointsOKBody {
	return o.Payload
}

func (o *GetDeviceMountpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetDeviceMountpointsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceMountpointsBadRequest creates a GetDeviceMountpointsBadRequest with default headers values
func NewGetDeviceMountpointsBadRequest() *GetDeviceMountpointsBadRequest {
	return &GetDeviceMountpointsBadRequest{}
}

/*GetDeviceMountpointsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetDeviceMountpointsBadRequest struct {
}

func (o *GetDeviceMountpointsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/device/mountpoints/][%d] getDeviceMountpointsBadRequest ", 400)
}

func (o *GetDeviceMountpointsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeviceMountpointsUnauthorized creates a GetDeviceMountpointsUnauthorized with default headers values
func NewGetDeviceMountpointsUnauthorized() *GetDeviceMountpointsUnauthorized {
	return &GetDeviceMountpointsUnauthorized{}
}

/*GetDeviceMountpointsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetDeviceMountpointsUnauthorized struct {
}

func (o *GetDeviceMountpointsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/device/mountpoints/][%d] getDeviceMountpointsUnauthorized ", 401)
}

func (o *GetDeviceMountpointsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeviceMountpointsForbidden creates a GetDeviceMountpointsForbidden with default headers values
func NewGetDeviceMountpointsForbidden() *GetDeviceMountpointsForbidden {
	return &GetDeviceMountpointsForbidden{}
}

/*GetDeviceMountpointsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetDeviceMountpointsForbidden struct {
}

func (o *GetDeviceMountpointsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/device/mountpoints/][%d] getDeviceMountpointsForbidden ", 403)
}

func (o *GetDeviceMountpointsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeviceMountpointsNotFound creates a GetDeviceMountpointsNotFound with default headers values
func NewGetDeviceMountpointsNotFound() *GetDeviceMountpointsNotFound {
	return &GetDeviceMountpointsNotFound{}
}

/*GetDeviceMountpointsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetDeviceMountpointsNotFound struct {
}

func (o *GetDeviceMountpointsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/device/mountpoints/][%d] getDeviceMountpointsNotFound ", 404)
}

func (o *GetDeviceMountpointsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeviceMountpointsMethodNotAllowed creates a GetDeviceMountpointsMethodNotAllowed with default headers values
func NewGetDeviceMountpointsMethodNotAllowed() *GetDeviceMountpointsMethodNotAllowed {
	return &GetDeviceMountpointsMethodNotAllowed{}
}

/*GetDeviceMountpointsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetDeviceMountpointsMethodNotAllowed struct {
}

func (o *GetDeviceMountpointsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/device/mountpoints/][%d] getDeviceMountpointsMethodNotAllowed ", 405)
}

func (o *GetDeviceMountpointsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeviceMountpointsGone creates a GetDeviceMountpointsGone with default headers values
func NewGetDeviceMountpointsGone() *GetDeviceMountpointsGone {
	return &GetDeviceMountpointsGone{}
}

/*GetDeviceMountpointsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetDeviceMountpointsGone struct {
}

func (o *GetDeviceMountpointsGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/device/mountpoints/][%d] getDeviceMountpointsGone ", 410)
}

func (o *GetDeviceMountpointsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeviceMountpointsInternalServerError creates a GetDeviceMountpointsInternalServerError with default headers values
func NewGetDeviceMountpointsInternalServerError() *GetDeviceMountpointsInternalServerError {
	return &GetDeviceMountpointsInternalServerError{}
}

/*GetDeviceMountpointsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type GetDeviceMountpointsInternalServerError struct {
}

func (o *GetDeviceMountpointsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/device/mountpoints/][%d] getDeviceMountpointsInternalServerError ", 500)
}

func (o *GetDeviceMountpointsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeviceMountpointsServiceUnavailable creates a GetDeviceMountpointsServiceUnavailable with default headers values
func NewGetDeviceMountpointsServiceUnavailable() *GetDeviceMountpointsServiceUnavailable {
	return &GetDeviceMountpointsServiceUnavailable{}
}

/*GetDeviceMountpointsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetDeviceMountpointsServiceUnavailable struct {
}

func (o *GetDeviceMountpointsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/device/mountpoints/][%d] getDeviceMountpointsServiceUnavailable ", 503)
}

func (o *GetDeviceMountpointsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetDeviceMountpointsOKBody get device mountpoints o k body
swagger:model GetDeviceMountpointsOKBody
*/
type GetDeviceMountpointsOKBody struct {

	// mountpoint details
	MountpointDetails []*MountpointDetailsItems0 `json:"mountpoint_details"`

	// total count
	TotalCount interface{} `json:"total_count,omitempty"`
}

// Validate validates this get device mountpoints o k body
func (o *GetDeviceMountpointsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMountpointDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDeviceMountpointsOKBody) validateMountpointDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.MountpointDetails) { // not required
		return nil
	}

	for i := 0; i < len(o.MountpointDetails); i++ {
		if swag.IsZero(o.MountpointDetails[i]) { // not required
			continue
		}

		if o.MountpointDetails[i] != nil {
			if err := o.MountpointDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDeviceMountpointsOK" + "." + "mountpoint_details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDeviceMountpointsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeviceMountpointsOKBody) UnmarshalBinary(b []byte) error {
	var res GetDeviceMountpointsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*MountpointDetailsItems0 mountpoint details items0
swagger:model MountpointDetailsItems0
*/
type MountpointDetailsItems0 struct {

	// capacity
	Capacity interface{} `json:"capacity,omitempty"`

	// device
	Device interface{} `json:"device,omitempty"`

	// filesystem
	Filesystem interface{} `json:"filesystem,omitempty"`

	// free capacity
	FreeCapacity interface{} `json:"free_capacity,omitempty"`

	// fstype
	Fstype interface{} `json:"fstype,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`

	// label
	Label interface{} `json:"label,omitempty"`

	// last updated
	LastUpdated interface{} `json:"last_updated,omitempty"`

	// mountpoint
	Mountpoint interface{} `json:"mountpoint,omitempty"`
}

// Validate validates this mountpoint details items0
func (o *MountpointDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MountpointDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MountpointDetailsItems0) UnmarshalBinary(b []byte) error {
	var res MountpointDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
