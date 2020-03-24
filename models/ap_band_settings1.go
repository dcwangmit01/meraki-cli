// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApBandSettings1 ApBandSettings1
//
// Settings that will be enabled if selectionType is set to 'ap'.
//
// swagger:model ApBandSettings1
type ApBandSettings1 struct {

	// band operation mode
	BandOperationMode BandOperationMode1 `json:"bandOperationMode,omitempty"`

	// Steers client to most open band. Can be either true or false.
	BandSteeringEnabled bool `json:"bandSteeringEnabled,omitempty"`
}

// Validate validates this ap band settings1
func (m *ApBandSettings1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBandOperationMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApBandSettings1) validateBandOperationMode(formats strfmt.Registry) error {

	if swag.IsZero(m.BandOperationMode) { // not required
		return nil
	}

	if err := m.BandOperationMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("bandOperationMode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApBandSettings1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApBandSettings1) UnmarshalBinary(b []byte) error {
	var res ApBandSettings1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}