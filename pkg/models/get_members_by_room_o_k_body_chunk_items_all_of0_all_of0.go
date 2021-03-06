// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0 State Event
//
// In addition to the Room Event fields, State Events have the following additional fields.
// swagger:model getMembersByRoomOKBodyChunkItemsAllOf0AllOf0
type GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0 struct {
	GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0AllOf0

	GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0AllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0AllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0AllOf0 = aO0

	// AO1
	var aO1 GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0AllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0AllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0AllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0AllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get members by room o k body chunk items all of0 all of0
func (m *GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0AllOf0
	if err := m.GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0AllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0AllOf1
	if err := m.GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0AllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0) UnmarshalBinary(b []byte) error {
	var res GetMembersByRoomOKBodyChunkItemsAllOf0AllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
