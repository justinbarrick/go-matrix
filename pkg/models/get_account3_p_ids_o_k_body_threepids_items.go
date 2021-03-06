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

// GetAccount3PIdsOKBodyThreepidsItems Third party identifier
// swagger:model getAccount3PIdsOKBodyThreepidsItems
type GetAccount3PIdsOKBodyThreepidsItems struct {

	// The timestamp, in milliseconds, when the homeserver associated the third party identifier with the user.
	// Required: true
	AddedAt *int64 `json:"added_at"`

	// The third party identifier address.
	// Required: true
	Address *string `json:"address"`

	// The medium of the third party identifier.
	// Required: true
	// Enum: [email msisdn]
	Medium *string `json:"medium"`

	// The timestamp, in milliseconds, when the identifier was
	// validated by the identity server.
	// Required: true
	ValidatedAt *int64 `json:"validated_at"`
}

// Validate validates this get account3 p ids o k body threepids items
func (m *GetAccount3PIdsOKBodyThreepidsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMedium(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetAccount3PIdsOKBodyThreepidsItems) validateAddedAt(formats strfmt.Registry) error {

	if err := validate.Required("added_at", "body", m.AddedAt); err != nil {
		return err
	}

	return nil
}

func (m *GetAccount3PIdsOKBodyThreepidsItems) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

var getAccount3PIdsOKBodyThreepidsItemsTypeMediumPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["email","msisdn"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getAccount3PIdsOKBodyThreepidsItemsTypeMediumPropEnum = append(getAccount3PIdsOKBodyThreepidsItemsTypeMediumPropEnum, v)
	}
}

const (

	// GetAccount3PIdsOKBodyThreepidsItemsMediumEmail captures enum value "email"
	GetAccount3PIdsOKBodyThreepidsItemsMediumEmail string = "email"

	// GetAccount3PIdsOKBodyThreepidsItemsMediumMsisdn captures enum value "msisdn"
	GetAccount3PIdsOKBodyThreepidsItemsMediumMsisdn string = "msisdn"
)

// prop value enum
func (m *GetAccount3PIdsOKBodyThreepidsItems) validateMediumEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getAccount3PIdsOKBodyThreepidsItemsTypeMediumPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetAccount3PIdsOKBodyThreepidsItems) validateMedium(formats strfmt.Registry) error {

	if err := validate.Required("medium", "body", m.Medium); err != nil {
		return err
	}

	// value enum
	if err := m.validateMediumEnum("medium", "body", *m.Medium); err != nil {
		return err
	}

	return nil
}

func (m *GetAccount3PIdsOKBodyThreepidsItems) validateValidatedAt(formats strfmt.Registry) error {

	if err := validate.Required("validated_at", "body", m.ValidatedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetAccount3PIdsOKBodyThreepidsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAccount3PIdsOKBodyThreepidsItems) UnmarshalBinary(b []byte) error {
	var res GetAccount3PIdsOKBodyThreepidsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
