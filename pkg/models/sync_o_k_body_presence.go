// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SyncOKBodyPresence Presence
//
// The updates to the presence status of other users.
// swagger:model syncOKBodyPresence
type SyncOKBodyPresence struct {
	SyncOKBodyPresenceAllOf0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SyncOKBodyPresence) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SyncOKBodyPresenceAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SyncOKBodyPresenceAllOf0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SyncOKBodyPresence) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.SyncOKBodyPresenceAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this sync o k body presence
func (m *SyncOKBodyPresence) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SyncOKBodyPresenceAllOf0
	if err := m.SyncOKBodyPresenceAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SyncOKBodyPresence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyncOKBodyPresence) UnmarshalBinary(b []byte) error {
	var res SyncOKBodyPresence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
