// Code generated by go-swagger; DO NOT EDIT.

package parts_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/poroping/libdevice42/models"
)

// GetPartmodelsReader is a Reader for the GetPartmodels structure.
type GetPartmodelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPartmodelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPartmodelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPartmodelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPartmodelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPartmodelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPartmodelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetPartmodelsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetPartmodelsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPartmodelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetPartmodelsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPartmodelsOK creates a GetPartmodelsOK with default headers values
func NewGetPartmodelsOK() *GetPartmodelsOK {
	return &GetPartmodelsOK{}
}

/*GetPartmodelsOK handles this case with default header values.

The above command returns results like this:
*/
type GetPartmodelsOK struct {
	Payload *models.Partmodels
}

func (o *GetPartmodelsOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/partmodels/][%d] getPartmodelsOK  %+v", 200, o.Payload)
}

func (o *GetPartmodelsOK) GetPayload() *models.Partmodels {
	return o.Payload
}

func (o *GetPartmodelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Partmodels)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPartmodelsBadRequest creates a GetPartmodelsBadRequest with default headers values
func NewGetPartmodelsBadRequest() *GetPartmodelsBadRequest {
	return &GetPartmodelsBadRequest{}
}

/*GetPartmodelsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetPartmodelsBadRequest struct {
}

func (o *GetPartmodelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/partmodels/][%d] getPartmodelsBadRequest ", 400)
}

func (o *GetPartmodelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPartmodelsUnauthorized creates a GetPartmodelsUnauthorized with default headers values
func NewGetPartmodelsUnauthorized() *GetPartmodelsUnauthorized {
	return &GetPartmodelsUnauthorized{}
}

/*GetPartmodelsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetPartmodelsUnauthorized struct {
}

func (o *GetPartmodelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/partmodels/][%d] getPartmodelsUnauthorized ", 401)
}

func (o *GetPartmodelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPartmodelsForbidden creates a GetPartmodelsForbidden with default headers values
func NewGetPartmodelsForbidden() *GetPartmodelsForbidden {
	return &GetPartmodelsForbidden{}
}

/*GetPartmodelsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetPartmodelsForbidden struct {
}

func (o *GetPartmodelsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/partmodels/][%d] getPartmodelsForbidden ", 403)
}

func (o *GetPartmodelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPartmodelsNotFound creates a GetPartmodelsNotFound with default headers values
func NewGetPartmodelsNotFound() *GetPartmodelsNotFound {
	return &GetPartmodelsNotFound{}
}

/*GetPartmodelsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetPartmodelsNotFound struct {
}

func (o *GetPartmodelsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/partmodels/][%d] getPartmodelsNotFound ", 404)
}

func (o *GetPartmodelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPartmodelsMethodNotAllowed creates a GetPartmodelsMethodNotAllowed with default headers values
func NewGetPartmodelsMethodNotAllowed() *GetPartmodelsMethodNotAllowed {
	return &GetPartmodelsMethodNotAllowed{}
}

/*GetPartmodelsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetPartmodelsMethodNotAllowed struct {
}

func (o *GetPartmodelsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/partmodels/][%d] getPartmodelsMethodNotAllowed ", 405)
}

func (o *GetPartmodelsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPartmodelsGone creates a GetPartmodelsGone with default headers values
func NewGetPartmodelsGone() *GetPartmodelsGone {
	return &GetPartmodelsGone{}
}

/*GetPartmodelsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetPartmodelsGone struct {
}

func (o *GetPartmodelsGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/partmodels/][%d] getPartmodelsGone ", 410)
}

func (o *GetPartmodelsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPartmodelsInternalServerError creates a GetPartmodelsInternalServerError with default headers values
func NewGetPartmodelsInternalServerError() *GetPartmodelsInternalServerError {
	return &GetPartmodelsInternalServerError{}
}

/*GetPartmodelsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetPartmodelsInternalServerError struct {
}

func (o *GetPartmodelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/partmodels/][%d] getPartmodelsInternalServerError ", 500)
}

func (o *GetPartmodelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPartmodelsServiceUnavailable creates a GetPartmodelsServiceUnavailable with default headers values
func NewGetPartmodelsServiceUnavailable() *GetPartmodelsServiceUnavailable {
	return &GetPartmodelsServiceUnavailable{}
}

/*GetPartmodelsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetPartmodelsServiceUnavailable struct {
}

func (o *GetPartmodelsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/partmodels/][%d] getPartmodelsServiceUnavailable ", 503)
}

func (o *GetPartmodelsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
