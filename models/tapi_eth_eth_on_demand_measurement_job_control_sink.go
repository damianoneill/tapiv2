// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthEthOnDemandMeasurementJobControlSink tapi eth eth on demand measurement job control sink
// swagger:model tapi.eth.EthOnDemandMeasurementJobControlSink
type TapiEthEthOnDemandMeasurementJobControlSink struct {
	TapiEthEthMeasurementJobControlCommon

	// none
	SinkMepID int32 `json:"sink-mep-id,omitempty"`

	// This attribute contains the MAC address of the peer MEP. See G.8013 for details.
	SourceAddress string `json:"source-address,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiEthEthOnDemandMeasurementJobControlSink) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiEthEthMeasurementJobControlCommon
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiEthEthMeasurementJobControlCommon = aO0

	// AO1
	var dataAO1 struct {
		SinkMepID int32 `json:"sink-mep-id,omitempty"`

		SourceAddress string `json:"source-address,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.SinkMepID = dataAO1.SinkMepID

	m.SourceAddress = dataAO1.SourceAddress

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiEthEthOnDemandMeasurementJobControlSink) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.TapiEthEthMeasurementJobControlCommon)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		SinkMepID int32 `json:"sink-mep-id,omitempty"`

		SourceAddress string `json:"source-address,omitempty"`
	}

	dataAO1.SinkMepID = m.SinkMepID

	dataAO1.SourceAddress = m.SourceAddress

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi eth eth on demand measurement job control sink
func (m *TapiEthEthOnDemandMeasurementJobControlSink) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiEthEthMeasurementJobControlCommon
	if err := m.TapiEthEthMeasurementJobControlCommon.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthEthOnDemandMeasurementJobControlSink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthEthOnDemandMeasurementJobControlSink) UnmarshalBinary(b []byte) error {
	var res TapiEthEthOnDemandMeasurementJobControlSink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}