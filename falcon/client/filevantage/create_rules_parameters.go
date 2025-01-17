// Code generated by go-swagger; DO NOT EDIT.

package filevantage

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

// NewCreateRulesParams creates a new CreateRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRulesParams() *CreateRulesParams {
	return &CreateRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRulesParamsWithTimeout creates a new CreateRulesParams object
// with the ability to set a timeout on a request.
func NewCreateRulesParamsWithTimeout(timeout time.Duration) *CreateRulesParams {
	return &CreateRulesParams{
		timeout: timeout,
	}
}

// NewCreateRulesParamsWithContext creates a new CreateRulesParams object
// with the ability to set a context for a request.
func NewCreateRulesParamsWithContext(ctx context.Context) *CreateRulesParams {
	return &CreateRulesParams{
		Context: ctx,
	}
}

// NewCreateRulesParamsWithHTTPClient creates a new CreateRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRulesParamsWithHTTPClient(client *http.Client) *CreateRulesParams {
	return &CreateRulesParams{
		HTTPClient: client,
	}
}

/*
CreateRulesParams contains all the parameters to send to the API endpoint

	for the create rules operation.

	Typically these are written to a http.Request.
*/
type CreateRulesParams struct {

	/* Body.

	     Create a new rule configuration for the specified rule group.

	 * `id` is not supported for creation of a rule, the new id of the created rule will be included in the response.

	 * `rule_group_id` to add the new rule configuration.

	 * `description` can be between 0 and 500 characters.

	 * `path` representing the file system or registry path to monitor.

	   * must be between 1 and 250 characters.

	   * All paths must end with the path separator, e.g. c:\windows\ /usr/bin/

	 * `severity` to categorize change events produced by this rule; must be one of: `Low`, `Medium`, `High` or `Critical`

	 * `depth` below the base path to monitor; must be one of: `1`, `2`, `3`, `4`, `5` or `ANY`

	 * `precedence` - is not supported for creation of a rule, new rules will be added last in precedence order.will result this rule being placed before that existing rule.

	Falcon GLOB syntax is supported for the following 6 properties. Allowed rule group configuration is based on the type of rule group the rule group is added to.

	 * `include` represents the files, directories, registry keys, or registry values that will be monitored.

	 * `exclude` represents the files, directories, registry keys, or registry values that will `NOT` be monitored.

	 * `include_users` represents the changes performed by specific users that will be monitored.

	 * `exclude_users` represents the changes performed by specific users that will `NOT` be monitored.

	 * `include_processes` represents the changes performed by specific processes that will be monitored.

	 * `exclude_processes` represents the changes performed by specific processes that will be `NOT` monitored.

	File system directory monitoring:

	 * `watch_delete_directory_changes`

	 * `watch_create_directory_changes`

	 * `watch_rename_directory_changes`

	 * `watch_attributes_directory_changes` (`macOS` is not supported at this time)

	 * `watch_permissions_directory_changes` (`macOS` is not supported at this time)

	File system file monitoring:

	 * `watch_rename_file_changes`

	 * `watch_write_file_changes`

	 * `watch_create_file_changes`

	 * `watch_delete_file_changes`

	 * `watch_attributes_file_changes` (`macOS` is not supported at this time)

	 * `watch_permissions_file_changes` (`macOS` is not supported at this time)

	Windows registry key and value monitoring:

	 * `watch_create_key_changes`

	 * `watch_delete_key_changes`

	 * `watch_rename_key_changes`

	 * `watch_set_value_changes`

	 * `watch_delete_value_changes`

	 * `watch_create_file_changes`
	*/
	Body *models.RulegroupsRule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRulesParams) WithDefaults() *CreateRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create rules params
func (o *CreateRulesParams) WithTimeout(timeout time.Duration) *CreateRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create rules params
func (o *CreateRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create rules params
func (o *CreateRulesParams) WithContext(ctx context.Context) *CreateRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create rules params
func (o *CreateRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create rules params
func (o *CreateRulesParams) WithHTTPClient(client *http.Client) *CreateRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create rules params
func (o *CreateRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create rules params
func (o *CreateRulesParams) WithBody(body *models.RulegroupsRule) *CreateRulesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create rules params
func (o *CreateRulesParams) SetBody(body *models.RulegroupsRule) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
