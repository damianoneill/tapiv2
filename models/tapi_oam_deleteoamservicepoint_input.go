// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiOamDeleteoamservicepointInput tapi oam deleteoamservicepoint input
// swagger:model tapi.oam.deleteoamservicepoint.Input
type TapiOamDeleteoamservicepointInput struct {

	// none
	OamServiceID string `json:"oam-service-id,omitempty"`

	// none
	OamServicePointID string `json:"oam-service-point-id,omitempty"`
}

// Validate validates this tapi oam deleteoamservicepoint input
func (m *TapiOamDeleteoamservicepointInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamDeleteoamservicepointInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamDeleteoamservicepointInput) UnmarshalBinary(b []byte) error {
	var res TapiOamDeleteoamservicepointInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
