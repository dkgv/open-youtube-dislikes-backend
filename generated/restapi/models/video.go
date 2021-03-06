// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Video video
//
// swagger:model Video
type Video struct {

	// comments
	Comments int64 `json:"comments,omitempty"`

	// dislikes
	Dislikes int64 `json:"dislikes,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// id hash
	IDHash string `json:"idHash,omitempty"`

	// likes
	Likes int64 `json:"likes,omitempty"`

	// published at
	PublishedAt int64 `json:"publishedAt,omitempty"`

	// subscribers
	Subscribers int64 `json:"subscribers,omitempty"`

	// views
	Views int64 `json:"views,omitempty"`
}

// Validate validates this video
func (m *Video) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this video based on context it is used
func (m *Video) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Video) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Video) UnmarshalBinary(b []byte) error {
	var res Video
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
