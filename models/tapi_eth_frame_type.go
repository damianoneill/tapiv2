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

// TapiEthFrameType tapi eth frame type
// swagger:model tapi.eth.FrameType
type TapiEthFrameType string

const (

	// TapiEthFrameTypeADMITONLYVLANTAGGEDFRAMES captures enum value "ADMIT_ONLY_VLAN_TAGGED_FRAMES"
	TapiEthFrameTypeADMITONLYVLANTAGGEDFRAMES TapiEthFrameType = "ADMIT_ONLY_VLAN_TAGGED_FRAMES"

	// TapiEthFrameTypeADMITONLYUNTAGGEDANDPRIORITYTAGGEDFRAMES captures enum value "ADMIT_ONLY_UNTAGGED_AND_PRIORITY_TAGGED_FRAMES"
	TapiEthFrameTypeADMITONLYUNTAGGEDANDPRIORITYTAGGEDFRAMES TapiEthFrameType = "ADMIT_ONLY_UNTAGGED_AND_PRIORITY_TAGGED_FRAMES"

	// TapiEthFrameTypeADMITALLFRAMES captures enum value "ADMIT_ALL_FRAMES"
	TapiEthFrameTypeADMITALLFRAMES TapiEthFrameType = "ADMIT_ALL_FRAMES"
)

// for schema
var tapiEthFrameTypeEnum []interface{}

func init() {
	var res []TapiEthFrameType
	if err := json.Unmarshal([]byte(`["ADMIT_ONLY_VLAN_TAGGED_FRAMES","ADMIT_ONLY_UNTAGGED_AND_PRIORITY_TAGGED_FRAMES","ADMIT_ALL_FRAMES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiEthFrameTypeEnum = append(tapiEthFrameTypeEnum, v)
	}
}

func (m TapiEthFrameType) validateTapiEthFrameTypeEnum(path, location string, value TapiEthFrameType) error {
	if err := validate.Enum(path, location, value, tapiEthFrameTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi eth frame type
func (m TapiEthFrameType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiEthFrameTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}