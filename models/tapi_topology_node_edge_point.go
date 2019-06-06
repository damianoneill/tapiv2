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

// TapiTopologyNodeEdgePoint tapi topology node edge point
// swagger:model tapi.topology.NodeEdgePoint
type TapiTopologyNodeEdgePoint struct {
	TapiCommonAdminStatePac

	TapiCommonCapacityPac

	TapiCommonGlobalClass

	TapiCommonTerminationPac

	// none
	AggregatedNodeEdgePoint []*TapiTopologyNodeEdgePointRef `json:"aggregated-node-edge-point"`

	// none
	AvailableCepLayerProtocol []*TapiTopologyNepLayerProtocolCapability `json:"available-cep-layer-protocol"`

	// none
	LayerProtocolName TapiCommonLayerProtocolName `json:"layer-protocol-name,omitempty"`

	// The orientation of defined flow at the LinkEnd.
	LinkPortDirection TapiCommonPortDirection `json:"link-port-direction,omitempty"`

	// Each LinkEnd of the Link has a role (e.g., symmetric, hub, spoke, leaf, root)  in the context of the Link with respect to the Link function.
	LinkPortRole TapiCommonPortRole `json:"link-port-role,omitempty"`

	// NodeEdgePoint mapped to more than ServiceInterfacePoint (slicing/virtualizing) or a ServiceInterfacePoint mapped to more than one NodeEdgePoint (load balancing/Resilience) should be considered experimental
	MappedServiceInterfacePoint []*TapiCommonServiceInterfacePointRef `json:"mapped-service-interface-point"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiTopologyNodeEdgePoint) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiCommonAdminStatePac
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiCommonAdminStatePac = aO0

	// AO1
	var aO1 TapiCommonCapacityPac
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TapiCommonCapacityPac = aO1

	// AO2
	var aO2 TapiCommonGlobalClass
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.TapiCommonGlobalClass = aO2

	// AO3
	var aO3 TapiCommonTerminationPac
	if err := swag.ReadJSON(raw, &aO3); err != nil {
		return err
	}
	m.TapiCommonTerminationPac = aO3

	// AO4
	var dataAO4 struct {
		AggregatedNodeEdgePoint []*TapiTopologyNodeEdgePointRef `json:"aggregated-node-edge-point"`

		AvailableCepLayerProtocol []*TapiTopologyNepLayerProtocolCapability `json:"available-cep-layer-protocol"`

		LayerProtocolName TapiCommonLayerProtocolName `json:"layer-protocol-name,omitempty"`

		LinkPortDirection TapiCommonPortDirection `json:"link-port-direction,omitempty"`

		LinkPortRole TapiCommonPortRole `json:"link-port-role,omitempty"`

		MappedServiceInterfacePoint []*TapiCommonServiceInterfacePointRef `json:"mapped-service-interface-point"`
	}
	if err := swag.ReadJSON(raw, &dataAO4); err != nil {
		return err
	}

	m.AggregatedNodeEdgePoint = dataAO4.AggregatedNodeEdgePoint

	m.AvailableCepLayerProtocol = dataAO4.AvailableCepLayerProtocol

	m.LayerProtocolName = dataAO4.LayerProtocolName

	m.LinkPortDirection = dataAO4.LinkPortDirection

	m.LinkPortRole = dataAO4.LinkPortRole

	m.MappedServiceInterfacePoint = dataAO4.MappedServiceInterfacePoint

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiTopologyNodeEdgePoint) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 5)

	aO0, err := swag.WriteJSON(m.TapiCommonAdminStatePac)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TapiCommonCapacityPac)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.TapiCommonGlobalClass)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	aO3, err := swag.WriteJSON(m.TapiCommonTerminationPac)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO3)

	var dataAO4 struct {
		AggregatedNodeEdgePoint []*TapiTopologyNodeEdgePointRef `json:"aggregated-node-edge-point"`

		AvailableCepLayerProtocol []*TapiTopologyNepLayerProtocolCapability `json:"available-cep-layer-protocol"`

		LayerProtocolName TapiCommonLayerProtocolName `json:"layer-protocol-name,omitempty"`

		LinkPortDirection TapiCommonPortDirection `json:"link-port-direction,omitempty"`

		LinkPortRole TapiCommonPortRole `json:"link-port-role,omitempty"`

		MappedServiceInterfacePoint []*TapiCommonServiceInterfacePointRef `json:"mapped-service-interface-point"`
	}

	dataAO4.AggregatedNodeEdgePoint = m.AggregatedNodeEdgePoint

	dataAO4.AvailableCepLayerProtocol = m.AvailableCepLayerProtocol

	dataAO4.LayerProtocolName = m.LayerProtocolName

	dataAO4.LinkPortDirection = m.LinkPortDirection

	dataAO4.LinkPortRole = m.LinkPortRole

	dataAO4.MappedServiceInterfacePoint = m.MappedServiceInterfacePoint

	jsonDataAO4, errAO4 := swag.WriteJSON(dataAO4)
	if errAO4 != nil {
		return nil, errAO4
	}
	_parts = append(_parts, jsonDataAO4)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi topology node edge point
func (m *TapiTopologyNodeEdgePoint) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiCommonAdminStatePac
	if err := m.TapiCommonAdminStatePac.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiCommonCapacityPac
	if err := m.TapiCommonCapacityPac.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiCommonGlobalClass
	if err := m.TapiCommonGlobalClass.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiCommonTerminationPac
	if err := m.TapiCommonTerminationPac.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAggregatedNodeEdgePoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvailableCepLayerProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLayerProtocolName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkPortDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkPortRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMappedServiceInterfacePoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiTopologyNodeEdgePoint) validateAggregatedNodeEdgePoint(formats strfmt.Registry) error {

	if swag.IsZero(m.AggregatedNodeEdgePoint) { // not required
		return nil
	}

	for i := 0; i < len(m.AggregatedNodeEdgePoint); i++ {
		if swag.IsZero(m.AggregatedNodeEdgePoint[i]) { // not required
			continue
		}

		if m.AggregatedNodeEdgePoint[i] != nil {
			if err := m.AggregatedNodeEdgePoint[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aggregated-node-edge-point" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiTopologyNodeEdgePoint) validateAvailableCepLayerProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.AvailableCepLayerProtocol) { // not required
		return nil
	}

	for i := 0; i < len(m.AvailableCepLayerProtocol); i++ {
		if swag.IsZero(m.AvailableCepLayerProtocol[i]) { // not required
			continue
		}

		if m.AvailableCepLayerProtocol[i] != nil {
			if err := m.AvailableCepLayerProtocol[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("available-cep-layer-protocol" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiTopologyNodeEdgePoint) validateLayerProtocolName(formats strfmt.Registry) error {

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

func (m *TapiTopologyNodeEdgePoint) validateLinkPortDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.LinkPortDirection) { // not required
		return nil
	}

	if err := m.LinkPortDirection.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("link-port-direction")
		}
		return err
	}

	return nil
}

func (m *TapiTopologyNodeEdgePoint) validateLinkPortRole(formats strfmt.Registry) error {

	if swag.IsZero(m.LinkPortRole) { // not required
		return nil
	}

	if err := m.LinkPortRole.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("link-port-role")
		}
		return err
	}

	return nil
}

func (m *TapiTopologyNodeEdgePoint) validateMappedServiceInterfacePoint(formats strfmt.Registry) error {

	if swag.IsZero(m.MappedServiceInterfacePoint) { // not required
		return nil
	}

	for i := 0; i < len(m.MappedServiceInterfacePoint); i++ {
		if swag.IsZero(m.MappedServiceInterfacePoint[i]) { // not required
			continue
		}

		if m.MappedServiceInterfacePoint[i] != nil {
			if err := m.MappedServiceInterfacePoint[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mapped-service-interface-point" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiTopologyNodeEdgePoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiTopologyNodeEdgePoint) UnmarshalBinary(b []byte) error {
	var res TapiTopologyNodeEdgePoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}