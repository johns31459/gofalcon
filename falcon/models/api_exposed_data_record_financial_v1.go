// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIExposedDataRecordFinancialV1 api exposed data record financial v1
//
// swagger:model api.ExposedDataRecordFinancialV1
type APIExposedDataRecordFinancialV1 struct {

	// bank account
	BankAccount string `json:"bank_account,omitempty"`

	// credit card
	CreditCard string `json:"credit_card,omitempty"`

	// crypto currency addresses
	CryptoCurrencyAddresses []string `json:"crypto_currency_addresses"`
}

// Validate validates this api exposed data record financial v1
func (m *APIExposedDataRecordFinancialV1) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this api exposed data record financial v1 based on context it is used
func (m *APIExposedDataRecordFinancialV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIExposedDataRecordFinancialV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIExposedDataRecordFinancialV1) UnmarshalBinary(b []byte) error {
	var res APIExposedDataRecordFinancialV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
