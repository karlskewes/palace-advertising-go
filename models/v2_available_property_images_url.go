// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V2AvailablePropertyImagesURL v2 available property images URL
//
// swagger:model v2AvailablePropertyImagesURL
type V2AvailablePropertyImagesURL struct {

	// Internal 'Image Code' (This is only unique to a specific image)
	PropertyImageCode string `json:"PropertyImageCode,omitempty"`

	// Global 'Image Code' (This is unique across ALL Palace databases)
	PropertyImageCodeGlobal string `json:"PropertyImageCodeGlobal,omitempty"`

	// Free form user defined description.
	PropertyImageDescription string `json:"PropertyImageDescription,omitempty"`

	// URL to Full Image (Not resized)
	PropertyImageURL string `json:"PropertyImageURL,omitempty"`

	// URL to Image (Resized to medium thumbnail)
	PropertyImageURLThumbnailMedium string `json:"PropertyImageURLThumbnailMedium,omitempty"`

	// URL to Image (Resized to small thumbnail)
	PropertyImageURLThumbnailSmall string `json:"PropertyImageURLThumbnailSmall,omitempty"`
}

// Validate validates this v2 available property images URL
func (m *V2AvailablePropertyImagesURL) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2AvailablePropertyImagesURL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2AvailablePropertyImagesURL) UnmarshalBinary(b []byte) error {
	var res V2AvailablePropertyImagesURL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
