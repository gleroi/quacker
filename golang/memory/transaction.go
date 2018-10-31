package memory

import (
	"quacker/event"
)

// Transaction is a in-memory transaction for EventStore
type Transaction struct {
	events []event.Event
	store  event.EventStore
}

func NewTransaction(store event.EventStore) *Transaction {
	return &Transaction{
		events: make([]event.Event, 0),
		store:  store,
	}
}

func (p *Transaction) Append(evt event.Event) {
	p.events = append(p.events, evt)
}

func (p *Transaction) Commit() {
	for _, evt := range p.events {
		p.store.Append(evt)
	}
}
