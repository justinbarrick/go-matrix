// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// SearchOKBodySearchCategoriesRoomEventsState Current state
//
// The current state for every room in the results.
// This is included if the request had the
// ``include_state`` key set with a value of ``true``.
//
// The ``string`` key is the room ID for which the ``State
// Event`` array belongs to.
// swagger:model searchOKBodySearchCategoriesRoomEventsState
type SearchOKBodySearchCategoriesRoomEventsState map[string][]SearchOKBodySearchCategoriesRoomEventsStateAdditionalPropertiesItems

// Validate validates this search o k body search categories room events state
func (m SearchOKBodySearchCategoriesRoomEventsState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", SearchOKBodySearchCategoriesRoomEventsState(m)); err != nil {
		return err
	}

	for k := range m {

		if err := validate.Required(k, "body", m[k]); err != nil {
			return err
		}

		for i := 0; i < len(m[k]); i++ {

			if err := m[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
