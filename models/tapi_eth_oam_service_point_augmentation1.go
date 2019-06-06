// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthOamServicePointAugmentation1 tapi eth oam service point augmentation1
// swagger:model tapi.eth.OamServicePointAugmentation1
type TapiEthOamServicePointAugmentation1 struct {

	// none
	EthOamMipServicePoint *TapiEthEthOamMipServicePoint `json:"eth-oam-mip-service-point,omitempty"`
}

// Validate validates this tapi eth oam service point augmentation1
func (m *TapiEthOamServicePointAugmentation1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthOamMipServicePoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthOamServicePointAugmentation1) validateEthOamMipServicePoint(formats strfmt.Registry) error {

	if swag.IsZero(m.EthOamMipServicePoint) { // not required
		return nil
	}

	if m.EthOamMipServicePoint != nil {
		if err := m.EthOamMipServicePoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-oam-mip-service-point")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthOamServicePointAugmentation1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthOamServicePointAugmentation1) UnmarshalBinary(b []byte) error {
	var res TapiEthOamServicePointAugmentation1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
