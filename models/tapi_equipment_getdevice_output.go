// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEquipmentGetdeviceOutput tapi equipment getdevice output
// swagger:model tapi.equipment.getdevice.Output
type TapiEquipmentGetdeviceOutput struct {

	// none
	Device *TapiEquipmentDevice `json:"device,omitempty"`
}

// Validate validates this tapi equipment getdevice output
func (m *TapiEquipmentGetdeviceOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEquipmentGetdeviceOutput) validateDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.Device) { // not required
		return nil
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEquipmentGetdeviceOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEquipmentGetdeviceOutput) UnmarshalBinary(b []byte) error {
	var res TapiEquipmentGetdeviceOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
