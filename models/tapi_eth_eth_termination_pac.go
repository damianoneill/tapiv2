// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthEthTerminationPac tapi eth eth termination pac
// swagger:model tapi.eth.EthTerminationPac
type TapiEthEthTerminationPac struct {

	// none
	EthTerminationCommonPac *TapiEthEthTerminationCommonPac `json:"eth-termination-common-pac,omitempty"`
}

// Validate validates this tapi eth eth termination pac
func (m *TapiEthEthTerminationPac) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthTerminationCommonPac(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthEthTerminationPac) validateEthTerminationCommonPac(formats strfmt.Registry) error {

	if swag.IsZero(m.EthTerminationCommonPac) { // not required
		return nil
	}

	if m.EthTerminationCommonPac != nil {
		if err := m.EthTerminationCommonPac.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-termination-common-pac")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthEthTerminationPac) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthEthTerminationPac) UnmarshalBinary(b []byte) error {
	var res TapiEthEthTerminationPac
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}