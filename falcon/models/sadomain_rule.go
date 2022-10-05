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

// SadomainRule sadomain rule
//
// swagger:model sadomain.Rule
type SadomainRule struct {

	// Whether to monitor for breach data. Available only for `Company Domains` and `Email addresses` rule topics. When enabled, ownership of the monitored domains or emails is required
	// Required: true
	BreachMonitoringEnabled *bool `json:"breach_monitoring_enabled"`

	// cid
	// Required: true
	Cid *string `json:"cid"`

	// The creation time for a given rule
	// Required: true
	// Format: date-time
	CreatedTimestamp *strfmt.DateTime `json:"created_timestamp"`

	// The FQL filter contained in a rule and used for searching. Parentheses may be added automatically for clarity
	// Required: true
	Filter *string `json:"filter"`

	// The ID of a given rule
	// Required: true
	ID *string `json:"id"`

	// The name of a given rule
	// Required: true
	Name *string `json:"name"`

	// The customer assets for which ownership must be verified, in order to monitor for breach data
	OwnershipAssets *SadomainCustomerAssets `json:"ownership_assets,omitempty"`

	// The permissions of a given rule
	// Required: true
	Permissions *string `json:"permissions"`

	// The priority of a given rule
	// Required: true
	Priority *string `json:"priority"`

	// The status of a given rule
	// Required: true
	Status *string `json:"status"`

	// The detailed status message of a given rule
	StatusMessage string `json:"status_message,omitempty"`

	// Whether to monitor for substring matches. Only available for the `Typosquatting` rule topic
	// Required: true
	SubstringMatchingEnabled *bool `json:"substring_matching_enabled"`

	// The topic of a given rule
	// Required: true
	Topic *string `json:"topic"`

	// The last updated time for a given rule
	// Required: true
	// Format: date-time
	UpdatedTimestamp *strfmt.DateTime `json:"updated_timestamp"`

	// The user ID of the user that created a given rule
	UserID string `json:"user_id,omitempty"`

	// The user name of the user that created a given rule
	UserName string `json:"user_name,omitempty"`

	// The UUID of the user that created a given rule
	// Required: true
	UserUUID *string `json:"user_uuid"`
}

// Validate validates this sadomain rule
func (m *SadomainRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBreachMonitoringEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnershipAssets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubstringMatchingEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SadomainRule) validateBreachMonitoringEnabled(formats strfmt.Registry) error {

	if err := validate.Required("breach_monitoring_enabled", "body", m.BreachMonitoringEnabled); err != nil {
		return err
	}

	return nil
}

func (m *SadomainRule) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *SadomainRule) validateCreatedTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("created_timestamp", "body", m.CreatedTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("created_timestamp", "body", "date-time", m.CreatedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SadomainRule) validateFilter(formats strfmt.Registry) error {

	if err := validate.Required("filter", "body", m.Filter); err != nil {
		return err
	}

	return nil
}

func (m *SadomainRule) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SadomainRule) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SadomainRule) validateOwnershipAssets(formats strfmt.Registry) error {
	if swag.IsZero(m.OwnershipAssets) { // not required
		return nil
	}

	if m.OwnershipAssets != nil {
		if err := m.OwnershipAssets.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ownership_assets")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ownership_assets")
			}
			return err
		}
	}

	return nil
}

func (m *SadomainRule) validatePermissions(formats strfmt.Registry) error {

	if err := validate.Required("permissions", "body", m.Permissions); err != nil {
		return err
	}

	return nil
}

func (m *SadomainRule) validatePriority(formats strfmt.Registry) error {

	if err := validate.Required("priority", "body", m.Priority); err != nil {
		return err
	}

	return nil
}

func (m *SadomainRule) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *SadomainRule) validateSubstringMatchingEnabled(formats strfmt.Registry) error {

	if err := validate.Required("substring_matching_enabled", "body", m.SubstringMatchingEnabled); err != nil {
		return err
	}

	return nil
}

func (m *SadomainRule) validateTopic(formats strfmt.Registry) error {

	if err := validate.Required("topic", "body", m.Topic); err != nil {
		return err
	}

	return nil
}

func (m *SadomainRule) validateUpdatedTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("updated_timestamp", "body", m.UpdatedTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_timestamp", "body", "date-time", m.UpdatedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SadomainRule) validateUserUUID(formats strfmt.Registry) error {

	if err := validate.Required("user_uuid", "body", m.UserUUID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this sadomain rule based on the context it is used
func (m *SadomainRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOwnershipAssets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SadomainRule) contextValidateOwnershipAssets(ctx context.Context, formats strfmt.Registry) error {

	if m.OwnershipAssets != nil {
		if err := m.OwnershipAssets.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ownership_assets")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ownership_assets")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SadomainRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SadomainRule) UnmarshalBinary(b []byte) error {
	var res SadomainRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
