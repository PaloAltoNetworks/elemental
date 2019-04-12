// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
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
	Entity    []byte       `json:"entity"`
	Identity  string       `json:"identity"`
	Type      EventType    `json:"type"`
	Timestamp time.Time    `json:"timestamp"`
	Encoding  EncodingType `json:"encoding"`
}

// NewEvent returns a new Event.
func NewEvent(t EventType, o Identifiable) *Event {
	return NewEventWithEncoding(t, o, EncodingTypeMSGPACK)
}

// NewEventWithEncoding returns a new Event using the given encoding
func NewEventWithEncoding(t EventType, o Identifiable, encoding EncodingType) *Event {

	data, err := Encode(encoding, o)
	if err != nil {
		panic(err)
	}

	return &Event{
		Type:      t,
		Entity:    data,
		Identity:  o.Identity().Name,
		Timestamp: time.Now(),
		Encoding:  encoding,
	}
}

// Decode decodes the data into the given destination.
func (e *Event) Decode(dst interface{}) error {

	return Decode(e.Encoding, e.Entity, dst)
}

func (e *Event) String() string {

	return fmt.Sprintf("<event type: %s identity: %s>", e.Type, e.Identity)
}

// Duplicate creates a copy of the event.
func (e *Event) Duplicate() *Event {

	return &Event{
		Type:      e.Type,
		Entity:    e.Entity[:],
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
