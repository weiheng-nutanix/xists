// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IamTokenResponse IAM token response
//
// IAM token response
// swagger:model iam_token_response
type IamTokenResponse struct {

	// access_token
	// Required: true
	AccessToken *string `json:"access_token"`

	// Token expiration time in seconds
	// Required: true
	ExpiresIn *int64 `json:"expires_in"`

	// The typical value is "bearer".
	// Required: true
	TokenType *string `json:"token_type"`
}

// Validate validates this iam token response
func (m *IamTokenResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IamTokenResponse) validateAccessToken(formats strfmt.Registry) error {

	if err := validate.Required("access_token", "body", m.AccessToken); err != nil {
		return err
	}

	return nil
}

func (m *IamTokenResponse) validateExpiresIn(formats strfmt.Registry) error {

	if err := validate.Required("expires_in", "body", m.ExpiresIn); err != nil {
		return err
	}

	return nil
}

func (m *IamTokenResponse) validateTokenType(formats strfmt.Registry) error {

	if err := validate.Required("token_type", "body", m.TokenType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IamTokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamTokenResponse) UnmarshalBinary(b []byte) error {
	var res IamTokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
