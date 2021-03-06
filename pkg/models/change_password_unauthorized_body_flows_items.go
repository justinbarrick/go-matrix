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

// ChangePasswordUnauthorizedBodyFlowsItems change password unauthorized body flows items
// swagger:model changePasswordUnauthorizedBodyFlowsItems
type ChangePasswordUnauthorizedBodyFlowsItems struct {

	// The login type of each of the stages required to complete this
	// authentication flow
	// Required: true
	Stages []string `json:"stages"`
}

// Validate validates this change password unauthorized body flows items
func (m *ChangePasswordUnauthorizedBodyFlowsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChangePasswordUnauthorizedBodyFlowsItems) validateStages(formats strfmt.Registry) error {

	if err := validate.Required("stages", "body", m.Stages); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChangePasswordUnauthorizedBodyFlowsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangePasswordUnauthorizedBodyFlowsItems) UnmarshalBinary(b []byte) error {
	var res ChangePasswordUnauthorizedBodyFlowsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
