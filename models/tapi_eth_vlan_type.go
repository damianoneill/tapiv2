// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// TapiEthVlanType tapi eth vlan type
// swagger:model tapi.eth.VlanType
type TapiEthVlanType string

const (

	// TapiEthVlanTypeCTag captures enum value "C_Tag"
	TapiEthVlanTypeCTag TapiEthVlanType = "C_Tag"

	// TapiEthVlanTypeSTag captures enum value "S_Tag"
	TapiEthVlanTypeSTag TapiEthVlanType = "S_Tag"

	// TapiEthVlanTypeITag captures enum value "I_Tag"
	TapiEthVlanTypeITag TapiEthVlanType = "I_Tag"
)

// for schema
var tapiEthVlanTypeEnum []interface{}

func init() {
	var res []TapiEthVlanType
	if err := json.Unmarshal([]byte(`["C_Tag","S_Tag","I_Tag"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiEthVlanTypeEnum = append(tapiEthVlanTypeEnum, v)
	}
}

func (m TapiEthVlanType) validateTapiEthVlanTypeEnum(path, location string, value TapiEthVlanType) error {
	if err := validate.Enum(path, location, value, tapiEthVlanTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi eth vlan type
func (m TapiEthVlanType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiEthVlanTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
