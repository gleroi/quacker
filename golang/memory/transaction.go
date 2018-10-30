package memory

import (
	"quacker"
)

// Transaction is a in-memory transaction for EventStore
type Transaction struct {
	events []quacker.Event
	store  quacker.EventStore
}

func NewTransaction(store quacker.EventStore) *Transaction {
	return &Transaction{
		events: make([]quacker.Event, 0),
		store:  store,
	}
}

func (p *Transaction) Append(evt quacker.Event) {
	p.events = append(p.events, evt)
}

func (p *Transaction) Commit() {
	for _, evt := range p.events {
		p.store.Append(evt)
	}
}
