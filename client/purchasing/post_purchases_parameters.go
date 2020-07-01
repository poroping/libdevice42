// Code generated by go-swagger; DO NOT EDIT.

package purchasing

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

// NewPostPurchasesParams creates a new PostPurchasesParams object
// with the default values initialized.
func NewPostPurchasesParams() *PostPurchasesParams {
	var ()
	return &PostPurchasesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPurchasesParamsWithTimeout creates a new PostPurchasesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPurchasesParamsWithTimeout(timeout time.Duration) *PostPurchasesParams {
	var ()
	return &PostPurchasesParams{

		timeout: timeout,
	}
}

// NewPostPurchasesParamsWithContext creates a new PostPurchasesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPurchasesParamsWithContext(ctx context.Context) *PostPurchasesParams {
	var ()
	return &PostPurchasesParams{

		Context: ctx,
	}
}

// NewPostPurchasesParamsWithHTTPClient creates a new PostPurchasesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPurchasesParamsWithHTTPClient(client *http.Client) *PostPurchasesParams {
	var ()
	return &PostPurchasesParams{
		HTTPClient: client,
	}
}

/*PostPurchasesParams contains all the parameters to send to the API endpoint
for the post purchases operation typically these are written to a http.Request
*/
type PostPurchasesParams struct {

	/*Completed*/
	Completed *string
	/*Cost*/
	Cost *string
	/*CostCenter*/
	CostCenter *string
	/*Groups
	  If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. If this parameter is present with no value, all groups are deleted.

	*/
	Groups *string
	/*LineAssetIds
	  Comma separated asset_id. Only applicable if line_item_type is asset.

	*/
	LineAssetIds *string
	/*LineBuildingIds
	  D42 ID for the building(s)

	*/
	LineBuildingIds *string
	/*LineCancelPolicy*/
	LineCancelPolicy *string
	/*LineCertificateIds
	  D42 ID for the certificate(s)

	*/
	LineCertificateIds *string
	/*LineCircuitIds
	  D42 ID for the circuit(s)

	*/
	LineCircuitIds *string
	/*LineCircuits
	  circuit ID name

	*/
	LineCircuits *string
	/*LineCompleted*/
	LineCompleted *string
	/*LineContractID
	  (added in v9.0.2)

	*/
	LineContractID *string
	/*LineContractType*/
	LineContractType *string
	/*LineCost
	  cost for single object / item.

	*/
	LineCost *string
	/*LineCostCenter*/
	LineCostCenter *string
	/*LineCustomer*/
	LineCustomer *string
	/*LineDeviceAssetNos
	  Comma separated asset numbers. Only applicable if line_item_type is device. Will only work on existing asset numbers.

	*/
	LineDeviceAssetNos *string
	/*LineDeviceOsIds
	  D42 ID for the Device OS

	*/
	LineDeviceOsIds *string
	/*LineDeviceSerialNos
	  Comma separated serial numbers. Only applicable if line_item_type is device. Will only work on existing serial numbers.

	*/
	LineDeviceSerialNos *string
	/*LineDevices
	  Comma separated device names. Only applicable if line_item_type is device. Will create new devices if device with name specific here does not exist.

	*/
	LineDevices *string
	/*LineEndDate
	  Date in YYYY-MM-DD format

	*/
	LineEndDate *string
	/*LineFrequency*/
	LineFrequency *string
	/*LineItemType
	  Default is device.

	*/
	LineItemType *string
	/*LineName*/
	LineName *string
	/*LineNo
	  required for existing line items, use existing line # to change existing line item.

	*/
	LineNo *string
	/*LineNotes*/
	LineNotes *string
	/*LinePartIds
	  D42 ID for the part(s)

	*/
	LinePartIds *string
	/*LineQuantity
	  can be calculated automatically from # of objects associated

	*/
	LineQuantity *string
	/*LineRackIds
	  D42 ID for the rack(s)

	*/
	LineRackIds *string
	/*LineRenewDate
	  Date in YYYY-MM-DD format

	*/
	LineRenewDate *string
	/*LineRoomIds
	  D42 ID for the room(s)

	*/
	LineRoomIds *string
	/*LineServiceType
	  new service type will be created, if it doesn’t exist (added in v9.0.2)

	*/
	LineServiceType *string
	/*LineSoftwareIds
	  D42 ID for the software

	*/
	LineSoftwareIds *string
	/*LineStartDate
	  Date in YYYY-MM-DD format

	*/
	LineStartDate *string
	/*LineType
	  required for any new line being added for both device or contract.

	*/
	LineType *string
	/*Notes
	  Any additional notes

	*/
	Notes *string
	/*OrderNo
	  order number / name for the purchase. Can be new or existing.

	*/
	OrderNo *string
	/*PoDate*/
	PoDate *string
	/*PurchaseID
	  Can be used instead of order_no to update existing purchase

	*/
	PurchaseID *string
	/*Vendor
	  The cloud vendor

	*/
	Vendor *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post purchases params
func (o *PostPurchasesParams) WithTimeout(timeout time.Duration) *PostPurchasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post purchases params
func (o *PostPurchasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post purchases params
func (o *PostPurchasesParams) WithContext(ctx context.Context) *PostPurchasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post purchases params
func (o *PostPurchasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post purchases params
func (o *PostPurchasesParams) WithHTTPClient(client *http.Client) *PostPurchasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post purchases params
func (o *PostPurchasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompleted adds the completed to the post purchases params
func (o *PostPurchasesParams) WithCompleted(completed *string) *PostPurchasesParams {
	o.SetCompleted(completed)
	return o
}

// SetCompleted adds the completed to the post purchases params
func (o *PostPurchasesParams) SetCompleted(completed *string) {
	o.Completed = completed
}

// WithCost adds the cost to the post purchases params
func (o *PostPurchasesParams) WithCost(cost *string) *PostPurchasesParams {
	o.SetCost(cost)
	return o
}

// SetCost adds the cost to the post purchases params
func (o *PostPurchasesParams) SetCost(cost *string) {
	o.Cost = cost
}

// WithCostCenter adds the costCenter to the post purchases params
func (o *PostPurchasesParams) WithCostCenter(costCenter *string) *PostPurchasesParams {
	o.SetCostCenter(costCenter)
	return o
}

// SetCostCenter adds the costCenter to the post purchases params
func (o *PostPurchasesParams) SetCostCenter(costCenter *string) {
	o.CostCenter = costCenter
}

// WithGroups adds the groups to the post purchases params
func (o *PostPurchasesParams) WithGroups(groups *string) *PostPurchasesParams {
	o.SetGroups(groups)
	return o
}

// SetGroups adds the groups to the post purchases params
func (o *PostPurchasesParams) SetGroups(groups *string) {
	o.Groups = groups
}

// WithLineAssetIds adds the lineAssetIds to the post purchases params
func (o *PostPurchasesParams) WithLineAssetIds(lineAssetIds *string) *PostPurchasesParams {
	o.SetLineAssetIds(lineAssetIds)
	return o
}

// SetLineAssetIds adds the lineAssetIds to the post purchases params
func (o *PostPurchasesParams) SetLineAssetIds(lineAssetIds *string) {
	o.LineAssetIds = lineAssetIds
}

// WithLineBuildingIds adds the lineBuildingIds to the post purchases params
func (o *PostPurchasesParams) WithLineBuildingIds(lineBuildingIds *string) *PostPurchasesParams {
	o.SetLineBuildingIds(lineBuildingIds)
	return o
}

// SetLineBuildingIds adds the lineBuildingIds to the post purchases params
func (o *PostPurchasesParams) SetLineBuildingIds(lineBuildingIds *string) {
	o.LineBuildingIds = lineBuildingIds
}

// WithLineCancelPolicy adds the lineCancelPolicy to the post purchases params
func (o *PostPurchasesParams) WithLineCancelPolicy(lineCancelPolicy *string) *PostPurchasesParams {
	o.SetLineCancelPolicy(lineCancelPolicy)
	return o
}

// SetLineCancelPolicy adds the lineCancelPolicy to the post purchases params
func (o *PostPurchasesParams) SetLineCancelPolicy(lineCancelPolicy *string) {
	o.LineCancelPolicy = lineCancelPolicy
}

// WithLineCertificateIds adds the lineCertificateIds to the post purchases params
func (o *PostPurchasesParams) WithLineCertificateIds(lineCertificateIds *string) *PostPurchasesParams {
	o.SetLineCertificateIds(lineCertificateIds)
	return o
}

// SetLineCertificateIds adds the lineCertificateIds to the post purchases params
func (o *PostPurchasesParams) SetLineCertificateIds(lineCertificateIds *string) {
	o.LineCertificateIds = lineCertificateIds
}

// WithLineCircuitIds adds the lineCircuitIds to the post purchases params
func (o *PostPurchasesParams) WithLineCircuitIds(lineCircuitIds *string) *PostPurchasesParams {
	o.SetLineCircuitIds(lineCircuitIds)
	return o
}

// SetLineCircuitIds adds the lineCircuitIds to the post purchases params
func (o *PostPurchasesParams) SetLineCircuitIds(lineCircuitIds *string) {
	o.LineCircuitIds = lineCircuitIds
}

// WithLineCircuits adds the lineCircuits to the post purchases params
func (o *PostPurchasesParams) WithLineCircuits(lineCircuits *string) *PostPurchasesParams {
	o.SetLineCircuits(lineCircuits)
	return o
}

// SetLineCircuits adds the lineCircuits to the post purchases params
func (o *PostPurchasesParams) SetLineCircuits(lineCircuits *string) {
	o.LineCircuits = lineCircuits
}

// WithLineCompleted adds the lineCompleted to the post purchases params
func (o *PostPurchasesParams) WithLineCompleted(lineCompleted *string) *PostPurchasesParams {
	o.SetLineCompleted(lineCompleted)
	return o
}

// SetLineCompleted adds the lineCompleted to the post purchases params
func (o *PostPurchasesParams) SetLineCompleted(lineCompleted *string) {
	o.LineCompleted = lineCompleted
}

// WithLineContractID adds the lineContractID to the post purchases params
func (o *PostPurchasesParams) WithLineContractID(lineContractID *string) *PostPurchasesParams {
	o.SetLineContractID(lineContractID)
	return o
}

// SetLineContractID adds the lineContractId to the post purchases params
func (o *PostPurchasesParams) SetLineContractID(lineContractID *string) {
	o.LineContractID = lineContractID
}

// WithLineContractType adds the lineContractType to the post purchases params
func (o *PostPurchasesParams) WithLineContractType(lineContractType *string) *PostPurchasesParams {
	o.SetLineContractType(lineContractType)
	return o
}

// SetLineContractType adds the lineContractType to the post purchases params
func (o *PostPurchasesParams) SetLineContractType(lineContractType *string) {
	o.LineContractType = lineContractType
}

// WithLineCost adds the lineCost to the post purchases params
func (o *PostPurchasesParams) WithLineCost(lineCost *string) *PostPurchasesParams {
	o.SetLineCost(lineCost)
	return o
}

// SetLineCost adds the lineCost to the post purchases params
func (o *PostPurchasesParams) SetLineCost(lineCost *string) {
	o.LineCost = lineCost
}

// WithLineCostCenter adds the lineCostCenter to the post purchases params
func (o *PostPurchasesParams) WithLineCostCenter(lineCostCenter *string) *PostPurchasesParams {
	o.SetLineCostCenter(lineCostCenter)
	return o
}

// SetLineCostCenter adds the lineCostCenter to the post purchases params
func (o *PostPurchasesParams) SetLineCostCenter(lineCostCenter *string) {
	o.LineCostCenter = lineCostCenter
}

// WithLineCustomer adds the lineCustomer to the post purchases params
func (o *PostPurchasesParams) WithLineCustomer(lineCustomer *string) *PostPurchasesParams {
	o.SetLineCustomer(lineCustomer)
	return o
}

// SetLineCustomer adds the lineCustomer to the post purchases params
func (o *PostPurchasesParams) SetLineCustomer(lineCustomer *string) {
	o.LineCustomer = lineCustomer
}

// WithLineDeviceAssetNos adds the lineDeviceAssetNos to the post purchases params
func (o *PostPurchasesParams) WithLineDeviceAssetNos(lineDeviceAssetNos *string) *PostPurchasesParams {
	o.SetLineDeviceAssetNos(lineDeviceAssetNos)
	return o
}

// SetLineDeviceAssetNos adds the lineDeviceAssetNos to the post purchases params
func (o *PostPurchasesParams) SetLineDeviceAssetNos(lineDeviceAssetNos *string) {
	o.LineDeviceAssetNos = lineDeviceAssetNos
}

// WithLineDeviceOsIds adds the lineDeviceOsIds to the post purchases params
func (o *PostPurchasesParams) WithLineDeviceOsIds(lineDeviceOsIds *string) *PostPurchasesParams {
	o.SetLineDeviceOsIds(lineDeviceOsIds)
	return o
}

// SetLineDeviceOsIds adds the lineDeviceOsIds to the post purchases params
func (o *PostPurchasesParams) SetLineDeviceOsIds(lineDeviceOsIds *string) {
	o.LineDeviceOsIds = lineDeviceOsIds
}

// WithLineDeviceSerialNos adds the lineDeviceSerialNos to the post purchases params
func (o *PostPurchasesParams) WithLineDeviceSerialNos(lineDeviceSerialNos *string) *PostPurchasesParams {
	o.SetLineDeviceSerialNos(lineDeviceSerialNos)
	return o
}

// SetLineDeviceSerialNos adds the lineDeviceSerialNos to the post purchases params
func (o *PostPurchasesParams) SetLineDeviceSerialNos(lineDeviceSerialNos *string) {
	o.LineDeviceSerialNos = lineDeviceSerialNos
}

// WithLineDevices adds the lineDevices to the post purchases params
func (o *PostPurchasesParams) WithLineDevices(lineDevices *string) *PostPurchasesParams {
	o.SetLineDevices(lineDevices)
	return o
}

// SetLineDevices adds the lineDevices to the post purchases params
func (o *PostPurchasesParams) SetLineDevices(lineDevices *string) {
	o.LineDevices = lineDevices
}

// WithLineEndDate adds the lineEndDate to the post purchases params
func (o *PostPurchasesParams) WithLineEndDate(lineEndDate *string) *PostPurchasesParams {
	o.SetLineEndDate(lineEndDate)
	return o
}

// SetLineEndDate adds the lineEndDate to the post purchases params
func (o *PostPurchasesParams) SetLineEndDate(lineEndDate *string) {
	o.LineEndDate = lineEndDate
}

// WithLineFrequency adds the lineFrequency to the post purchases params
func (o *PostPurchasesParams) WithLineFrequency(lineFrequency *string) *PostPurchasesParams {
	o.SetLineFrequency(lineFrequency)
	return o
}

// SetLineFrequency adds the lineFrequency to the post purchases params
func (o *PostPurchasesParams) SetLineFrequency(lineFrequency *string) {
	o.LineFrequency = lineFrequency
}

// WithLineItemType adds the lineItemType to the post purchases params
func (o *PostPurchasesParams) WithLineItemType(lineItemType *string) *PostPurchasesParams {
	o.SetLineItemType(lineItemType)
	return o
}

// SetLineItemType adds the lineItemType to the post purchases params
func (o *PostPurchasesParams) SetLineItemType(lineItemType *string) {
	o.LineItemType = lineItemType
}

// WithLineName adds the lineName to the post purchases params
func (o *PostPurchasesParams) WithLineName(lineName *string) *PostPurchasesParams {
	o.SetLineName(lineName)
	return o
}

// SetLineName adds the lineName to the post purchases params
func (o *PostPurchasesParams) SetLineName(lineName *string) {
	o.LineName = lineName
}

// WithLineNo adds the lineNo to the post purchases params
func (o *PostPurchasesParams) WithLineNo(lineNo *string) *PostPurchasesParams {
	o.SetLineNo(lineNo)
	return o
}

// SetLineNo adds the lineNo to the post purchases params
func (o *PostPurchasesParams) SetLineNo(lineNo *string) {
	o.LineNo = lineNo
}

// WithLineNotes adds the lineNotes to the post purchases params
func (o *PostPurchasesParams) WithLineNotes(lineNotes *string) *PostPurchasesParams {
	o.SetLineNotes(lineNotes)
	return o
}

// SetLineNotes adds the lineNotes to the post purchases params
func (o *PostPurchasesParams) SetLineNotes(lineNotes *string) {
	o.LineNotes = lineNotes
}

// WithLinePartIds adds the linePartIds to the post purchases params
func (o *PostPurchasesParams) WithLinePartIds(linePartIds *string) *PostPurchasesParams {
	o.SetLinePartIds(linePartIds)
	return o
}

// SetLinePartIds adds the linePartIds to the post purchases params
func (o *PostPurchasesParams) SetLinePartIds(linePartIds *string) {
	o.LinePartIds = linePartIds
}

// WithLineQuantity adds the lineQuantity to the post purchases params
func (o *PostPurchasesParams) WithLineQuantity(lineQuantity *string) *PostPurchasesParams {
	o.SetLineQuantity(lineQuantity)
	return o
}

// SetLineQuantity adds the lineQuantity to the post purchases params
func (o *PostPurchasesParams) SetLineQuantity(lineQuantity *string) {
	o.LineQuantity = lineQuantity
}

// WithLineRackIds adds the lineRackIds to the post purchases params
func (o *PostPurchasesParams) WithLineRackIds(lineRackIds *string) *PostPurchasesParams {
	o.SetLineRackIds(lineRackIds)
	return o
}

// SetLineRackIds adds the lineRackIds to the post purchases params
func (o *PostPurchasesParams) SetLineRackIds(lineRackIds *string) {
	o.LineRackIds = lineRackIds
}

// WithLineRenewDate adds the lineRenewDate to the post purchases params
func (o *PostPurchasesParams) WithLineRenewDate(lineRenewDate *string) *PostPurchasesParams {
	o.SetLineRenewDate(lineRenewDate)
	return o
}

// SetLineRenewDate adds the lineRenewDate to the post purchases params
func (o *PostPurchasesParams) SetLineRenewDate(lineRenewDate *string) {
	o.LineRenewDate = lineRenewDate
}

// WithLineRoomIds adds the lineRoomIds to the post purchases params
func (o *PostPurchasesParams) WithLineRoomIds(lineRoomIds *string) *PostPurchasesParams {
	o.SetLineRoomIds(lineRoomIds)
	return o
}

// SetLineRoomIds adds the lineRoomIds to the post purchases params
func (o *PostPurchasesParams) SetLineRoomIds(lineRoomIds *string) {
	o.LineRoomIds = lineRoomIds
}

// WithLineServiceType adds the lineServiceType to the post purchases params
func (o *PostPurchasesParams) WithLineServiceType(lineServiceType *string) *PostPurchasesParams {
	o.SetLineServiceType(lineServiceType)
	return o
}

// SetLineServiceType adds the lineServiceType to the post purchases params
func (o *PostPurchasesParams) SetLineServiceType(lineServiceType *string) {
	o.LineServiceType = lineServiceType
}

// WithLineSoftwareIds adds the lineSoftwareIds to the post purchases params
func (o *PostPurchasesParams) WithLineSoftwareIds(lineSoftwareIds *string) *PostPurchasesParams {
	o.SetLineSoftwareIds(lineSoftwareIds)
	return o
}

// SetLineSoftwareIds adds the lineSoftwareIds to the post purchases params
func (o *PostPurchasesParams) SetLineSoftwareIds(lineSoftwareIds *string) {
	o.LineSoftwareIds = lineSoftwareIds
}

// WithLineStartDate adds the lineStartDate to the post purchases params
func (o *PostPurchasesParams) WithLineStartDate(lineStartDate *string) *PostPurchasesParams {
	o.SetLineStartDate(lineStartDate)
	return o
}

// SetLineStartDate adds the lineStartDate to the post purchases params
func (o *PostPurchasesParams) SetLineStartDate(lineStartDate *string) {
	o.LineStartDate = lineStartDate
}

// WithLineType adds the lineType to the post purchases params
func (o *PostPurchasesParams) WithLineType(lineType *string) *PostPurchasesParams {
	o.SetLineType(lineType)
	return o
}

// SetLineType adds the lineType to the post purchases params
func (o *PostPurchasesParams) SetLineType(lineType *string) {
	o.LineType = lineType
}

// WithNotes adds the notes to the post purchases params
func (o *PostPurchasesParams) WithNotes(notes *string) *PostPurchasesParams {
	o.SetNotes(notes)
	return o
}

// SetNotes adds the notes to the post purchases params
func (o *PostPurchasesParams) SetNotes(notes *string) {
	o.Notes = notes
}

// WithOrderNo adds the orderNo to the post purchases params
func (o *PostPurchasesParams) WithOrderNo(orderNo *string) *PostPurchasesParams {
	o.SetOrderNo(orderNo)
	return o
}

// SetOrderNo adds the orderNo to the post purchases params
func (o *PostPurchasesParams) SetOrderNo(orderNo *string) {
	o.OrderNo = orderNo
}

// WithPoDate adds the poDate to the post purchases params
func (o *PostPurchasesParams) WithPoDate(poDate *string) *PostPurchasesParams {
	o.SetPoDate(poDate)
	return o
}

// SetPoDate adds the poDate to the post purchases params
func (o *PostPurchasesParams) SetPoDate(poDate *string) {
	o.PoDate = poDate
}

// WithPurchaseID adds the purchaseID to the post purchases params
func (o *PostPurchasesParams) WithPurchaseID(purchaseID *string) *PostPurchasesParams {
	o.SetPurchaseID(purchaseID)
	return o
}

// SetPurchaseID adds the purchaseId to the post purchases params
func (o *PostPurchasesParams) SetPurchaseID(purchaseID *string) {
	o.PurchaseID = purchaseID
}

// WithVendor adds the vendor to the post purchases params
func (o *PostPurchasesParams) WithVendor(vendor *string) *PostPurchasesParams {
	o.SetVendor(vendor)
	return o
}

// SetVendor adds the vendor to the post purchases params
func (o *PostPurchasesParams) SetVendor(vendor *string) {
	o.Vendor = vendor
}

// WriteToRequest writes these params to a swagger request
func (o *PostPurchasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Completed != nil {

		// form param completed
		var frCompleted string
		if o.Completed != nil {
			frCompleted = *o.Completed
		}
		fCompleted := frCompleted
		if fCompleted != "" {
			if err := r.SetFormParam("completed", fCompleted); err != nil {
				return err
			}
		}

	}

	if o.Cost != nil {

		// form param cost
		var frCost string
		if o.Cost != nil {
			frCost = *o.Cost
		}
		fCost := frCost
		if fCost != "" {
			if err := r.SetFormParam("cost", fCost); err != nil {
				return err
			}
		}

	}

	if o.CostCenter != nil {

		// form param cost_center
		var frCostCenter string
		if o.CostCenter != nil {
			frCostCenter = *o.CostCenter
		}
		fCostCenter := frCostCenter
		if fCostCenter != "" {
			if err := r.SetFormParam("cost_center", fCostCenter); err != nil {
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

	if o.LineAssetIds != nil {

		// form param line_asset_ids
		var frLineAssetIds string
		if o.LineAssetIds != nil {
			frLineAssetIds = *o.LineAssetIds
		}
		fLineAssetIds := frLineAssetIds
		if fLineAssetIds != "" {
			if err := r.SetFormParam("line_asset_ids", fLineAssetIds); err != nil {
				return err
			}
		}

	}

	if o.LineBuildingIds != nil {

		// form param line_building_ids
		var frLineBuildingIds string
		if o.LineBuildingIds != nil {
			frLineBuildingIds = *o.LineBuildingIds
		}
		fLineBuildingIds := frLineBuildingIds
		if fLineBuildingIds != "" {
			if err := r.SetFormParam("line_building_ids", fLineBuildingIds); err != nil {
				return err
			}
		}

	}

	if o.LineCancelPolicy != nil {

		// form param line_cancel_policy
		var frLineCancelPolicy string
		if o.LineCancelPolicy != nil {
			frLineCancelPolicy = *o.LineCancelPolicy
		}
		fLineCancelPolicy := frLineCancelPolicy
		if fLineCancelPolicy != "" {
			if err := r.SetFormParam("line_cancel_policy", fLineCancelPolicy); err != nil {
				return err
			}
		}

	}

	if o.LineCertificateIds != nil {

		// form param line_certificate_ids
		var frLineCertificateIds string
		if o.LineCertificateIds != nil {
			frLineCertificateIds = *o.LineCertificateIds
		}
		fLineCertificateIds := frLineCertificateIds
		if fLineCertificateIds != "" {
			if err := r.SetFormParam("line_certificate_ids", fLineCertificateIds); err != nil {
				return err
			}
		}

	}

	if o.LineCircuitIds != nil {

		// form param line_circuit_ids
		var frLineCircuitIds string
		if o.LineCircuitIds != nil {
			frLineCircuitIds = *o.LineCircuitIds
		}
		fLineCircuitIds := frLineCircuitIds
		if fLineCircuitIds != "" {
			if err := r.SetFormParam("line_circuit_ids", fLineCircuitIds); err != nil {
				return err
			}
		}

	}

	if o.LineCircuits != nil {

		// form param line_circuits
		var frLineCircuits string
		if o.LineCircuits != nil {
			frLineCircuits = *o.LineCircuits
		}
		fLineCircuits := frLineCircuits
		if fLineCircuits != "" {
			if err := r.SetFormParam("line_circuits", fLineCircuits); err != nil {
				return err
			}
		}

	}

	if o.LineCompleted != nil {

		// form param line_completed
		var frLineCompleted string
		if o.LineCompleted != nil {
			frLineCompleted = *o.LineCompleted
		}
		fLineCompleted := frLineCompleted
		if fLineCompleted != "" {
			if err := r.SetFormParam("line_completed", fLineCompleted); err != nil {
				return err
			}
		}

	}

	if o.LineContractID != nil {

		// form param line_contract_id
		var frLineContractID string
		if o.LineContractID != nil {
			frLineContractID = *o.LineContractID
		}
		fLineContractID := frLineContractID
		if fLineContractID != "" {
			if err := r.SetFormParam("line_contract_id", fLineContractID); err != nil {
				return err
			}
		}

	}

	if o.LineContractType != nil {

		// form param line_contract_type
		var frLineContractType string
		if o.LineContractType != nil {
			frLineContractType = *o.LineContractType
		}
		fLineContractType := frLineContractType
		if fLineContractType != "" {
			if err := r.SetFormParam("line_contract_type", fLineContractType); err != nil {
				return err
			}
		}

	}

	if o.LineCost != nil {

		// form param line_cost
		var frLineCost string
		if o.LineCost != nil {
			frLineCost = *o.LineCost
		}
		fLineCost := frLineCost
		if fLineCost != "" {
			if err := r.SetFormParam("line_cost", fLineCost); err != nil {
				return err
			}
		}

	}

	if o.LineCostCenter != nil {

		// form param line_cost_center
		var frLineCostCenter string
		if o.LineCostCenter != nil {
			frLineCostCenter = *o.LineCostCenter
		}
		fLineCostCenter := frLineCostCenter
		if fLineCostCenter != "" {
			if err := r.SetFormParam("line_cost_center", fLineCostCenter); err != nil {
				return err
			}
		}

	}

	if o.LineCustomer != nil {

		// form param line_customer
		var frLineCustomer string
		if o.LineCustomer != nil {
			frLineCustomer = *o.LineCustomer
		}
		fLineCustomer := frLineCustomer
		if fLineCustomer != "" {
			if err := r.SetFormParam("line_customer", fLineCustomer); err != nil {
				return err
			}
		}

	}

	if o.LineDeviceAssetNos != nil {

		// form param line_device_asset_nos
		var frLineDeviceAssetNos string
		if o.LineDeviceAssetNos != nil {
			frLineDeviceAssetNos = *o.LineDeviceAssetNos
		}
		fLineDeviceAssetNos := frLineDeviceAssetNos
		if fLineDeviceAssetNos != "" {
			if err := r.SetFormParam("line_device_asset_nos", fLineDeviceAssetNos); err != nil {
				return err
			}
		}

	}

	if o.LineDeviceOsIds != nil {

		// form param line_device_os_ids
		var frLineDeviceOsIds string
		if o.LineDeviceOsIds != nil {
			frLineDeviceOsIds = *o.LineDeviceOsIds
		}
		fLineDeviceOsIds := frLineDeviceOsIds
		if fLineDeviceOsIds != "" {
			if err := r.SetFormParam("line_device_os_ids", fLineDeviceOsIds); err != nil {
				return err
			}
		}

	}

	if o.LineDeviceSerialNos != nil {

		// form param line_device_serial_nos
		var frLineDeviceSerialNos string
		if o.LineDeviceSerialNos != nil {
			frLineDeviceSerialNos = *o.LineDeviceSerialNos
		}
		fLineDeviceSerialNos := frLineDeviceSerialNos
		if fLineDeviceSerialNos != "" {
			if err := r.SetFormParam("line_device_serial_nos", fLineDeviceSerialNos); err != nil {
				return err
			}
		}

	}

	if o.LineDevices != nil {

		// form param line_devices
		var frLineDevices string
		if o.LineDevices != nil {
			frLineDevices = *o.LineDevices
		}
		fLineDevices := frLineDevices
		if fLineDevices != "" {
			if err := r.SetFormParam("line_devices", fLineDevices); err != nil {
				return err
			}
		}

	}

	if o.LineEndDate != nil {

		// form param line_end_date
		var frLineEndDate string
		if o.LineEndDate != nil {
			frLineEndDate = *o.LineEndDate
		}
		fLineEndDate := frLineEndDate
		if fLineEndDate != "" {
			if err := r.SetFormParam("line_end_date", fLineEndDate); err != nil {
				return err
			}
		}

	}

	if o.LineFrequency != nil {

		// form param line_frequency
		var frLineFrequency string
		if o.LineFrequency != nil {
			frLineFrequency = *o.LineFrequency
		}
		fLineFrequency := frLineFrequency
		if fLineFrequency != "" {
			if err := r.SetFormParam("line_frequency", fLineFrequency); err != nil {
				return err
			}
		}

	}

	if o.LineItemType != nil {

		// form param line_item_type
		var frLineItemType string
		if o.LineItemType != nil {
			frLineItemType = *o.LineItemType
		}
		fLineItemType := frLineItemType
		if fLineItemType != "" {
			if err := r.SetFormParam("line_item_type", fLineItemType); err != nil {
				return err
			}
		}

	}

	if o.LineName != nil {

		// form param line_name
		var frLineName string
		if o.LineName != nil {
			frLineName = *o.LineName
		}
		fLineName := frLineName
		if fLineName != "" {
			if err := r.SetFormParam("line_name", fLineName); err != nil {
				return err
			}
		}

	}

	if o.LineNo != nil {

		// form param line_no
		var frLineNo string
		if o.LineNo != nil {
			frLineNo = *o.LineNo
		}
		fLineNo := frLineNo
		if fLineNo != "" {
			if err := r.SetFormParam("line_no", fLineNo); err != nil {
				return err
			}
		}

	}

	if o.LineNotes != nil {

		// form param line_notes
		var frLineNotes string
		if o.LineNotes != nil {
			frLineNotes = *o.LineNotes
		}
		fLineNotes := frLineNotes
		if fLineNotes != "" {
			if err := r.SetFormParam("line_notes", fLineNotes); err != nil {
				return err
			}
		}

	}

	if o.LinePartIds != nil {

		// form param line_part_ids
		var frLinePartIds string
		if o.LinePartIds != nil {
			frLinePartIds = *o.LinePartIds
		}
		fLinePartIds := frLinePartIds
		if fLinePartIds != "" {
			if err := r.SetFormParam("line_part_ids", fLinePartIds); err != nil {
				return err
			}
		}

	}

	if o.LineQuantity != nil {

		// form param line_quantity
		var frLineQuantity string
		if o.LineQuantity != nil {
			frLineQuantity = *o.LineQuantity
		}
		fLineQuantity := frLineQuantity
		if fLineQuantity != "" {
			if err := r.SetFormParam("line_quantity", fLineQuantity); err != nil {
				return err
			}
		}

	}

	if o.LineRackIds != nil {

		// form param line_rack_ids
		var frLineRackIds string
		if o.LineRackIds != nil {
			frLineRackIds = *o.LineRackIds
		}
		fLineRackIds := frLineRackIds
		if fLineRackIds != "" {
			if err := r.SetFormParam("line_rack_ids", fLineRackIds); err != nil {
				return err
			}
		}

	}

	if o.LineRenewDate != nil {

		// form param line_renew_date
		var frLineRenewDate string
		if o.LineRenewDate != nil {
			frLineRenewDate = *o.LineRenewDate
		}
		fLineRenewDate := frLineRenewDate
		if fLineRenewDate != "" {
			if err := r.SetFormParam("line_renew_date", fLineRenewDate); err != nil {
				return err
			}
		}

	}

	if o.LineRoomIds != nil {

		// form param line_room_ids
		var frLineRoomIds string
		if o.LineRoomIds != nil {
			frLineRoomIds = *o.LineRoomIds
		}
		fLineRoomIds := frLineRoomIds
		if fLineRoomIds != "" {
			if err := r.SetFormParam("line_room_ids", fLineRoomIds); err != nil {
				return err
			}
		}

	}

	if o.LineServiceType != nil {

		// form param line_service_type
		var frLineServiceType string
		if o.LineServiceType != nil {
			frLineServiceType = *o.LineServiceType
		}
		fLineServiceType := frLineServiceType
		if fLineServiceType != "" {
			if err := r.SetFormParam("line_service_type", fLineServiceType); err != nil {
				return err
			}
		}

	}

	if o.LineSoftwareIds != nil {

		// form param line_software_ids
		var frLineSoftwareIds string
		if o.LineSoftwareIds != nil {
			frLineSoftwareIds = *o.LineSoftwareIds
		}
		fLineSoftwareIds := frLineSoftwareIds
		if fLineSoftwareIds != "" {
			if err := r.SetFormParam("line_software_ids", fLineSoftwareIds); err != nil {
				return err
			}
		}

	}

	if o.LineStartDate != nil {

		// form param line_start_date
		var frLineStartDate string
		if o.LineStartDate != nil {
			frLineStartDate = *o.LineStartDate
		}
		fLineStartDate := frLineStartDate
		if fLineStartDate != "" {
			if err := r.SetFormParam("line_start_date", fLineStartDate); err != nil {
				return err
			}
		}

	}

	if o.LineType != nil {

		// form param line_type
		var frLineType string
		if o.LineType != nil {
			frLineType = *o.LineType
		}
		fLineType := frLineType
		if fLineType != "" {
			if err := r.SetFormParam("line_type", fLineType); err != nil {
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

	if o.OrderNo != nil {

		// form param order_no
		var frOrderNo string
		if o.OrderNo != nil {
			frOrderNo = *o.OrderNo
		}
		fOrderNo := frOrderNo
		if fOrderNo != "" {
			if err := r.SetFormParam("order_no", fOrderNo); err != nil {
				return err
			}
		}

	}

	if o.PoDate != nil {

		// form param po_date
		var frPoDate string
		if o.PoDate != nil {
			frPoDate = *o.PoDate
		}
		fPoDate := frPoDate
		if fPoDate != "" {
			if err := r.SetFormParam("po_date", fPoDate); err != nil {
				return err
			}
		}

	}

	if o.PurchaseID != nil {

		// form param purchase_id
		var frPurchaseID string
		if o.PurchaseID != nil {
			frPurchaseID = *o.PurchaseID
		}
		fPurchaseID := frPurchaseID
		if fPurchaseID != "" {
			if err := r.SetFormParam("purchase_id", fPurchaseID); err != nil {
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