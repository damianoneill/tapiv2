// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiOamDeleteoamprofileInput tapi oam deleteoamprofile input
// swagger:model tapi.oam.deleteoamprofile.Input
type TapiOamDeleteoamprofileInput struct {

	// none
	OamProfileID string `json:"oam-profile-id,omitempty"`
}

// Validate validates this tapi oam deleteoamprofile input
func (m *TapiOamDeleteoamprofileInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamDeleteoamprofileInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamDeleteoamprofileInput) UnmarshalBinary(b []byte) error {
	var res TapiOamDeleteoamprofileInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
