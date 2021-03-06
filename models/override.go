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

// Override Override
//
// swagger:model Override
type Override struct {

	// MTU size for the switches or switch profiles.
	// Required: true
	MtuSize *int32 `json:"mtuSize"`

	// List of switch profile IDs. Applicable only for template network.
	SwitchProfiles []string `json:"switchProfiles"`

	// List of switch serials. Applicable only for switch network.
	Switches []string `json:"switches"`
}

// Validate validates this override
func (m *Override) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMtuSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Override) validateMtuSize(formats strfmt.Registry) error {

	if err := validate.Required("mtuSize", "body", m.MtuSize); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Override) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Override) UnmarshalBinary(b []byte) error {
	var res Override
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
