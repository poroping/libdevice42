// Code generated by go-swagger; DO NOT EDIT.

package buildings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostBuildingsReader is a Reader for the PostBuildings structure.
type PostBuildingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostBuildingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostBuildingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostBuildingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostBuildingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostBuildingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostBuildingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostBuildingsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostBuildingsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostBuildingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostBuildingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostBuildingsOK creates a PostBuildingsOK with default headers values
func NewPostBuildingsOK() *PostBuildingsOK {
	return &PostBuildingsOK{}
}

/*PostBuildingsOK handles this case with default header values.

The above command returns results like this:
*/
type PostBuildingsOK struct {
	Payload *PostBuildingsOKBody
}

func (o *PostBuildingsOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/buildings/][%d] postBuildingsOK  %+v", 200, o.Payload)
}

func (o *PostBuildingsOK) GetPayload() *PostBuildingsOKBody {
	return o.Payload
}

func (o *PostBuildingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostBuildingsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostBuildingsBadRequest creates a PostBuildingsBadRequest with default headers values
func NewPostBuildingsBadRequest() *PostBuildingsBadRequest {
	return &PostBuildingsBadRequest{}
}

/*PostBuildingsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostBuildingsBadRequest struct {
}

func (o *PostBuildingsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/buildings/][%d] postBuildingsBadRequest ", 400)
}

func (o *PostBuildingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostBuildingsUnauthorized creates a PostBuildingsUnauthorized with default headers values
func NewPostBuildingsUnauthorized() *PostBuildingsUnauthorized {
	return &PostBuildingsUnauthorized{}
}

/*PostBuildingsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostBuildingsUnauthorized struct {
}

func (o *PostBuildingsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/buildings/][%d] postBuildingsUnauthorized ", 401)
}

func (o *PostBuildingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostBuildingsForbidden creates a PostBuildingsForbidden with default headers values
func NewPostBuildingsForbidden() *PostBuildingsForbidden {
	return &PostBuildingsForbidden{}
}

/*PostBuildingsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostBuildingsForbidden struct {
}

func (o *PostBuildingsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/buildings/][%d] postBuildingsForbidden ", 403)
}

func (o *PostBuildingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostBuildingsNotFound creates a PostBuildingsNotFound with default headers values
func NewPostBuildingsNotFound() *PostBuildingsNotFound {
	return &PostBuildingsNotFound{}
}

/*PostBuildingsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostBuildingsNotFound struct {
}

func (o *PostBuildingsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/buildings/][%d] postBuildingsNotFound ", 404)
}

func (o *PostBuildingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostBuildingsMethodNotAllowed creates a PostBuildingsMethodNotAllowed with default headers values
func NewPostBuildingsMethodNotAllowed() *PostBuildingsMethodNotAllowed {
	return &PostBuildingsMethodNotAllowed{}
}

/*PostBuildingsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostBuildingsMethodNotAllowed struct {
}

func (o *PostBuildingsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/buildings/][%d] postBuildingsMethodNotAllowed ", 405)
}

func (o *PostBuildingsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostBuildingsGone creates a PostBuildingsGone with default headers values
func NewPostBuildingsGone() *PostBuildingsGone {
	return &PostBuildingsGone{}
}

/*PostBuildingsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostBuildingsGone struct {
}

func (o *PostBuildingsGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/buildings/][%d] postBuildingsGone ", 410)
}

func (o *PostBuildingsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostBuildingsInternalServerError creates a PostBuildingsInternalServerError with default headers values
func NewPostBuildingsInternalServerError() *PostBuildingsInternalServerError {
	return &PostBuildingsInternalServerError{}
}

/*PostBuildingsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type PostBuildingsInternalServerError struct {
}

func (o *PostBuildingsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/buildings/][%d] postBuildingsInternalServerError ", 500)
}

func (o *PostBuildingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostBuildingsServiceUnavailable creates a PostBuildingsServiceUnavailable with default headers values
func NewPostBuildingsServiceUnavailable() *PostBuildingsServiceUnavailable {
	return &PostBuildingsServiceUnavailable{}
}

/*PostBuildingsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostBuildingsServiceUnavailable struct {
}

func (o *PostBuildingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/buildings/][%d] postBuildingsServiceUnavailable ", 503)
}

func (o *PostBuildingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostBuildingsOKBody post buildings o k body
swagger:model PostBuildingsOKBody
*/
type PostBuildingsOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post buildings o k body
func (o *PostBuildingsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostBuildingsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostBuildingsOKBody) UnmarshalBinary(b []byte) error {
	var res PostBuildingsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
