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
	"github.com/go-openapi/swag"
)

// NewPostDeviceRackParams creates a new PostDeviceRackParams object
// with the default values initialized.
func NewPostDeviceRackParams() *PostDeviceRackParams {
	var ()
	return &PostDeviceRackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDeviceRackParamsWithTimeout creates a new PostDeviceRackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDeviceRackParamsWithTimeout(timeout time.Duration) *PostDeviceRackParams {
	var ()
	return &PostDeviceRackParams{

		timeout: timeout,
	}
}

// NewPostDeviceRackParamsWithContext creates a new PostDeviceRackParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDeviceRackParamsWithContext(ctx context.Context) *PostDeviceRackParams {
	var ()
	return &PostDeviceRackParams{

		Context: ctx,
	}
}

// NewPostDeviceRackParamsWithHTTPClient creates a new PostDeviceRackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDeviceRackParamsWithHTTPClient(client *http.Client) *PostDeviceRackParams {
	var ()
	return &PostDeviceRackParams{
		HTTPClient: client,
	}
}

/*PostDeviceRackParams contains all the parameters to send to the API endpoint
for the post device rack operation typically these are written to a http.Request
*/
type PostDeviceRackParams struct {

	/*AssetNo
	  asset # of the existing device

	*/
	AssetNo *string
	/*Building
	  building is building name, room is room name, rack is rack name and row is optional. This is used if rack_id is not provided and a unique rack is found with that combination. This could be just rack for rack name, if the rack name is unique. Otherwise add row, room or building to identify a unique rack.

	*/
	Building *string
	/*Device
	  name of the new or existing device

	*/
	Device *string
	/*DeviceID
	  Device ID of existing device

	*/
	DeviceID *int64
	/*HwModel
	  If the hw_model doesn’t exist or doesn’t have a type, we will add it as type “regular (rack mountable)” (changed in v6.6.0)

	*/
	HwModel *string
	/*Manufacturer
	  manufacturer of the hardware model. Only for new hardware model being added(added in v6.6.0)

	*/
	Manufacturer *string
	/*Rack*/
	Rack *string
	/*RackID
	  required if building name, room name or rack name are not provided. This is the id of the rack. It can be obtained from Tools > Import > Import Racked Devices.

	*/
	RackID *string
	/*Room*/
	Room *string
	/*Row*/
	Row *string
	/*SerialNo
	  serial # of the existing device

	*/
	SerialNo *string
	/*Size
	  size of the hardware model, only for new hardware model or if hardware model doesn’t have size. required for new hardware model (added in v6.6.0)

	*/
	Size *string
	/*StartAt
	  This is the starting U location for the device in the rack. Starting with v535, you can use “auto” as value to automatically mount the device in next available slot.

	*/
	StartAt string
	/*Where
	  location in rack, one of ‘above’, ‘below’, ‘left’, ‘right’, ‘mounted’. Note: If mounted a size must be provided or available from the hardware model.

	*/
	Where *string
	/*XPos
	  A number between 0 and 2520 representing the position within the u slot. 0 is flush left. 2520 is flush right.

	*/
	XPos *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post device rack params
func (o *PostDeviceRackParams) WithTimeout(timeout time.Duration) *PostDeviceRackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post device rack params
func (o *PostDeviceRackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post device rack params
func (o *PostDeviceRackParams) WithContext(ctx context.Context) *PostDeviceRackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post device rack params
func (o *PostDeviceRackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post device rack params
func (o *PostDeviceRackParams) WithHTTPClient(client *http.Client) *PostDeviceRackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post device rack params
func (o *PostDeviceRackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssetNo adds the assetNo to the post device rack params
func (o *PostDeviceRackParams) WithAssetNo(assetNo *string) *PostDeviceRackParams {
	o.SetAssetNo(assetNo)
	return o
}

// SetAssetNo adds the assetNo to the post device rack params
func (o *PostDeviceRackParams) SetAssetNo(assetNo *string) {
	o.AssetNo = assetNo
}

// WithBuilding adds the building to the post device rack params
func (o *PostDeviceRackParams) WithBuilding(building *string) *PostDeviceRackParams {
	o.SetBuilding(building)
	return o
}

// SetBuilding adds the building to the post device rack params
func (o *PostDeviceRackParams) SetBuilding(building *string) {
	o.Building = building
}

// WithDevice adds the device to the post device rack params
func (o *PostDeviceRackParams) WithDevice(device *string) *PostDeviceRackParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the post device rack params
func (o *PostDeviceRackParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the post device rack params
func (o *PostDeviceRackParams) WithDeviceID(deviceID *int64) *PostDeviceRackParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the post device rack params
func (o *PostDeviceRackParams) SetDeviceID(deviceID *int64) {
	o.DeviceID = deviceID
}

// WithHwModel adds the hwModel to the post device rack params
func (o *PostDeviceRackParams) WithHwModel(hwModel *string) *PostDeviceRackParams {
	o.SetHwModel(hwModel)
	return o
}

// SetHwModel adds the hwModel to the post device rack params
func (o *PostDeviceRackParams) SetHwModel(hwModel *string) {
	o.HwModel = hwModel
}

// WithManufacturer adds the manufacturer to the post device rack params
func (o *PostDeviceRackParams) WithManufacturer(manufacturer *string) *PostDeviceRackParams {
	o.SetManufacturer(manufacturer)
	return o
}

// SetManufacturer adds the manufacturer to the post device rack params
func (o *PostDeviceRackParams) SetManufacturer(manufacturer *string) {
	o.Manufacturer = manufacturer
}

// WithRack adds the rack to the post device rack params
func (o *PostDeviceRackParams) WithRack(rack *string) *PostDeviceRackParams {
	o.SetRack(rack)
	return o
}

// SetRack adds the rack to the post device rack params
func (o *PostDeviceRackParams) SetRack(rack *string) {
	o.Rack = rack
}

// WithRackID adds the rackID to the post device rack params
func (o *PostDeviceRackParams) WithRackID(rackID *string) *PostDeviceRackParams {
	o.SetRackID(rackID)
	return o
}

// SetRackID adds the rackId to the post device rack params
func (o *PostDeviceRackParams) SetRackID(rackID *string) {
	o.RackID = rackID
}

// WithRoom adds the room to the post device rack params
func (o *PostDeviceRackParams) WithRoom(room *string) *PostDeviceRackParams {
	o.SetRoom(room)
	return o
}

// SetRoom adds the room to the post device rack params
func (o *PostDeviceRackParams) SetRoom(room *string) {
	o.Room = room
}

// WithRow adds the row to the post device rack params
func (o *PostDeviceRackParams) WithRow(row *string) *PostDeviceRackParams {
	o.SetRow(row)
	return o
}

// SetRow adds the row to the post device rack params
func (o *PostDeviceRackParams) SetRow(row *string) {
	o.Row = row
}

// WithSerialNo adds the serialNo to the post device rack params
func (o *PostDeviceRackParams) WithSerialNo(serialNo *string) *PostDeviceRackParams {
	o.SetSerialNo(serialNo)
	return o
}

// SetSerialNo adds the serialNo to the post device rack params
func (o *PostDeviceRackParams) SetSerialNo(serialNo *string) {
	o.SerialNo = serialNo
}

// WithSize adds the size to the post device rack params
func (o *PostDeviceRackParams) WithSize(size *string) *PostDeviceRackParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the post device rack params
func (o *PostDeviceRackParams) SetSize(size *string) {
	o.Size = size
}

// WithStartAt adds the startAt to the post device rack params
func (o *PostDeviceRackParams) WithStartAt(startAt string) *PostDeviceRackParams {
	o.SetStartAt(startAt)
	return o
}

// SetStartAt adds the startAt to the post device rack params
func (o *PostDeviceRackParams) SetStartAt(startAt string) {
	o.StartAt = startAt
}

// WithWhere adds the where to the post device rack params
func (o *PostDeviceRackParams) WithWhere(where *string) *PostDeviceRackParams {
	o.SetWhere(where)
	return o
}

// SetWhere adds the where to the post device rack params
func (o *PostDeviceRackParams) SetWhere(where *string) {
	o.Where = where
}

// WithXPos adds the xPos to the post device rack params
func (o *PostDeviceRackParams) WithXPos(xPos *string) *PostDeviceRackParams {
	o.SetXPos(xPos)
	return o
}

// SetXPos adds the xPos to the post device rack params
func (o *PostDeviceRackParams) SetXPos(xPos *string) {
	o.XPos = xPos
}

// WriteToRequest writes these params to a swagger request
func (o *PostDeviceRackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AssetNo != nil {

		// form param asset_no
		var frAssetNo string
		if o.AssetNo != nil {
			frAssetNo = *o.AssetNo
		}
		fAssetNo := frAssetNo
		if fAssetNo != "" {
			if err := r.SetFormParam("asset_no", fAssetNo); err != nil {
				return err
			}
		}

	}

	if o.Building != nil {

		// form param building
		var frBuilding string
		if o.Building != nil {
			frBuilding = *o.Building
		}
		fBuilding := frBuilding
		if fBuilding != "" {
			if err := r.SetFormParam("building", fBuilding); err != nil {
				return err
			}
		}

	}

	if o.Device != nil {

		// form param device
		var frDevice string
		if o.Device != nil {
			frDevice = *o.Device
		}
		fDevice := frDevice
		if fDevice != "" {
			if err := r.SetFormParam("device", fDevice); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// form param device_id
		var frDeviceID int64
		if o.DeviceID != nil {
			frDeviceID = *o.DeviceID
		}
		fDeviceID := swag.FormatInt64(frDeviceID)
		if fDeviceID != "" {
			if err := r.SetFormParam("device_id", fDeviceID); err != nil {
				return err
			}
		}

	}

	if o.HwModel != nil {

		// form param hw_model
		var frHwModel string
		if o.HwModel != nil {
			frHwModel = *o.HwModel
		}
		fHwModel := frHwModel
		if fHwModel != "" {
			if err := r.SetFormParam("hw_model", fHwModel); err != nil {
				return err
			}
		}

	}

	if o.Manufacturer != nil {

		// form param manufacturer
		var frManufacturer string
		if o.Manufacturer != nil {
			frManufacturer = *o.Manufacturer
		}
		fManufacturer := frManufacturer
		if fManufacturer != "" {
			if err := r.SetFormParam("manufacturer", fManufacturer); err != nil {
				return err
			}
		}

	}

	if o.Rack != nil {

		// form param rack
		var frRack string
		if o.Rack != nil {
			frRack = *o.Rack
		}
		fRack := frRack
		if fRack != "" {
			if err := r.SetFormParam("rack", fRack); err != nil {
				return err
			}
		}

	}

	if o.RackID != nil {

		// form param rack_id
		var frRackID string
		if o.RackID != nil {
			frRackID = *o.RackID
		}
		fRackID := frRackID
		if fRackID != "" {
			if err := r.SetFormParam("rack_id", fRackID); err != nil {
				return err
			}
		}

	}

	if o.Room != nil {

		// form param room
		var frRoom string
		if o.Room != nil {
			frRoom = *o.Room
		}
		fRoom := frRoom
		if fRoom != "" {
			if err := r.SetFormParam("room", fRoom); err != nil {
				return err
			}
		}

	}

	if o.Row != nil {

		// form param row
		var frRow string
		if o.Row != nil {
			frRow = *o.Row
		}
		fRow := frRow
		if fRow != "" {
			if err := r.SetFormParam("row", fRow); err != nil {
				return err
			}
		}

	}

	if o.SerialNo != nil {

		// form param serial_no
		var frSerialNo string
		if o.SerialNo != nil {
			frSerialNo = *o.SerialNo
		}
		fSerialNo := frSerialNo
		if fSerialNo != "" {
			if err := r.SetFormParam("serial_no", fSerialNo); err != nil {
				return err
			}
		}

	}

	if o.Size != nil {

		// form param size
		var frSize string
		if o.Size != nil {
			frSize = *o.Size
		}
		fSize := frSize
		if fSize != "" {
			if err := r.SetFormParam("size", fSize); err != nil {
				return err
			}
		}

	}

	// form param start_at
	frStartAt := o.StartAt
	fStartAt := frStartAt
	if fStartAt != "" {
		if err := r.SetFormParam("start_at", fStartAt); err != nil {
			return err
		}
	}

	if o.Where != nil {

		// form param where
		var frWhere string
		if o.Where != nil {
			frWhere = *o.Where
		}
		fWhere := frWhere
		if fWhere != "" {
			if err := r.SetFormParam("where", fWhere); err != nil {
				return err
			}
		}

	}

	if o.XPos != nil {

		// form param x_pos
		var frXPos string
		if o.XPos != nil {
			frXPos = *o.XPos
		}
		fXPos := frXPos
		if fXPos != "" {
			if err := r.SetFormParam("x_pos", fXPos); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}