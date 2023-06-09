// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StandardError standard error
//
// swagger:model StandardError
type StandardError struct {

	// error
	// Example: Something went wrong while writing the file to disk, the server may be out of storage space.
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this standard error
func (m *StandardError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StandardError) validateError(formats strfmt.Registry) error {

	if err := validate.Required("error", "body", m.Error); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this standard error based on context it is used
func (m *StandardError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StandardError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StandardError) UnmarshalBinary(b []byte) error {
	var res StandardError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
