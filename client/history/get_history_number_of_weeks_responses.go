// Code generated by go-swagger; DO NOT EDIT.

package history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/libdevice42/models"
)

// GetHistoryNumberOfWeeksReader is a Reader for the GetHistoryNumberOfWeeks structure.
type GetHistoryNumberOfWeeksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHistoryNumberOfWeeksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHistoryNumberOfWeeksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetHistoryNumberOfWeeksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetHistoryNumberOfWeeksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetHistoryNumberOfWeeksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetHistoryNumberOfWeeksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetHistoryNumberOfWeeksMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetHistoryNumberOfWeeksGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetHistoryNumberOfWeeksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetHistoryNumberOfWeeksServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetHistoryNumberOfWeeksOK creates a GetHistoryNumberOfWeeksOK with default headers values
func NewGetHistoryNumberOfWeeksOK() *GetHistoryNumberOfWeeksOK {
	return &GetHistoryNumberOfWeeksOK{}
}

/*GetHistoryNumberOfWeeksOK handles this case with default header values.

The above command returns results like this:
*/
type GetHistoryNumberOfWeeksOK struct {
	Payload []*models.GetHistory
}

func (o *GetHistoryNumberOfWeeksOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/history/<number_of_weeks>/][%d] getHistoryNumberOfWeeksOK  %+v", 200, o.Payload)
}

func (o *GetHistoryNumberOfWeeksOK) GetPayload() []*models.GetHistory {
	return o.Payload
}

func (o *GetHistoryNumberOfWeeksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHistoryNumberOfWeeksBadRequest creates a GetHistoryNumberOfWeeksBadRequest with default headers values
func NewGetHistoryNumberOfWeeksBadRequest() *GetHistoryNumberOfWeeksBadRequest {
	return &GetHistoryNumberOfWeeksBadRequest{}
}

/*GetHistoryNumberOfWeeksBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetHistoryNumberOfWeeksBadRequest struct {
}

func (o *GetHistoryNumberOfWeeksBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/history/<number_of_weeks>/][%d] getHistoryNumberOfWeeksBadRequest ", 400)
}

func (o *GetHistoryNumberOfWeeksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHistoryNumberOfWeeksUnauthorized creates a GetHistoryNumberOfWeeksUnauthorized with default headers values
func NewGetHistoryNumberOfWeeksUnauthorized() *GetHistoryNumberOfWeeksUnauthorized {
	return &GetHistoryNumberOfWeeksUnauthorized{}
}

/*GetHistoryNumberOfWeeksUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetHistoryNumberOfWeeksUnauthorized struct {
}

func (o *GetHistoryNumberOfWeeksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/history/<number_of_weeks>/][%d] getHistoryNumberOfWeeksUnauthorized ", 401)
}

func (o *GetHistoryNumberOfWeeksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHistoryNumberOfWeeksForbidden creates a GetHistoryNumberOfWeeksForbidden with default headers values
func NewGetHistoryNumberOfWeeksForbidden() *GetHistoryNumberOfWeeksForbidden {
	return &GetHistoryNumberOfWeeksForbidden{}
}

/*GetHistoryNumberOfWeeksForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetHistoryNumberOfWeeksForbidden struct {
}

func (o *GetHistoryNumberOfWeeksForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/history/<number_of_weeks>/][%d] getHistoryNumberOfWeeksForbidden ", 403)
}

func (o *GetHistoryNumberOfWeeksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHistoryNumberOfWeeksNotFound creates a GetHistoryNumberOfWeeksNotFound with default headers values
func NewGetHistoryNumberOfWeeksNotFound() *GetHistoryNumberOfWeeksNotFound {
	return &GetHistoryNumberOfWeeksNotFound{}
}

/*GetHistoryNumberOfWeeksNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetHistoryNumberOfWeeksNotFound struct {
}

func (o *GetHistoryNumberOfWeeksNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/history/<number_of_weeks>/][%d] getHistoryNumberOfWeeksNotFound ", 404)
}

func (o *GetHistoryNumberOfWeeksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHistoryNumberOfWeeksMethodNotAllowed creates a GetHistoryNumberOfWeeksMethodNotAllowed with default headers values
func NewGetHistoryNumberOfWeeksMethodNotAllowed() *GetHistoryNumberOfWeeksMethodNotAllowed {
	return &GetHistoryNumberOfWeeksMethodNotAllowed{}
}

/*GetHistoryNumberOfWeeksMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetHistoryNumberOfWeeksMethodNotAllowed struct {
}

func (o *GetHistoryNumberOfWeeksMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/history/<number_of_weeks>/][%d] getHistoryNumberOfWeeksMethodNotAllowed ", 405)
}

func (o *GetHistoryNumberOfWeeksMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHistoryNumberOfWeeksGone creates a GetHistoryNumberOfWeeksGone with default headers values
func NewGetHistoryNumberOfWeeksGone() *GetHistoryNumberOfWeeksGone {
	return &GetHistoryNumberOfWeeksGone{}
}

/*GetHistoryNumberOfWeeksGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetHistoryNumberOfWeeksGone struct {
}

func (o *GetHistoryNumberOfWeeksGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/history/<number_of_weeks>/][%d] getHistoryNumberOfWeeksGone ", 410)
}

func (o *GetHistoryNumberOfWeeksGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHistoryNumberOfWeeksInternalServerError creates a GetHistoryNumberOfWeeksInternalServerError with default headers values
func NewGetHistoryNumberOfWeeksInternalServerError() *GetHistoryNumberOfWeeksInternalServerError {
	return &GetHistoryNumberOfWeeksInternalServerError{}
}

/*GetHistoryNumberOfWeeksInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetHistoryNumberOfWeeksInternalServerError struct {
}

func (o *GetHistoryNumberOfWeeksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/history/<number_of_weeks>/][%d] getHistoryNumberOfWeeksInternalServerError ", 500)
}

func (o *GetHistoryNumberOfWeeksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHistoryNumberOfWeeksServiceUnavailable creates a GetHistoryNumberOfWeeksServiceUnavailable with default headers values
func NewGetHistoryNumberOfWeeksServiceUnavailable() *GetHistoryNumberOfWeeksServiceUnavailable {
	return &GetHistoryNumberOfWeeksServiceUnavailable{}
}

/*GetHistoryNumberOfWeeksServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetHistoryNumberOfWeeksServiceUnavailable struct {
}

func (o *GetHistoryNumberOfWeeksServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/history/<number_of_weeks>/][%d] getHistoryNumberOfWeeksServiceUnavailable ", 503)
}

func (o *GetHistoryNumberOfWeeksServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}