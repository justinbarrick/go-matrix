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

// InitialSyncOKBodyRoomsItemsInviteAllOf0ContentThirdPartyInvite Invite
// swagger:model initialSyncOKBodyRoomsItemsInviteAllOf0ContentThirdPartyInvite
type InitialSyncOKBodyRoomsItemsInviteAllOf0ContentThirdPartyInvite struct {

	// A name which can be displayed to represent the user instead of their third party identifier
	// Required: true
	DisplayName *string `json:"display_name"`

	// signed
	// Required: true
	Signed *InitialSyncOKBodyRoomsItemsInviteAllOf0ContentThirdPartyInviteSigned `json:"signed"`
}

// Validate validates this initial sync o k body rooms items invite all of0 content third party invite
func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0ContentThirdPartyInvite) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSigned(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0ContentThirdPartyInvite) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("display_name", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0ContentThirdPartyInvite) validateSigned(formats strfmt.Registry) error {

	if err := validate.Required("signed", "body", m.Signed); err != nil {
		return err
	}

	if m.Signed != nil {
		if err := m.Signed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signed")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0ContentThirdPartyInvite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0ContentThirdPartyInvite) UnmarshalBinary(b []byte) error {
	var res InitialSyncOKBodyRoomsItemsInviteAllOf0ContentThirdPartyInvite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
