// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiPathComputationPathObjectiveFunction tapi path computation path objective function
// swagger:model tapi.path.computation.PathObjectiveFunction
type TapiPathComputationPathObjectiveFunction struct {
	TapiCommonLocalClass

	// none
	BandwidthOptimization TapiCommonDirectiveValue `json:"bandwidth-optimization,omitempty"`

	// none
	ConcurrentPaths TapiCommonDirectiveValue `json:"concurrent-paths,omitempty"`

	// none
	CostOptimization TapiCommonDirectiveValue `json:"cost-optimization,omitempty"`

	// none
	LinkUtilization TapiCommonDirectiveValue `json:"link-utilization,omitempty"`

	// none
	ResourceSharing TapiCommonDirectiveValue `json:"resource-sharing,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiPathComputationPathObjectiveFunction) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiCommonLocalClass
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiCommonLocalClass = aO0

	// AO1
	var dataAO1 struct {
		BandwidthOptimization TapiCommonDirectiveValue `json:"bandwidth-optimization,omitempty"`

		ConcurrentPaths TapiCommonDirectiveValue `json:"concurrent-paths,omitempty"`

		CostOptimization TapiCommonDirectiveValue `json:"cost-optimization,omitempty"`

		LinkUtilization TapiCommonDirectiveValue `json:"link-utilization,omitempty"`

		ResourceSharing TapiCommonDirectiveValue `json:"resource-sharing,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.BandwidthOptimization = dataAO1.BandwidthOptimization

	m.ConcurrentPaths = dataAO1.ConcurrentPaths

	m.CostOptimization = dataAO1.CostOptimization

	m.LinkUtilization = dataAO1.LinkUtilization

	m.ResourceSharing = dataAO1.ResourceSharing

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiPathComputationPathObjectiveFunction) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.TapiCommonLocalClass)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		BandwidthOptimization TapiCommonDirectiveValue `json:"bandwidth-optimization,omitempty"`

		ConcurrentPaths TapiCommonDirectiveValue `json:"concurrent-paths,omitempty"`

		CostOptimization TapiCommonDirectiveValue `json:"cost-optimization,omitempty"`

		LinkUtilization TapiCommonDirectiveValue `json:"link-utilization,omitempty"`

		ResourceSharing TapiCommonDirectiveValue `json:"resource-sharing,omitempty"`
	}

	dataAO1.BandwidthOptimization = m.BandwidthOptimization

	dataAO1.ConcurrentPaths = m.ConcurrentPaths

	dataAO1.CostOptimization = m.CostOptimization

	dataAO1.LinkUtilization = m.LinkUtilization

	dataAO1.ResourceSharing = m.ResourceSharing

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi path computation path objective function
func (m *TapiPathComputationPathObjectiveFunction) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiCommonLocalClass
	if err := m.TapiCommonLocalClass.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBandwidthOptimization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConcurrentPaths(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCostOptimization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkUtilization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceSharing(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiPathComputationPathObjectiveFunction) validateBandwidthOptimization(formats strfmt.Registry) error {

	if swag.IsZero(m.BandwidthOptimization) { // not required
		return nil
	}

	if err := m.BandwidthOptimization.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("bandwidth-optimization")
		}
		return err
	}

	return nil
}

func (m *TapiPathComputationPathObjectiveFunction) validateConcurrentPaths(formats strfmt.Registry) error {

	if swag.IsZero(m.ConcurrentPaths) { // not required
		return nil
	}

	if err := m.ConcurrentPaths.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("concurrent-paths")
		}
		return err
	}

	return nil
}

func (m *TapiPathComputationPathObjectiveFunction) validateCostOptimization(formats strfmt.Registry) error {

	if swag.IsZero(m.CostOptimization) { // not required
		return nil
	}

	if err := m.CostOptimization.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cost-optimization")
		}
		return err
	}

	return nil
}

func (m *TapiPathComputationPathObjectiveFunction) validateLinkUtilization(formats strfmt.Registry) error {

	if swag.IsZero(m.LinkUtilization) { // not required
		return nil
	}

	if err := m.LinkUtilization.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("link-utilization")
		}
		return err
	}

	return nil
}

func (m *TapiPathComputationPathObjectiveFunction) validateResourceSharing(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceSharing) { // not required
		return nil
	}

	if err := m.ResourceSharing.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("resource-sharing")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiPathComputationPathObjectiveFunction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiPathComputationPathObjectiveFunction) UnmarshalBinary(b []byte) error {
	var res TapiPathComputationPathObjectiveFunction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
