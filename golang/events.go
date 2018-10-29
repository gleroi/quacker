package quacker

import (
	"github.com/google/uuid"
)

// Event is an event that was decided by its aggregate.
type Event interface {
	AggregateID() uuid.UUID
}

// EventPublisher publish events from aggregate to others.
type EventPublisher interface {
	Publish(evt Event)
}
