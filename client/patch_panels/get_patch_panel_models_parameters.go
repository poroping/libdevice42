// Code generated by go-swagger; DO NOT EDIT.

package patch_panels

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

// NewGetPatchPanelModelsParams creates a new GetPatchPanelModelsParams object
// with the default values initialized.
func NewGetPatchPanelModelsParams() *GetPatchPanelModelsParams {
	var ()
	return &GetPatchPanelModelsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPatchPanelModelsParamsWithTimeout creates a new GetPatchPanelModelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPatchPanelModelsParamsWithTimeout(timeout time.Duration) *GetPatchPanelModelsParams {
	var ()
	return &GetPatchPanelModelsParams{

		timeout: timeout,
	}
}

// NewGetPatchPanelModelsParamsWithContext creates a new GetPatchPanelModelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPatchPanelModelsParamsWithContext(ctx context.Context) *GetPatchPanelModelsParams {
	var ()
	return &GetPatchPanelModelsParams{

		Context: ctx,
	}
}

// NewGetPatchPanelModelsParamsWithHTTPClient creates a new GetPatchPanelModelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPatchPanelModelsParamsWithHTTPClient(client *http.Client) *GetPatchPanelModelsParams {
	var ()
	return &GetPatchPanelModelsParams{
		HTTPClient: client,
	}
}

/*GetPatchPanelModelsParams contains all the parameters to send to the API endpoint
for the get patch panel models operation typically these are written to a http.Request
*/
type GetPatchPanelModelsParams struct {

	/*Name
	  filter by name

	*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get patch panel models params
func (o *GetPatchPanelModelsParams) WithTimeout(timeout time.Duration) *GetPatchPanelModelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get patch panel models params
func (o *GetPatchPanelModelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get patch panel models params
func (o *GetPatchPanelModelsParams) WithContext(ctx context.Context) *GetPatchPanelModelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get patch panel models params
func (o *GetPatchPanelModelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get patch panel models params
func (o *GetPatchPanelModelsParams) WithHTTPClient(client *http.Client) *GetPatchPanelModelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get patch panel models params
func (o *GetPatchPanelModelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get patch panel models params
func (o *GetPatchPanelModelsParams) WithName(name *string) *GetPatchPanelModelsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get patch panel models params
func (o *GetPatchPanelModelsParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetPatchPanelModelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}