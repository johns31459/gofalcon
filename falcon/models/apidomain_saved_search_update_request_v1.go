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

// ApidomainSavedSearchUpdateRequestV1 apidomain saved search update request v1
//
// swagger:model apidomain.SavedSearchUpdateRequestV1
type ApidomainSavedSearchUpdateRequestV1 struct {

	// description
	Description string `json:"description,omitempty"`

	// end
	// Required: true
	End *string `json:"end"`

	// id
	ID string `json:"id,omitempty"`

	// permissions
	Permissions string `json:"permissions,omitempty"`

	// repo or view
	RepoOrView string `json:"repo_or_view,omitempty"`

	// request schema
	RequestSchema string `json:"request_schema,omitempty"`

	// response schema
	ResponseSchema string `json:"response_schema,omitempty"`

	// search name
	SearchName string `json:"search_name,omitempty"`

	// search query
	SearchQuery string `json:"search_query,omitempty"`

	// search query args
	SearchQueryArgs ApidomainQueryArgs `json:"search_query_args,omitempty"`

	// start
	Start string `json:"start,omitempty"`

	// status
	// Required: true
	Status *string `json:"status"`

	// version
	Version string `json:"version,omitempty"`

	// version numeric
	VersionNumeric int32 `json:"version_numeric,omitempty"`

	// workflow meta
	WorkflowMeta *ApidomainWorkflowMetadataV1 `json:"workflow_meta,omitempty"`
}

// Validate validates this apidomain saved search update request v1
func (m *ApidomainSavedSearchUpdateRequestV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchQueryArgs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *ApidomainSavedSearchUpdateRequestV1) validateEnd(formats strfmt.Registry) error {

	if err := validate.Required("end", "body", m.End); err != nil {
		return err
	}

	return nil
}

func (m *ApidomainSavedSearchUpdateRequestV1) validateSearchQueryArgs(formats strfmt.Registry) error {
	if swag.IsZero(m.SearchQueryArgs) { // not required
		return nil
	}

	if m.SearchQueryArgs != nil {
		if err := m.SearchQueryArgs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("search_query_args")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("search_query_args")
			}
			return err
		}
	}

	return nil
}

func (m *ApidomainSavedSearchUpdateRequestV1) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ApidomainSavedSearchUpdateRequestV1) validateWorkflowMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkflowMeta) { // not required
		return nil
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

// ContextValidate validate this apidomain saved search update request v1 based on the context it is used
func (m *ApidomainSavedSearchUpdateRequestV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSearchQueryArgs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkflowMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApidomainSavedSearchUpdateRequestV1) contextValidateSearchQueryArgs(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.SearchQueryArgs) { // not required
		return nil
	}

	if err := m.SearchQueryArgs.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("search_query_args")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("search_query_args")
		}
		return err
	}

	return nil
}

func (m *ApidomainSavedSearchUpdateRequestV1) contextValidateWorkflowMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.WorkflowMeta != nil {

		if swag.IsZero(m.WorkflowMeta) { // not required
			return nil
		}

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
func (m *ApidomainSavedSearchUpdateRequestV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApidomainSavedSearchUpdateRequestV1) UnmarshalBinary(b []byte) error {
	var res ApidomainSavedSearchUpdateRequestV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
