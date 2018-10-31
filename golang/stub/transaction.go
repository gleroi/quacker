package stub

import (
	"quacker/event"
)

// Transaction is a in-memory event store transactions
type Transaction struct {
	Events []event.Event
}

func NewTransaction() *Transaction {
	return &Transaction{
		Events: make([]event.Event, 0),
	}
}

func (p *Transaction) Append(evt event.Event) {
	p.Events = append(p.Events, evt)
}

func (p *Transaction) Commit() {

}
