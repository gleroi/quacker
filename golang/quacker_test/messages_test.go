package quacker_test

import (
	"quacker"
	"quacker/quacker_test/stub"
	"testing"

	"github.com/google/uuid"
)

var author quacker.UserID = "author1@quack.com"
var requacker quacker.UserID = "requacker@quack.com"
var defaultContent = "Some content"

func TestWhenQuackOneMessage(t *testing.T) {
	var tr *stub.Transaction = stub.NewTransaction()

	quacker.Quack(tr, author, defaultContent)

	if len(tr.Events) != 1 {
		t.Errorf("expected one event")
	}
	if evt, ok := tr.Events[0].(quacker.MessageQuacked); !ok {
		t.Errorf("event must be a MessageQuacked event")
	} else {
		if evt.Content != defaultContent {
			t.Errorf("expected %s, got %s", defaultContent, evt.Content)
		}
		if evt.Author != author {
			t.Errorf("expected %s, got %s", author, evt.Author)
		}
	}
}

func TestWhenQuackTwoMessagesThenIDsAreDifferents(t *testing.T) {
	var tr *stub.Transaction = stub.NewTransaction()

	quacker.Quack(tr, author, defaultContent)
	quacker.Quack(tr, author, defaultContent)

	if len(tr.Events) != 2 {
		t.Errorf("expected 2 events, got %d", len(tr.Events))
	}
	var evt1, evt2 quacker.MessageQuacked
	var ok bool
	if evt1, ok = tr.Events[0].(quacker.MessageQuacked); !ok {
		t.Errorf("expected a MessageQuacked event, got %T", tr.Events[0])
	}
	if evt2, ok = tr.Events[1].(quacker.MessageQuacked); !ok {
		t.Errorf("expected a MessageQuacked event, got %T", tr.Events[1])
	}
	if evt1.ID == evt2.ID {
		t.Errorf("events ids should be differents")
	}
}

func TestGivenAQuackedMessage(t *testing.T) {
	m := &quacker.Message{}
	mID := quacker.NewMessageID()

	m.Apply(quacker.MessageQuacked{
		ID:      mID,
		Author:  author,
		Content: defaultContent,
	})

	t.Run("WhenRequackBySomeElseThenMessageRequacked", func(t *testing.T) {
		var tr *stub.Transaction = stub.NewTransaction()
		m.Requack(tr, requacker)
		if len(tr.Events) != 1 {
			t.Errorf("expected one event only")
		}
		if requack, ok := tr.Events[0].(quacker.MessageRequacked); !ok {
			t.Errorf("expected event MessageRequacked, got %T", tr.Events[0])
		} else {
			if uuid.UUID(mID) != requack.AggregateID() {
				t.Errorf("expected ID %s, got %s", mID, requack.AggregateID())
			}
			if requack.Requacker != requacker {
				t.Errorf("expected requacker %s, got %s", requacker, requack.Requacker)
			}
		}
	})

	t.Run("WhenRequackTwoTimesBySomeElseThenNoMessageRequacked", func(t *testing.T) {
		var tr *stub.Transaction = stub.NewTransaction()
		var mRequacked = *m
		mRequacked.Apply(quacker.MessageRequacked{
			ID:        mID,
			Requacker: requacker,
		})

		m.Requack(tr, requacker)

		if len(tr.Events) != 0 {
			t.Errorf("expected no event")
		}
	})

	t.Run("WhenRequackByMySelfThenMessageNotRequacked", func(t *testing.T) {
		var tr *stub.Transaction = stub.NewTransaction()
		m.Requack(tr, author)
		if len(tr.Events) != 0 {
			t.Errorf("expected no event")
		}
	})
}
