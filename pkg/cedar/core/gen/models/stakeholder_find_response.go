// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StakeholderFindResponse stakeholder find response
//
// swagger:model StakeholderFindResponse
type StakeholderFindResponse struct {

	// stakeholders
	Stakeholders []*Stakeholder `json:"Stakeholders"`

	// count
	// Required: true
	Count *int32 `json:"count"`
}

// Validate validates this stakeholder find response
func (m *StakeholderFindResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStakeholders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StakeholderFindResponse) validateStakeholders(formats strfmt.Registry) error {
	if swag.IsZero(m.Stakeholders) { // not required
		return nil
	}

	for i := 0; i < len(m.Stakeholders); i++ {
		if swag.IsZero(m.Stakeholders[i]) { // not required
			continue
		}

		if m.Stakeholders[i] != nil {
			if err := m.Stakeholders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Stakeholders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Stakeholders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StakeholderFindResponse) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this stakeholder find response based on the context it is used
func (m *StakeholderFindResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStakeholders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StakeholderFindResponse) contextValidateStakeholders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Stakeholders); i++ {

		if m.Stakeholders[i] != nil {
			if err := m.Stakeholders[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Stakeholders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Stakeholders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StakeholderFindResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StakeholderFindResponse) UnmarshalBinary(b []byte) error {
	var res StakeholderFindResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}