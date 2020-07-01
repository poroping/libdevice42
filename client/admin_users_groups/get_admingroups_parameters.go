// Code generated by go-swagger; DO NOT EDIT.

package admin_users_groups

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
)

// NewGetAdmingroupsParams creates a new GetAdmingroupsParams object
// with the default values initialized.
func NewGetAdmingroupsParams() *GetAdmingroupsParams {

	return &GetAdmingroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAdmingroupsParamsWithTimeout creates a new GetAdmingroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAdmingroupsParamsWithTimeout(timeout time.Duration) *GetAdmingroupsParams {

	return &GetAdmingroupsParams{

		timeout: timeout,
	}
}

// NewGetAdmingroupsParamsWithContext creates a new GetAdmingroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAdmingroupsParamsWithContext(ctx context.Context) *GetAdmingroupsParams {

	return &GetAdmingroupsParams{

		Context: ctx,
	}
}

// NewGetAdmingroupsParamsWithHTTPClient creates a new GetAdmingroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAdmingroupsParamsWithHTTPClient(client *http.Client) *GetAdmingroupsParams {

	return &GetAdmingroupsParams{
		HTTPClient: client,
	}
}

/*GetAdmingroupsParams contains all the parameters to send to the API endpoint
for the get admingroups operation typically these are written to a http.Request
*/
type GetAdmingroupsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get admingroups params
func (o *GetAdmingroupsParams) WithTimeout(timeout time.Duration) *GetAdmingroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get admingroups params
func (o *GetAdmingroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get admingroups params
func (o *GetAdmingroupsParams) WithContext(ctx context.Context) *GetAdmingroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get admingroups params
func (o *GetAdmingroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get admingroups params
func (o *GetAdmingroupsParams) WithHTTPClient(client *http.Client) *GetAdmingroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get admingroups params
func (o *GetAdmingroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAdmingroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}