// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiOduOduPoolPac tapi odu odu pool pac
// swagger:model tapi.odu.OduPoolPac
type TapiOduOduPoolPac struct {

	// none
	ClientCapacity int32 `json:"client-capacity,omitempty"`

	// none
	MaxClientInstances int32 `json:"max-client-instances,omitempty"`

	// none
	MaxClientSize int32 `json:"max-client-size,omitempty"`
}

// Validate validates this tapi odu odu pool pac
func (m *TapiOduOduPoolPac) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiOduOduPoolPac) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOduOduPoolPac) UnmarshalBinary(b []byte) error {
	var res TapiOduOduPoolPac
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}