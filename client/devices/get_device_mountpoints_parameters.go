// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// NewGetDeviceMountpointsParams creates a new GetDeviceMountpointsParams object
// with the default values initialized.
func NewGetDeviceMountpointsParams() *GetDeviceMountpointsParams {
	var ()
	return &GetDeviceMountpointsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceMountpointsParamsWithTimeout creates a new GetDeviceMountpointsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceMountpointsParamsWithTimeout(timeout time.Duration) *GetDeviceMountpointsParams {
	var ()
	return &GetDeviceMountpointsParams{

		timeout: timeout,
	}
}

// NewGetDeviceMountpointsParamsWithContext creates a new GetDeviceMountpointsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceMountpointsParamsWithContext(ctx context.Context) *GetDeviceMountpointsParams {
	var ()
	return &GetDeviceMountpointsParams{

		Context: ctx,
	}
}

// NewGetDeviceMountpointsParamsWithHTTPClient creates a new GetDeviceMountpointsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceMountpointsParamsWithHTTPClient(client *http.Client) *GetDeviceMountpointsParams {
	var ()
	return &GetDeviceMountpointsParams{
		HTTPClient: client,
	}
}

/*GetDeviceMountpointsParams contains all the parameters to send to the API endpoint
for the get device mountpoints operation typically these are written to a http.Request
*/
type GetDeviceMountpointsParams struct {

	/*DeviceID
	  id of the device

	*/
	DeviceID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device mountpoints params
func (o *GetDeviceMountpointsParams) WithTimeout(timeout time.Duration) *GetDeviceMountpointsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device mountpoints params
func (o *GetDeviceMountpointsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device mountpoints params
func (o *GetDeviceMountpointsParams) WithContext(ctx context.Context) *GetDeviceMountpointsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device mountpoints params
func (o *GetDeviceMountpointsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device mountpoints params
func (o *GetDeviceMountpointsParams) WithHTTPClient(client *http.Client) *GetDeviceMountpointsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device mountpoints params
func (o *GetDeviceMountpointsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get device mountpoints params
func (o *GetDeviceMountpointsParams) WithDeviceID(deviceID *string) *GetDeviceMountpointsParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get device mountpoints params
func (o *GetDeviceMountpointsParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceMountpointsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID string
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := qrDeviceID
		if qDeviceID != "" {
			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}