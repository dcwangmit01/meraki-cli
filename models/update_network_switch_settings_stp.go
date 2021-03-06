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

// UpdateNetworkSwitchSettingsStp updateNetworkSwitchSettingsStp
//
// swagger:model updateNetworkSwitchSettingsStp
type UpdateNetworkSwitchSettingsStp struct {

	// The spanning tree protocol status in network
	RstpEnabled bool `json:"rstpEnabled,omitempty"`

	// STP bridge priority for switches/stacks or switch profiles. An empty array will clear the STP bridge priority settings.
	StpBridgePriority []*StpBridgePriority `json:"stpBridgePriority"`
}

// Validate validates this update network switch settings stp
func (m *UpdateNetworkSwitchSettingsStp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStpBridgePriority(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateNetworkSwitchSettingsStp) validateStpBridgePriority(formats strfmt.Registry) error {

	if swag.IsZero(m.StpBridgePriority) { // not required
		return nil
	}

	for i := 0; i < len(m.StpBridgePriority); i++ {
		if swag.IsZero(m.StpBridgePriority[i]) { // not required
			continue
		}

		if m.StpBridgePriority[i] != nil {
			if err := m.StpBridgePriority[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stpBridgePriority" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateNetworkSwitchSettingsStp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNetworkSwitchSettingsStp) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchSettingsStp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
