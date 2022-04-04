// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EnrolmentInfo enrolment info
//
// swagger:model enrolment-info
type EnrolmentInfo struct {

	// features
	Features *EnrolmentInfoFeatures `json:"features,omitempty"`

	// target namespace
	TargetNamespace *string `json:"target-namespace,omitempty"`
}

// Validate validates this enrolment info
func (m *EnrolmentInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnrolmentInfo) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	if m.Features != nil {
		if err := m.Features.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnrolmentInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnrolmentInfo) UnmarshalBinary(b []byte) error {
	var res EnrolmentInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EnrolmentInfoFeatures enrolment info features
//
// swagger:model EnrolmentInfoFeatures
type EnrolmentInfoFeatures struct {

	// hardware
	Hardware *HardwareInfo `json:"hardware,omitempty"`

	// model name
	ModelName string `json:"modelName,omitempty"`
}

// Validate validates this enrolment info features
func (m *EnrolmentInfoFeatures) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHardware(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnrolmentInfoFeatures) validateHardware(formats strfmt.Registry) error {

	if swag.IsZero(m.Hardware) { // not required
		return nil
	}

	if m.Hardware != nil {
		if err := m.Hardware.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features" + "." + "hardware")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnrolmentInfoFeatures) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnrolmentInfoFeatures) UnmarshalBinary(b []byte) error {
	var res EnrolmentInfoFeatures
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
