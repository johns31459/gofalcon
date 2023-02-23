// Code generated by go-swagger; DO NOT EDIT.

package intel

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

// NewGetMitreReportParams creates a new GetMitreReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMitreReportParams() *GetMitreReportParams {
	return &GetMitreReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMitreReportParamsWithTimeout creates a new GetMitreReportParams object
// with the ability to set a timeout on a request.
func NewGetMitreReportParamsWithTimeout(timeout time.Duration) *GetMitreReportParams {
	return &GetMitreReportParams{
		timeout: timeout,
	}
}

// NewGetMitreReportParamsWithContext creates a new GetMitreReportParams object
// with the ability to set a context for a request.
func NewGetMitreReportParamsWithContext(ctx context.Context) *GetMitreReportParams {
	return &GetMitreReportParams{
		Context: ctx,
	}
}

// NewGetMitreReportParamsWithHTTPClient creates a new GetMitreReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMitreReportParamsWithHTTPClient(client *http.Client) *GetMitreReportParams {
	return &GetMitreReportParams{
		HTTPClient: client,
	}
}

/*
GetMitreReportParams contains all the parameters to send to the API endpoint

	for the get mitre report operation.

	Typically these are written to a http.Request.
*/
type GetMitreReportParams struct {

	/* ActorID.

	   Actor ID(derived from the actor's name)
	*/
	ActorID string

	/* Format.

	   Supported report formats: CSV or JSON
	*/
	Format string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get mitre report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMitreReportParams) WithDefaults() *GetMitreReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get mitre report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMitreReportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get mitre report params
func (o *GetMitreReportParams) WithTimeout(timeout time.Duration) *GetMitreReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get mitre report params
func (o *GetMitreReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get mitre report params
func (o *GetMitreReportParams) WithContext(ctx context.Context) *GetMitreReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get mitre report params
func (o *GetMitreReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get mitre report params
func (o *GetMitreReportParams) WithHTTPClient(client *http.Client) *GetMitreReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get mitre report params
func (o *GetMitreReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActorID adds the actorID to the get mitre report params
func (o *GetMitreReportParams) WithActorID(actorID string) *GetMitreReportParams {
	o.SetActorID(actorID)
	return o
}

// SetActorID adds the actorId to the get mitre report params
func (o *GetMitreReportParams) SetActorID(actorID string) {
	o.ActorID = actorID
}

// WithFormat adds the format to the get mitre report params
func (o *GetMitreReportParams) WithFormat(format string) *GetMitreReportParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get mitre report params
func (o *GetMitreReportParams) SetFormat(format string) {
	o.Format = format
}

// WriteToRequest writes these params to a swagger request
func (o *GetMitreReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param actor_id
	qrActorID := o.ActorID
	qActorID := qrActorID
	if qActorID != "" {

		if err := r.SetQueryParam("actor_id", qActorID); err != nil {
			return err
		}
	}

	// query param format
	qrFormat := o.Format
	qFormat := qrFormat
	if qFormat != "" {

		if err := r.SetQueryParam("format", qFormat); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}