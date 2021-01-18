// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainDevice domain device
//
// swagger:model domain.Device
type DomainDevice struct {

	// platform ID numeric
	// Required: true
	PlatformIDNumeric *int32 `json:"PlatformIDNumeric"`

	// agent version
	AgentVersion string `json:"agent_version,omitempty"`

	// config id base
	ConfigIDBase string `json:"config_id_base,omitempty"`

	// config id build
	ConfigIDBuild string `json:"config_id_build,omitempty"`

	// config id platform
	ConfigIDPlatform string `json:"config_id_platform,omitempty"`

	// device id
	// Required: true
	DeviceID *string `json:"device_id"`

	// external ip
	ExternalIP string `json:"external_ip,omitempty"`

	// first login timestamp
	FirstLoginTimestamp string `json:"first_login_timestamp,omitempty"`

	// first login user
	FirstLoginUser string `json:"first_login_user,omitempty"`

	// first seen
	FirstSeen string `json:"first_seen,omitempty"`

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// last login timestamp
	LastLoginTimestamp string `json:"last_login_timestamp,omitempty"`

	// last login user
	LastLoginUser string `json:"last_login_user,omitempty"`

	// last seen
	LastSeen string `json:"last_seen,omitempty"`

	// last seen ago seconds
	LastSeenAgoSeconds int64 `json:"last_seen_ago_seconds,omitempty"`

	// local ip
	LocalIP string `json:"local_ip,omitempty"`

	// mac address
	MacAddress string `json:"mac_address,omitempty"`

	// machine domain
	MachineDomain string `json:"machine_domain,omitempty"`

	// major version
	MajorVersion string `json:"major_version,omitempty"`

	// minor version
	MinorVersion string `json:"minor_version,omitempty"`

	// modified timestamp
	ModifiedTimestamp string `json:"modified_timestamp,omitempty"`

	// notes
	Notes []string `json:"notes"`

	// os version
	OsVersion string `json:"os_version,omitempty"`

	// ou
	Ou []string `json:"ou"`

	// platform id
	PlatformID string `json:"platform_id,omitempty"`

	// platform name
	PlatformName string `json:"platform_name,omitempty"`

	// product type
	ProductType string `json:"product_type,omitempty"`

	// product type desc
	ProductTypeDesc string `json:"product_type_desc,omitempty"`

	// release group
	ReleaseGroup string `json:"release_group,omitempty"`

	// site name
	SiteName string `json:"site_name,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// system manufacturer
	SystemManufacturer string `json:"system_manufacturer,omitempty"`

	// system product name
	SystemProductName string `json:"system_product_name,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this domain device
func (m *DomainDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlatformIDNumeric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainDevice) validatePlatformIDNumeric(formats strfmt.Registry) error {

	if err := validate.Required("PlatformIDNumeric", "body", m.PlatformIDNumeric); err != nil {
		return err
	}

	return nil
}

func (m *DomainDevice) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("device_id", "body", m.DeviceID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainDevice) UnmarshalBinary(b []byte) error {
	var res DomainDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}