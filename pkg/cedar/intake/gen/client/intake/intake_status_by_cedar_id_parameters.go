// Code generated by go-swagger; DO NOT EDIT.

package intake

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

// NewIntakeStatusByCedarIDParams creates a new IntakeStatusByCedarIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIntakeStatusByCedarIDParams() *IntakeStatusByCedarIDParams {
	return &IntakeStatusByCedarIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIntakeStatusByCedarIDParamsWithTimeout creates a new IntakeStatusByCedarIDParams object
// with the ability to set a timeout on a request.
func NewIntakeStatusByCedarIDParamsWithTimeout(timeout time.Duration) *IntakeStatusByCedarIDParams {
	return &IntakeStatusByCedarIDParams{
		timeout: timeout,
	}
}

// NewIntakeStatusByCedarIDParamsWithContext creates a new IntakeStatusByCedarIDParams object
// with the ability to set a context for a request.
func NewIntakeStatusByCedarIDParamsWithContext(ctx context.Context) *IntakeStatusByCedarIDParams {
	return &IntakeStatusByCedarIDParams{
		Context: ctx,
	}
}

// NewIntakeStatusByCedarIDParamsWithHTTPClient creates a new IntakeStatusByCedarIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewIntakeStatusByCedarIDParamsWithHTTPClient(client *http.Client) *IntakeStatusByCedarIDParams {
	return &IntakeStatusByCedarIDParams{
		HTTPClient: client,
	}
}

/* IntakeStatusByCedarIDParams contains all the parameters to send to the API endpoint
   for the intake status by cedar Id operation.

   Typically these are written to a http.Request.
*/
type IntakeStatusByCedarIDParams struct {

	/* ID.

	   An intake ID assigned by CEDAR
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the intake status by cedar Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IntakeStatusByCedarIDParams) WithDefaults() *IntakeStatusByCedarIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the intake status by cedar Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IntakeStatusByCedarIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the intake status by cedar Id params
func (o *IntakeStatusByCedarIDParams) WithTimeout(timeout time.Duration) *IntakeStatusByCedarIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the intake status by cedar Id params
func (o *IntakeStatusByCedarIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the intake status by cedar Id params
func (o *IntakeStatusByCedarIDParams) WithContext(ctx context.Context) *IntakeStatusByCedarIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the intake status by cedar Id params
func (o *IntakeStatusByCedarIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the intake status by cedar Id params
func (o *IntakeStatusByCedarIDParams) WithHTTPClient(client *http.Client) *IntakeStatusByCedarIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the intake status by cedar Id params
func (o *IntakeStatusByCedarIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the intake status by cedar Id params
func (o *IntakeStatusByCedarIDParams) WithID(id string) *IntakeStatusByCedarIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the intake status by cedar Id params
func (o *IntakeStatusByCedarIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IntakeStatusByCedarIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}