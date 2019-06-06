// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOamGetmegOutputMegMip tapi oam getmeg output meg mip
// swagger:model tapi.oam.getmeg.output.meg.Mip
type TapiOamGetmegOutputMegMip struct {
	TapiEthMipAugmentation1

	TapiOamMip
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiOamGetmegOutputMegMip) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiEthMipAugmentation1
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiEthMipAugmentation1 = aO0

	// AO1
	var aO1 TapiOamMip
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TapiOamMip = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiOamGetmegOutputMegMip) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.TapiEthMipAugmentation1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TapiOamMip)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi oam getmeg output meg mip
func (m *TapiOamGetmegOutputMegMip) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiEthMipAugmentation1
	if err := m.TapiEthMipAugmentation1.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiOamMip
	if err := m.TapiOamMip.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamGetmegOutputMegMip) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamGetmegOutputMegMip) UnmarshalBinary(b []byte) error {
	var res TapiOamGetmegOutputMegMip
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}