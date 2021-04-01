// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewHealthCheckGetParams creates a new HealthCheckGetParams object
// with the default values initialized.
func NewHealthCheckGetParams() *HealthCheckGetParams {

	return &HealthCheckGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHealthCheckGetParamsWithTimeout creates a new HealthCheckGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHealthCheckGetParamsWithTimeout(timeout time.Duration) *HealthCheckGetParams {

	return &HealthCheckGetParams{

		timeout: timeout,
	}
}

// NewHealthCheckGetParamsWithContext creates a new HealthCheckGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewHealthCheckGetParamsWithContext(ctx context.Context) *HealthCheckGetParams {

	return &HealthCheckGetParams{

		Context: ctx,
	}
}

// NewHealthCheckGetParamsWithHTTPClient creates a new HealthCheckGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHealthCheckGetParamsWithHTTPClient(client *http.Client) *HealthCheckGetParams {

	return &HealthCheckGetParams{
		HTTPClient: client,
	}
}

/*HealthCheckGetParams contains all the parameters to send to the API endpoint
for the health check get operation typically these are written to a http.Request
*/
type HealthCheckGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the health check get params
func (o *HealthCheckGetParams) WithTimeout(timeout time.Duration) *HealthCheckGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the health check get params
func (o *HealthCheckGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the health check get params
func (o *HealthCheckGetParams) WithContext(ctx context.Context) *HealthCheckGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the health check get params
func (o *HealthCheckGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the health check get params
func (o *HealthCheckGetParams) WithHTTPClient(client *http.Client) *HealthCheckGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the health check get params
func (o *HealthCheckGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HealthCheckGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}