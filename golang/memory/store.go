package memory

import (
	"quacker"

	"github.com/google/uuid"
)

type Store struct {
	events []quacker.Event
}

func NewStore() *Store {
	return &Store{
		events: make([]quacker.Event, 0, 1024),
	}
}

func (s *Store) Append(evt quacker.Event) {
	s.events = append(s.events, evt)
}

func (s *Store) Get(id uuid.UUID) []quacker.Event {
	agg := make([]quacker.Event, 0, 16)
	for _, evt := range s.events {
		if evt.AggregateID() == id {
			agg = append(agg, evt)
		}
	}
	return agg
}

func (s *Store) All() []quacker.Event {
	return s.events //TODO: replace with pub/sub ?
}
