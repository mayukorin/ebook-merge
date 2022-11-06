// Code generated by go-swagger; DO NOT EDIT.

package generated_swagger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EbookService EbookService
//
// swagger:model EbookService
type EbookService struct {

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this ebook service
func (m *EbookService) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ebook service based on context it is used
func (m *EbookService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EbookService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EbookService) UnmarshalBinary(b []byte) error {
	var res EbookService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
