// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"encoding/json"
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

// UpdateMechanism is the mechanism of an event
type UpdateMechanism string

// An Event represents a computational event.
type Event struct {
	Entity    json.RawMessage `json:"entity"`
	Identity  string          `json:"identity"`
	Type      EventType       `json:"type"`
	Timestamp time.Time       `json:"timestamp"`
}

// NewEvent returns a new Event.
func NewEvent(t EventType, o Identifiable) *Event {

	data, err := json.Marshal(o)
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

// An Events represents a list of Event.
type Events []*Event

// NewEvents retutns a new Events.
func NewEvents(events ...*Event) Events {

	return append(Events{}, events...)
}
