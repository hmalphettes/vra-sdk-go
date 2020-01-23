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

// Workspace Workspace
//
// Model which contains details about container and host for executing continuous integration tasks.
// swagger:discriminator Workspace Model which contains details about container and host for executing continuous integration tasks.
type Workspace interface {
	runtime.Validatable

	// Indicates that git clone will be performed automatically.
	AutoCloneForTrigger() bool
	SetAutoCloneForTrigger(bool)

	// List of paths to store artifacts and logs.
	Cache() []string
	SetCache([]string)

	// Name of the endpoint is docker.
	Endpoint() string
	SetEndpoint(string)

	// Image to create CI container.
	Image() string
	SetImage(string)

	// Working directory for executing commands.
	Path() string
	SetPath(string)

	// Name of the docker registry.
	Registry() string
	SetRegistry(string)
}

type workspace struct {
	autoCloneForTriggerField bool

	cacheField []string

	endpointField string

	imageField string

	pathField string

	registryField string
}

// AutoCloneForTrigger gets the auto clone for trigger of this polymorphic type
func (m *workspace) AutoCloneForTrigger() bool {
	return m.autoCloneForTriggerField
}

// SetAutoCloneForTrigger sets the auto clone for trigger of this polymorphic type
func (m *workspace) SetAutoCloneForTrigger(val bool) {
	m.autoCloneForTriggerField = val
}

// Cache gets the cache of this polymorphic type
func (m *workspace) Cache() []string {
	return m.cacheField
}

// SetCache sets the cache of this polymorphic type
func (m *workspace) SetCache(val []string) {
	m.cacheField = val
}

// Endpoint gets the endpoint of this polymorphic type
func (m *workspace) Endpoint() string {
	return m.endpointField
}

// SetEndpoint sets the endpoint of this polymorphic type
func (m *workspace) SetEndpoint(val string) {
	m.endpointField = val
}

// Image gets the image of this polymorphic type
func (m *workspace) Image() string {
	return m.imageField
}

// SetImage sets the image of this polymorphic type
func (m *workspace) SetImage(val string) {
	m.imageField = val
}

// Path gets the path of this polymorphic type
func (m *workspace) Path() string {
	return m.pathField
}

// SetPath sets the path of this polymorphic type
func (m *workspace) SetPath(val string) {
	m.pathField = val
}

// Registry gets the registry of this polymorphic type
func (m *workspace) Registry() string {
	return m.registryField
}

// SetRegistry sets the registry of this polymorphic type
func (m *workspace) SetRegistry(val string) {
	m.registryField = val
}

// UnmarshalWorkspaceSlice unmarshals polymorphic slices of Workspace
func UnmarshalWorkspaceSlice(reader io.Reader, consumer runtime.Consumer) ([]Workspace, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Workspace
	for _, element := range elements {
		obj, err := unmarshalWorkspace(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalWorkspace unmarshals polymorphic Workspace
func UnmarshalWorkspace(reader io.Reader, consumer runtime.Consumer) (Workspace, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalWorkspace(data, consumer)
}

func unmarshalWorkspace(data []byte, consumer runtime.Consumer) (Workspace, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Model which contains details about container and host for executing continuous integration tasks. property.
	var getType struct {
		ModelWhichContainsDetailsAboutContainerAndHostForExecutingContinuousIntegrationTasks string `json:"Model which contains details about container and host for executing continuous integration tasks."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Model which contains details about container and host for executing continuous integration tasks.", "body", getType.ModelWhichContainsDetailsAboutContainerAndHostForExecutingContinuousIntegrationTasks); err != nil {
		return nil, err
	}

	// The value of Model which contains details about container and host for executing continuous integration tasks. is used to determine which type to create and unmarshal the data into
	switch getType.ModelWhichContainsDetailsAboutContainerAndHostForExecutingContinuousIntegrationTasks {
	case "Workspace":
		var result workspace
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid Model which contains details about container and host for executing continuous integration tasks. value: %q", getType.ModelWhichContainsDetailsAboutContainerAndHostForExecutingContinuousIntegrationTasks)

}

// Validate validates this workspace
func (m *workspace) Validate(formats strfmt.Registry) error {
	return nil
}