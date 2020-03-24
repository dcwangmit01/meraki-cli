// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AllowedURLPatterns AllowedUrlPatterns
//
// Settings for whitelisted URL patterns
//
// swagger:model AllowedUrlPatterns
type AllowedURLPatterns struct {

	// A whitelist of URL patterns to allow
	Patterns []string `json:"patterns"`

	// settings
	Settings Settings2 `json:"settings,omitempty"`
}

// Validate validates this allowed Url patterns
func (m *AllowedURLPatterns) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllowedURLPatterns) validateSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	if err := m.Settings.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("settings")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AllowedURLPatterns) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllowedURLPatterns) UnmarshalBinary(b []byte) error {
	var res AllowedURLPatterns
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}