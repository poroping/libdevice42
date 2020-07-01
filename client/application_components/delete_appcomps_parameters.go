// Code generated by go-swagger; DO NOT EDIT.

package application_components

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

// NewDeleteAppcompsParams creates a new DeleteAppcompsParams object
// with the default values initialized.
func NewDeleteAppcompsParams() *DeleteAppcompsParams {
	var ()
	return &DeleteAppcompsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAppcompsParamsWithTimeout creates a new DeleteAppcompsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAppcompsParamsWithTimeout(timeout time.Duration) *DeleteAppcompsParams {
	var ()
	return &DeleteAppcompsParams{

		timeout: timeout,
	}
}

// NewDeleteAppcompsParamsWithContext creates a new DeleteAppcompsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAppcompsParamsWithContext(ctx context.Context) *DeleteAppcompsParams {
	var ()
	return &DeleteAppcompsParams{

		Context: ctx,
	}
}

// NewDeleteAppcompsParamsWithHTTPClient creates a new DeleteAppcompsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAppcompsParamsWithHTTPClient(client *http.Client) *DeleteAppcompsParams {
	var ()
	return &DeleteAppcompsParams{
		HTTPClient: client,
	}
}

/*DeleteAppcompsParams contains all the parameters to send to the API endpoint
for the delete appcomps operation typically these are written to a http.Request
*/
type DeleteAppcompsParams struct {

	/*AppcompID
	  IP Address id

	*/
	AppcompID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete appcomps params
func (o *DeleteAppcompsParams) WithTimeout(timeout time.Duration) *DeleteAppcompsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete appcomps params
func (o *DeleteAppcompsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete appcomps params
func (o *DeleteAppcompsParams) WithContext(ctx context.Context) *DeleteAppcompsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete appcomps params
func (o *DeleteAppcompsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete appcomps params
func (o *DeleteAppcompsParams) WithHTTPClient(client *http.Client) *DeleteAppcompsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete appcomps params
func (o *DeleteAppcompsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppcompID adds the appcompID to the delete appcomps params
func (o *DeleteAppcompsParams) WithAppcompID(appcompID int64) *DeleteAppcompsParams {
	o.SetAppcompID(appcompID)
	return o
}

// SetAppcompID adds the appcompId to the delete appcomps params
func (o *DeleteAppcompsParams) SetAppcompID(appcompID int64) {
	o.AppcompID = appcompID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAppcompsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appcomp_id
	if err := r.SetPathParam("appcomp_id", swag.FormatInt64(o.AppcompID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}