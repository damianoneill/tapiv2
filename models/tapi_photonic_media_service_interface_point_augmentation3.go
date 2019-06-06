// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiPhotonicMediaServiceInterfacePointAugmentation3 tapi photonic media service interface point augmentation3
// swagger:model tapi.photonic.media.ServiceInterfacePointAugmentation3
type TapiPhotonicMediaServiceInterfacePointAugmentation3 struct {

	// none
	OtsiServiceInterfacePointSpec *TapiPhotonicMediaOtsiServiceInterfacePointSpec `json:"otsi-service-interface-point-spec,omitempty"`
}

// Validate validates this tapi photonic media service interface point augmentation3
func (m *TapiPhotonicMediaServiceInterfacePointAugmentation3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOtsiServiceInterfacePointSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiPhotonicMediaServiceInterfacePointAugmentation3) validateOtsiServiceInterfacePointSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.OtsiServiceInterfacePointSpec) { // not required
		return nil
	}

	if m.OtsiServiceInterfacePointSpec != nil {
		if err := m.OtsiServiceInterfacePointSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("otsi-service-interface-point-spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiPhotonicMediaServiceInterfacePointAugmentation3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiPhotonicMediaServiceInterfacePointAugmentation3) UnmarshalBinary(b []byte) error {
	var res TapiPhotonicMediaServiceInterfacePointAugmentation3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
