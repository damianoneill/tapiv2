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

// TapiConnectivityServiceType tapi connectivity service type
// swagger:model tapi.connectivity.ServiceType
type TapiConnectivityServiceType string

const (

	// TapiConnectivityServiceTypePOINTTOPOINTCONNECTIVITY captures enum value "POINT_TO_POINT_CONNECTIVITY"
	TapiConnectivityServiceTypePOINTTOPOINTCONNECTIVITY TapiConnectivityServiceType = "POINT_TO_POINT_CONNECTIVITY"

	// TapiConnectivityServiceTypePOINTTOMULTIPOINTCONNECTIVITY captures enum value "POINT_TO_MULTIPOINT_CONNECTIVITY"
	TapiConnectivityServiceTypePOINTTOMULTIPOINTCONNECTIVITY TapiConnectivityServiceType = "POINT_TO_MULTIPOINT_CONNECTIVITY"

	// TapiConnectivityServiceTypeMULTIPOINTCONNECTIVITY captures enum value "MULTIPOINT_CONNECTIVITY"
	TapiConnectivityServiceTypeMULTIPOINTCONNECTIVITY TapiConnectivityServiceType = "MULTIPOINT_CONNECTIVITY"

	// TapiConnectivityServiceTypeROOTEDMULTIPOINTCONNECTIVITY captures enum value "ROOTED_MULTIPOINT_CONNECTIVITY"
	TapiConnectivityServiceTypeROOTEDMULTIPOINTCONNECTIVITY TapiConnectivityServiceType = "ROOTED_MULTIPOINT_CONNECTIVITY"
)

// for schema
var tapiConnectivityServiceTypeEnum []interface{}

func init() {
	var res []TapiConnectivityServiceType
	if err := json.Unmarshal([]byte(`["POINT_TO_POINT_CONNECTIVITY","POINT_TO_MULTIPOINT_CONNECTIVITY","MULTIPOINT_CONNECTIVITY","ROOTED_MULTIPOINT_CONNECTIVITY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiConnectivityServiceTypeEnum = append(tapiConnectivityServiceTypeEnum, v)
	}
}

func (m TapiConnectivityServiceType) validateTapiConnectivityServiceTypeEnum(path, location string, value TapiConnectivityServiceType) error {
	if err := validate.Enum(path, location, value, tapiConnectivityServiceTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi connectivity service type
func (m TapiConnectivityServiceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiConnectivityServiceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
