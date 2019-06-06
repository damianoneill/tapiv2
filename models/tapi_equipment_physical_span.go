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

// TapiEquipmentPhysicalSpan tapi equipment physical span
// swagger:model tapi.equipment.PhysicalSpan
type TapiEquipmentPhysicalSpan struct {
	TapiCommonGlobalClass

	// Both the serial segments that form an end-end strand and the parallel end-end strands.
	AbstractStrand []*TapiEquipmentAbstractStrand `json:"abstract-strand"`

	// none
	AccessPort []*TapiEquipmentAccessPortRef `json:"access-port"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiEquipmentPhysicalSpan) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiCommonGlobalClass
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiCommonGlobalClass = aO0

	// AO1
	var dataAO1 struct {
		AbstractStrand []*TapiEquipmentAbstractStrand `json:"abstract-strand"`

		AccessPort []*TapiEquipmentAccessPortRef `json:"access-port"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AbstractStrand = dataAO1.AbstractStrand

	m.AccessPort = dataAO1.AccessPort

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiEquipmentPhysicalSpan) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.TapiCommonGlobalClass)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		AbstractStrand []*TapiEquipmentAbstractStrand `json:"abstract-strand"`

		AccessPort []*TapiEquipmentAccessPortRef `json:"access-port"`
	}

	dataAO1.AbstractStrand = m.AbstractStrand

	dataAO1.AccessPort = m.AccessPort

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi equipment physical span
func (m *TapiEquipmentPhysicalSpan) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiCommonGlobalClass
	if err := m.TapiCommonGlobalClass.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAbstractStrand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccessPort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEquipmentPhysicalSpan) validateAbstractStrand(formats strfmt.Registry) error {

	if swag.IsZero(m.AbstractStrand) { // not required
		return nil
	}

	for i := 0; i < len(m.AbstractStrand); i++ {
		if swag.IsZero(m.AbstractStrand[i]) { // not required
			continue
		}

		if m.AbstractStrand[i] != nil {
			if err := m.AbstractStrand[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("abstract-strand" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiEquipmentPhysicalSpan) validateAccessPort(formats strfmt.Registry) error {

	if swag.IsZero(m.AccessPort) { // not required
		return nil
	}

	for i := 0; i < len(m.AccessPort); i++ {
		if swag.IsZero(m.AccessPort[i]) { // not required
			continue
		}

		if m.AccessPort[i] != nil {
			if err := m.AccessPort[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("access-port" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEquipmentPhysicalSpan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEquipmentPhysicalSpan) UnmarshalBinary(b []byte) error {
	var res TapiEquipmentPhysicalSpan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}