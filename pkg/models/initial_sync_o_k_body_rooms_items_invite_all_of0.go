// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InitialSyncOKBodyRoomsItemsInviteAllOf0 The current membership state of a user in the room.
//
// Adjusts the membership state for a user in a room. It is preferable to use the membership APIs (``/rooms/<room id>/invite`` etc) when performing membership actions rather than adjusting the state directly as there are a restricted set of valid transformations. For example, user A cannot force user B to join a room, and trying to force this state change directly will fail.
//
// The following membership states are specified:
//
// - ``invite`` - The user has been invited to join a room, but has not yet joined it. They may not participate in the room until they join.
//
// - ``join`` - The user has joined the room (possibly after accepting an invite), and may participate in it.
//
// - ``leave`` - The user was once joined to the room, but has since left (possibly by choice, or possibly by being kicked).
//
// - ``ban`` - The user has been banned from the room, and is no longer allowed to join it until they are un-banned from the room (by having their membership state set to a value other than ``ban``).
//
// - ``knock`` - This is a reserved word, which currently has no meaning.
//
// The ``third_party_invite`` property will be set if this invite is an ``invite`` event and is the successor of an ``m.room.third_party_invite`` event, and absent otherwise.
//
// This event may also include an ``invite_room_state`` key inside the event's ``unsigned`` data.
// If present, this contains an array of ``StrippedState`` Events. These events provide information
// on a subset of state events such as the room name.
// swagger:model initialSyncOKBodyRoomsItemsInviteAllOf0
type InitialSyncOKBodyRoomsItemsInviteAllOf0 struct {
	InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0

	// content
	Content *InitialSyncOKBodyRoomsItemsInviteAllOf0Content `json:"content,omitempty"`

	// The ``user_id`` this membership event relates to. In all cases except for when ``membership`` is
	// ``join``, the user ID sending the event does not need to match the user ID in the ``state_key``,
	// unlike other events. Regular authorisation rules still apply.
	StateKey string `json:"state_key,omitempty"`

	// type
	// Enum: [m.room.member]
	Type string `json:"type,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m InitialSyncOKBodyRoomsItemsInviteAllOf0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this initial sync o k body rooms items invite all of0
func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0
	if err := m.InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0) validateContent(formats strfmt.Registry) error {

	if swag.IsZero(m.Content) { // not required
		return nil
	}

	if m.Content != nil {
		if err := m.Content.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

var initialSyncOKBodyRoomsItemsInviteAllOf0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["m.room.member"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		initialSyncOKBodyRoomsItemsInviteAllOf0TypeTypePropEnum = append(initialSyncOKBodyRoomsItemsInviteAllOf0TypeTypePropEnum, v)
	}
}

const (

	// InitialSyncOKBodyRoomsItemsInviteAllOf0TypeMRoomMember captures enum value "m.room.member"
	InitialSyncOKBodyRoomsItemsInviteAllOf0TypeMRoomMember string = "m.room.member"
)

// prop value enum
func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, initialSyncOKBodyRoomsItemsInviteAllOf0TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0) UnmarshalBinary(b []byte) error {
	var res InitialSyncOKBodyRoomsItemsInviteAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
