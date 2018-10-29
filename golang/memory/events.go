package memory

import (
	"quacker"
)

// Publisher is a in-memory event store and publisher
type Publisher struct {
	Events []quacker.Event
}

func NewPublisher() *Publisher {
	return &Publisher{
		Events: make([]quacker.Event, 0),
	}
}

func (p *Publisher) Publish(evt quacker.Event) {
	p.Events = append(p.Events, evt)
}
