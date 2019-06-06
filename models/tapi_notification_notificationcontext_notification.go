// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiNotificationNotificationcontextNotification tapi notification notificationcontext notification
// swagger:model tapi.notification.notificationcontext.Notification
type TapiNotificationNotificationcontextNotification struct {
	TapiNotificationNotification

	TapiOamNotificationAugmentation1

	TapiOamNotificationAugmentation2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiNotificationNotificationcontextNotification) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiNotificationNotification
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiNotificationNotification = aO0

	// AO1
	var aO1 TapiOamNotificationAugmentation1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TapiOamNotificationAugmentation1 = aO1

	// AO2
	var aO2 TapiOamNotificationAugmentation2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.TapiOamNotificationAugmentation2 = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiNotificationNotificationcontextNotification) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.TapiNotificationNotification)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TapiOamNotificationAugmentation1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.TapiOamNotificationAugmentation2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi notification notificationcontext notification
func (m *TapiNotificationNotificationcontextNotification) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiNotificationNotification
	if err := m.TapiNotificationNotification.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiOamNotificationAugmentation1
	if err := m.TapiOamNotificationAugmentation1.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiOamNotificationAugmentation2
	if err := m.TapiOamNotificationAugmentation2.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TapiNotificationNotificationcontextNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiNotificationNotificationcontextNotification) UnmarshalBinary(b []byte) error {
	var res TapiNotificationNotificationcontextNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}