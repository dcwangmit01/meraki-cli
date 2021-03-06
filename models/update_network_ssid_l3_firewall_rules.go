// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateNetworkSsidL3FirewallRules updateNetworkSsidL3FirewallRules
//
// swagger:model updateNetworkSsidL3FirewallRules
type UpdateNetworkSsidL3FirewallRules struct {

	// Allow wireless client access to local LAN (boolean value - true allows access and false denies access) (optional)
	AllowLanAccess bool `json:"allowLanAccess,omitempty"`

	// An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule)
	Rules []*Rule10 `json:"rules"`
}

// Validate validates this update network ssid l3 firewall rules
func (m *UpdateNetworkSsidL3FirewallRules) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateNetworkSsidL3FirewallRules) validateRules(formats strfmt.Registry) error {

	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	for i := 0; i < len(m.Rules); i++ {
		if swag.IsZero(m.Rules[i]) { // not required
			continue
		}

		if m.Rules[i] != nil {
			if err := m.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateNetworkSsidL3FirewallRules) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNetworkSsidL3FirewallRules) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSsidL3FirewallRules
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
