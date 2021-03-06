// Code generated by go-swagger; DO NOT EDIT.

package cables

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

// NewPostCablesParams creates a new PostCablesParams object
// with the default values initialized.
func NewPostCablesParams() *PostCablesParams {
	var ()
	return &PostCablesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCablesParamsWithTimeout creates a new PostCablesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCablesParamsWithTimeout(timeout time.Duration) *PostCablesParams {
	var ()
	return &PostCablesParams{

		timeout: timeout,
	}
}

// NewPostCablesParamsWithContext creates a new PostCablesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCablesParamsWithContext(ctx context.Context) *PostCablesParams {
	var ()
	return &PostCablesParams{

		Context: ctx,
	}
}

// NewPostCablesParamsWithHTTPClient creates a new PostCablesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCablesParamsWithHTTPClient(client *http.Client) *PostCablesParams {
	var ()
	return &PostCablesParams{
		HTTPClient: client,
	}
}

/*PostCablesParams contains all the parameters to send to the API endpoint
for the post cables operation typically these are written to a http.Request
*/
type PostCablesParams struct {

	/*ID
	  Device42 ID of cable

	*/
	ID string
	/*CableID
	  Cable ID/Name

	*/
	CableID string
	/*CableLength
	  Length of Cable

	*/
	CableLength *string
	/*CableLengthUnits
	  Units for Cable Length (âmâ or âftâ)

	*/
	CableLengthUnits *string
	/*EndPointBackPachPanel*/
	EndPointBackPachPanel *string
	/*EndPointCable
	  Endpoint Cable Type (User Definable)

	*/
	EndPointCable *string
	/*EndPointCableColor
	  Endpoint Cable Color

	*/
	EndPointCableColor *string
	/*EndPointConnectorType
	  Connector Type (User Definable)

	*/
	EndPointConnectorType *string
	/*EndPointID
	  ID of the end point

	*/
	EndPointID *string
	/*EndPointMultiple
	  yes to allow multiple endpoints

	*/
	EndPointMultiple *string
	/*EndPointOpticType
	  Optic Type (Definable, ie multimode)

	*/
	EndPointOpticType *string
	/*EndPointType
	  Type of end point.

	*/
	EndPointType *string
	/*Groups
	  If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. If this parameter is present with no value, all groups are deleted.

	*/
	Groups *string
	/*Notes
	  Any additional notes

	*/
	Notes *string
	/*OriginBackPatchPanel*/
	OriginBackPatchPanel *string
	/*OriginCable
	  Cable Type (User definable)

	*/
	OriginCable *string
	/*OriginCableColor
	  Origin Cable Color

	*/
	OriginCableColor *string
	/*OriginConnectorType
	  Connector Type (User Definable)

	*/
	OriginConnectorType *string
	/*OriginID
	  ID of the origin point

	*/
	OriginID *string
	/*OriginOpticType
	  Optic Type (Definable, ie multimode)

	*/
	OriginOpticType *string
	/*OriginType
	  Type of origin point.

	*/
	OriginType *string
	/*Room
	  Room name

	*/
	Room *string
	/*RoomID
	  Room ID

	*/
	RoomID *string
	/*Vendor
	  Cable vendor

	*/
	Vendor *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post cables params
func (o *PostCablesParams) WithTimeout(timeout time.Duration) *PostCablesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cables params
func (o *PostCablesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cables params
func (o *PostCablesParams) WithContext(ctx context.Context) *PostCablesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cables params
func (o *PostCablesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cables params
func (o *PostCablesParams) WithHTTPClient(client *http.Client) *PostCablesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cables params
func (o *PostCablesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post cables params
func (o *PostCablesParams) WithID(id string) *PostCablesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post cables params
func (o *PostCablesParams) SetID(id string) {
	o.ID = id
}

// WithCableID adds the cableID to the post cables params
func (o *PostCablesParams) WithCableID(cableID string) *PostCablesParams {
	o.SetCableID(cableID)
	return o
}

// SetCableID adds the cableId to the post cables params
func (o *PostCablesParams) SetCableID(cableID string) {
	o.CableID = cableID
}

// WithCableLength adds the cableLength to the post cables params
func (o *PostCablesParams) WithCableLength(cableLength *string) *PostCablesParams {
	o.SetCableLength(cableLength)
	return o
}

// SetCableLength adds the cableLength to the post cables params
func (o *PostCablesParams) SetCableLength(cableLength *string) {
	o.CableLength = cableLength
}

// WithCableLengthUnits adds the cableLengthUnits to the post cables params
func (o *PostCablesParams) WithCableLengthUnits(cableLengthUnits *string) *PostCablesParams {
	o.SetCableLengthUnits(cableLengthUnits)
	return o
}

// SetCableLengthUnits adds the cableLengthUnits to the post cables params
func (o *PostCablesParams) SetCableLengthUnits(cableLengthUnits *string) {
	o.CableLengthUnits = cableLengthUnits
}

// WithEndPointBackPachPanel adds the endPointBackPachPanel to the post cables params
func (o *PostCablesParams) WithEndPointBackPachPanel(endPointBackPachPanel *string) *PostCablesParams {
	o.SetEndPointBackPachPanel(endPointBackPachPanel)
	return o
}

// SetEndPointBackPachPanel adds the endPointBackPachPanel to the post cables params
func (o *PostCablesParams) SetEndPointBackPachPanel(endPointBackPachPanel *string) {
	o.EndPointBackPachPanel = endPointBackPachPanel
}

// WithEndPointCable adds the endPointCable to the post cables params
func (o *PostCablesParams) WithEndPointCable(endPointCable *string) *PostCablesParams {
	o.SetEndPointCable(endPointCable)
	return o
}

// SetEndPointCable adds the endPointCable to the post cables params
func (o *PostCablesParams) SetEndPointCable(endPointCable *string) {
	o.EndPointCable = endPointCable
}

// WithEndPointCableColor adds the endPointCableColor to the post cables params
func (o *PostCablesParams) WithEndPointCableColor(endPointCableColor *string) *PostCablesParams {
	o.SetEndPointCableColor(endPointCableColor)
	return o
}

// SetEndPointCableColor adds the endPointCableColor to the post cables params
func (o *PostCablesParams) SetEndPointCableColor(endPointCableColor *string) {
	o.EndPointCableColor = endPointCableColor
}

// WithEndPointConnectorType adds the endPointConnectorType to the post cables params
func (o *PostCablesParams) WithEndPointConnectorType(endPointConnectorType *string) *PostCablesParams {
	o.SetEndPointConnectorType(endPointConnectorType)
	return o
}

// SetEndPointConnectorType adds the endPointConnectorType to the post cables params
func (o *PostCablesParams) SetEndPointConnectorType(endPointConnectorType *string) {
	o.EndPointConnectorType = endPointConnectorType
}

// WithEndPointID adds the endPointID to the post cables params
func (o *PostCablesParams) WithEndPointID(endPointID *string) *PostCablesParams {
	o.SetEndPointID(endPointID)
	return o
}

// SetEndPointID adds the endPointId to the post cables params
func (o *PostCablesParams) SetEndPointID(endPointID *string) {
	o.EndPointID = endPointID
}

// WithEndPointMultiple adds the endPointMultiple to the post cables params
func (o *PostCablesParams) WithEndPointMultiple(endPointMultiple *string) *PostCablesParams {
	o.SetEndPointMultiple(endPointMultiple)
	return o
}

// SetEndPointMultiple adds the endPointMultiple to the post cables params
func (o *PostCablesParams) SetEndPointMultiple(endPointMultiple *string) {
	o.EndPointMultiple = endPointMultiple
}

// WithEndPointOpticType adds the endPointOpticType to the post cables params
func (o *PostCablesParams) WithEndPointOpticType(endPointOpticType *string) *PostCablesParams {
	o.SetEndPointOpticType(endPointOpticType)
	return o
}

// SetEndPointOpticType adds the endPointOpticType to the post cables params
func (o *PostCablesParams) SetEndPointOpticType(endPointOpticType *string) {
	o.EndPointOpticType = endPointOpticType
}

// WithEndPointType adds the endPointType to the post cables params
func (o *PostCablesParams) WithEndPointType(endPointType *string) *PostCablesParams {
	o.SetEndPointType(endPointType)
	return o
}

// SetEndPointType adds the endPointType to the post cables params
func (o *PostCablesParams) SetEndPointType(endPointType *string) {
	o.EndPointType = endPointType
}

// WithGroups adds the groups to the post cables params
func (o *PostCablesParams) WithGroups(groups *string) *PostCablesParams {
	o.SetGroups(groups)
	return o
}

// SetGroups adds the groups to the post cables params
func (o *PostCablesParams) SetGroups(groups *string) {
	o.Groups = groups
}

// WithNotes adds the notes to the post cables params
func (o *PostCablesParams) WithNotes(notes *string) *PostCablesParams {
	o.SetNotes(notes)
	return o
}

// SetNotes adds the notes to the post cables params
func (o *PostCablesParams) SetNotes(notes *string) {
	o.Notes = notes
}

// WithOriginBackPatchPanel adds the originBackPatchPanel to the post cables params
func (o *PostCablesParams) WithOriginBackPatchPanel(originBackPatchPanel *string) *PostCablesParams {
	o.SetOriginBackPatchPanel(originBackPatchPanel)
	return o
}

// SetOriginBackPatchPanel adds the originBackPatchPanel to the post cables params
func (o *PostCablesParams) SetOriginBackPatchPanel(originBackPatchPanel *string) {
	o.OriginBackPatchPanel = originBackPatchPanel
}

// WithOriginCable adds the originCable to the post cables params
func (o *PostCablesParams) WithOriginCable(originCable *string) *PostCablesParams {
	o.SetOriginCable(originCable)
	return o
}

// SetOriginCable adds the originCable to the post cables params
func (o *PostCablesParams) SetOriginCable(originCable *string) {
	o.OriginCable = originCable
}

// WithOriginCableColor adds the originCableColor to the post cables params
func (o *PostCablesParams) WithOriginCableColor(originCableColor *string) *PostCablesParams {
	o.SetOriginCableColor(originCableColor)
	return o
}

// SetOriginCableColor adds the originCableColor to the post cables params
func (o *PostCablesParams) SetOriginCableColor(originCableColor *string) {
	o.OriginCableColor = originCableColor
}

// WithOriginConnectorType adds the originConnectorType to the post cables params
func (o *PostCablesParams) WithOriginConnectorType(originConnectorType *string) *PostCablesParams {
	o.SetOriginConnectorType(originConnectorType)
	return o
}

// SetOriginConnectorType adds the originConnectorType to the post cables params
func (o *PostCablesParams) SetOriginConnectorType(originConnectorType *string) {
	o.OriginConnectorType = originConnectorType
}

// WithOriginID adds the originID to the post cables params
func (o *PostCablesParams) WithOriginID(originID *string) *PostCablesParams {
	o.SetOriginID(originID)
	return o
}

// SetOriginID adds the originId to the post cables params
func (o *PostCablesParams) SetOriginID(originID *string) {
	o.OriginID = originID
}

// WithOriginOpticType adds the originOpticType to the post cables params
func (o *PostCablesParams) WithOriginOpticType(originOpticType *string) *PostCablesParams {
	o.SetOriginOpticType(originOpticType)
	return o
}

// SetOriginOpticType adds the originOpticType to the post cables params
func (o *PostCablesParams) SetOriginOpticType(originOpticType *string) {
	o.OriginOpticType = originOpticType
}

// WithOriginType adds the originType to the post cables params
func (o *PostCablesParams) WithOriginType(originType *string) *PostCablesParams {
	o.SetOriginType(originType)
	return o
}

// SetOriginType adds the originType to the post cables params
func (o *PostCablesParams) SetOriginType(originType *string) {
	o.OriginType = originType
}

// WithRoom adds the room to the post cables params
func (o *PostCablesParams) WithRoom(room *string) *PostCablesParams {
	o.SetRoom(room)
	return o
}

// SetRoom adds the room to the post cables params
func (o *PostCablesParams) SetRoom(room *string) {
	o.Room = room
}

// WithRoomID adds the roomID to the post cables params
func (o *PostCablesParams) WithRoomID(roomID *string) *PostCablesParams {
	o.SetRoomID(roomID)
	return o
}

// SetRoomID adds the roomId to the post cables params
func (o *PostCablesParams) SetRoomID(roomID *string) {
	o.RoomID = roomID
}

// WithVendor adds the vendor to the post cables params
func (o *PostCablesParams) WithVendor(vendor *string) *PostCablesParams {
	o.SetVendor(vendor)
	return o
}

// SetVendor adds the vendor to the post cables params
func (o *PostCablesParams) SetVendor(vendor *string) {
	o.Vendor = vendor
}

// WriteToRequest writes these params to a swagger request
func (o *PostCablesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param ID
	frID := o.ID
	fID := frID
	if fID != "" {
		if err := r.SetFormParam("ID", fID); err != nil {
			return err
		}
	}

	// form param cable_id
	frCableID := o.CableID
	fCableID := frCableID
	if fCableID != "" {
		if err := r.SetFormParam("cable_id", fCableID); err != nil {
			return err
		}
	}

	if o.CableLength != nil {

		// form param cable_length
		var frCableLength string
		if o.CableLength != nil {
			frCableLength = *o.CableLength
		}
		fCableLength := frCableLength
		if fCableLength != "" {
			if err := r.SetFormParam("cable_length", fCableLength); err != nil {
				return err
			}
		}

	}

	if o.CableLengthUnits != nil {

		// form param cable_length_units
		var frCableLengthUnits string
		if o.CableLengthUnits != nil {
			frCableLengthUnits = *o.CableLengthUnits
		}
		fCableLengthUnits := frCableLengthUnits
		if fCableLengthUnits != "" {
			if err := r.SetFormParam("cable_length_units", fCableLengthUnits); err != nil {
				return err
			}
		}

	}

	if o.EndPointBackPachPanel != nil {

		// form param end_point_back_pach_panel
		var frEndPointBackPachPanel string
		if o.EndPointBackPachPanel != nil {
			frEndPointBackPachPanel = *o.EndPointBackPachPanel
		}
		fEndPointBackPachPanel := frEndPointBackPachPanel
		if fEndPointBackPachPanel != "" {
			if err := r.SetFormParam("end_point_back_pach_panel", fEndPointBackPachPanel); err != nil {
				return err
			}
		}

	}

	if o.EndPointCable != nil {

		// form param end_point_cable
		var frEndPointCable string
		if o.EndPointCable != nil {
			frEndPointCable = *o.EndPointCable
		}
		fEndPointCable := frEndPointCable
		if fEndPointCable != "" {
			if err := r.SetFormParam("end_point_cable", fEndPointCable); err != nil {
				return err
			}
		}

	}

	if o.EndPointCableColor != nil {

		// form param end_point_cable_color
		var frEndPointCableColor string
		if o.EndPointCableColor != nil {
			frEndPointCableColor = *o.EndPointCableColor
		}
		fEndPointCableColor := frEndPointCableColor
		if fEndPointCableColor != "" {
			if err := r.SetFormParam("end_point_cable_color", fEndPointCableColor); err != nil {
				return err
			}
		}

	}

	if o.EndPointConnectorType != nil {

		// form param end_point_connector_type
		var frEndPointConnectorType string
		if o.EndPointConnectorType != nil {
			frEndPointConnectorType = *o.EndPointConnectorType
		}
		fEndPointConnectorType := frEndPointConnectorType
		if fEndPointConnectorType != "" {
			if err := r.SetFormParam("end_point_connector_type", fEndPointConnectorType); err != nil {
				return err
			}
		}

	}

	if o.EndPointID != nil {

		// form param end_point_id
		var frEndPointID string
		if o.EndPointID != nil {
			frEndPointID = *o.EndPointID
		}
		fEndPointID := frEndPointID
		if fEndPointID != "" {
			if err := r.SetFormParam("end_point_id", fEndPointID); err != nil {
				return err
			}
		}

	}

	if o.EndPointMultiple != nil {

		// form param end_point_multiple
		var frEndPointMultiple string
		if o.EndPointMultiple != nil {
			frEndPointMultiple = *o.EndPointMultiple
		}
		fEndPointMultiple := frEndPointMultiple
		if fEndPointMultiple != "" {
			if err := r.SetFormParam("end_point_multiple", fEndPointMultiple); err != nil {
				return err
			}
		}

	}

	if o.EndPointOpticType != nil {

		// form param end_point_optic_type
		var frEndPointOpticType string
		if o.EndPointOpticType != nil {
			frEndPointOpticType = *o.EndPointOpticType
		}
		fEndPointOpticType := frEndPointOpticType
		if fEndPointOpticType != "" {
			if err := r.SetFormParam("end_point_optic_type", fEndPointOpticType); err != nil {
				return err
			}
		}

	}

	if o.EndPointType != nil {

		// form param end_point_type
		var frEndPointType string
		if o.EndPointType != nil {
			frEndPointType = *o.EndPointType
		}
		fEndPointType := frEndPointType
		if fEndPointType != "" {
			if err := r.SetFormParam("end_point_type", fEndPointType); err != nil {
				return err
			}
		}

	}

	if o.Groups != nil {

		// form param groups
		var frGroups string
		if o.Groups != nil {
			frGroups = *o.Groups
		}
		fGroups := frGroups
		if fGroups != "" {
			if err := r.SetFormParam("groups", fGroups); err != nil {
				return err
			}
		}

	}

	if o.Notes != nil {

		// form param notes
		var frNotes string
		if o.Notes != nil {
			frNotes = *o.Notes
		}
		fNotes := frNotes
		if fNotes != "" {
			if err := r.SetFormParam("notes", fNotes); err != nil {
				return err
			}
		}

	}

	if o.OriginBackPatchPanel != nil {

		// form param origin_back_patch_panel
		var frOriginBackPatchPanel string
		if o.OriginBackPatchPanel != nil {
			frOriginBackPatchPanel = *o.OriginBackPatchPanel
		}
		fOriginBackPatchPanel := frOriginBackPatchPanel
		if fOriginBackPatchPanel != "" {
			if err := r.SetFormParam("origin_back_patch_panel", fOriginBackPatchPanel); err != nil {
				return err
			}
		}

	}

	if o.OriginCable != nil {

		// form param origin_cable
		var frOriginCable string
		if o.OriginCable != nil {
			frOriginCable = *o.OriginCable
		}
		fOriginCable := frOriginCable
		if fOriginCable != "" {
			if err := r.SetFormParam("origin_cable", fOriginCable); err != nil {
				return err
			}
		}

	}

	if o.OriginCableColor != nil {

		// form param origin_cable_color
		var frOriginCableColor string
		if o.OriginCableColor != nil {
			frOriginCableColor = *o.OriginCableColor
		}
		fOriginCableColor := frOriginCableColor
		if fOriginCableColor != "" {
			if err := r.SetFormParam("origin_cable_color", fOriginCableColor); err != nil {
				return err
			}
		}

	}

	if o.OriginConnectorType != nil {

		// form param origin_connector_type
		var frOriginConnectorType string
		if o.OriginConnectorType != nil {
			frOriginConnectorType = *o.OriginConnectorType
		}
		fOriginConnectorType := frOriginConnectorType
		if fOriginConnectorType != "" {
			if err := r.SetFormParam("origin_connector_type", fOriginConnectorType); err != nil {
				return err
			}
		}

	}

	if o.OriginID != nil {

		// form param origin_id
		var frOriginID string
		if o.OriginID != nil {
			frOriginID = *o.OriginID
		}
		fOriginID := frOriginID
		if fOriginID != "" {
			if err := r.SetFormParam("origin_id", fOriginID); err != nil {
				return err
			}
		}

	}

	if o.OriginOpticType != nil {

		// form param origin_optic_type
		var frOriginOpticType string
		if o.OriginOpticType != nil {
			frOriginOpticType = *o.OriginOpticType
		}
		fOriginOpticType := frOriginOpticType
		if fOriginOpticType != "" {
			if err := r.SetFormParam("origin_optic_type", fOriginOpticType); err != nil {
				return err
			}
		}

	}

	if o.OriginType != nil {

		// form param origin_type
		var frOriginType string
		if o.OriginType != nil {
			frOriginType = *o.OriginType
		}
		fOriginType := frOriginType
		if fOriginType != "" {
			if err := r.SetFormParam("origin_type", fOriginType); err != nil {
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

	if o.RoomID != nil {

		// form param room_id
		var frRoomID string
		if o.RoomID != nil {
			frRoomID = *o.RoomID
		}
		fRoomID := frRoomID
		if fRoomID != "" {
			if err := r.SetFormParam("room_id", fRoomID); err != nil {
				return err
			}
		}

	}

	if o.Vendor != nil {

		// form param vendor
		var frVendor string
		if o.Vendor != nil {
			frVendor = *o.Vendor
		}
		fVendor := frVendor
		if fVendor != "" {
			if err := r.SetFormParam("vendor", fVendor); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
