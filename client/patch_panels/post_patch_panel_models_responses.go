// Code generated by go-swagger; DO NOT EDIT.

package patch_panels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostPatchPanelModelsReader is a Reader for the PostPatchPanelModels structure.
type PostPatchPanelModelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPatchPanelModelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostPatchPanelModelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPatchPanelModelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostPatchPanelModelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostPatchPanelModelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostPatchPanelModelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostPatchPanelModelsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostPatchPanelModelsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostPatchPanelModelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostPatchPanelModelsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPatchPanelModelsOK creates a PostPatchPanelModelsOK with default headers values
func NewPostPatchPanelModelsOK() *PostPatchPanelModelsOK {
	return &PostPatchPanelModelsOK{}
}

/*PostPatchPanelModelsOK handles this case with default header values.

The above command returns results like this:
*/
type PostPatchPanelModelsOK struct {
	Payload *PostPatchPanelModelsOKBody
}

func (o *PostPatchPanelModelsOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/patch_panel_models/][%d] postPatchPanelModelsOK  %+v", 200, o.Payload)
}

func (o *PostPatchPanelModelsOK) GetPayload() *PostPatchPanelModelsOKBody {
	return o.Payload
}

func (o *PostPatchPanelModelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostPatchPanelModelsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPatchPanelModelsBadRequest creates a PostPatchPanelModelsBadRequest with default headers values
func NewPostPatchPanelModelsBadRequest() *PostPatchPanelModelsBadRequest {
	return &PostPatchPanelModelsBadRequest{}
}

/*PostPatchPanelModelsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostPatchPanelModelsBadRequest struct {
}

func (o *PostPatchPanelModelsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/patch_panel_models/][%d] postPatchPanelModelsBadRequest ", 400)
}

func (o *PostPatchPanelModelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPatchPanelModelsUnauthorized creates a PostPatchPanelModelsUnauthorized with default headers values
func NewPostPatchPanelModelsUnauthorized() *PostPatchPanelModelsUnauthorized {
	return &PostPatchPanelModelsUnauthorized{}
}

/*PostPatchPanelModelsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostPatchPanelModelsUnauthorized struct {
}

func (o *PostPatchPanelModelsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/patch_panel_models/][%d] postPatchPanelModelsUnauthorized ", 401)
}

func (o *PostPatchPanelModelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPatchPanelModelsForbidden creates a PostPatchPanelModelsForbidden with default headers values
func NewPostPatchPanelModelsForbidden() *PostPatchPanelModelsForbidden {
	return &PostPatchPanelModelsForbidden{}
}

/*PostPatchPanelModelsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostPatchPanelModelsForbidden struct {
}

func (o *PostPatchPanelModelsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/patch_panel_models/][%d] postPatchPanelModelsForbidden ", 403)
}

func (o *PostPatchPanelModelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPatchPanelModelsNotFound creates a PostPatchPanelModelsNotFound with default headers values
func NewPostPatchPanelModelsNotFound() *PostPatchPanelModelsNotFound {
	return &PostPatchPanelModelsNotFound{}
}

/*PostPatchPanelModelsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostPatchPanelModelsNotFound struct {
}

func (o *PostPatchPanelModelsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/patch_panel_models/][%d] postPatchPanelModelsNotFound ", 404)
}

func (o *PostPatchPanelModelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPatchPanelModelsMethodNotAllowed creates a PostPatchPanelModelsMethodNotAllowed with default headers values
func NewPostPatchPanelModelsMethodNotAllowed() *PostPatchPanelModelsMethodNotAllowed {
	return &PostPatchPanelModelsMethodNotAllowed{}
}

/*PostPatchPanelModelsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostPatchPanelModelsMethodNotAllowed struct {
}

func (o *PostPatchPanelModelsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/patch_panel_models/][%d] postPatchPanelModelsMethodNotAllowed ", 405)
}

func (o *PostPatchPanelModelsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPatchPanelModelsGone creates a PostPatchPanelModelsGone with default headers values
func NewPostPatchPanelModelsGone() *PostPatchPanelModelsGone {
	return &PostPatchPanelModelsGone{}
}

/*PostPatchPanelModelsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostPatchPanelModelsGone struct {
}

func (o *PostPatchPanelModelsGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/patch_panel_models/][%d] postPatchPanelModelsGone ", 410)
}

func (o *PostPatchPanelModelsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPatchPanelModelsInternalServerError creates a PostPatchPanelModelsInternalServerError with default headers values
func NewPostPatchPanelModelsInternalServerError() *PostPatchPanelModelsInternalServerError {
	return &PostPatchPanelModelsInternalServerError{}
}

/*PostPatchPanelModelsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned ???msg??? from the call.)
*/
type PostPatchPanelModelsInternalServerError struct {
}

func (o *PostPatchPanelModelsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/patch_panel_models/][%d] postPatchPanelModelsInternalServerError ", 500)
}

func (o *PostPatchPanelModelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPatchPanelModelsServiceUnavailable creates a PostPatchPanelModelsServiceUnavailable with default headers values
func NewPostPatchPanelModelsServiceUnavailable() *PostPatchPanelModelsServiceUnavailable {
	return &PostPatchPanelModelsServiceUnavailable{}
}

/*PostPatchPanelModelsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostPatchPanelModelsServiceUnavailable struct {
}

func (o *PostPatchPanelModelsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/patch_panel_models/][%d] postPatchPanelModelsServiceUnavailable ", 503)
}

func (o *PostPatchPanelModelsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostPatchPanelModelsOKBody post patch panel models o k body
swagger:model PostPatchPanelModelsOKBody
*/
type PostPatchPanelModelsOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post patch panel models o k body
func (o *PostPatchPanelModelsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostPatchPanelModelsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostPatchPanelModelsOKBody) UnmarshalBinary(b []byte) error {
	var res PostPatchPanelModelsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
