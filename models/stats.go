// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Stats stats
//
// swagger:model Stats
type Stats struct {

	// comment count
	CommentCount int64 `json:"commentCount,omitempty"`

	// download count
	DownloadCount int64 `json:"downloadCount,omitempty"`

	// favorite count
	FavoriteCount int64 `json:"favoriteCount,omitempty"`

	// rating
	Rating float64 `json:"rating,omitempty"`

	// rating count
	RatingCount int64 `json:"ratingCount,omitempty"`

	// thumbs down count
	ThumbsDownCount int64 `json:"thumbsDownCount,omitempty"`

	// thumbs up count
	ThumbsUpCount int64 `json:"thumbsUpCount,omitempty"`

	// tipped amount count
	TippedAmountCount int64 `json:"tippedAmountCount,omitempty"`
}

// Validate validates this stats
func (m *Stats) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this stats based on context it is used
func (m *Stats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Stats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Stats) UnmarshalBinary(b []byte) error {
	var res Stats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
