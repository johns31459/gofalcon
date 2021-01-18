// Code generated by go-swagger; DO NOT EDIT.

package custom_ioa

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

// NewGetRulesMixin0Params creates a new GetRulesMixin0Params object
// with the default values initialized.
func NewGetRulesMixin0Params() *GetRulesMixin0Params {
	var ()
	return &GetRulesMixin0Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRulesMixin0ParamsWithTimeout creates a new GetRulesMixin0Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRulesMixin0ParamsWithTimeout(timeout time.Duration) *GetRulesMixin0Params {
	var ()
	return &GetRulesMixin0Params{

		timeout: timeout,
	}
}

// NewGetRulesMixin0ParamsWithContext creates a new GetRulesMixin0Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetRulesMixin0ParamsWithContext(ctx context.Context) *GetRulesMixin0Params {
	var ()
	return &GetRulesMixin0Params{

		Context: ctx,
	}
}

// NewGetRulesMixin0ParamsWithHTTPClient creates a new GetRulesMixin0Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRulesMixin0ParamsWithHTTPClient(client *http.Client) *GetRulesMixin0Params {
	var ()
	return &GetRulesMixin0Params{
		HTTPClient: client,
	}
}

/*GetRulesMixin0Params contains all the parameters to send to the API endpoint
for the get rules mixin0 operation typically these are written to a http.Request
*/
type GetRulesMixin0Params struct {

	/*Ids
	  The IDs of the entities

	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get rules mixin0 params
func (o *GetRulesMixin0Params) WithTimeout(timeout time.Duration) *GetRulesMixin0Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get rules mixin0 params
func (o *GetRulesMixin0Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get rules mixin0 params
func (o *GetRulesMixin0Params) WithContext(ctx context.Context) *GetRulesMixin0Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get rules mixin0 params
func (o *GetRulesMixin0Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get rules mixin0 params
func (o *GetRulesMixin0Params) WithHTTPClient(client *http.Client) *GetRulesMixin0Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get rules mixin0 params
func (o *GetRulesMixin0Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get rules mixin0 params
func (o *GetRulesMixin0Params) WithIds(ids []string) *GetRulesMixin0Params {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get rules mixin0 params
func (o *GetRulesMixin0Params) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetRulesMixin0Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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