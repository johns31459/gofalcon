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

// RegistrationAWSAccountPatch registration a w s account patch
//
// swagger:model registration.AWSAccountPatch
type RegistrationAWSAccountPatch struct {

	// account id
	// Required: true
	AccountID *string `json:"account_id"`

	// behavior assessment enabled
	BehaviorAssessmentEnabled bool `json:"behavior_assessment_enabled,omitempty"`

	// cloudtrail region
	CloudtrailRegion string `json:"cloudtrail_region,omitempty"`

	// iam role arn
	// Required: true
	IamRoleArn *string `json:"iam_role_arn"`

	// remediation region
	RemediationRegion string `json:"remediation_region,omitempty"`

	// remediation tou accepted
	// Format: date-time
	RemediationTouAccepted strfmt.DateTime `json:"remediation_tou_accepted,omitempty"`

	// sensor management enabled
	SensorManagementEnabled bool `json:"sensor_management_enabled,omitempty"`
}

// Validate validates this registration a w s account patch
func (m *RegistrationAWSAccountPatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIamRoleArn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemediationTouAccepted(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistrationAWSAccountPatch) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("account_id", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *RegistrationAWSAccountPatch) validateIamRoleArn(formats strfmt.Registry) error {

	if err := validate.Required("iam_role_arn", "body", m.IamRoleArn); err != nil {
		return err
	}

	return nil
}

func (m *RegistrationAWSAccountPatch) validateRemediationTouAccepted(formats strfmt.Registry) error {
	if swag.IsZero(m.RemediationTouAccepted) { // not required
		return nil
	}

	if err := validate.FormatOf("remediation_tou_accepted", "body", "date-time", m.RemediationTouAccepted.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this registration a w s account patch based on context it is used
func (m *RegistrationAWSAccountPatch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegistrationAWSAccountPatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistrationAWSAccountPatch) UnmarshalBinary(b []byte) error {
	var res RegistrationAWSAccountPatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
