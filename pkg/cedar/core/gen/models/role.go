// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Role role
//
// swagger:model Role
type Role struct {

	// Application where the role assignment exists
	// Required: true
	Application *string `json:"application"`

	// assignee desc
	AssigneeDesc string `json:"assigneeDesc,omitempty"`

	// assignee email
	AssigneeEmail string `json:"assigneeEmail,omitempty"`

	// assignee first name
	AssigneeFirstName string `json:"assigneeFirstName,omitempty"`

	// ID of the role assignee, if a person
	AssigneeID string `json:"assigneeId,omitempty"`

	// assignee last name
	AssigneeLastName string `json:"assigneeLastName,omitempty"`

	// ID of the role assignee, if an organization
	AssigneeOrgID string `json:"assigneeOrgId,omitempty"`

	// assignee org name
	AssigneeOrgName string `json:"assigneeOrgName,omitempty"`

	// assignee phone
	AssigneePhone string `json:"assigneePhone,omitempty"`

	// assignee type
	// Enum: [organization person]
	AssigneeType string `json:"assigneeType,omitempty"`

	// Username of the role assignee, if a person
	AssigneeUserName string `json:"assigneeUserName,omitempty"`

	// ID of the object the role is assiged to
	// Required: true
	ObjectID *string `json:"objectId"`

	// The type of object the role is assigned to
	ObjectType string `json:"objectType,omitempty"`

	// ID of the role assignment
	RoleID string `json:"roleId,omitempty"`

	// Description of the role type
	RoleTypeDesc string `json:"roleTypeDesc,omitempty"`

	// ID of the role type
	// Required: true
	RoleTypeID *string `json:"roleTypeId"`

	// Name of the role type
	RoleTypeName string `json:"roleTypeName,omitempty"`
}

// Validate validates this role
func (m *Role) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssigneeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Role) validateApplication(formats strfmt.Registry) error {

	if err := validate.Required("application", "body", m.Application); err != nil {
		return err
	}

	return nil
}

var roleTypeAssigneeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["organization","person"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		roleTypeAssigneeTypePropEnum = append(roleTypeAssigneeTypePropEnum, v)
	}
}

const (

	// RoleAssigneeTypeOrganization captures enum value "organization"
	RoleAssigneeTypeOrganization string = "organization"

	// RoleAssigneeTypePerson captures enum value "person"
	RoleAssigneeTypePerson string = "person"
)

// prop value enum
func (m *Role) validateAssigneeTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, roleTypeAssigneeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Role) validateAssigneeType(formats strfmt.Registry) error {
	if swag.IsZero(m.AssigneeType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAssigneeTypeEnum("assigneeType", "body", m.AssigneeType); err != nil {
		return err
	}

	return nil
}

func (m *Role) validateObjectID(formats strfmt.Registry) error {

	if err := validate.Required("objectId", "body", m.ObjectID); err != nil {
		return err
	}

	return nil
}

func (m *Role) validateRoleTypeID(formats strfmt.Registry) error {

	if err := validate.Required("roleTypeId", "body", m.RoleTypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this role based on context it is used
func (m *Role) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Role) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Role) UnmarshalBinary(b []byte) error {
	var res Role
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}