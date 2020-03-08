// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StepReply step reply
// swagger:model step-reply
type StepReply struct {

	// data
	Data string `json:"data,omitempty"`

	// return code
	ReturnCode int64 `json:"return-code,omitempty"`

	// step type
	StepType StepType `json:"step-type,omitempty"`
}

// Validate validates this step reply
func (m *StepReply) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStepType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StepReply) validateStepType(formats strfmt.Registry) error {

	if swag.IsZero(m.StepType) { // not required
		return nil
	}

	if err := m.StepType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("step-type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StepReply) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StepReply) UnmarshalBinary(b []byte) error {
	var res StepReply
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
