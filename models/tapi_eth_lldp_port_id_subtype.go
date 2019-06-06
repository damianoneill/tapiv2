// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiEthLldpPortIDSubtype tapi eth lldp port Id subtype
// swagger:model tapi.eth.LldpPortIdSubtype
type TapiEthLldpPortIDSubtype struct {

	// Represents a port identifier based on the agent-local identifier of the circuit (defined in RFC 3046), detected by the agent and associated with a particular port.
	AgentCircuitID string `json:"agent-circuit-id,omitempty"`

	// String length '0..64'
	//                 Represents a port identifier based on the ifAlias MIB object, defined in IETF RFC 2863.
	InterfaceAlias string `json:"interface-alias,omitempty"`

	// String length '0..64'
	//                 Represents a port identifier based on the ifName MIB object, defined in IETF RFC 2863.
	InterfaceName string `json:"interface-name,omitempty"`

	// Represents a port identifier based on a value locally assigned.
	Local string `json:"local,omitempty"`

	// Represents a port identifier based on a unicast source address (encoded in network byte order and IEEE 802.3 canonical bit order), which has been detected by the agent and associated with a particular port (IEEE Std 802-2001).
	MacAddress string `json:"mac-address,omitempty"`

	// Represents a port identifier based on a network address, detected by the agent and associated with a particular port.
	//                 Octet string that identifies a particular network address family and an associated network address that are encoded in network octet order.
	//                 An IP address, for example, would be encoded with the first octet containing the IANA Address Family Numbers enumeration value for the specific address type and octets 2 through n containing the address value.
	//
	NetworkAddress string `json:"network-address,omitempty"`

	// String length '0..32'
	//                 Represents a port identifier based on the value of entPhysicalAlias (defined in IETF RFC 2737) for a port component (i.e., entPhysicalClass value of port(10)), within the containing chassis.
	PortComponent string `json:"port-component,omitempty"`
}

// Validate validates this tapi eth lldp port Id subtype
func (m *TapiEthLldpPortIDSubtype) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthLldpPortIDSubtype) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthLldpPortIDSubtype) UnmarshalBinary(b []byte) error {
	var res TapiEthLldpPortIDSubtype
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
