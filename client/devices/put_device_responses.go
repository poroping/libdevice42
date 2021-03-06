// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/poroping/libdevice42/models"
)

// PutDeviceReader is a Reader for the PutDevice structure.
type PutDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutDeviceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutDeviceOK creates a PutDeviceOK with default headers values
func NewPutDeviceOK() *PutDeviceOK {
	return &PutDeviceOK{}
}

/*PutDeviceOK handles this case with default header values.

The above command returns results like this:
*/
type PutDeviceOK struct {
	Payload *PutDeviceOKBody
}

func (o *PutDeviceOK) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/device/][%d] putDeviceOK  %+v", 200, o.Payload)
}

func (o *PutDeviceOK) GetPayload() *PutDeviceOKBody {
	return o.Payload
}

func (o *PutDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutDeviceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDeviceBadRequest creates a PutDeviceBadRequest with default headers values
func NewPutDeviceBadRequest() *PutDeviceBadRequest {
	return &PutDeviceBadRequest{}
}

/*PutDeviceBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PutDeviceBadRequest struct {
}

func (o *PutDeviceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/1.0/device/][%d] putDeviceBadRequest ", 400)
}

func (o *PutDeviceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PutDeviceOKBody put device o k body
swagger:model PutDeviceOKBody
*/
type PutDeviceOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg models.DeviceName `json:"msg,omitempty"`
}

// Validate validates this put device o k body
func (o *PutDeviceOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutDeviceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutDeviceOKBody) UnmarshalBinary(b []byte) error {
	var res PutDeviceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
