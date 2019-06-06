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

// TapiConnectivityUpdateconnectivityserviceOutputService tapi connectivity updateconnectivityservice output service
// swagger:model tapi.connectivity.updateconnectivityservice.output.Service
type TapiConnectivityUpdateconnectivityserviceOutputService struct {
	TapiCommonAdminStatePac

	TapiCommonGlobalClass

	TapiEthServiceAugmentation1

	// none
	Connection []*TapiConnectivityConnectionRef `json:"connection"`

	// none
	ConnectivityConstraint *TapiConnectivityConnectivityConstraint `json:"connectivity-constraint,omitempty"`

	// none
	Direction TapiCommonForwardingDirection `json:"direction,omitempty"`

	// none
	EndPoint []*TapiConnectivityCreateconnectivityserviceInputEndPoint `json:"end-point"`

	// none
	LayerProtocolName TapiCommonLayerProtocolName `json:"layer-protocol-name,omitempty"`

	// none
	ResilienceConstraint *TapiConnectivityResilienceConstraint `json:"resilience-constraint,omitempty"`

	// none
	RoutingConstraint *TapiPathComputationRoutingConstraint `json:"routing-constraint,omitempty"`

	// none
	TopologyConstraint []*TapiPathComputationTopologyConstraint `json:"topology-constraint"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiConnectivityUpdateconnectivityserviceOutputService) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiCommonAdminStatePac
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiCommonAdminStatePac = aO0

	// AO1
	var aO1 TapiCommonGlobalClass
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TapiCommonGlobalClass = aO1

	// AO2
	var aO2 TapiEthServiceAugmentation1
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.TapiEthServiceAugmentation1 = aO2

	// AO3
	var dataAO3 struct {
		Connection []*TapiConnectivityConnectionRef `json:"connection"`

		ConnectivityConstraint *TapiConnectivityConnectivityConstraint `json:"connectivity-constraint,omitempty"`

		Direction TapiCommonForwardingDirection `json:"direction,omitempty"`

		EndPoint []*TapiConnectivityCreateconnectivityserviceInputEndPoint `json:"end-point"`

		LayerProtocolName TapiCommonLayerProtocolName `json:"layer-protocol-name,omitempty"`

		ResilienceConstraint *TapiConnectivityResilienceConstraint `json:"resilience-constraint,omitempty"`

		RoutingConstraint *TapiPathComputationRoutingConstraint `json:"routing-constraint,omitempty"`

		TopologyConstraint []*TapiPathComputationTopologyConstraint `json:"topology-constraint"`
	}
	if err := swag.ReadJSON(raw, &dataAO3); err != nil {
		return err
	}

	m.Connection = dataAO3.Connection

	m.ConnectivityConstraint = dataAO3.ConnectivityConstraint

	m.Direction = dataAO3.Direction

	m.EndPoint = dataAO3.EndPoint

	m.LayerProtocolName = dataAO3.LayerProtocolName

	m.ResilienceConstraint = dataAO3.ResilienceConstraint

	m.RoutingConstraint = dataAO3.RoutingConstraint

	m.TopologyConstraint = dataAO3.TopologyConstraint

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiConnectivityUpdateconnectivityserviceOutputService) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 4)

	aO0, err := swag.WriteJSON(m.TapiCommonAdminStatePac)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TapiCommonGlobalClass)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.TapiEthServiceAugmentation1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	var dataAO3 struct {
		Connection []*TapiConnectivityConnectionRef `json:"connection"`

		ConnectivityConstraint *TapiConnectivityConnectivityConstraint `json:"connectivity-constraint,omitempty"`

		Direction TapiCommonForwardingDirection `json:"direction,omitempty"`

		EndPoint []*TapiConnectivityCreateconnectivityserviceInputEndPoint `json:"end-point"`

		LayerProtocolName TapiCommonLayerProtocolName `json:"layer-protocol-name,omitempty"`

		ResilienceConstraint *TapiConnectivityResilienceConstraint `json:"resilience-constraint,omitempty"`

		RoutingConstraint *TapiPathComputationRoutingConstraint `json:"routing-constraint,omitempty"`

		TopologyConstraint []*TapiPathComputationTopologyConstraint `json:"topology-constraint"`
	}

	dataAO3.Connection = m.Connection

	dataAO3.ConnectivityConstraint = m.ConnectivityConstraint

	dataAO3.Direction = m.Direction

	dataAO3.EndPoint = m.EndPoint

	dataAO3.LayerProtocolName = m.LayerProtocolName

	dataAO3.ResilienceConstraint = m.ResilienceConstraint

	dataAO3.RoutingConstraint = m.RoutingConstraint

	dataAO3.TopologyConstraint = m.TopologyConstraint

	jsonDataAO3, errAO3 := swag.WriteJSON(dataAO3)
	if errAO3 != nil {
		return nil, errAO3
	}
	_parts = append(_parts, jsonDataAO3)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi connectivity updateconnectivityservice output service
func (m *TapiConnectivityUpdateconnectivityserviceOutputService) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiCommonAdminStatePac
	if err := m.TapiCommonAdminStatePac.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiCommonGlobalClass
	if err := m.TapiCommonGlobalClass.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiEthServiceAugmentation1
	if err := m.TapiEthServiceAugmentation1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectivityConstraint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndPoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLayerProtocolName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResilienceConstraint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutingConstraint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopologyConstraint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiConnectivityUpdateconnectivityserviceOutputService) validateConnection(formats strfmt.Registry) error {

	if swag.IsZero(m.Connection) { // not required
		return nil
	}

	for i := 0; i < len(m.Connection); i++ {
		if swag.IsZero(m.Connection[i]) { // not required
			continue
		}

		if m.Connection[i] != nil {
			if err := m.Connection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiConnectivityUpdateconnectivityserviceOutputService) validateConnectivityConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectivityConstraint) { // not required
		return nil
	}

	if m.ConnectivityConstraint != nil {
		if err := m.ConnectivityConstraint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectivity-constraint")
			}
			return err
		}
	}

	return nil
}

func (m *TapiConnectivityUpdateconnectivityserviceOutputService) validateDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	if err := m.Direction.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("direction")
		}
		return err
	}

	return nil
}

func (m *TapiConnectivityUpdateconnectivityserviceOutputService) validateEndPoint(formats strfmt.Registry) error {

	if swag.IsZero(m.EndPoint) { // not required
		return nil
	}

	for i := 0; i < len(m.EndPoint); i++ {
		if swag.IsZero(m.EndPoint[i]) { // not required
			continue
		}

		if m.EndPoint[i] != nil {
			if err := m.EndPoint[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("end-point" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiConnectivityUpdateconnectivityserviceOutputService) validateLayerProtocolName(formats strfmt.Registry) error {

	if swag.IsZero(m.LayerProtocolName) { // not required
		return nil
	}

	if err := m.LayerProtocolName.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("layer-protocol-name")
		}
		return err
	}

	return nil
}

func (m *TapiConnectivityUpdateconnectivityserviceOutputService) validateResilienceConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.ResilienceConstraint) { // not required
		return nil
	}

	if m.ResilienceConstraint != nil {
		if err := m.ResilienceConstraint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resilience-constraint")
			}
			return err
		}
	}

	return nil
}

func (m *TapiConnectivityUpdateconnectivityserviceOutputService) validateRoutingConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.RoutingConstraint) { // not required
		return nil
	}

	if m.RoutingConstraint != nil {
		if err := m.RoutingConstraint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("routing-constraint")
			}
			return err
		}
	}

	return nil
}

func (m *TapiConnectivityUpdateconnectivityserviceOutputService) validateTopologyConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.TopologyConstraint) { // not required
		return nil
	}

	for i := 0; i < len(m.TopologyConstraint); i++ {
		if swag.IsZero(m.TopologyConstraint[i]) { // not required
			continue
		}

		if m.TopologyConstraint[i] != nil {
			if err := m.TopologyConstraint[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("topology-constraint" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivityUpdateconnectivityserviceOutputService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivityUpdateconnectivityserviceOutputService) UnmarshalBinary(b []byte) error {
	var res TapiConnectivityUpdateconnectivityserviceOutputService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
