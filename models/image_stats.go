// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImageStats image stats
//
// swagger:model ImageStats
type ImageStats struct {

	// The number of comment reactions.
	CommentCount int64 `json:"commentCount,omitempty"`

	// The number of cry reactions.
	CryCount int64 `json:"cryCount,omitempty"`

	// dislike count
	DislikeCount int64 `json:"dislikeCount,omitempty"`

	// The number of heart reactions.
	HeartCount int64 `json:"heartCount,omitempty"`

	// The number of laugh reactions.
	LaughCount int64 `json:"laughCount,omitempty"`

	// The number of like reactions.
	LikeCount int64 `json:"likeCount,omitempty"`
}

// Validate validates this image stats
func (m *ImageStats) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image stats based on context it is used
func (m *ImageStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImageStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageStats) UnmarshalBinary(b []byte) error {
	var res ImageStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
