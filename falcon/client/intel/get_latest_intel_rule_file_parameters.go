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

// NewGetLatestIntelRuleFileParams creates a new GetLatestIntelRuleFileParams object
// with the default values initialized.
func NewGetLatestIntelRuleFileParams() *GetLatestIntelRuleFileParams {
	var ()
	return &GetLatestIntelRuleFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLatestIntelRuleFileParamsWithTimeout creates a new GetLatestIntelRuleFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLatestIntelRuleFileParamsWithTimeout(timeout time.Duration) *GetLatestIntelRuleFileParams {
	var ()
	return &GetLatestIntelRuleFileParams{

		timeout: timeout,
	}
}

// NewGetLatestIntelRuleFileParamsWithContext creates a new GetLatestIntelRuleFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLatestIntelRuleFileParamsWithContext(ctx context.Context) *GetLatestIntelRuleFileParams {
	var ()
	return &GetLatestIntelRuleFileParams{

		Context: ctx,
	}
}

// NewGetLatestIntelRuleFileParamsWithHTTPClient creates a new GetLatestIntelRuleFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLatestIntelRuleFileParamsWithHTTPClient(client *http.Client) *GetLatestIntelRuleFileParams {
	var ()
	return &GetLatestIntelRuleFileParams{
		HTTPClient: client,
	}
}

/*GetLatestIntelRuleFileParams contains all the parameters to send to the API endpoint
for the get latest intel rule file operation typically these are written to a http.Request
*/
type GetLatestIntelRuleFileParams struct {

	/*Accept
	  Choose the format you want the rule set in.

	*/
	Accept *string
	/*Format
	  Choose the format you want the rule set in. Valid formats are zip and gzip. Defaults to zip.

	*/
	Format *string
	/*Type
	  The rule news report type. Accepted values:

	snort-suricata-master

	snort-suricata-update

	snort-suricata-changelog

	yara-master

	yara-update

	yara-changelog

	common-event-format

	netwitness

	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get latest intel rule file params
func (o *GetLatestIntelRuleFileParams) WithTimeout(timeout time.Duration) *GetLatestIntelRuleFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get latest intel rule file params
func (o *GetLatestIntelRuleFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get latest intel rule file params
func (o *GetLatestIntelRuleFileParams) WithContext(ctx context.Context) *GetLatestIntelRuleFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get latest intel rule file params
func (o *GetLatestIntelRuleFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get latest intel rule file params
func (o *GetLatestIntelRuleFileParams) WithHTTPClient(client *http.Client) *GetLatestIntelRuleFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get latest intel rule file params
func (o *GetLatestIntelRuleFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccept adds the accept to the get latest intel rule file params
func (o *GetLatestIntelRuleFileParams) WithAccept(accept *string) *GetLatestIntelRuleFileParams {
	o.SetAccept(accept)
	return o
}

// SetAccept adds the accept to the get latest intel rule file params
func (o *GetLatestIntelRuleFileParams) SetAccept(accept *string) {
	o.Accept = accept
}

// WithFormat adds the format to the get latest intel rule file params
func (o *GetLatestIntelRuleFileParams) WithFormat(format *string) *GetLatestIntelRuleFileParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get latest intel rule file params
func (o *GetLatestIntelRuleFileParams) SetFormat(format *string) {
	o.Format = format
}

// WithType adds the typeVar to the get latest intel rule file params
func (o *GetLatestIntelRuleFileParams) WithType(typeVar string) *GetLatestIntelRuleFileParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get latest intel rule file params
func (o *GetLatestIntelRuleFileParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetLatestIntelRuleFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Accept != nil {

		// header param Accept
		if err := r.SetHeaderParam("Accept", *o.Accept); err != nil {
			return err
		}

	}

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}

	}

	// query param type
	qrType := o.Type
	qType := qrType
	if qType != "" {
		if err := r.SetQueryParam("type", qType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}