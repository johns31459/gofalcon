// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClientExtractionFileParametersV1 client extraction file parameters v1
//
// swagger:model client.ExtractionFileParametersV1
type ClientExtractionFileParametersV1 struct {

	// comment
	Comment string `json:"comment,omitempty"`

	// is confidential
	IsConfidential bool `json:"is_confidential,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this client extraction file parameters v1
func (m *ClientExtractionFileParametersV1) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this client extraction file parameters v1 based on context it is used
func (m *ClientExtractionFileParametersV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClientExtractionFileParametersV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientExtractionFileParametersV1) UnmarshalBinary(b []byte) error {
	var res ClientExtractionFileParametersV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}