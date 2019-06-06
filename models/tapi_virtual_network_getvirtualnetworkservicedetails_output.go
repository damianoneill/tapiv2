// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiVirtualNetworkGetvirtualnetworkservicedetailsOutput tapi virtual network getvirtualnetworkservicedetails output
// swagger:model tapi.virtual.network.getvirtualnetworkservicedetails.Output
type TapiVirtualNetworkGetvirtualnetworkservicedetailsOutput struct {

	// none
	Service *TapiVirtualNetworkVirtualNetworkService `json:"service,omitempty"`
}

// Validate validates this tapi virtual network getvirtualnetworkservicedetails output
func (m *TapiVirtualNetworkGetvirtualnetworkservicedetailsOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiVirtualNetworkGetvirtualnetworkservicedetailsOutput) validateService(formats strfmt.Registry) error {

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
func (m *TapiVirtualNetworkGetvirtualnetworkservicedetailsOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiVirtualNetworkGetvirtualnetworkservicedetailsOutput) UnmarshalBinary(b []byte) error {
	var res TapiVirtualNetworkGetvirtualnetworkservicedetailsOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}