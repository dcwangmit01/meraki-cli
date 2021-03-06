// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateNetworkSwitchAccessControlLists updateNetworkSwitchAccessControlLists
//
// swagger:model updateNetworkSwitchAccessControlLists
type UpdateNetworkSwitchAccessControlLists struct {

	// An ordered array of the access control list rules (not including the default rule). An empty array will clear the rules.
	// Required: true
	Rules []*Rule12 `json:"rules"`
}

// Validate validates this update network switch access control lists
func (m *UpdateNetworkSwitchAccessControlLists) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateNetworkSwitchAccessControlLists) validateRules(formats strfmt.Registry) error {

	if err := validate.Required("rules", "body", m.Rules); err != nil {
		return err
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
func (m *UpdateNetworkSwitchAccessControlLists) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNetworkSwitchAccessControlLists) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchAccessControlLists
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
