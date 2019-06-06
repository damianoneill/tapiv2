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

// TapiPhotonicMediaGridType tapi photonic media grid type
// swagger:model tapi.photonic.media.GridType
type TapiPhotonicMediaGridType string

const (

	// TapiPhotonicMediaGridTypeDWDM captures enum value "DWDM"
	TapiPhotonicMediaGridTypeDWDM TapiPhotonicMediaGridType = "DWDM"

	// TapiPhotonicMediaGridTypeCWDM captures enum value "CWDM"
	TapiPhotonicMediaGridTypeCWDM TapiPhotonicMediaGridType = "CWDM"

	// TapiPhotonicMediaGridTypeFLEX captures enum value "FLEX"
	TapiPhotonicMediaGridTypeFLEX TapiPhotonicMediaGridType = "FLEX"

	// TapiPhotonicMediaGridTypeGRIDLESS captures enum value "GRIDLESS"
	TapiPhotonicMediaGridTypeGRIDLESS TapiPhotonicMediaGridType = "GRIDLESS"

	// TapiPhotonicMediaGridTypeUNSPECIFIED captures enum value "UNSPECIFIED"
	TapiPhotonicMediaGridTypeUNSPECIFIED TapiPhotonicMediaGridType = "UNSPECIFIED"
)

// for schema
var tapiPhotonicMediaGridTypeEnum []interface{}

func init() {
	var res []TapiPhotonicMediaGridType
	if err := json.Unmarshal([]byte(`["DWDM","CWDM","FLEX","GRIDLESS","UNSPECIFIED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiPhotonicMediaGridTypeEnum = append(tapiPhotonicMediaGridTypeEnum, v)
	}
}

func (m TapiPhotonicMediaGridType) validateTapiPhotonicMediaGridTypeEnum(path, location string, value TapiPhotonicMediaGridType) error {
	if err := validate.Enum(path, location, value, tapiPhotonicMediaGridTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi photonic media grid type
func (m TapiPhotonicMediaGridType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiPhotonicMediaGridTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
