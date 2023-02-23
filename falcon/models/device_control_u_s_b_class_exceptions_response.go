// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeviceControlUSBClassExceptionsResponse device control u s b class exceptions response
//
// swagger:model device_control.USBClassExceptionsResponse
type DeviceControlUSBClassExceptionsResponse struct {

	// Policy action
	// Required: true
	// Enum: [FULL_ACCESS FULL_BLOCK READ_ONLY]
	Action *string `json:"action"`

	// Exceptions to the rules of this policy setting
	// Required: true
	Exceptions []*DeviceControlExceptionRespV1 `json:"exceptions"`

	// USB Class id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this device control u s b class exceptions response
func (m *DeviceControlUSBClassExceptionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExceptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var deviceControlUSBClassExceptionsResponseTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FULL_ACCESS","FULL_BLOCK","READ_ONLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceControlUSBClassExceptionsResponseTypeActionPropEnum = append(deviceControlUSBClassExceptionsResponseTypeActionPropEnum, v)
	}
}

const (

	// DeviceControlUSBClassExceptionsResponseActionFULLACCESS captures enum value "FULL_ACCESS"
	DeviceControlUSBClassExceptionsResponseActionFULLACCESS string = "FULL_ACCESS"

	// DeviceControlUSBClassExceptionsResponseActionFULLBLOCK captures enum value "FULL_BLOCK"
	DeviceControlUSBClassExceptionsResponseActionFULLBLOCK string = "FULL_BLOCK"

	// DeviceControlUSBClassExceptionsResponseActionREADONLY captures enum value "READ_ONLY"
	DeviceControlUSBClassExceptionsResponseActionREADONLY string = "READ_ONLY"
)

// prop value enum
func (m *DeviceControlUSBClassExceptionsResponse) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deviceControlUSBClassExceptionsResponseTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeviceControlUSBClassExceptionsResponse) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *DeviceControlUSBClassExceptionsResponse) validateExceptions(formats strfmt.Registry) error {

	if err := validate.Required("exceptions", "body", m.Exceptions); err != nil {
		return err
	}

	for i := 0; i < len(m.Exceptions); i++ {
		if swag.IsZero(m.Exceptions[i]) { // not required
			continue
		}

		if m.Exceptions[i] != nil {
			if err := m.Exceptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("exceptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("exceptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceControlUSBClassExceptionsResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this device control u s b class exceptions response based on the context it is used
func (m *DeviceControlUSBClassExceptionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExceptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceControlUSBClassExceptionsResponse) contextValidateExceptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Exceptions); i++ {

		if m.Exceptions[i] != nil {
			if err := m.Exceptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("exceptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("exceptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceControlUSBClassExceptionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceControlUSBClassExceptionsResponse) UnmarshalBinary(b []byte) error {
	var res DeviceControlUSBClassExceptionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}