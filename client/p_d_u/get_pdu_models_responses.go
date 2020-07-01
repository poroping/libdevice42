// Code generated by go-swagger; DO NOT EDIT.

package p_d_u

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

	"github.com/mjpitz/libdevice42/models"
)

// GetPduModelsReader is a Reader for the GetPduModels structure.
type GetPduModelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPduModelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPduModelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPduModelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPduModelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPduModelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPduModelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetPduModelsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetPduModelsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPduModelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetPduModelsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPduModelsOK creates a GetPduModelsOK with default headers values
func NewGetPduModelsOK() *GetPduModelsOK {
	return &GetPduModelsOK{}
}

/*GetPduModelsOK handles this case with default header values.

The above command returns results like this:
*/
type GetPduModelsOK struct {
	Payload *GetPduModelsOKBody
}

func (o *GetPduModelsOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/pdu_models/][%d] getPduModelsOK  %+v", 200, o.Payload)
}

func (o *GetPduModelsOK) GetPayload() *GetPduModelsOKBody {
	return o.Payload
}

func (o *GetPduModelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPduModelsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPduModelsBadRequest creates a GetPduModelsBadRequest with default headers values
func NewGetPduModelsBadRequest() *GetPduModelsBadRequest {
	return &GetPduModelsBadRequest{}
}

/*GetPduModelsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetPduModelsBadRequest struct {
}

func (o *GetPduModelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/pdu_models/][%d] getPduModelsBadRequest ", 400)
}

func (o *GetPduModelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPduModelsUnauthorized creates a GetPduModelsUnauthorized with default headers values
func NewGetPduModelsUnauthorized() *GetPduModelsUnauthorized {
	return &GetPduModelsUnauthorized{}
}

/*GetPduModelsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetPduModelsUnauthorized struct {
}

func (o *GetPduModelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/pdu_models/][%d] getPduModelsUnauthorized ", 401)
}

func (o *GetPduModelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPduModelsForbidden creates a GetPduModelsForbidden with default headers values
func NewGetPduModelsForbidden() *GetPduModelsForbidden {
	return &GetPduModelsForbidden{}
}

/*GetPduModelsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetPduModelsForbidden struct {
}

func (o *GetPduModelsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/pdu_models/][%d] getPduModelsForbidden ", 403)
}

func (o *GetPduModelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPduModelsNotFound creates a GetPduModelsNotFound with default headers values
func NewGetPduModelsNotFound() *GetPduModelsNotFound {
	return &GetPduModelsNotFound{}
}

/*GetPduModelsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetPduModelsNotFound struct {
}

func (o *GetPduModelsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/pdu_models/][%d] getPduModelsNotFound ", 404)
}

func (o *GetPduModelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPduModelsMethodNotAllowed creates a GetPduModelsMethodNotAllowed with default headers values
func NewGetPduModelsMethodNotAllowed() *GetPduModelsMethodNotAllowed {
	return &GetPduModelsMethodNotAllowed{}
}

/*GetPduModelsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetPduModelsMethodNotAllowed struct {
}

func (o *GetPduModelsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/pdu_models/][%d] getPduModelsMethodNotAllowed ", 405)
}

func (o *GetPduModelsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPduModelsGone creates a GetPduModelsGone with default headers values
func NewGetPduModelsGone() *GetPduModelsGone {
	return &GetPduModelsGone{}
}

/*GetPduModelsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetPduModelsGone struct {
}

func (o *GetPduModelsGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/pdu_models/][%d] getPduModelsGone ", 410)
}

func (o *GetPduModelsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPduModelsInternalServerError creates a GetPduModelsInternalServerError with default headers values
func NewGetPduModelsInternalServerError() *GetPduModelsInternalServerError {
	return &GetPduModelsInternalServerError{}
}

/*GetPduModelsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetPduModelsInternalServerError struct {
}

func (o *GetPduModelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/pdu_models/][%d] getPduModelsInternalServerError ", 500)
}

func (o *GetPduModelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPduModelsServiceUnavailable creates a GetPduModelsServiceUnavailable with default headers values
func NewGetPduModelsServiceUnavailable() *GetPduModelsServiceUnavailable {
	return &GetPduModelsServiceUnavailable{}
}

/*GetPduModelsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetPduModelsServiceUnavailable struct {
}

func (o *GetPduModelsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/pdu_models/][%d] getPduModelsServiceUnavailable ", 503)
}

func (o *GetPduModelsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetPduModelsOKBody get pdu models o k body
swagger:model GetPduModelsOKBody
*/
type GetPduModelsOKBody struct {

	// pdu models
	PduModels []*models.PduModels `json:"pdu_models"`
}

// Validate validates this get pdu models o k body
func (o *GetPduModelsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePduModels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPduModelsOKBody) validatePduModels(formats strfmt.Registry) error {

	if swag.IsZero(o.PduModels) { // not required
		return nil
	}

	for i := 0; i < len(o.PduModels); i++ {
		if swag.IsZero(o.PduModels[i]) { // not required
			continue
		}

		if o.PduModels[i] != nil {
			if err := o.PduModels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPduModelsOK" + "." + "pdu_models" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPduModelsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPduModelsOKBody) UnmarshalBinary(b []byte) error {
	var res GetPduModelsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}