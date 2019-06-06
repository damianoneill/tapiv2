// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiPathComputationPathRef tapi path computation path ref
// swagger:model tapi.path.computation.PathRef
type TapiPathComputationPathRef struct {

	// none
	PathUUID string `json:"path-uuid,omitempty"`
}

// Validate validates this tapi path computation path ref
func (m *TapiPathComputationPathRef) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiPathComputationPathRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiPathComputationPathRef) UnmarshalBinary(b []byte) error {
	var res TapiPathComputationPathRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}