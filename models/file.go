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

// File file
//
// swagger:model File
type File struct {

	// download Url
	DownloadURL string `json:"downloadUrl,omitempty"`

	// The format of the file ('pickle' or 'safetensor').
	Format string `json:"format,omitempty"`

	// hashes
	Hashes *Hash `json:"hashes,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// metadata
	Metadata *FileMetadata `json:"metadata,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// pickle scan message
	PickleScanMessage string `json:"pickleScanMessage,omitempty"`

	// Status of the pickle scan ('Pending', 'Success', 'Danger', 'Error').
	PickleScanResult string `json:"pickleScanResult,omitempty"`

	// If the file is the primary file for the model version.
	Primary bool `json:"primary,omitempty"`

	// The date in which the file was scanned.
	// Format: date-time
	ScannedAt strfmt.DateTime `json:"scannedAt,omitempty"`

	// The size of the model file.
	SizeKB float64 `json:"sizeKB,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// virus scan message
	VirusScanMessage string `json:"virusScanMessage,omitempty"`

	// Status of the virus scan ('Pending', 'Success', 'Danger', 'Error').
	VirusScanResult string `json:"virusScanResult,omitempty"`
}

// Validate validates this file
func (m *File) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHashes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScannedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *File) validateHashes(formats strfmt.Registry) error {
	if swag.IsZero(m.Hashes) { // not required
		return nil
	}

	if m.Hashes != nil {
		if err := m.Hashes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hashes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hashes")
			}
			return err
		}
	}

	return nil
}

func (m *File) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *File) validateScannedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ScannedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("scannedAt", "body", "date-time", m.ScannedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this file based on the context it is used
func (m *File) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHashes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *File) contextValidateHashes(ctx context.Context, formats strfmt.Registry) error {

	if m.Hashes != nil {

		if swag.IsZero(m.Hashes) { // not required
			return nil
		}

		if err := m.Hashes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hashes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hashes")
			}
			return err
		}
	}

	return nil
}

func (m *File) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {

		if swag.IsZero(m.Metadata) { // not required
			return nil
		}

		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *File) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *File) UnmarshalBinary(b []byte) error {
	var res File
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
