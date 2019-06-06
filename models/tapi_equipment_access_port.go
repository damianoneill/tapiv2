// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEquipmentAccessPort tapi equipment access port
// swagger:model tapi.equipment.AccessPort
type TapiEquipmentAccessPort struct {
	TapiCommonGlobalClass

	// The list of Pins that support the AccessPort.
	ConnectorPin []*TapiEquipmentConnectorPinAddress `json:"connector-pin"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiEquipmentAccessPort) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiCommonGlobalClass
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiCommonGlobalClass = aO0

	// AO1
	var dataAO1 struct {
		ConnectorPin []*TapiEquipmentConnectorPinAddress `json:"connector-pin"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ConnectorPin = dataAO1.ConnectorPin

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiEquipmentAccessPort) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.TapiCommonGlobalClass)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ConnectorPin []*TapiEquipmentConnectorPinAddress `json:"connector-pin"`
	}

	dataAO1.ConnectorPin = m.ConnectorPin

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi equipment access port
func (m *TapiEquipmentAccessPort) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiCommonGlobalClass
	if err := m.TapiCommonGlobalClass.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectorPin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEquipmentAccessPort) validateConnectorPin(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectorPin) { // not required
		return nil
	}

	for i := 0; i < len(m.ConnectorPin); i++ {
		if swag.IsZero(m.ConnectorPin[i]) { // not required
			continue
		}

		if m.ConnectorPin[i] != nil {
			if err := m.ConnectorPin[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connector-pin" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEquipmentAccessPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEquipmentAccessPort) UnmarshalBinary(b []byte) error {
	var res TapiEquipmentAccessPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
