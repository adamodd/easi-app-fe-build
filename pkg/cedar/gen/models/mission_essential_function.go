// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MissionEssentialFunction mission essential function
// swagger:model mission_essential_function
type MissionEssentialFunction struct {

	// mission essential function level
	// Required: true
	MissionEssentialFunctionLevel *string `json:"mission_essential_function_level"`

	// mission essential function name
	// Required: true
	MissionEssentialFunctionName *string `json:"mission_essential_function_name"`
}

// Validate validates this mission essential function
func (m *MissionEssentialFunction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMissionEssentialFunctionLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMissionEssentialFunctionName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MissionEssentialFunction) validateMissionEssentialFunctionLevel(formats strfmt.Registry) error {

	if err := validate.Required("mission_essential_function_level", "body", m.MissionEssentialFunctionLevel); err != nil {
		return err
	}

	return nil
}

func (m *MissionEssentialFunction) validateMissionEssentialFunctionName(formats strfmt.Registry) error {

	if err := validate.Required("mission_essential_function_name", "body", m.MissionEssentialFunctionName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MissionEssentialFunction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MissionEssentialFunction) UnmarshalBinary(b []byte) error {
	var res MissionEssentialFunction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
