package quacker

import (
	"fmt"

	"github.com/google/uuid"
)

// MessageID is a unique key identifying a Message.
type MessageID uuid.UUID

func NewMessageID() MessageID {
	return MessageID(uuid.New())
}

func (id MessageID) String() string {
	return uuid.UUID(id).String()
}

// UserID identify a user by its email address.
type UserID string

// Message quacked or requacked by a user.
type Message struct {
	id       MessageID
	author   UserID
	quackers map[UserID]bool
}

func (m *Message) Apply(evt Event) error {
	switch mevt := evt.(type) {
	case MessageQuacked:
		m.id = mevt.ID
		m.author = mevt.Author
		m.quackers = make(map[UserID]bool)
		m.quackers[mevt.Author] = true
	case MessageRequacked:
		m.quackers[mevt.Requacker] = true
	default:
		return fmt.Errorf("unexpected event of type %t", evt)
	}
	return nil
}

type MessageQuacked struct {
	ID      MessageID
	Author  UserID
	Content string
}

func (m MessageQuacked) AggregateID() uuid.UUID {
	return uuid.UUID(m.ID)
}

func Quack(tr Transaction, author UserID, content string) {
	tr.Append(MessageQuacked{
		ID:      NewMessageID(),
		Author:  author,
		Content: content,
	})
}

type MessageRequacked struct {
	ID        MessageID
	Requacker UserID
}

func (m MessageRequacked) AggregateID() uuid.UUID {
	return uuid.UUID(m.ID)
}

func (m Message) Requack(tr Transaction, requacker UserID) {
	if m.quackers[requacker] == true {
		return
	}

	tr.Append(MessageRequacked{
		ID:        m.id,
		Requacker: requacker,
	})
}
