// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApidomainWorkflowMetadataV1 apidomain workflow metadata v1
//
// swagger:model apidomain.WorkflowMetadataV1
type ApidomainWorkflowMetadataV1 struct {

	// is system
	IsSystem bool `json:"is_system,omitempty"`

	// polling interval
	PollingInterval int32 `json:"polling_interval,omitempty"`

	// polling timeout
	PollingTimeout int32 `json:"polling_timeout,omitempty"`

	// tag ids
	TagIds []string `json:"tag_ids"`
}

// Validate validates this apidomain workflow metadata v1
func (m *ApidomainWorkflowMetadataV1) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this apidomain workflow metadata v1 based on context it is used
func (m *ApidomainWorkflowMetadataV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApidomainWorkflowMetadataV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApidomainWorkflowMetadataV1) UnmarshalBinary(b []byte) error {
	var res ApidomainWorkflowMetadataV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
