// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClientExtractionWithFilesV1 client extraction with files v1
//
// swagger:model client.ExtractionWithFilesV1
type ClientExtractionWithFilesV1 struct {

	// error
	Error string `json:"error,omitempty"`

	// extract timestamp
	// Required: true
	ExtractTimestamp *string `json:"extract_timestamp"`

	// files
	Files []*ClientExtractionFileResultV1 `json:"files"`

	// id
	ID string `json:"id,omitempty"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this client extraction with files v1
func (m *ClientExtractionWithFilesV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtractTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientExtractionWithFilesV1) validateExtractTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("extract_timestamp", "body", m.ExtractTimestamp); err != nil {
		return err
	}

	return nil
}

func (m *ClientExtractionWithFilesV1) validateFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.Files) { // not required
		return nil
	}

	for i := 0; i < len(m.Files); i++ {
		if swag.IsZero(m.Files[i]) { // not required
			continue
		}

		if m.Files[i] != nil {
			if err := m.Files[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClientExtractionWithFilesV1) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this client extraction with files v1 based on the context it is used
func (m *ClientExtractionWithFilesV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientExtractionWithFilesV1) contextValidateFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Files); i++ {

		if m.Files[i] != nil {
			if err := m.Files[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClientExtractionWithFilesV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientExtractionWithFilesV1) UnmarshalBinary(b []byte) error {
	var res ClientExtractionWithFilesV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
