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
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WebhookEventConfig WebhookEventConfig
//
// Represents a Webhook Event configuration.
// swagger:discriminator WebhookEventConfig Represents a Webhook Event configuration.
type WebhookEventConfig interface {
	runtime.Validatable

	// Action Type.
	Action() string
	SetAction(string)

	// endpoint
	Endpoint() string
	SetEndpoint(string)

	// event
	// Enum: [SUCCESS FAILURE WAITING CANCELED]
	Event() string
	SetEvent(string)

	// The headers required to perform HTTP action on the provided url.
	Headers() map[string]string
	SetHeaders(map[string]string)

	// Payload to be sent to the webhook.
	Payload() string
	SetPayload(string)

	// stage
	Stage() string
	SetStage(string)

	// task
	Task() string
	SetTask(string)

	// WebhookURL to send payload.
	URL() string
	SetURL(string)
}

type webhookEventConfig struct {
	actionField string

	endpointField string

	eventField string

	headersField map[string]string

	payloadField string

	stageField string

	taskField string

	urlField string
}

// Action gets the action of this polymorphic type
func (m *webhookEventConfig) Action() string {
	return m.actionField
}

// SetAction sets the action of this polymorphic type
func (m *webhookEventConfig) SetAction(val string) {
	m.actionField = val
}

// Endpoint gets the endpoint of this polymorphic type
func (m *webhookEventConfig) Endpoint() string {
	return m.endpointField
}

// SetEndpoint sets the endpoint of this polymorphic type
func (m *webhookEventConfig) SetEndpoint(val string) {
	m.endpointField = val
}

// Event gets the event of this polymorphic type
func (m *webhookEventConfig) Event() string {
	return m.eventField
}

// SetEvent sets the event of this polymorphic type
func (m *webhookEventConfig) SetEvent(val string) {
	m.eventField = val
}

// Headers gets the headers of this polymorphic type
func (m *webhookEventConfig) Headers() map[string]string {
	return m.headersField
}

// SetHeaders sets the headers of this polymorphic type
func (m *webhookEventConfig) SetHeaders(val map[string]string) {
	m.headersField = val
}

// Payload gets the payload of this polymorphic type
func (m *webhookEventConfig) Payload() string {
	return m.payloadField
}

// SetPayload sets the payload of this polymorphic type
func (m *webhookEventConfig) SetPayload(val string) {
	m.payloadField = val
}

// Stage gets the stage of this polymorphic type
func (m *webhookEventConfig) Stage() string {
	return m.stageField
}

// SetStage sets the stage of this polymorphic type
func (m *webhookEventConfig) SetStage(val string) {
	m.stageField = val
}

// Task gets the task of this polymorphic type
func (m *webhookEventConfig) Task() string {
	return m.taskField
}

// SetTask sets the task of this polymorphic type
func (m *webhookEventConfig) SetTask(val string) {
	m.taskField = val
}

// URL gets the url of this polymorphic type
func (m *webhookEventConfig) URL() string {
	return m.urlField
}

// SetURL sets the url of this polymorphic type
func (m *webhookEventConfig) SetURL(val string) {
	m.urlField = val
}

// UnmarshalWebhookEventConfigSlice unmarshals polymorphic slices of WebhookEventConfig
func UnmarshalWebhookEventConfigSlice(reader io.Reader, consumer runtime.Consumer) ([]WebhookEventConfig, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []WebhookEventConfig
	for _, element := range elements {
		obj, err := unmarshalWebhookEventConfig(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalWebhookEventConfig unmarshals polymorphic WebhookEventConfig
func UnmarshalWebhookEventConfig(reader io.Reader, consumer runtime.Consumer) (WebhookEventConfig, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalWebhookEventConfig(data, consumer)
}

func unmarshalWebhookEventConfig(data []byte, consumer runtime.Consumer) (WebhookEventConfig, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Represents a Webhook Event configuration. property.
	var getType struct {
		RepresentsaWebhookEventConfiguration string `json:"Represents a Webhook Event configuration."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Represents a Webhook Event configuration.", "body", getType.RepresentsaWebhookEventConfiguration); err != nil {
		return nil, err
	}

	// The value of Represents a Webhook Event configuration. is used to determine which type to create and unmarshal the data into
	switch getType.RepresentsaWebhookEventConfiguration {
	case "WebhookEventConfig":
		var result webhookEventConfig
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid Represents a Webhook Event configuration. value: %q", getType.RepresentsaWebhookEventConfiguration)

}

// Validate validates this webhook event config
func (m *webhookEventConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var webhookEventConfigTypeEventPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUCCESS","FAILURE","WAITING","CANCELED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		webhookEventConfigTypeEventPropEnum = append(webhookEventConfigTypeEventPropEnum, v)
	}
}

const (

	// WebhookEventConfigEventSUCCESS captures enum value "SUCCESS"
	WebhookEventConfigEventSUCCESS string = "SUCCESS"

	// WebhookEventConfigEventFAILURE captures enum value "FAILURE"
	WebhookEventConfigEventFAILURE string = "FAILURE"

	// WebhookEventConfigEventWAITING captures enum value "WAITING"
	WebhookEventConfigEventWAITING string = "WAITING"

	// WebhookEventConfigEventCANCELED captures enum value "CANCELED"
	WebhookEventConfigEventCANCELED string = "CANCELED"
)

// prop value enum
func (m *webhookEventConfig) validateEventEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, webhookEventConfigTypeEventPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *webhookEventConfig) validateEvent(formats strfmt.Registry) error {

	if swag.IsZero(m.Event()) { // not required
		return nil
	}

	// value enum
	if err := m.validateEventEnum("event", "body", m.Event()); err != nil {
		return err
	}

	return nil
}