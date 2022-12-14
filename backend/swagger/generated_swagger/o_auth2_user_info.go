// Code generated by go-swagger; DO NOT EDIT.

package generated_swagger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OAuth2UserInfo OAuth2UserInfo
//
// swagger:model OAuth2UserInfo
type OAuth2UserInfo struct {

	// email
	Email string `json:"email,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// picture
	Picture string `json:"picture,omitempty"`

	// verified email
	VerifiedEmail bool `json:"verified_email,omitempty"`
}

// Validate validates this o auth2 user info
func (m *OAuth2UserInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this o auth2 user info based on context it is used
func (m *OAuth2UserInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OAuth2UserInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OAuth2UserInfo) UnmarshalBinary(b []byte) error {
	var res OAuth2UserInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
