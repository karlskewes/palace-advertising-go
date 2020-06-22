// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V2AvailablePropertyPrevious v2 available property previous
//
// swagger:model v2AvailablePropertyPrevious
type V2AvailablePropertyPrevious struct {

	// The 'Global Code' of the 'Property'.  This code is unique across ALL Palace databases.  It is recommended that you use 'PropertyCodeGlobal' if you are making the data public or consuming property information from multiple client databases.
	PropertyCodeGlobal string `json:"PropertyCodeGlobal,omitempty"`

	// When the 'Date Available' was removed from the 'Property' (Marking the property as no longer available).  This will be formatted as yyyy-MM-dd
	PropertyDateAvailableRemoved string `json:"PropertyDateAvailableRemoved,omitempty"`

	// Descripton of the 'Reason' the 'Date Available' was removed.  This is separated by a colon (See example above)
	PropertyDateAvailableRemovedReason string `json:"PropertyDateAvailableRemovedReason,omitempty"`
}

// Validate validates this v2 available property previous
func (m *V2AvailablePropertyPrevious) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2AvailablePropertyPrevious) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2AvailablePropertyPrevious) UnmarshalBinary(b []byte) error {
	var res V2AvailablePropertyPrevious
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}