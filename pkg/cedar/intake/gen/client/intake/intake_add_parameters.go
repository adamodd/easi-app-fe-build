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

	"github.com/cmsgov/easi-app/pkg/cedar/intake/gen/models"
)

// NewIntakeAddParams creates a new IntakeAddParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIntakeAddParams() *IntakeAddParams {
	return &IntakeAddParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIntakeAddParamsWithTimeout creates a new IntakeAddParams object
// with the ability to set a timeout on a request.
func NewIntakeAddParamsWithTimeout(timeout time.Duration) *IntakeAddParams {
	return &IntakeAddParams{
		timeout: timeout,
	}
}

// NewIntakeAddParamsWithContext creates a new IntakeAddParams object
// with the ability to set a context for a request.
func NewIntakeAddParamsWithContext(ctx context.Context) *IntakeAddParams {
	return &IntakeAddParams{
		Context: ctx,
	}
}

// NewIntakeAddParamsWithHTTPClient creates a new IntakeAddParams object
// with the ability to set a custom HTTPClient for a request.
func NewIntakeAddParamsWithHTTPClient(client *http.Client) *IntakeAddParams {
	return &IntakeAddParams{
		HTTPClient: client,
	}
}

/* IntakeAddParams contains all the parameters to send to the API endpoint
   for the intake add operation.

   Typically these are written to a http.Request.
*/
type IntakeAddParams struct {

	// Body.
	Body *models.IntakeInput

	/* ValidatePayload.

	   Determines if schema validation of the payload is performed syncronously before persisting the record or asyncronously after the record has been persisted
	*/
	ValidatePayload string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the intake add params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IntakeAddParams) WithDefaults() *IntakeAddParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the intake add params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IntakeAddParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the intake add params
func (o *IntakeAddParams) WithTimeout(timeout time.Duration) *IntakeAddParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the intake add params
func (o *IntakeAddParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the intake add params
func (o *IntakeAddParams) WithContext(ctx context.Context) *IntakeAddParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the intake add params
func (o *IntakeAddParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the intake add params
func (o *IntakeAddParams) WithHTTPClient(client *http.Client) *IntakeAddParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the intake add params
func (o *IntakeAddParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the intake add params
func (o *IntakeAddParams) WithBody(body *models.IntakeInput) *IntakeAddParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the intake add params
func (o *IntakeAddParams) SetBody(body *models.IntakeInput) {
	o.Body = body
}

// WithValidatePayload adds the validatePayload to the intake add params
func (o *IntakeAddParams) WithValidatePayload(validatePayload string) *IntakeAddParams {
	o.SetValidatePayload(validatePayload)
	return o
}

// SetValidatePayload adds the validatePayload to the intake add params
func (o *IntakeAddParams) SetValidatePayload(validatePayload string) {
	o.ValidatePayload = validatePayload
}

// WriteToRequest writes these params to a swagger request
func (o *IntakeAddParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param validatePayload
	qrValidatePayload := o.ValidatePayload
	qValidatePayload := qrValidatePayload
	if qValidatePayload != "" {

		if err := r.SetQueryParam("validatePayload", qValidatePayload); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}