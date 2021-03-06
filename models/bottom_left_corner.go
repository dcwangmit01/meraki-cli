// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BottomLeftCorner BottomLeftCorner
//
// The longitude and latitude of the bottom left corner of your floor plan.
//
// swagger:model BottomLeftCorner
type BottomLeftCorner struct {

	// Latitude
	Lat float64 `json:"lat,omitempty"`

	// Longitude
	Lng float64 `json:"lng,omitempty"`
}

// Validate validates this bottom left corner
func (m *BottomLeftCorner) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BottomLeftCorner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BottomLeftCorner) UnmarshalBinary(b []byte) error {
	var res BottomLeftCorner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
