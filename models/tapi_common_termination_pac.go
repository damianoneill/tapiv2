// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiCommonTerminationPac tapi common termination pac
// swagger:model tapi.common.TerminationPac
type TapiCommonTerminationPac struct {

	// The overall directionality of the LP.
	//                 - A BIDIRECTIONAL LP will have some SINK and/or SOURCE flowss.
	//                 - A SINK LP can only contain elements with SINK flows or CONTRA_DIRECTION_SOURCE flows
	//                 - A SOURCE LP can only contain SOURCE flows or CONTRA_DIRECTION_SINK flows
	TerminationDirection TapiCommonTerminationDirection `json:"termination-direction,omitempty"`

	// Indicates whether the layer is terminated and if so how.
	TerminationState TapiCommonTerminationState `json:"termination-state,omitempty"`
}

// Validate validates this tapi common termination pac
func (m *TapiCommonTerminationPac) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTerminationDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiCommonTerminationPac) validateTerminationDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.TerminationDirection) { // not required
		return nil
	}

	if err := m.TerminationDirection.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("termination-direction")
		}
		return err
	}

	return nil
}

func (m *TapiCommonTerminationPac) validateTerminationState(formats strfmt.Registry) error {

	if swag.IsZero(m.TerminationState) { // not required
		return nil
	}

	if err := m.TerminationState.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("termination-state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiCommonTerminationPac) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiCommonTerminationPac) UnmarshalBinary(b []byte) error {
	var res TapiCommonTerminationPac
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
