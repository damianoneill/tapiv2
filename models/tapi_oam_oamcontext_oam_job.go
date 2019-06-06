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

// TapiOamOamcontextOamJob tapi oam oamcontext oam job
// swagger:model tapi.oam.oamcontext.OamJob
type TapiOamOamcontextOamJob struct {
	TapiCommonAdminStatePac

	TapiCommonGlobalClass

	TapiEthOamJobAugmentation1

	TapiEthOamJobAugmentation2

	TapiEthOamJobAugmentation3

	TapiEthOamJobAugmentation4

	TapiEthOamJobAugmentation5

	TapiEthOamJobAugmentation6

	TapiEthOamJobAugmentation7

	// none
	CreationTime string `json:"creation-time,omitempty"`

	// Granularity period of the CurrentData identifies the specific CurrentData instance in the scope of this OamJob.
	//                 For example, typically at least
	//                 one 15min and
	//                 one 24hr;
	//                 optionally one additional configurable (< 15min)
	CurrentData []*TapiOamOamjobCurrentData `json:"current-data"`

	// none
	OamJobType string `json:"oam-job-type,omitempty"`

	// none
	OamProfile *TapiOamOamProfileRef `json:"oam-profile,omitempty"`

	// none
	OamServicePoint []*TapiOamOamServicePointRef `json:"oam-service-point"`

	// none
	Schedule *TapiCommonTimeRange `json:"schedule,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiOamOamcontextOamJob) UnmarshalJSON(raw []byte) error {
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
	var aO2 TapiEthOamJobAugmentation1
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.TapiEthOamJobAugmentation1 = aO2

	// AO3
	var aO3 TapiEthOamJobAugmentation2
	if err := swag.ReadJSON(raw, &aO3); err != nil {
		return err
	}
	m.TapiEthOamJobAugmentation2 = aO3

	// AO4
	var aO4 TapiEthOamJobAugmentation3
	if err := swag.ReadJSON(raw, &aO4); err != nil {
		return err
	}
	m.TapiEthOamJobAugmentation3 = aO4

	// AO5
	var aO5 TapiEthOamJobAugmentation4
	if err := swag.ReadJSON(raw, &aO5); err != nil {
		return err
	}
	m.TapiEthOamJobAugmentation4 = aO5

	// AO6
	var aO6 TapiEthOamJobAugmentation5
	if err := swag.ReadJSON(raw, &aO6); err != nil {
		return err
	}
	m.TapiEthOamJobAugmentation5 = aO6

	// AO7
	var aO7 TapiEthOamJobAugmentation6
	if err := swag.ReadJSON(raw, &aO7); err != nil {
		return err
	}
	m.TapiEthOamJobAugmentation6 = aO7

	// AO8
	var aO8 TapiEthOamJobAugmentation7
	if err := swag.ReadJSON(raw, &aO8); err != nil {
		return err
	}
	m.TapiEthOamJobAugmentation7 = aO8

	// AO9
	var dataAO9 struct {
		CreationTime string `json:"creation-time,omitempty"`

		CurrentData []*TapiOamOamjobCurrentData `json:"current-data"`

		OamJobType string `json:"oam-job-type,omitempty"`

		OamProfile *TapiOamOamProfileRef `json:"oam-profile,omitempty"`

		OamServicePoint []*TapiOamOamServicePointRef `json:"oam-service-point"`

		Schedule *TapiCommonTimeRange `json:"schedule,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO9); err != nil {
		return err
	}

	m.CreationTime = dataAO9.CreationTime

	m.CurrentData = dataAO9.CurrentData

	m.OamJobType = dataAO9.OamJobType

	m.OamProfile = dataAO9.OamProfile

	m.OamServicePoint = dataAO9.OamServicePoint

	m.Schedule = dataAO9.Schedule

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiOamOamcontextOamJob) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 10)

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

	aO2, err := swag.WriteJSON(m.TapiEthOamJobAugmentation1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	aO3, err := swag.WriteJSON(m.TapiEthOamJobAugmentation2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO3)

	aO4, err := swag.WriteJSON(m.TapiEthOamJobAugmentation3)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO4)

	aO5, err := swag.WriteJSON(m.TapiEthOamJobAugmentation4)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO5)

	aO6, err := swag.WriteJSON(m.TapiEthOamJobAugmentation5)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO6)

	aO7, err := swag.WriteJSON(m.TapiEthOamJobAugmentation6)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO7)

	aO8, err := swag.WriteJSON(m.TapiEthOamJobAugmentation7)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO8)

	var dataAO9 struct {
		CreationTime string `json:"creation-time,omitempty"`

		CurrentData []*TapiOamOamjobCurrentData `json:"current-data"`

		OamJobType string `json:"oam-job-type,omitempty"`

		OamProfile *TapiOamOamProfileRef `json:"oam-profile,omitempty"`

		OamServicePoint []*TapiOamOamServicePointRef `json:"oam-service-point"`

		Schedule *TapiCommonTimeRange `json:"schedule,omitempty"`
	}

	dataAO9.CreationTime = m.CreationTime

	dataAO9.CurrentData = m.CurrentData

	dataAO9.OamJobType = m.OamJobType

	dataAO9.OamProfile = m.OamProfile

	dataAO9.OamServicePoint = m.OamServicePoint

	dataAO9.Schedule = m.Schedule

	jsonDataAO9, errAO9 := swag.WriteJSON(dataAO9)
	if errAO9 != nil {
		return nil, errAO9
	}
	_parts = append(_parts, jsonDataAO9)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi oam oamcontext oam job
func (m *TapiOamOamcontextOamJob) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiCommonAdminStatePac
	if err := m.TapiCommonAdminStatePac.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiCommonGlobalClass
	if err := m.TapiCommonGlobalClass.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiEthOamJobAugmentation1
	if err := m.TapiEthOamJobAugmentation1.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiEthOamJobAugmentation2
	if err := m.TapiEthOamJobAugmentation2.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiEthOamJobAugmentation3
	if err := m.TapiEthOamJobAugmentation3.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiEthOamJobAugmentation4
	if err := m.TapiEthOamJobAugmentation4.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiEthOamJobAugmentation5
	if err := m.TapiEthOamJobAugmentation5.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiEthOamJobAugmentation6
	if err := m.TapiEthOamJobAugmentation6.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiEthOamJobAugmentation7
	if err := m.TapiEthOamJobAugmentation7.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOamProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOamServicePoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamOamcontextOamJob) validateCurrentData(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentData) { // not required
		return nil
	}

	for i := 0; i < len(m.CurrentData); i++ {
		if swag.IsZero(m.CurrentData[i]) { // not required
			continue
		}

		if m.CurrentData[i] != nil {
			if err := m.CurrentData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("current-data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiOamOamcontextOamJob) validateOamProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.OamProfile) { // not required
		return nil
	}

	if m.OamProfile != nil {
		if err := m.OamProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oam-profile")
			}
			return err
		}
	}

	return nil
}

func (m *TapiOamOamcontextOamJob) validateOamServicePoint(formats strfmt.Registry) error {

	if swag.IsZero(m.OamServicePoint) { // not required
		return nil
	}

	for i := 0; i < len(m.OamServicePoint); i++ {
		if swag.IsZero(m.OamServicePoint[i]) { // not required
			continue
		}

		if m.OamServicePoint[i] != nil {
			if err := m.OamServicePoint[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("oam-service-point" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiOamOamcontextOamJob) validateSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamOamcontextOamJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamOamcontextOamJob) UnmarshalBinary(b []byte) error {
	var res TapiOamOamcontextOamJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
