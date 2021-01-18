// Code generated by go-swagger; DO NOT EDIT.

package falconx_sandbox

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewQuerySampleV1Params creates a new QuerySampleV1Params object
// with the default values initialized.
func NewQuerySampleV1Params() *QuerySampleV1Params {
	var ()
	return &QuerySampleV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuerySampleV1ParamsWithTimeout creates a new QuerySampleV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuerySampleV1ParamsWithTimeout(timeout time.Duration) *QuerySampleV1Params {
	var ()
	return &QuerySampleV1Params{

		timeout: timeout,
	}
}

// NewQuerySampleV1ParamsWithContext creates a new QuerySampleV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewQuerySampleV1ParamsWithContext(ctx context.Context) *QuerySampleV1Params {
	var ()
	return &QuerySampleV1Params{

		Context: ctx,
	}
}

// NewQuerySampleV1ParamsWithHTTPClient creates a new QuerySampleV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuerySampleV1ParamsWithHTTPClient(client *http.Client) *QuerySampleV1Params {
	var ()
	return &QuerySampleV1Params{
		HTTPClient: client,
	}
}

/*QuerySampleV1Params contains all the parameters to send to the API endpoint
for the query sample v1 operation typically these are written to a http.Request
*/
type QuerySampleV1Params struct {

	/*XCSUSERUUID
	  User UUID

	*/
	XCSUSERUUID *string
	/*Body
	  Pass a list of sha256s to check if the exist. It will be returned the list of existing hashes.

	*/
	Body *models.SamplestoreQuerySamplesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query sample v1 params
func (o *QuerySampleV1Params) WithTimeout(timeout time.Duration) *QuerySampleV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query sample v1 params
func (o *QuerySampleV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query sample v1 params
func (o *QuerySampleV1Params) WithContext(ctx context.Context) *QuerySampleV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query sample v1 params
func (o *QuerySampleV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query sample v1 params
func (o *QuerySampleV1Params) WithHTTPClient(client *http.Client) *QuerySampleV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query sample v1 params
func (o *QuerySampleV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCSUSERUUID adds the xCSUSERUUID to the query sample v1 params
func (o *QuerySampleV1Params) WithXCSUSERUUID(xCSUSERUUID *string) *QuerySampleV1Params {
	o.SetXCSUSERUUID(xCSUSERUUID)
	return o
}

// SetXCSUSERUUID adds the xCSUSERUuid to the query sample v1 params
func (o *QuerySampleV1Params) SetXCSUSERUUID(xCSUSERUUID *string) {
	o.XCSUSERUUID = xCSUSERUUID
}

// WithBody adds the body to the query sample v1 params
func (o *QuerySampleV1Params) WithBody(body *models.SamplestoreQuerySamplesRequest) *QuerySampleV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the query sample v1 params
func (o *QuerySampleV1Params) SetBody(body *models.SamplestoreQuerySamplesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *QuerySampleV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XCSUSERUUID != nil {

		// header param X-CS-USERUUID
		if err := r.SetHeaderParam("X-CS-USERUUID", *o.XCSUSERUUID); err != nil {
			return err
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}