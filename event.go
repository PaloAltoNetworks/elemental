// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package elemental

import (
	"encoding/json"
	"fmt"
	"time"
)

// EventType is the type of an event.
type EventType string

const (
	// EventCreate is the type of creation events.
	EventCreate EventType = "create"

	// EventUpdate is the type of update events.
	EventUpdate EventType = "update"

	// EventDelete is the type of delete events.
	EventDelete EventType = "delete"

	// EventError is the type of error events.
	EventError EventType = "error"
)

// An Event represents a computational event.
type Event struct {
	RawData   []byte          `msgpack:"entity" json:"-"`
	JSONData  json.RawMessage `msgpack:"-" json:"entity"`
	Identity  string          `msgpack:"identity" json:"identity"`
	Type      EventType       `msgpack:"type" json:"type"`
	Timestamp time.Time       `msgpack:"timestamp" json:"timestamp"`
	Encoding  EncodingType    `msgpack:"encoding" json:"encoding"`
}

// NewEvent returns a new Event.
func NewEvent(t EventType, o Identifiable) *Event {
	return NewEventWithEncoding(t, o, EncodingTypeMSGPACK)
}

// NewErrorEvent returns a new (error) Event embedded with the provided elemental.Error
func NewErrorEvent(ee Error, encoding EncodingType) *Event {

	data, err := Encode(encoding, ee)
	if err != nil {
		panic(fmt.Sprintf("unable to create new error event: %s", err))
	}

	event := &Event{
		Type:      EventError,
		Timestamp: time.Now(),
		Encoding:  encoding,
	}

	event.configureData(encoding, data)
	return event
}

// NewEventWithEncoding returns a new Event using the given encoding
func NewEventWithEncoding(t EventType, o Identifiable, encoding EncodingType) *Event {

	data, err := Encode(encoding, o)
	if err != nil {
		panic(fmt.Sprintf("unable to create new event: %s", err))
	}

	event := &Event{
		Type:      t,
		Identity:  o.Identity().Name,
		Timestamp: time.Now(),
		Encoding:  encoding,
	}

	event.configureData(encoding, data)
	return event
}

func (e *Event) configureData(encoding EncodingType, data []byte) {
	switch encoding {
	case EncodingTypeJSON:
		e.JSONData = json.RawMessage(data)
	case EncodingTypeMSGPACK:
		e.RawData = data
	}
}

// GetEncoding returns the encoding used to encode the entity.
func (e *Event) GetEncoding() EncodingType {
	return e.Encoding
}

// Decode decodes the data into the given destination.
func (e *Event) Decode(dst interface{}) error {
	return Decode(e.GetEncoding(), e.Entity(), dst)
}

// Convert converts the internal encoded data to the given
// encoding.
func (e *Event) Convert(encoding EncodingType) error {

	switch e.Encoding {

	case encoding:
		return nil

	case EncodingTypeMSGPACK:

		d, err := Convert(e.Encoding, encoding, e.RawData)
		if err != nil {
			return err
		}
		e.JSONData = json.RawMessage(d)
		e.RawData = nil

	default:

		d, err := Convert(e.Encoding, encoding, []byte(e.JSONData))
		if err != nil {
			return err
		}
		e.JSONData = nil
		e.RawData = d
	}

	e.Encoding = encoding

	return nil
}

// Entity returns the byte encoded entity.
func (e *Event) Entity() []byte {

	switch e.Encoding {
	case EncodingTypeMSGPACK:
		return e.RawData
	default:
		return []byte(e.JSONData)
	}
}

func (e *Event) String() string {

	return fmt.Sprintf("<event type: %s identity: %s>", e.Type, e.Identity)
}

// Duplicate creates a copy of the event.
func (e *Event) Duplicate() *Event {

	var jd json.RawMessage
	var rd []byte

	if e.JSONData != nil {
		jd = append(json.RawMessage{}, e.JSONData...)
	}

	if e.RawData != nil {
		rd = append([]byte{}, e.RawData...)
	}

	return &Event{
		Type:      e.Type,
		JSONData:  jd,
		RawData:   rd,
		Identity:  e.Identity,
		Timestamp: e.Timestamp,
		Encoding:  e.Encoding,
	}
}

// An Events represents a list of Event.
type Events []*Event

// NewEvents retutns a new Events.
func NewEvents(events ...*Event) Events {

	return append(Events{}, events...)
}
