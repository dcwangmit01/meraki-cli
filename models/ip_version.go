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

// IPVersion IpVersion
//
// IP address version (must be 'any', 'ipv4' or 'ipv6'). Applicable only if network supports IPv6. Default value is 'ipv4'.
//
// swagger:model IpVersion
type IPVersion string

const (

	// IPVersionAny captures enum value "any"
	IPVersionAny IPVersion = "any"

	// IPVersionIPV4 captures enum value "ipv4"
	IPVersionIPV4 IPVersion = "ipv4"

	// IPVersionIPV6 captures enum value "ipv6"
	IPVersionIPV6 IPVersion = "ipv6"
)

// for schema
var ipVersionEnum []interface{}

func init() {
	var res []IPVersion
	if err := json.Unmarshal([]byte(`["any","ipv4","ipv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipVersionEnum = append(ipVersionEnum, v)
	}
}

func (m IPVersion) validateIPVersionEnum(path, location string, value IPVersion) error {
	if err := validate.Enum(path, location, value, ipVersionEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this Ip version
func (m IPVersion) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIPVersionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
