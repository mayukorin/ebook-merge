// Code generated by go-swagger; DO NOT EDIT.

package generated_swagger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListEbook ListEbook
//
// swagger:model ListEbook
type ListEbook struct {

	// ebooks
	Ebooks []*Ebook `json:"ebooks"`
}

// Validate validates this list ebook
func (m *ListEbook) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEbooks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListEbook) validateEbooks(formats strfmt.Registry) error {
	if swag.IsZero(m.Ebooks) { // not required
		return nil
	}

	for i := 0; i < len(m.Ebooks); i++ {
		if swag.IsZero(m.Ebooks[i]) { // not required
			continue
		}

		if m.Ebooks[i] != nil {
			if err := m.Ebooks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ebooks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ebooks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list ebook based on the context it is used
func (m *ListEbook) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEbooks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListEbook) contextValidateEbooks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ebooks); i++ {

		if m.Ebooks[i] != nil {
			if err := m.Ebooks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ebooks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ebooks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListEbook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListEbook) UnmarshalBinary(b []byte) error {
	var res ListEbook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
