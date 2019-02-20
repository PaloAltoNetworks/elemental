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
	Entity    json.RawMessage `json:"entity"`
	Identity  string          `json:"identity"`
	Type      EventType       `json:"type"`
	Timestamp time.Time       `json:"timestamp"`
}

// NewEvent returns a new Event.
func NewEvent(t EventType, o Identifiable) *Event {

	return NewEventWithMarshaler(t, o, json.Marshal)
}

// NewEventWithMarshaler returns a new Event with the identifiable encoded with the provided marshalFunc.
func NewEventWithMarshaler(t EventType, o Identifiable, marshalFunc func(interface{}) ([]byte, error)) *Event {

	data, err := marshalFunc(o)
	if err != nil {
		panic(err)
	}

	return &Event{
		Type:      t,
		Entity:    data,
		Identity:  o.Identity().Name,
		Timestamp: time.Now(),
	}
}

// Decode decodes the data into the given destination.
func (e *Event) Decode(dst interface{}) error {

	return e.DecodeWithUnmarshaler(dst, json.Unmarshal)
}

// DecodeWithUnmarshaler decodes the data into the given destination using the given unmarshaller
func (e *Event) DecodeWithUnmarshaler(dst interface{}, unmarshalFunc func([]byte, interface{}) error) error {

	return unmarshalFunc(e.Entity, &dst)
}

func (e *Event) String() string {

	return fmt.Sprintf("<event type: %s identity: %s>", e.Type, e.Identity)
}

// Duplicate creates a copy of the event.
func (e *Event) Duplicate() *Event {

	return &Event{
		Type:      e.Type,
		Entity:    e.Entity,
		Identity:  e.Identity,
		Timestamp: e.Timestamp,
	}
}

// An Events represents a list of Event.
type Events []*Event

// NewEvents retutns a new Events.
func NewEvents(events ...*Event) Events {

	return append(Events{}, events...)
}
