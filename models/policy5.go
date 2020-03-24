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

// Policy5 Policy5
//
// 'Deny' traffic specified by this rule
//
// swagger:model Policy5
type Policy5 string

const (

	// Policy5Deny captures enum value "deny"
	Policy5Deny Policy5 = "deny"
)

// for schema
var policy5Enum []interface{}

func init() {
	var res []Policy5
	if err := json.Unmarshal([]byte(`["deny"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policy5Enum = append(policy5Enum, v)
	}
}

func (m Policy5) validatePolicy5Enum(path, location string, value Policy5) error {
	if err := validate.Enum(path, location, value, policy5Enum); err != nil {
		return err
	}
	return nil
}

// Validate validates this policy5
func (m Policy5) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePolicy5Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}