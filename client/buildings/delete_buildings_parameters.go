// Code generated by go-swagger; DO NOT EDIT.

package buildings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteBuildingsParams creates a new DeleteBuildingsParams object
// with the default values initialized.
func NewDeleteBuildingsParams() *DeleteBuildingsParams {
	var ()
	return &DeleteBuildingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBuildingsParamsWithTimeout creates a new DeleteBuildingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteBuildingsParamsWithTimeout(timeout time.Duration) *DeleteBuildingsParams {
	var ()
	return &DeleteBuildingsParams{

		timeout: timeout,
	}
}

// NewDeleteBuildingsParamsWithContext creates a new DeleteBuildingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteBuildingsParamsWithContext(ctx context.Context) *DeleteBuildingsParams {
	var ()
	return &DeleteBuildingsParams{

		Context: ctx,
	}
}

// NewDeleteBuildingsParamsWithHTTPClient creates a new DeleteBuildingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteBuildingsParamsWithHTTPClient(client *http.Client) *DeleteBuildingsParams {
	var ()
	return &DeleteBuildingsParams{
		HTTPClient: client,
	}
}

/*DeleteBuildingsParams contains all the parameters to send to the API endpoint
for the delete buildings operation typically these are written to a http.Request
*/
type DeleteBuildingsParams struct {

	/*ID
	  building id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete buildings params
func (o *DeleteBuildingsParams) WithTimeout(timeout time.Duration) *DeleteBuildingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete buildings params
func (o *DeleteBuildingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete buildings params
func (o *DeleteBuildingsParams) WithContext(ctx context.Context) *DeleteBuildingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete buildings params
func (o *DeleteBuildingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete buildings params
func (o *DeleteBuildingsParams) WithHTTPClient(client *http.Client) *DeleteBuildingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete buildings params
func (o *DeleteBuildingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete buildings params
func (o *DeleteBuildingsParams) WithID(id int64) *DeleteBuildingsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete buildings params
func (o *DeleteBuildingsParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBuildingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ID
	if err := r.SetPathParam("ID", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}