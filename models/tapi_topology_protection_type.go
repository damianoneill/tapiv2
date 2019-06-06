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

// TapiTopologyProtectionType tapi topology protection type
// swagger:model tapi.topology.ProtectionType
type TapiTopologyProtectionType string

const (

	// TapiTopologyProtectionTypeNOPROTECTION captures enum value "NO_PROTECTION"
	TapiTopologyProtectionTypeNOPROTECTION TapiTopologyProtectionType = "NO_PROTECTION"

	// TapiTopologyProtectionTypeONEPLUSONEPROTECTION captures enum value "ONE_PLUS_ONE_PROTECTION"
	TapiTopologyProtectionTypeONEPLUSONEPROTECTION TapiTopologyProtectionType = "ONE_PLUS_ONE_PROTECTION"

	// TapiTopologyProtectionTypeONEFORONEPROTECTION captures enum value "ONE_FOR_ONE_PROTECTION"
	TapiTopologyProtectionTypeONEFORONEPROTECTION TapiTopologyProtectionType = "ONE_FOR_ONE_PROTECTION"

	// TapiTopologyProtectionTypeONEFORNPROTECTION captures enum value "ONE_FOR_N_PROTECTION"
	TapiTopologyProtectionTypeONEFORNPROTECTION TapiTopologyProtectionType = "ONE_FOR_N_PROTECTION"

	// TapiTopologyProtectionTypeMFORNPROTECTION captures enum value "M_FOR_N_PROTECTION"
	TapiTopologyProtectionTypeMFORNPROTECTION TapiTopologyProtectionType = "M_FOR_N_PROTECTION"

	// TapiTopologyProtectionTypeONEFORONEBYN captures enum value "ONE_FOR_ONE_BY_N"
	TapiTopologyProtectionTypeONEFORONEBYN TapiTopologyProtectionType = "ONE_FOR_ONE_BY_N"
)

// for schema
var tapiTopologyProtectionTypeEnum []interface{}

func init() {
	var res []TapiTopologyProtectionType
	if err := json.Unmarshal([]byte(`["NO_PROTECTION","ONE_PLUS_ONE_PROTECTION","ONE_FOR_ONE_PROTECTION","ONE_FOR_N_PROTECTION","M_FOR_N_PROTECTION","ONE_FOR_ONE_BY_N"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiTopologyProtectionTypeEnum = append(tapiTopologyProtectionTypeEnum, v)
	}
}

func (m TapiTopologyProtectionType) validateTapiTopologyProtectionTypeEnum(path, location string, value TapiTopologyProtectionType) error {
	if err := validate.Enum(path, location, value, tapiTopologyProtectionTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi topology protection type
func (m TapiTopologyProtectionType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiTopologyProtectionTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}