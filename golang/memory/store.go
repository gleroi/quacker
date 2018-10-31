package memory

import (
	"quacker/event"

	"github.com/google/uuid"
)

type Store struct {
	events []event.Event
}

func NewStore() *Store {
	return &Store{
		events: make([]event.Event, 0, 1024),
	}
}

func (s *Store) Append(evt event.Event) {
	s.events = append(s.events, evt)
}

func (s *Store) Get(id uuid.UUID) []event.Event {
	agg := make([]event.Event, 0, 16)
	for _, evt := range s.events {
		if evt.AggregateID() == id {
			agg = append(agg, evt)
		}
	}
	return agg
}

func (s *Store) All() []event.Event {
	return s.events //TODO: replace with pub/sub ?
}
