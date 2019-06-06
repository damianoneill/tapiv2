// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOduDegThr tapi odu deg thr
// swagger:model tapi.odu.DegThr
type TapiOduDegThr struct {

	// Number of errored blocks
	DegThrType TapiOduDegThrType `json:"deg-thr-type,omitempty"`

	// Percentage of detected errored blocks
	DegThrValue int32 `json:"deg-thr-value,omitempty"`

	// none
	PercentageGranularity TapiOduPercentageGranularity `json:"percentage-granularity,omitempty"`
}

// Validate validates this tapi odu deg thr
func (m *TapiOduDegThr) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDegThrType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePercentageGranularity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOduDegThr) validateDegThrType(formats strfmt.Registry) error {

	if swag.IsZero(m.DegThrType) { // not required
		return nil
	}

	if err := m.DegThrType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deg-thr-type")
		}
		return err
	}

	return nil
}

func (m *TapiOduDegThr) validatePercentageGranularity(formats strfmt.Registry) error {

	if swag.IsZero(m.PercentageGranularity) { // not required
		return nil
	}

	if err := m.PercentageGranularity.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("percentage-granularity")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiOduDegThr) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOduDegThr) UnmarshalBinary(b []byte) error {
	var res TapiOduDegThr
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
