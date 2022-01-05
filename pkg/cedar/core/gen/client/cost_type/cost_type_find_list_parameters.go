// Code generated by go-swagger; DO NOT EDIT.

package cost_type

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

// NewCostTypeFindListParams creates a new CostTypeFindListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCostTypeFindListParams() *CostTypeFindListParams {
	return &CostTypeFindListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCostTypeFindListParamsWithTimeout creates a new CostTypeFindListParams object
// with the ability to set a timeout on a request.
func NewCostTypeFindListParamsWithTimeout(timeout time.Duration) *CostTypeFindListParams {
	return &CostTypeFindListParams{
		timeout: timeout,
	}
}

// NewCostTypeFindListParamsWithContext creates a new CostTypeFindListParams object
// with the ability to set a context for a request.
func NewCostTypeFindListParamsWithContext(ctx context.Context) *CostTypeFindListParams {
	return &CostTypeFindListParams{
		Context: ctx,
	}
}

// NewCostTypeFindListParamsWithHTTPClient creates a new CostTypeFindListParams object
// with the ability to set a custom HTTPClient for a request.
func NewCostTypeFindListParamsWithHTTPClient(client *http.Client) *CostTypeFindListParams {
	return &CostTypeFindListParams{
		HTTPClient: client,
	}
}

/* CostTypeFindListParams contains all the parameters to send to the API endpoint
   for the cost type find list operation.

   Typically these are written to a http.Request.
*/
type CostTypeFindListParams struct {

	/* Application.

	   Application where the object or role exists.
	*/
	Application string

	/* Name.

	   The name of a specific group of cost types
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cost type find list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CostTypeFindListParams) WithDefaults() *CostTypeFindListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cost type find list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CostTypeFindListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cost type find list params
func (o *CostTypeFindListParams) WithTimeout(timeout time.Duration) *CostTypeFindListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cost type find list params
func (o *CostTypeFindListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cost type find list params
func (o *CostTypeFindListParams) WithContext(ctx context.Context) *CostTypeFindListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cost type find list params
func (o *CostTypeFindListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cost type find list params
func (o *CostTypeFindListParams) WithHTTPClient(client *http.Client) *CostTypeFindListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cost type find list params
func (o *CostTypeFindListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplication adds the application to the cost type find list params
func (o *CostTypeFindListParams) WithApplication(application string) *CostTypeFindListParams {
	o.SetApplication(application)
	return o
}

// SetApplication adds the application to the cost type find list params
func (o *CostTypeFindListParams) SetApplication(application string) {
	o.Application = application
}

// WithName adds the name to the cost type find list params
func (o *CostTypeFindListParams) WithName(name string) *CostTypeFindListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the cost type find list params
func (o *CostTypeFindListParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *CostTypeFindListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param application
	qrApplication := o.Application
	qApplication := qrApplication
	if qApplication != "" {

		if err := r.SetQueryParam("application", qApplication); err != nil {
			return err
		}
	}

	// query param name
	qrName := o.Name
	qName := qrName
	if qName != "" {

		if err := r.SetQueryParam("name", qName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}