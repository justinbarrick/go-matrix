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

// SearchUserDirectoryParamsBody search user directory params body
// swagger:model searchUserDirectoryParamsBody
type SearchUserDirectoryParamsBody struct {

	// The maximum number of results to return. Defaults to 10.
	Limit int64 `json:"limit,omitempty"`

	// The term to search for
	// Required: true
	SearchTerm *string `json:"search_term"`
}

// Validate validates this search user directory params body
func (m *SearchUserDirectoryParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSearchTerm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchUserDirectoryParamsBody) validateSearchTerm(formats strfmt.Registry) error {

	if err := validate.Required("search_term", "body", m.SearchTerm); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchUserDirectoryParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchUserDirectoryParamsBody) UnmarshalBinary(b []byte) error {
	var res SearchUserDirectoryParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
