// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"
)

// PipelinePatchRequest PipelinePatchRequest
//
// Patch Request for a Pipeline
// swagger:discriminator PipelinePatchRequest Patch Request for a Pipeline
type PipelinePatchRequest interface {
	runtime.Validatable

	// A human-friendly description for the Pipeline.
	Description() string
	SetDescription(string)

	// Indicates if the Pipeline is in enabled state.
	Enabled() bool
	SetEnabled(bool)

	// A human-friendly name used as an identifier for the Pipeline.
	Name() string
	SetName(string)

	// A set of tag keys and optional values that need to be set on the Pipeline.
	Tags() []string
	SetTags([]string)
}

type pipelinePatchRequest struct {
	descriptionField string

	enabledField bool

	nameField string

	tagsField []string
}

// Description gets the description of this polymorphic type
func (m *pipelinePatchRequest) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this polymorphic type
func (m *pipelinePatchRequest) SetDescription(val string) {
	m.descriptionField = val
}

// Enabled gets the enabled of this polymorphic type
func (m *pipelinePatchRequest) Enabled() bool {
	return m.enabledField
}

// SetEnabled sets the enabled of this polymorphic type
func (m *pipelinePatchRequest) SetEnabled(val bool) {
	m.enabledField = val
}

// Name gets the name of this polymorphic type
func (m *pipelinePatchRequest) Name() string {
	return m.nameField
}

// SetName sets the name of this polymorphic type
func (m *pipelinePatchRequest) SetName(val string) {
	m.nameField = val
}

// Tags gets the tags of this polymorphic type
func (m *pipelinePatchRequest) Tags() []string {
	return m.tagsField
}

// SetTags sets the tags of this polymorphic type
func (m *pipelinePatchRequest) SetTags(val []string) {
	m.tagsField = val
}

// UnmarshalPipelinePatchRequestSlice unmarshals polymorphic slices of PipelinePatchRequest
func UnmarshalPipelinePatchRequestSlice(reader io.Reader, consumer runtime.Consumer) ([]PipelinePatchRequest, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []PipelinePatchRequest
	for _, element := range elements {
		obj, err := unmarshalPipelinePatchRequest(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalPipelinePatchRequest unmarshals polymorphic PipelinePatchRequest
func UnmarshalPipelinePatchRequest(reader io.Reader, consumer runtime.Consumer) (PipelinePatchRequest, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalPipelinePatchRequest(data, consumer)
}

func unmarshalPipelinePatchRequest(data []byte, consumer runtime.Consumer) (PipelinePatchRequest, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Patch Request for a Pipeline property.
	var getType struct {
		PatchRequestForaPipeline string `json:"Patch Request for a Pipeline"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Patch Request for a Pipeline", "body", getType.PatchRequestForaPipeline); err != nil {
		return nil, err
	}

	// The value of Patch Request for a Pipeline is used to determine which type to create and unmarshal the data into
	switch getType.PatchRequestForaPipeline {
	case "PipelinePatchRequest":
		var result pipelinePatchRequest
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid Patch Request for a Pipeline value: %q", getType.PatchRequestForaPipeline)

}

// Validate validates this pipeline patch request
func (m *pipelinePatchRequest) Validate(formats strfmt.Registry) error {
	return nil
}