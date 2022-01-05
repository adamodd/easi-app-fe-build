// Code generated by go-swagger; DO NOT EDIT.

package budget

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

// NewBudgetFindParams creates a new BudgetFindParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBudgetFindParams() *BudgetFindParams {
	return &BudgetFindParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBudgetFindParamsWithTimeout creates a new BudgetFindParams object
// with the ability to set a timeout on a request.
func NewBudgetFindParamsWithTimeout(timeout time.Duration) *BudgetFindParams {
	return &BudgetFindParams{
		timeout: timeout,
	}
}

// NewBudgetFindParamsWithContext creates a new BudgetFindParams object
// with the ability to set a context for a request.
func NewBudgetFindParamsWithContext(ctx context.Context) *BudgetFindParams {
	return &BudgetFindParams{
		Context: ctx,
	}
}

// NewBudgetFindParamsWithHTTPClient creates a new BudgetFindParams object
// with the ability to set a custom HTTPClient for a request.
func NewBudgetFindParamsWithHTTPClient(client *http.Client) *BudgetFindParams {
	return &BudgetFindParams{
		HTTPClient: client,
	}
}

/* BudgetFindParams contains all the parameters to send to the API endpoint
   for the budget find operation.

   Typically these are written to a http.Request.
*/
type BudgetFindParams struct {

	/* IdsOnly.

	   Whether the call will return simply the IDs.  If not set, the response will also include projectId, systemId, fundingId and funding.
	*/
	IdsOnly *bool

	/* OnlyIds.

	   Whether the call will return simply the IDs. If not set, the response will also include projectId, systemId, fundingId and funding.
	*/
	OnlyIds *bool

	/* ProjectID.

	   Project Id string to search. Adding this parameter will  instruct the interface to only return the system matching the projectId.
	*/
	ProjectID *string

	/* ProjectTitle.

	   Project Title string to search. Adding this parameter will  instruct the interface to only return the system matching the systemId.
	*/
	ProjectTitle *string

	/* SystemID.

	   ID of the system that the budget is associated with. Adding this parameter will instruct the interface to only return the system matching the systemId.
	*/
	SystemID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the budget find params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BudgetFindParams) WithDefaults() *BudgetFindParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the budget find params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BudgetFindParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the budget find params
func (o *BudgetFindParams) WithTimeout(timeout time.Duration) *BudgetFindParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the budget find params
func (o *BudgetFindParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the budget find params
func (o *BudgetFindParams) WithContext(ctx context.Context) *BudgetFindParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the budget find params
func (o *BudgetFindParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the budget find params
func (o *BudgetFindParams) WithHTTPClient(client *http.Client) *BudgetFindParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the budget find params
func (o *BudgetFindParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdsOnly adds the idsOnly to the budget find params
func (o *BudgetFindParams) WithIdsOnly(idsOnly *bool) *BudgetFindParams {
	o.SetIdsOnly(idsOnly)
	return o
}

// SetIdsOnly adds the idsOnly to the budget find params
func (o *BudgetFindParams) SetIdsOnly(idsOnly *bool) {
	o.IdsOnly = idsOnly
}

// WithOnlyIds adds the onlyIds to the budget find params
func (o *BudgetFindParams) WithOnlyIds(onlyIds *bool) *BudgetFindParams {
	o.SetOnlyIds(onlyIds)
	return o
}

// SetOnlyIds adds the onlyIds to the budget find params
func (o *BudgetFindParams) SetOnlyIds(onlyIds *bool) {
	o.OnlyIds = onlyIds
}

// WithProjectID adds the projectID to the budget find params
func (o *BudgetFindParams) WithProjectID(projectID *string) *BudgetFindParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the budget find params
func (o *BudgetFindParams) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WithProjectTitle adds the projectTitle to the budget find params
func (o *BudgetFindParams) WithProjectTitle(projectTitle *string) *BudgetFindParams {
	o.SetProjectTitle(projectTitle)
	return o
}

// SetProjectTitle adds the projectTitle to the budget find params
func (o *BudgetFindParams) SetProjectTitle(projectTitle *string) {
	o.ProjectTitle = projectTitle
}

// WithSystemID adds the systemID to the budget find params
func (o *BudgetFindParams) WithSystemID(systemID *string) *BudgetFindParams {
	o.SetSystemID(systemID)
	return o
}

// SetSystemID adds the systemId to the budget find params
func (o *BudgetFindParams) SetSystemID(systemID *string) {
	o.SystemID = systemID
}

// WriteToRequest writes these params to a swagger request
func (o *BudgetFindParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IdsOnly != nil {

		// query param idsOnly
		var qrIdsOnly bool

		if o.IdsOnly != nil {
			qrIdsOnly = *o.IdsOnly
		}
		qIdsOnly := swag.FormatBool(qrIdsOnly)
		if qIdsOnly != "" {

			if err := r.SetQueryParam("idsOnly", qIdsOnly); err != nil {
				return err
			}
		}
	}

	if o.OnlyIds != nil {

		// query param onlyIds
		var qrOnlyIds bool

		if o.OnlyIds != nil {
			qrOnlyIds = *o.OnlyIds
		}
		qOnlyIds := swag.FormatBool(qrOnlyIds)
		if qOnlyIds != "" {

			if err := r.SetQueryParam("onlyIds", qOnlyIds); err != nil {
				return err
			}
		}
	}

	if o.ProjectID != nil {

		// query param projectId
		var qrProjectID string

		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := qrProjectID
		if qProjectID != "" {

			if err := r.SetQueryParam("projectId", qProjectID); err != nil {
				return err
			}
		}
	}

	if o.ProjectTitle != nil {

		// query param projectTitle
		var qrProjectTitle string

		if o.ProjectTitle != nil {
			qrProjectTitle = *o.ProjectTitle
		}
		qProjectTitle := qrProjectTitle
		if qProjectTitle != "" {

			if err := r.SetQueryParam("projectTitle", qProjectTitle); err != nil {
				return err
			}
		}
	}

	if o.SystemID != nil {

		// query param systemId
		var qrSystemID string

		if o.SystemID != nil {
			qrSystemID = *o.SystemID
		}
		qSystemID := qrSystemID
		if qSystemID != "" {

			if err := r.SetQueryParam("systemId", qSystemID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}