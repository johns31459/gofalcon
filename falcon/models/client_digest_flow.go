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

// ClientDigestFlow client digest flow
//
// swagger:model client.DigestFlow
type ClientDigestFlow struct {

	// ingest time known good
	// Required: true
	IngestTimeKnownGood *int64 `json:"ingestTimeKnownGood"`

	// max ingest latency
	// Required: true
	MaxIngestLatency *int32 `json:"maxIngestLatency"`

	// min ingest time included
	// Required: true
	MinIngestTimeIncluded *int64 `json:"minIngestTimeIncluded"`
}

// Validate validates this client digest flow
func (m *ClientDigestFlow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIngestTimeKnownGood(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxIngestLatency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinIngestTimeIncluded(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientDigestFlow) validateIngestTimeKnownGood(formats strfmt.Registry) error {

	if err := validate.Required("ingestTimeKnownGood", "body", m.IngestTimeKnownGood); err != nil {
		return err
	}

	return nil
}

func (m *ClientDigestFlow) validateMaxIngestLatency(formats strfmt.Registry) error {

	if err := validate.Required("maxIngestLatency", "body", m.MaxIngestLatency); err != nil {
		return err
	}

	return nil
}

func (m *ClientDigestFlow) validateMinIngestTimeIncluded(formats strfmt.Registry) error {

	if err := validate.Required("minIngestTimeIncluded", "body", m.MinIngestTimeIncluded); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this client digest flow based on context it is used
func (m *ClientDigestFlow) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClientDigestFlow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientDigestFlow) UnmarshalBinary(b []byte) error {
	var res ClientDigestFlow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
