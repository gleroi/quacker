package quacker

import (
	"github.com/google/uuid"
)

// Event is an event that was decided by its aggregate.
type Event interface {
	AggregateID() uuid.UUID
}

// Transaction store events in before commit or rollback.
type Transaction interface {
	Append(evt Event)
	Commit()
}

// EventStore stores events for all aggregates
type EventStore interface {
	Append(evt Event)
	Get(aggregateID uuid.UUID) []Event
	All() []Event
}
