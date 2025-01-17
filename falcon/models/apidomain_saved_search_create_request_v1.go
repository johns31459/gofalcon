// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ApidomainSavedSearchCreateRequestV1 apidomain saved search create request v1
//
// swagger:model apidomain.SavedSearchCreateRequestV1
type ApidomainSavedSearchCreateRequestV1 struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// end
	End string `json:"end,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// permissions
	// Required: true
	Permissions *string `json:"permissions"`

	// repo or view
	// Required: true
	RepoOrView *string `json:"repo_or_view"`

	// request schema
	// Required: true
	RequestSchema *string `json:"request_schema"`

	// response schema
	// Required: true
	ResponseSchema *string `json:"response_schema"`

	// search name
	// Required: true
	SearchName *string `json:"search_name"`

	// search query
	// Required: true
	SearchQuery *string `json:"search_query"`

	// search query args
	// Required: true
	SearchQueryArgs map[string]string `json:"search_query_args"`

	// start
	Start string `json:"start,omitempty"`

	// workflow meta
	// Required: true
	WorkflowMeta *ApidomainWorkflowMetadataV1 `json:"workflow_meta"`
}

// Validate validates this apidomain saved search create request v1
func (m *ApidomainSavedSearchCreateRequestV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepoOrView(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchQueryArgs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowMeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApidomainSavedSearchCreateRequestV1) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ApidomainSavedSearchCreateRequestV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ApidomainSavedSearchCreateRequestV1) validatePermissions(formats strfmt.Registry) error {

	if err := validate.Required("permissions", "body", m.Permissions); err != nil {
		return err
	}

	return nil
}

func (m *ApidomainSavedSearchCreateRequestV1) validateRepoOrView(formats strfmt.Registry) error {

	if err := validate.Required("repo_or_view", "body", m.RepoOrView); err != nil {
		return err
	}

	return nil
}

func (m *ApidomainSavedSearchCreateRequestV1) validateRequestSchema(formats strfmt.Registry) error {

	if err := validate.Required("request_schema", "body", m.RequestSchema); err != nil {
		return err
	}

	return nil
}

func (m *ApidomainSavedSearchCreateRequestV1) validateResponseSchema(formats strfmt.Registry) error {

	if err := validate.Required("response_schema", "body", m.ResponseSchema); err != nil {
		return err
	}

	return nil
}

func (m *ApidomainSavedSearchCreateRequestV1) validateSearchName(formats strfmt.Registry) error {

	if err := validate.Required("search_name", "body", m.SearchName); err != nil {
		return err
	}

	return nil
}

func (m *ApidomainSavedSearchCreateRequestV1) validateSearchQuery(formats strfmt.Registry) error {

	if err := validate.Required("search_query", "body", m.SearchQuery); err != nil {
		return err
	}

	return nil
}

func (m *ApidomainSavedSearchCreateRequestV1) validateSearchQueryArgs(formats strfmt.Registry) error {

	if err := validate.Required("search_query_args", "body", m.SearchQueryArgs); err != nil {
		return err
	}

	return nil
}

func (m *ApidomainSavedSearchCreateRequestV1) validateWorkflowMeta(formats strfmt.Registry) error {

	if err := validate.Required("workflow_meta", "body", m.WorkflowMeta); err != nil {
		return err
	}

	if m.WorkflowMeta != nil {
		if err := m.WorkflowMeta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workflow_meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workflow_meta")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this apidomain saved search create request v1 based on the context it is used
func (m *ApidomainSavedSearchCreateRequestV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWorkflowMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApidomainSavedSearchCreateRequestV1) contextValidateWorkflowMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.WorkflowMeta != nil {

		if err := m.WorkflowMeta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workflow_meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workflow_meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApidomainSavedSearchCreateRequestV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApidomainSavedSearchCreateRequestV1) UnmarshalBinary(b []byte) error {
	var res ApidomainSavedSearchCreateRequestV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
