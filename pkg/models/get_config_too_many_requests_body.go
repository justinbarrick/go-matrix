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

// GetConfigTooManyRequestsBody A Matrix-level Error
// swagger:model getConfigTooManyRequestsBody
type GetConfigTooManyRequestsBody struct {

	// An error code.
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`
}

// Validate validates this get config too many requests body
func (m *GetConfigTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetConfigTooManyRequestsBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("errcode", "body", m.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetConfigTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetConfigTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res GetConfigTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
