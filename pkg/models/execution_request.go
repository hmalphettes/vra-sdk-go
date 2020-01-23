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

// ExecutionRequest ExecutionRequest
//
// Execution Request for a pipeline
// swagger:discriminator ExecutionRequest Execution Request for a pipeline
type ExecutionRequest interface {
	runtime.Validatable

	// comments
	Comments() string
	SetComments(string)

	// execution Id
	ExecutionID() string
	SetExecutionID(string)

	// execution link
	ExecutionLink() string
	SetExecutionLink(string)

	// input
	Input() interface{}
	SetInput(interface{})

	// source
	Source() string
	SetSource(string)

	// tags
	Tags() []string
	SetTags([]string)
}

type executionRequest struct {
	commentsField string

	executionIdField string

	executionLinkField string

	inputField interface{}

	sourceField string

	tagsField []string
}

// Comments gets the comments of this polymorphic type
func (m *executionRequest) Comments() string {
	return m.commentsField
}

// SetComments sets the comments of this polymorphic type
func (m *executionRequest) SetComments(val string) {
	m.commentsField = val
}

// ExecutionID gets the execution Id of this polymorphic type
func (m *executionRequest) ExecutionID() string {
	return m.executionIdField
}

// SetExecutionID sets the execution Id of this polymorphic type
func (m *executionRequest) SetExecutionID(val string) {
	m.executionIdField = val
}

// ExecutionLink gets the execution link of this polymorphic type
func (m *executionRequest) ExecutionLink() string {
	return m.executionLinkField
}

// SetExecutionLink sets the execution link of this polymorphic type
func (m *executionRequest) SetExecutionLink(val string) {
	m.executionLinkField = val
}

// Input gets the input of this polymorphic type
func (m *executionRequest) Input() interface{} {
	return m.inputField
}

// SetInput sets the input of this polymorphic type
func (m *executionRequest) SetInput(val interface{}) {
	m.inputField = val
}

// Source gets the source of this polymorphic type
func (m *executionRequest) Source() string {
	return m.sourceField
}

// SetSource sets the source of this polymorphic type
func (m *executionRequest) SetSource(val string) {
	m.sourceField = val
}

// Tags gets the tags of this polymorphic type
func (m *executionRequest) Tags() []string {
	return m.tagsField
}

// SetTags sets the tags of this polymorphic type
func (m *executionRequest) SetTags(val []string) {
	m.tagsField = val
}

// UnmarshalExecutionRequestSlice unmarshals polymorphic slices of ExecutionRequest
func UnmarshalExecutionRequestSlice(reader io.Reader, consumer runtime.Consumer) ([]ExecutionRequest, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []ExecutionRequest
	for _, element := range elements {
		obj, err := unmarshalExecutionRequest(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalExecutionRequest unmarshals polymorphic ExecutionRequest
func UnmarshalExecutionRequest(reader io.Reader, consumer runtime.Consumer) (ExecutionRequest, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalExecutionRequest(data, consumer)
}

func unmarshalExecutionRequest(data []byte, consumer runtime.Consumer) (ExecutionRequest, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Execution Request for a pipeline property.
	var getType struct {
		ExecutionRequestForaPipeline string `json:"Execution Request for a pipeline"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Execution Request for a pipeline", "body", getType.ExecutionRequestForaPipeline); err != nil {
		return nil, err
	}

	// The value of Execution Request for a pipeline is used to determine which type to create and unmarshal the data into
	switch getType.ExecutionRequestForaPipeline {
	case "ExecutionRequest":
		var result executionRequest
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid Execution Request for a pipeline value: %q", getType.ExecutionRequestForaPipeline)

}

// Validate validates this execution request
func (m *executionRequest) Validate(formats strfmt.Registry) error {
	return nil
}