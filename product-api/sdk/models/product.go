// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Product Product defines the structure for an API product.
//
// Product defines the structure for an API product
//
// swagger:model Product
type Product struct {

	// the description for this poduct
	// Max Length: 10000
	Description string `json:"description,omitempty"`

	// the id for the product
	// Minimum: 1
	ID int64 `json:"id,omitempty"`

	// the name for this poduct
	// Required: true
	// Max Length: 255
	Name *string `json:"name"`

	// the price for the product
	// Required: true
	// Minimum: 0.01
	Price *float32 `json:"price"`

	// the SKU for the product
	// Required: true
	// Pattern: [a-z]+-[a-z]+-[a-z]+
	SKU *string `json:"sku"`
}

// Validate validates this product
func (m *Product) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSKU(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Product) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 10000); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinimumInt("id", "body", int64(m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 255); err != nil {
		return err
	}

	return nil
}

func (m *Product) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	if err := validate.Minimum("price", "body", float64(*m.Price), 0.01, false); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateSKU(formats strfmt.Registry) error {

	if err := validate.Required("sku", "body", m.SKU); err != nil {
		return err
	}

	if err := validate.Pattern("sku", "body", string(*m.SKU), `[a-z]+-[a-z]+-[a-z]+`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Product) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Product) UnmarshalBinary(b []byte) error {
	var res Product
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
