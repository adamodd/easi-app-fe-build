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

// BudgetUpdateRequest budget update request
//
// swagger:model BudgetUpdateRequest
type BudgetUpdateRequest struct {

	// budgets
	// Required: true
	Budgets []*Budget `json:"Budgets"`
}

// Validate validates this budget update request
func (m *BudgetUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBudgets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BudgetUpdateRequest) validateBudgets(formats strfmt.Registry) error {

	if err := validate.Required("Budgets", "body", m.Budgets); err != nil {
		return err
	}

	for i := 0; i < len(m.Budgets); i++ {
		if swag.IsZero(m.Budgets[i]) { // not required
			continue
		}

		if m.Budgets[i] != nil {
			if err := m.Budgets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Budgets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Budgets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this budget update request based on the context it is used
func (m *BudgetUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBudgets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BudgetUpdateRequest) contextValidateBudgets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Budgets); i++ {

		if m.Budgets[i] != nil {
			if err := m.Budgets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Budgets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Budgets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BudgetUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BudgetUpdateRequest) UnmarshalBinary(b []byte) error {
	var res BudgetUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
