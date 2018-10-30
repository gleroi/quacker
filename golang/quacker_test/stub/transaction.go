package stub

import (
	"quacker"
)

// Transaction is a in-memory event store transactions
type Transaction struct {
	Events []quacker.Event
}

func NewTransaction() *Transaction {
	return &Transaction{
		Events: make([]quacker.Event, 0),
	}
}

func (p *Transaction) Append(evt quacker.Event) {
	p.Events = append(p.Events, evt)
}

func (p *Transaction) Commit() {

}
