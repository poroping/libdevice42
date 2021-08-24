// Code generated by go-swagger; DO NOT EDIT.

package services

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

	"github.com/poroping/libdevice42/models"
)

// GetListenerConnectionStatsV1Reader is a Reader for the GetListenerConnectionStatsV1 structure.
type GetListenerConnectionStatsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetListenerConnectionStatsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetListenerConnectionStatsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetListenerConnectionStatsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetListenerConnectionStatsV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetListenerConnectionStatsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetListenerConnectionStatsV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetListenerConnectionStatsV1MethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetListenerConnectionStatsV1Gone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetListenerConnectionStatsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetListenerConnectionStatsV1ServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetListenerConnectionStatsV1OK creates a GetListenerConnectionStatsV1OK with default headers values
func NewGetListenerConnectionStatsV1OK() *GetListenerConnectionStatsV1OK {
	return &GetListenerConnectionStatsV1OK{}
}

/*GetListenerConnectionStatsV1OK handles this case with default header values.

The above command returns results like this:
*/
type GetListenerConnectionStatsV1OK struct {
	Payload *GetListenerConnectionStatsV1OKBody
}

func (o *GetListenerConnectionStatsV1OK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/listener_connection_stats/][%d] getListenerConnectionStatsV1OK  %+v", 200, o.Payload)
}

func (o *GetListenerConnectionStatsV1OK) GetPayload() *GetListenerConnectionStatsV1OKBody {
	return o.Payload
}

func (o *GetListenerConnectionStatsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetListenerConnectionStatsV1OKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListenerConnectionStatsV1BadRequest creates a GetListenerConnectionStatsV1BadRequest with default headers values
func NewGetListenerConnectionStatsV1BadRequest() *GetListenerConnectionStatsV1BadRequest {
	return &GetListenerConnectionStatsV1BadRequest{}
}

/*GetListenerConnectionStatsV1BadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetListenerConnectionStatsV1BadRequest struct {
}

func (o *GetListenerConnectionStatsV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/listener_connection_stats/][%d] getListenerConnectionStatsV1BadRequest ", 400)
}

func (o *GetListenerConnectionStatsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListenerConnectionStatsV1Unauthorized creates a GetListenerConnectionStatsV1Unauthorized with default headers values
func NewGetListenerConnectionStatsV1Unauthorized() *GetListenerConnectionStatsV1Unauthorized {
	return &GetListenerConnectionStatsV1Unauthorized{}
}

/*GetListenerConnectionStatsV1Unauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetListenerConnectionStatsV1Unauthorized struct {
}

func (o *GetListenerConnectionStatsV1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/listener_connection_stats/][%d] getListenerConnectionStatsV1Unauthorized ", 401)
}

func (o *GetListenerConnectionStatsV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListenerConnectionStatsV1Forbidden creates a GetListenerConnectionStatsV1Forbidden with default headers values
func NewGetListenerConnectionStatsV1Forbidden() *GetListenerConnectionStatsV1Forbidden {
	return &GetListenerConnectionStatsV1Forbidden{}
}

/*GetListenerConnectionStatsV1Forbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetListenerConnectionStatsV1Forbidden struct {
}

func (o *GetListenerConnectionStatsV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/listener_connection_stats/][%d] getListenerConnectionStatsV1Forbidden ", 403)
}

func (o *GetListenerConnectionStatsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListenerConnectionStatsV1NotFound creates a GetListenerConnectionStatsV1NotFound with default headers values
func NewGetListenerConnectionStatsV1NotFound() *GetListenerConnectionStatsV1NotFound {
	return &GetListenerConnectionStatsV1NotFound{}
}

/*GetListenerConnectionStatsV1NotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetListenerConnectionStatsV1NotFound struct {
}

func (o *GetListenerConnectionStatsV1NotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/listener_connection_stats/][%d] getListenerConnectionStatsV1NotFound ", 404)
}

func (o *GetListenerConnectionStatsV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListenerConnectionStatsV1MethodNotAllowed creates a GetListenerConnectionStatsV1MethodNotAllowed with default headers values
func NewGetListenerConnectionStatsV1MethodNotAllowed() *GetListenerConnectionStatsV1MethodNotAllowed {
	return &GetListenerConnectionStatsV1MethodNotAllowed{}
}

/*GetListenerConnectionStatsV1MethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetListenerConnectionStatsV1MethodNotAllowed struct {
}

func (o *GetListenerConnectionStatsV1MethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/listener_connection_stats/][%d] getListenerConnectionStatsV1MethodNotAllowed ", 405)
}

func (o *GetListenerConnectionStatsV1MethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListenerConnectionStatsV1Gone creates a GetListenerConnectionStatsV1Gone with default headers values
func NewGetListenerConnectionStatsV1Gone() *GetListenerConnectionStatsV1Gone {
	return &GetListenerConnectionStatsV1Gone{}
}

/*GetListenerConnectionStatsV1Gone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetListenerConnectionStatsV1Gone struct {
}

func (o *GetListenerConnectionStatsV1Gone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/listener_connection_stats/][%d] getListenerConnectionStatsV1Gone ", 410)
}

func (o *GetListenerConnectionStatsV1Gone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListenerConnectionStatsV1InternalServerError creates a GetListenerConnectionStatsV1InternalServerError with default headers values
func NewGetListenerConnectionStatsV1InternalServerError() *GetListenerConnectionStatsV1InternalServerError {
	return &GetListenerConnectionStatsV1InternalServerError{}
}

/*GetListenerConnectionStatsV1InternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetListenerConnectionStatsV1InternalServerError struct {
}

func (o *GetListenerConnectionStatsV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/listener_connection_stats/][%d] getListenerConnectionStatsV1InternalServerError ", 500)
}

func (o *GetListenerConnectionStatsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListenerConnectionStatsV1ServiceUnavailable creates a GetListenerConnectionStatsV1ServiceUnavailable with default headers values
func NewGetListenerConnectionStatsV1ServiceUnavailable() *GetListenerConnectionStatsV1ServiceUnavailable {
	return &GetListenerConnectionStatsV1ServiceUnavailable{}
}

/*GetListenerConnectionStatsV1ServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetListenerConnectionStatsV1ServiceUnavailable struct {
}

func (o *GetListenerConnectionStatsV1ServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/listener_connection_stats/][%d] getListenerConnectionStatsV1ServiceUnavailable ", 503)
}

func (o *GetListenerConnectionStatsV1ServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetListenerConnectionStatsV1OKBody get listener connection stats v1 o k body
swagger:model GetListenerConnectionStatsV1OKBody
*/
type GetListenerConnectionStatsV1OKBody struct {

	// client ips
	ClientIps interface{} `json:"client_ips,omitempty"`

	// client stats
	ClientStats []*models.ClientStats `json:"client_stats"`

	// description
	Description interface{} `json:"description,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`

	// listener device id
	ListenerDeviceID interface{} `json:"listener_device_id,omitempty"`

	// listener device name
	ListenerDeviceName interface{} `json:"listener_device_name,omitempty"`

	// listener service
	ListenerService interface{} `json:"listener_service,omitempty"`

	// listening ip
	ListeningIP interface{} `json:"listening_ip,omitempty"`

	// port
	Port interface{} `json:"port,omitempty"`
}

// Validate validates this get listener connection stats v1 o k body
func (o *GetListenerConnectionStatsV1OKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClientStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetListenerConnectionStatsV1OKBody) validateClientStats(formats strfmt.Registry) error {

	if swag.IsZero(o.ClientStats) { // not required
		return nil
	}

	for i := 0; i < len(o.ClientStats); i++ {
		if swag.IsZero(o.ClientStats[i]) { // not required
			continue
		}

		if o.ClientStats[i] != nil {
			if err := o.ClientStats[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getListenerConnectionStatsV1OK" + "." + "client_stats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetListenerConnectionStatsV1OKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetListenerConnectionStatsV1OKBody) UnmarshalBinary(b []byte) error {
	var res GetListenerConnectionStatsV1OKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
