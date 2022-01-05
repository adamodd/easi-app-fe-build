// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ComponentFindResponse component find response
//
// swagger:model ComponentFindResponse
type ComponentFindResponse struct {

	// catalog
	Catalog string `json:"catalog,omitempty"`

	// cms end of support date
	CmsEndOfSupportDate string `json:"cms_end_of_support_date,omitempty"`

	// cms technopedia build version
	CmsTechnopediaBuildVersion string `json:"cms_technopedia_build_version,omitempty"`

	// cms technopedia component
	CmsTechnopediaComponent string `json:"cms_technopedia_component,omitempty"`

	// cms technopedia edition
	CmsTechnopediaEdition string `json:"cms_technopedia_edition,omitempty"`

	// cms technopedia licensable
	CmsTechnopediaLicensable string `json:"cms_technopedia_licensable,omitempty"`

	// cms technopedia release
	CmsTechnopediaRelease string `json:"cms_technopedia_release,omitempty"`

	// cms technopedia release id
	CmsTechnopediaReleaseID string `json:"cms_technopedia_release_id,omitempty"`

	// cms technopedia servicepack
	CmsTechnopediaServicepack string `json:"cms_technopedia_servicepack,omitempty"`

	// cms technopedia version
	CmsTechnopediaVersion string `json:"cms_technopedia_version,omitempty"`

	// cms technopedia versiongroup
	CmsTechnopediaVersiongroup string `json:"cms_technopedia_versiongroup,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// ict object
	IctObject string `json:"ictObject,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// platform
	Platform string `json:"platform,omitempty"`

	// responsible user
	ResponsibleUser string `json:"responsibleUser,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// vendor
	Vendor string `json:"vendor,omitempty"`

	// vendor product
	VendorProduct string `json:"vendorProduct,omitempty"`
}

// Validate validates this component find response
func (m *ComponentFindResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this component find response based on context it is used
func (m *ComponentFindResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComponentFindResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComponentFindResponse) UnmarshalBinary(b []byte) error {
	var res ComponentFindResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}