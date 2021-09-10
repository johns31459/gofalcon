// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DomainQueriesPatchRequest domain queries patch request
//
// swagger:model domain.QueriesPatchRequest
type DomainQueriesPatchRequest struct {

	// action
	Action string `json:"action,omitempty"`

	// comment
	Comment string `json:"comment,omitempty"`

	// filter
	Filter string `json:"filter,omitempty"`

	// q
	Q string `json:"q,omitempty"`
}

// Validate validates this domain queries patch request
func (m *DomainQueriesPatchRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this domain queries patch request based on context it is used
func (m *DomainQueriesPatchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainQueriesPatchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainQueriesPatchRequest) UnmarshalBinary(b []byte) error {
	var res DomainQueriesPatchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}