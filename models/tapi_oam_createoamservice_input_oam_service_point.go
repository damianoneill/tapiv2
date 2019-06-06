// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOamCreateoamserviceInputOamServicePoint tapi oam createoamservice input oam service point
// swagger:model tapi.oam.createoamservice.input.OamServicePoint
type TapiOamCreateoamserviceInputOamServicePoint struct {
	TapiEthOamServicePointAugmentation1

	TapiEthOamServicePointAugmentation2

	TapiEthOamServicePointAugmentation3

	TapiOamOamServicePoint
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiOamCreateoamserviceInputOamServicePoint) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiEthOamServicePointAugmentation1
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiEthOamServicePointAugmentation1 = aO0

	// AO1
	var aO1 TapiEthOamServicePointAugmentation2
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TapiEthOamServicePointAugmentation2 = aO1

	// AO2
	var aO2 TapiEthOamServicePointAugmentation3
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.TapiEthOamServicePointAugmentation3 = aO2

	// AO3
	var aO3 TapiOamOamServicePoint
	if err := swag.ReadJSON(raw, &aO3); err != nil {
		return err
	}
	m.TapiOamOamServicePoint = aO3

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiOamCreateoamserviceInputOamServicePoint) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 4)

	aO0, err := swag.WriteJSON(m.TapiEthOamServicePointAugmentation1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TapiEthOamServicePointAugmentation2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.TapiEthOamServicePointAugmentation3)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	aO3, err := swag.WriteJSON(m.TapiOamOamServicePoint)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO3)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi oam createoamservice input oam service point
func (m *TapiOamCreateoamserviceInputOamServicePoint) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiEthOamServicePointAugmentation1
	if err := m.TapiEthOamServicePointAugmentation1.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiEthOamServicePointAugmentation2
	if err := m.TapiEthOamServicePointAugmentation2.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiEthOamServicePointAugmentation3
	if err := m.TapiEthOamServicePointAugmentation3.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiOamOamServicePoint
	if err := m.TapiOamOamServicePoint.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamCreateoamserviceInputOamServicePoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamCreateoamserviceInputOamServicePoint) UnmarshalBinary(b []byte) error {
	var res TapiOamCreateoamserviceInputOamServicePoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
