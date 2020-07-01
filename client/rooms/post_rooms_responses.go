// Code generated by go-swagger; DO NOT EDIT.

package rooms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostRoomsReader is a Reader for the PostRooms structure.
type PostRoomsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRoomsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRoomsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRoomsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRoomsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRoomsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRoomsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostRoomsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostRoomsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRoomsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostRoomsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRoomsOK creates a PostRoomsOK with default headers values
func NewPostRoomsOK() *PostRoomsOK {
	return &PostRoomsOK{}
}

/*PostRoomsOK handles this case with default header values.

The above command returns results like this:
*/
type PostRoomsOK struct {
	Payload *PostRoomsOKBody
}

func (o *PostRoomsOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/rooms/][%d] postRoomsOK  %+v", 200, o.Payload)
}

func (o *PostRoomsOK) GetPayload() *PostRoomsOKBody {
	return o.Payload
}

func (o *PostRoomsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostRoomsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoomsBadRequest creates a PostRoomsBadRequest with default headers values
func NewPostRoomsBadRequest() *PostRoomsBadRequest {
	return &PostRoomsBadRequest{}
}

/*PostRoomsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostRoomsBadRequest struct {
}

func (o *PostRoomsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/rooms/][%d] postRoomsBadRequest ", 400)
}

func (o *PostRoomsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRoomsUnauthorized creates a PostRoomsUnauthorized with default headers values
func NewPostRoomsUnauthorized() *PostRoomsUnauthorized {
	return &PostRoomsUnauthorized{}
}

/*PostRoomsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostRoomsUnauthorized struct {
}

func (o *PostRoomsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/rooms/][%d] postRoomsUnauthorized ", 401)
}

func (o *PostRoomsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRoomsForbidden creates a PostRoomsForbidden with default headers values
func NewPostRoomsForbidden() *PostRoomsForbidden {
	return &PostRoomsForbidden{}
}

/*PostRoomsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostRoomsForbidden struct {
}

func (o *PostRoomsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/rooms/][%d] postRoomsForbidden ", 403)
}

func (o *PostRoomsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRoomsNotFound creates a PostRoomsNotFound with default headers values
func NewPostRoomsNotFound() *PostRoomsNotFound {
	return &PostRoomsNotFound{}
}

/*PostRoomsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostRoomsNotFound struct {
}

func (o *PostRoomsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/rooms/][%d] postRoomsNotFound ", 404)
}

func (o *PostRoomsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRoomsMethodNotAllowed creates a PostRoomsMethodNotAllowed with default headers values
func NewPostRoomsMethodNotAllowed() *PostRoomsMethodNotAllowed {
	return &PostRoomsMethodNotAllowed{}
}

/*PostRoomsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostRoomsMethodNotAllowed struct {
}

func (o *PostRoomsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/rooms/][%d] postRoomsMethodNotAllowed ", 405)
}

func (o *PostRoomsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRoomsGone creates a PostRoomsGone with default headers values
func NewPostRoomsGone() *PostRoomsGone {
	return &PostRoomsGone{}
}

/*PostRoomsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostRoomsGone struct {
}

func (o *PostRoomsGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/rooms/][%d] postRoomsGone ", 410)
}

func (o *PostRoomsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRoomsInternalServerError creates a PostRoomsInternalServerError with default headers values
func NewPostRoomsInternalServerError() *PostRoomsInternalServerError {
	return &PostRoomsInternalServerError{}
}

/*PostRoomsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type PostRoomsInternalServerError struct {
}

func (o *PostRoomsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/rooms/][%d] postRoomsInternalServerError ", 500)
}

func (o *PostRoomsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRoomsServiceUnavailable creates a PostRoomsServiceUnavailable with default headers values
func NewPostRoomsServiceUnavailable() *PostRoomsServiceUnavailable {
	return &PostRoomsServiceUnavailable{}
}

/*PostRoomsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostRoomsServiceUnavailable struct {
}

func (o *PostRoomsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/rooms/][%d] postRoomsServiceUnavailable ", 503)
}

func (o *PostRoomsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostRoomsOKBody post rooms o k body
swagger:model PostRoomsOKBody
*/
type PostRoomsOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post rooms o k body
func (o *PostRoomsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostRoomsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRoomsOKBody) UnmarshalBinary(b []byte) error {
	var res PostRoomsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}