// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V2AvailablePropertyDetails v2 available property details
//
// swagger:model v2AvailablePropertyDetails
type V2AvailablePropertyDetails struct {

	// 'Street No.' in in a 'Street Address'
	PropertyAddress1 string `json:"PropertyAddress1,omitempty"`

	// 'Street Name' in in a 'Street Address'
	PropertyAddress2 string `json:"PropertyAddress2,omitempty"`

	// 'Suburb' in in a 'Street Address'
	PropertyAddress3 string `json:"PropertyAddress3,omitempty"`

	// 'City' in in a 'Street Address'
	PropertyAddress4 string `json:"PropertyAddress4,omitempty"`

	// property agent
	PropertyAgent *V2AvailablePropertyAgentDetails `json:"PropertyAgent,omitempty"`

	// Code indicating how many times this record has been updated.  The code is incremented by 1 on each update.  This can also be used to track if the record has been updated since the last request.
	PropertyChangeCode float64 `json:"PropertyChangeCode,omitempty"`

	// The 'Internal Code' of the 'Property'.  This code is only unique within the client database.  It is recommended that you use 'PropertyCodeGlobal' if you are making the data public or consuming property information from multiple client databases.
	PropertyCode string `json:"PropertyCode,omitempty"`

	// The 'Global Code' of the 'Property'.  This code is unique across ALL Palace databases.  It is recommended that you use 'PropertyCodeGlobal' if you are making the data public or consuming property information from multiple client databases.
	PropertyCodeGlobal string `json:"PropertyCodeGlobal,omitempty"`

	// property custom list
	PropertyCustomList []*V2AvailablePropertyCustomList `json:"PropertyCustomList"`

	// Date the property is available from.  This will be formatted as yyyy-MM-dd
	PropertyDateAvailable string `json:"PropertyDateAvailable,omitempty"`

	// property features
	PropertyFeatures *V2AvailablePropertyFeatures `json:"PropertyFeatures,omitempty"`

	// Generally is used as in 'Internal' reference for a 'Map Grid' for grouping purposes.
	PropertyGrid string `json:"PropertyGrid,omitempty"`

	// A description of the 'Management Type' of the property.  Although these descriptions can be defined by the user, Palace has a 'Pattern Criteria' definiting 4 specific types of 'Managements'. Descriptions containing the word 'Manage' specifies a residential management.  Descriptions containing the word 'Casual' specifies a casual property. Descriptions containing the word 'Commercial' specifies a commercial management.  Descriptions containing the word 'Holiday' specifies a holiday management
	PropertyManagementType string `json:"PropertyManagementType,omitempty"`

	// This is user defined 'Free Form' text giving a name to a property (NOTE: This is not the same as the address of the property)
	PropertyName string `json:"PropertyName,omitempty"`

	// Internal code assigned to the 'Owner' of the property.
	PropertyOwnerCode string `json:"PropertyOwnerCode,omitempty"`

	// This is the amount of the rent.
	PropertyRentAmount float64 `json:"PropertyRentAmount,omitempty"`

	// This is related to the the 'PropertyRentAmount' field (e.g. 200 / Week).  The 'Rental Period' indicates how often the rent should be paid on this property.  Options are Day, Week, Fortnight, Month
	PropertyRentalPeriod string `json:"PropertyRentalPeriod,omitempty"`

	// This is a user defined code for sorting.  Generally, this will be the same as the 'Street Addres' name
	PropertySortCode string `json:"PropertySortCode,omitempty"`

	// This value will be 'Active' or 'Inactive' and indicates if the 'Property' is an active management at this time (NOTE: An 'Inactive' management may suggest the 'Property' is having work done at this time and cannot currently be rented)
	PropertyStatus string `json:"PropertyStatus,omitempty"`

	// property suburb
	PropertySuburb []*V2AvailablePropertySuburbs `json:"PropertySuburb"`

	// The 'Unit' of the 'Property'.  If displaying the full 'Property Address' and a 'Unit' exists, the system should display 'Unit / PropertyAddress1' (e.g. 1/15 Stuart Street).
	PropertyUnit string `json:"PropertyUnit,omitempty"`
}

// Validate validates this v2 available property details
func (m *V2AvailablePropertyDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePropertyAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropertyCustomList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropertyFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropertySuburb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2AvailablePropertyDetails) validatePropertyAgent(formats strfmt.Registry) error {

	if swag.IsZero(m.PropertyAgent) { // not required
		return nil
	}

	if m.PropertyAgent != nil {
		if err := m.PropertyAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PropertyAgent")
			}
			return err
		}
	}

	return nil
}

func (m *V2AvailablePropertyDetails) validatePropertyCustomList(formats strfmt.Registry) error {

	if swag.IsZero(m.PropertyCustomList) { // not required
		return nil
	}

	for i := 0; i < len(m.PropertyCustomList); i++ {
		if swag.IsZero(m.PropertyCustomList[i]) { // not required
			continue
		}

		if m.PropertyCustomList[i] != nil {
			if err := m.PropertyCustomList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PropertyCustomList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V2AvailablePropertyDetails) validatePropertyFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.PropertyFeatures) { // not required
		return nil
	}

	if m.PropertyFeatures != nil {
		if err := m.PropertyFeatures.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PropertyFeatures")
			}
			return err
		}
	}

	return nil
}

func (m *V2AvailablePropertyDetails) validatePropertySuburb(formats strfmt.Registry) error {

	if swag.IsZero(m.PropertySuburb) { // not required
		return nil
	}

	for i := 0; i < len(m.PropertySuburb); i++ {
		if swag.IsZero(m.PropertySuburb[i]) { // not required
			continue
		}

		if m.PropertySuburb[i] != nil {
			if err := m.PropertySuburb[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PropertySuburb" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2AvailablePropertyDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2AvailablePropertyDetails) UnmarshalBinary(b []byte) error {
	var res V2AvailablePropertyDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
