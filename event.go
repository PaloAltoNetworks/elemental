// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"encoding/json"
	"fmt"
	"time"
)

// EventType is the type of an event.
type EventType string

const (
	// EventCreate is the type of creation event.
	EventCreate EventType = "create"

	// EventUpdate is the type of update event.
	EventUpdate EventType = "update"

	// EventDelete is the type of delete event.
	EventDelete EventType = "delete"
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

// NewEventWithEncoding returns a new Event using the given encoding
func NewEventWithEncoding(t EventType, o Identifiable, encoding EncodingType) *Event {

	data, err := Encode(encoding, o)
	if err != nil {
		panic(fmt.Sprintf("unable to create new event: %s", err))
	}

	evt := &Event{
		Type:      t,
		Identity:  o.Identity().Name,
		Timestamp: time.Now(),
		Encoding:  encoding,
	}

	if encoding == EncodingTypeJSON {
		evt.JSONData = json.RawMessage(data)
	} else {
		evt.RawData = data
	}

	return evt
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

	return &Event{
		Type:      e.Type,
		JSONData:  e.JSONData[:],
		RawData:   e.RawData[:],
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
