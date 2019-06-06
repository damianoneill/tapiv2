// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiEquipmentGetphysicalspanInput tapi equipment getphysicalspan input
// swagger:model tapi.equipment.getphysicalspan.Input
type TapiEquipmentGetphysicalspanInput struct {

	// none
	SpanID string `json:"span-id,omitempty"`
}

// Validate validates this tapi equipment getphysicalspan input
func (m *TapiEquipmentGetphysicalspanInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiEquipmentGetphysicalspanInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEquipmentGetphysicalspanInput) UnmarshalBinary(b []byte) error {
	var res TapiEquipmentGetphysicalspanInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}