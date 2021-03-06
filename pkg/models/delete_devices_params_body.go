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

// DeleteDevicesParamsBody delete devices params body
// swagger:model deleteDevicesParamsBody
type DeleteDevicesParamsBody struct {

	// auth
	Auth *DeleteDevicesParamsBodyAuth `json:"auth,omitempty"`

	// The list of device IDs to delete.
	// Required: true
	Devices []string `json:"devices"`
}

// Validate validates this delete devices params body
func (m *DeleteDevicesParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteDevicesParamsBody) validateAuth(formats strfmt.Registry) error {

	if swag.IsZero(m.Auth) { // not required
		return nil
	}

	if m.Auth != nil {
		if err := m.Auth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth")
			}
			return err
		}
	}

	return nil
}

func (m *DeleteDevicesParamsBody) validateDevices(formats strfmt.Registry) error {

	if err := validate.Required("devices", "body", m.Devices); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeleteDevicesParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteDevicesParamsBody) UnmarshalBinary(b []byte) error {
	var res DeleteDevicesParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
