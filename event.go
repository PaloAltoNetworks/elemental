// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import "time"

// EventsList represents a list of *Event.
type EventsList []*Event

// EventHandler is prototype of a Push Center Handler.
type EventHandler func(*Event)

// EventType is the type of an event
type EventType string

const (
	// EventCreate is the type of creation event.
	EventCreate = "create"

	// EventUpdate is the type of update event.
	EventUpdate = "update"

	// EventDelete is the type of delete event.
	EventDelete = "delete"
)

// UpdateMechanism is the mechanism of an event
type UpdateMechanism string

// Event represents one item of a Notification.
type Event struct {
	Entity    interface{} `json:"entity"`
	Identity  string      `json:"identity"`
	Type      EventType   `json:"type"`
	Timestamp time.Time   `json:"timestamp"`
}

// NewEvent returns a new *Notification.
func NewEvent(t EventType, o Identifiable) *Event {

	return &Event{
		Type:      t,
		Entity:    o,
		Identity:  o.Identity().Name,
		Timestamp: time.Now(),
	}
}
