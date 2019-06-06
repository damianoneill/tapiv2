// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiVirtualNetworkDeletevirtualnetworkserviceOutput tapi virtual network deletevirtualnetworkservice output
// swagger:model tapi.virtual.network.deletevirtualnetworkservice.Output
type TapiVirtualNetworkDeletevirtualnetworkserviceOutput struct {

	// none
	Service *TapiVirtualNetworkVirtualNetworkService `json:"service,omitempty"`
}

// Validate validates this tapi virtual network deletevirtualnetworkservice output
func (m *TapiVirtualNetworkDeletevirtualnetworkserviceOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiVirtualNetworkDeletevirtualnetworkserviceOutput) validateService(formats strfmt.Registry) error {

	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiVirtualNetworkDeletevirtualnetworkserviceOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiVirtualNetworkDeletevirtualnetworkserviceOutput) UnmarshalBinary(b []byte) error {
	var res TapiVirtualNetworkDeletevirtualnetworkserviceOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
