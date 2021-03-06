// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Uplink uplink
//
// swagger:model uplink
type Uplink string

const (

	// UplinkWan1 captures enum value "wan1"
	UplinkWan1 Uplink = "wan1"

	// UplinkWan2 captures enum value "wan2"
	UplinkWan2 Uplink = "wan2"

	// UplinkCellular captures enum value "cellular"
	UplinkCellular Uplink = "cellular"
)

// for schema
var uplinkEnum []interface{}

func init() {
	var res []Uplink
	if err := json.Unmarshal([]byte(`["wan1","wan2","cellular"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		uplinkEnum = append(uplinkEnum, v)
	}
}

func (m Uplink) validateUplinkEnum(path, location string, value Uplink) error {
	if err := validate.Enum(path, location, value, uplinkEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this uplink
func (m Uplink) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUplinkEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
