// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// StepType step type
// swagger:model step-type
type StepType string

const (

	// StepTypeHardawareInfo captures enum value "hardaware-info"
	StepTypeHardawareInfo StepType = "hardaware-info"

	// StepTypeConnectivityCheck captures enum value "connectivity-check"
	StepTypeConnectivityCheck StepType = "connectivity-check"
)

// for schema
var stepTypeEnum []interface{}

func init() {
	var res []StepType
	if err := json.Unmarshal([]byte(`["hardaware-info","connectivity-check"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stepTypeEnum = append(stepTypeEnum, v)
	}
}

func (m StepType) validateStepTypeEnum(path, location string, value StepType) error {
	if err := validate.Enum(path, location, value, stepTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this step type
func (m StepType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateStepTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
