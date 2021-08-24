// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/poroping/libdevice42/models"
)

// GetServiceInstancesReader is a Reader for the GetServiceInstances structure.
type GetServiceInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetServiceInstancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetServiceInstancesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetServiceInstancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetServiceInstancesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetServiceInstancesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetServiceInstancesGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetServiceInstancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetServiceInstancesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceInstancesOK creates a GetServiceInstancesOK with default headers values
func NewGetServiceInstancesOK() *GetServiceInstancesOK {
	return &GetServiceInstancesOK{}
}

/*GetServiceInstancesOK handles this case with default header values.

The above command returns results like this:
*/
type GetServiceInstancesOK struct {
	Payload *GetServiceInstancesOKBody
}

func (o *GetServiceInstancesOK) Error() string {
	return fmt.Sprintf("[GET /api/2.0/service_instances/][%d] getServiceInstancesOK  %+v", 200, o.Payload)
}

func (o *GetServiceInstancesOK) GetPayload() *GetServiceInstancesOKBody {
	return o.Payload
}

func (o *GetServiceInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetServiceInstancesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceInstancesBadRequest creates a GetServiceInstancesBadRequest with default headers values
func NewGetServiceInstancesBadRequest() *GetServiceInstancesBadRequest {
	return &GetServiceInstancesBadRequest{}
}

/*GetServiceInstancesBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetServiceInstancesBadRequest struct {
}

func (o *GetServiceInstancesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.0/service_instances/][%d] getServiceInstancesBadRequest ", 400)
}

func (o *GetServiceInstancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceInstancesUnauthorized creates a GetServiceInstancesUnauthorized with default headers values
func NewGetServiceInstancesUnauthorized() *GetServiceInstancesUnauthorized {
	return &GetServiceInstancesUnauthorized{}
}

/*GetServiceInstancesUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetServiceInstancesUnauthorized struct {
}

func (o *GetServiceInstancesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/2.0/service_instances/][%d] getServiceInstancesUnauthorized ", 401)
}

func (o *GetServiceInstancesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceInstancesForbidden creates a GetServiceInstancesForbidden with default headers values
func NewGetServiceInstancesForbidden() *GetServiceInstancesForbidden {
	return &GetServiceInstancesForbidden{}
}

/*GetServiceInstancesForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetServiceInstancesForbidden struct {
}

func (o *GetServiceInstancesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/2.0/service_instances/][%d] getServiceInstancesForbidden ", 403)
}

func (o *GetServiceInstancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceInstancesNotFound creates a GetServiceInstancesNotFound with default headers values
func NewGetServiceInstancesNotFound() *GetServiceInstancesNotFound {
	return &GetServiceInstancesNotFound{}
}

/*GetServiceInstancesNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetServiceInstancesNotFound struct {
}

func (o *GetServiceInstancesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/2.0/service_instances/][%d] getServiceInstancesNotFound ", 404)
}

func (o *GetServiceInstancesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceInstancesMethodNotAllowed creates a GetServiceInstancesMethodNotAllowed with default headers values
func NewGetServiceInstancesMethodNotAllowed() *GetServiceInstancesMethodNotAllowed {
	return &GetServiceInstancesMethodNotAllowed{}
}

/*GetServiceInstancesMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetServiceInstancesMethodNotAllowed struct {
}

func (o *GetServiceInstancesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/2.0/service_instances/][%d] getServiceInstancesMethodNotAllowed ", 405)
}

func (o *GetServiceInstancesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceInstancesGone creates a GetServiceInstancesGone with default headers values
func NewGetServiceInstancesGone() *GetServiceInstancesGone {
	return &GetServiceInstancesGone{}
}

/*GetServiceInstancesGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetServiceInstancesGone struct {
}

func (o *GetServiceInstancesGone) Error() string {
	return fmt.Sprintf("[GET /api/2.0/service_instances/][%d] getServiceInstancesGone ", 410)
}

func (o *GetServiceInstancesGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceInstancesInternalServerError creates a GetServiceInstancesInternalServerError with default headers values
func NewGetServiceInstancesInternalServerError() *GetServiceInstancesInternalServerError {
	return &GetServiceInstancesInternalServerError{}
}

/*GetServiceInstancesInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetServiceInstancesInternalServerError struct {
}

func (o *GetServiceInstancesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/2.0/service_instances/][%d] getServiceInstancesInternalServerError ", 500)
}

func (o *GetServiceInstancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceInstancesServiceUnavailable creates a GetServiceInstancesServiceUnavailable with default headers values
func NewGetServiceInstancesServiceUnavailable() *GetServiceInstancesServiceUnavailable {
	return &GetServiceInstancesServiceUnavailable{}
}

/*GetServiceInstancesServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetServiceInstancesServiceUnavailable struct {
}

func (o *GetServiceInstancesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/2.0/service_instances/][%d] getServiceInstancesServiceUnavailable ", 503)
}

func (o *GetServiceInstancesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetServiceInstancesOKBody get service instances o k body
swagger:model GetServiceInstancesOKBody
*/
type GetServiceInstancesOKBody struct {

	// limit
	Limit interface{} `json:"limit,omitempty"`

	// offset
	Offset interface{} `json:"offset,omitempty"`

	// service details
	ServiceDetails models.ServiceInstances `json:"service_details"`

	// total count
	TotalCount interface{} `json:"total_count,omitempty"`
}

// Validate validates this get service instances o k body
func (o *GetServiceInstancesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateServiceDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServiceInstancesOKBody) validateServiceDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.ServiceDetails) { // not required
		return nil
	}

	if err := o.ServiceDetails.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getServiceInstancesOK" + "." + "service_details")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceInstancesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceInstancesOKBody) UnmarshalBinary(b []byte) error {
	var res GetServiceInstancesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
