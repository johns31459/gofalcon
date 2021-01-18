// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

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

// NewGetPolicyContainersParams creates a new GetPolicyContainersParams object
// with the default values initialized.
func NewGetPolicyContainersParams() *GetPolicyContainersParams {
	var ()
	return &GetPolicyContainersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPolicyContainersParamsWithTimeout creates a new GetPolicyContainersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPolicyContainersParamsWithTimeout(timeout time.Duration) *GetPolicyContainersParams {
	var ()
	return &GetPolicyContainersParams{

		timeout: timeout,
	}
}

// NewGetPolicyContainersParamsWithContext creates a new GetPolicyContainersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPolicyContainersParamsWithContext(ctx context.Context) *GetPolicyContainersParams {
	var ()
	return &GetPolicyContainersParams{

		Context: ctx,
	}
}

// NewGetPolicyContainersParamsWithHTTPClient creates a new GetPolicyContainersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPolicyContainersParamsWithHTTPClient(client *http.Client) *GetPolicyContainersParams {
	var ()
	return &GetPolicyContainersParams{
		HTTPClient: client,
	}
}

/*GetPolicyContainersParams contains all the parameters to send to the API endpoint
for the get policy containers operation typically these are written to a http.Request
*/
type GetPolicyContainersParams struct {

	/*Ids
	  The policy container(s) to retrieve, identified by policy ID

	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get policy containers params
func (o *GetPolicyContainersParams) WithTimeout(timeout time.Duration) *GetPolicyContainersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get policy containers params
func (o *GetPolicyContainersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get policy containers params
func (o *GetPolicyContainersParams) WithContext(ctx context.Context) *GetPolicyContainersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get policy containers params
func (o *GetPolicyContainersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get policy containers params
func (o *GetPolicyContainersParams) WithHTTPClient(client *http.Client) *GetPolicyContainersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get policy containers params
func (o *GetPolicyContainersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get policy containers params
func (o *GetPolicyContainersParams) WithIds(ids []string) *GetPolicyContainersParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get policy containers params
func (o *GetPolicyContainersParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetPolicyContainersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesIds := o.Ids

	joinedIds := swag.JoinByFormat(valuesIds, "multi")
	// query array param ids
	if err := r.SetQueryParam("ids", joinedIds...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}