// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImageTag image tag
//
// swagger:model ImageTag
type ImageTag struct {

	// automated
	Automated bool `json:"automated,omitempty"`

	// needs review
	NeedsReview bool `json:"needsReview,omitempty"`

	// tag
	Tag *ImageTagData `json:"tag,omitempty"`
}

// Validate validates this image tag
func (m *ImageTag) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageTag) validateTag(formats strfmt.Registry) error {
	if swag.IsZero(m.Tag) { // not required
		return nil
	}

	if m.Tag != nil {
		if err := m.Tag.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tag")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tag")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this image tag based on the context it is used
func (m *ImageTag) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageTag) contextValidateTag(ctx context.Context, formats strfmt.Registry) error {

	if m.Tag != nil {
		if err := m.Tag.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tag")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tag")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageTag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageTag) UnmarshalBinary(b []byte) error {
	var res ImageTag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
