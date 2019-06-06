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

// TapiEthEthServiceIntefacePointSpec tapi eth eth service inteface point spec
// swagger:model tapi.eth.EthServiceIntefacePointSpec
type TapiEthEthServiceIntefacePointSpec struct {

	// This attribute identifies the PHY type of the ETY trail termination. See IEEE 802.3 clause 30.3.2.1.2.
	PhyType TapiEthEtyPhyType `json:"phy-type,omitempty"`

	// This attribute identifies the possible PHY types that could be supported at the ETY trail termination. See IEEE 802.3 clause 30.3.2.1.3.
	PhyTypeList []TapiEthEtyPhyType `json:"phy-type-list"`
}

// Validate validates this tapi eth eth service inteface point spec
func (m *TapiEthEthServiceIntefacePointSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhyType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhyTypeList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthEthServiceIntefacePointSpec) validatePhyType(formats strfmt.Registry) error {

	if swag.IsZero(m.PhyType) { // not required
		return nil
	}

	if err := m.PhyType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("phy-type")
		}
		return err
	}

	return nil
}

func (m *TapiEthEthServiceIntefacePointSpec) validatePhyTypeList(formats strfmt.Registry) error {

	if swag.IsZero(m.PhyTypeList) { // not required
		return nil
	}

	for i := 0; i < len(m.PhyTypeList); i++ {

		if err := m.PhyTypeList[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phy-type-list" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthEthServiceIntefacePointSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthEthServiceIntefacePointSpec) UnmarshalBinary(b []byte) error {
	var res TapiEthEthServiceIntefacePointSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
