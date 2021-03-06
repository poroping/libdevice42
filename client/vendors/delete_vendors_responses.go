// Code generated by go-swagger; DO NOT EDIT.

package vendors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteVendorsReader is a Reader for the DeleteVendors structure.
type DeleteVendorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVendorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVendorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVendorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteVendorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteVendorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteVendorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeleteVendorsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewDeleteVendorsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteVendorsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteVendorsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteVendorsOK creates a DeleteVendorsOK with default headers values
func NewDeleteVendorsOK() *DeleteVendorsOK {
	return &DeleteVendorsOK{}
}

/*DeleteVendorsOK handles this case with default header values.

The above command returns results like this:
*/
type DeleteVendorsOK struct {
	Payload *DeleteVendorsOKBody
}

func (o *DeleteVendorsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/vendors/{id}/][%d] deleteVendorsOK  %+v", 200, o.Payload)
}

func (o *DeleteVendorsOK) GetPayload() *DeleteVendorsOKBody {
	return o.Payload
}

func (o *DeleteVendorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteVendorsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVendorsBadRequest creates a DeleteVendorsBadRequest with default headers values
func NewDeleteVendorsBadRequest() *DeleteVendorsBadRequest {
	return &DeleteVendorsBadRequest{}
}

/*DeleteVendorsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type DeleteVendorsBadRequest struct {
}

func (o *DeleteVendorsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/vendors/{id}/][%d] deleteVendorsBadRequest ", 400)
}

func (o *DeleteVendorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVendorsUnauthorized creates a DeleteVendorsUnauthorized with default headers values
func NewDeleteVendorsUnauthorized() *DeleteVendorsUnauthorized {
	return &DeleteVendorsUnauthorized{}
}

/*DeleteVendorsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type DeleteVendorsUnauthorized struct {
}

func (o *DeleteVendorsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/vendors/{id}/][%d] deleteVendorsUnauthorized ", 401)
}

func (o *DeleteVendorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVendorsForbidden creates a DeleteVendorsForbidden with default headers values
func NewDeleteVendorsForbidden() *DeleteVendorsForbidden {
	return &DeleteVendorsForbidden{}
}

/*DeleteVendorsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type DeleteVendorsForbidden struct {
}

func (o *DeleteVendorsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/vendors/{id}/][%d] deleteVendorsForbidden ", 403)
}

func (o *DeleteVendorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVendorsNotFound creates a DeleteVendorsNotFound with default headers values
func NewDeleteVendorsNotFound() *DeleteVendorsNotFound {
	return &DeleteVendorsNotFound{}
}

/*DeleteVendorsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type DeleteVendorsNotFound struct {
}

func (o *DeleteVendorsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/vendors/{id}/][%d] deleteVendorsNotFound ", 404)
}

func (o *DeleteVendorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVendorsMethodNotAllowed creates a DeleteVendorsMethodNotAllowed with default headers values
func NewDeleteVendorsMethodNotAllowed() *DeleteVendorsMethodNotAllowed {
	return &DeleteVendorsMethodNotAllowed{}
}

/*DeleteVendorsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type DeleteVendorsMethodNotAllowed struct {
}

func (o *DeleteVendorsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/vendors/{id}/][%d] deleteVendorsMethodNotAllowed ", 405)
}

func (o *DeleteVendorsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVendorsGone creates a DeleteVendorsGone with default headers values
func NewDeleteVendorsGone() *DeleteVendorsGone {
	return &DeleteVendorsGone{}
}

/*DeleteVendorsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type DeleteVendorsGone struct {
}

func (o *DeleteVendorsGone) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/vendors/{id}/][%d] deleteVendorsGone ", 410)
}

func (o *DeleteVendorsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVendorsInternalServerError creates a DeleteVendorsInternalServerError with default headers values
func NewDeleteVendorsInternalServerError() *DeleteVendorsInternalServerError {
	return &DeleteVendorsInternalServerError{}
}

/*DeleteVendorsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type DeleteVendorsInternalServerError struct {
}

func (o *DeleteVendorsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/vendors/{id}/][%d] deleteVendorsInternalServerError ", 500)
}

func (o *DeleteVendorsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVendorsServiceUnavailable creates a DeleteVendorsServiceUnavailable with default headers values
func NewDeleteVendorsServiceUnavailable() *DeleteVendorsServiceUnavailable {
	return &DeleteVendorsServiceUnavailable{}
}

/*DeleteVendorsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type DeleteVendorsServiceUnavailable struct {
}

func (o *DeleteVendorsServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/1.0/vendors/{id}/][%d] deleteVendorsServiceUnavailable ", 503)
}

func (o *DeleteVendorsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeleteVendorsOKBody delete vendors o k body
swagger:model DeleteVendorsOKBody
*/
type DeleteVendorsOKBody struct {

	// deleted
	Deleted interface{} `json:"deleted,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`
}

// Validate validates this delete vendors o k body
func (o *DeleteVendorsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteVendorsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteVendorsOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteVendorsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
