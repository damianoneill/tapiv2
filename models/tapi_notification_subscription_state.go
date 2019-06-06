// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// TapiNotificationSubscriptionState tapi notification subscription state
// swagger:model tapi.notification.SubscriptionState
type TapiNotificationSubscriptionState string

const (

	// TapiNotificationSubscriptionStateSUSPENDED captures enum value "SUSPENDED"
	TapiNotificationSubscriptionStateSUSPENDED TapiNotificationSubscriptionState = "SUSPENDED"

	// TapiNotificationSubscriptionStateACTIVE captures enum value "ACTIVE"
	TapiNotificationSubscriptionStateACTIVE TapiNotificationSubscriptionState = "ACTIVE"
)

// for schema
var tapiNotificationSubscriptionStateEnum []interface{}

func init() {
	var res []TapiNotificationSubscriptionState
	if err := json.Unmarshal([]byte(`["SUSPENDED","ACTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiNotificationSubscriptionStateEnum = append(tapiNotificationSubscriptionStateEnum, v)
	}
}

func (m TapiNotificationSubscriptionState) validateTapiNotificationSubscriptionStateEnum(path, location string, value TapiNotificationSubscriptionState) error {
	if err := validate.Enum(path, location, value, tapiNotificationSubscriptionStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi notification subscription state
func (m TapiNotificationSubscriptionState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiNotificationSubscriptionStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
