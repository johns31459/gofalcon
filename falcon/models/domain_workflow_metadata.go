// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DomainWorkflowMetadata domain workflow metadata
//
// swagger:model domain.WorkflowMetadata
type DomainWorkflowMetadata struct {

	// is system
	IsSystem bool `json:"is_system,omitempty"`

	// polling interval
	PollingInterval int32 `json:"polling_interval,omitempty"`

	// polling timeout
	PollingTimeout int32 `json:"polling_timeout,omitempty"`

	// tag ids
	TagIds []string `json:"tag_ids"`
}

// Validate validates this domain workflow metadata
func (m *DomainWorkflowMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this domain workflow metadata based on context it is used
func (m *DomainWorkflowMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainWorkflowMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainWorkflowMetadata) UnmarshalBinary(b []byte) error {
	var res DomainWorkflowMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
