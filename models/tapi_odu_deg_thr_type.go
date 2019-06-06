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

// TapiOduDegThrType tapi odu deg thr type
// swagger:model tapi.odu.DegThrType
type TapiOduDegThrType string

const (

	// TapiOduDegThrTypePERCENTAGE captures enum value "PERCENTAGE"
	TapiOduDegThrTypePERCENTAGE TapiOduDegThrType = "PERCENTAGE"

	// TapiOduDegThrTypeNUMBERERROREDBLOCKS captures enum value "NUMBER_ERRORED_BLOCKS"
	TapiOduDegThrTypeNUMBERERROREDBLOCKS TapiOduDegThrType = "NUMBER_ERRORED_BLOCKS"
)

// for schema
var tapiOduDegThrTypeEnum []interface{}

func init() {
	var res []TapiOduDegThrType
	if err := json.Unmarshal([]byte(`["PERCENTAGE","NUMBER_ERRORED_BLOCKS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiOduDegThrTypeEnum = append(tapiOduDegThrTypeEnum, v)
	}
}

func (m TapiOduDegThrType) validateTapiOduDegThrTypeEnum(path, location string, value TapiOduDegThrType) error {
	if err := validate.Enum(path, location, value, tapiOduDegThrTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi odu deg thr type
func (m TapiOduDegThrType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiOduDegThrTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
