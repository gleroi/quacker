package quacker

import (
	"quacker/stub"
	"testing"
)

var follower = UserID("follower@mail.com")
var followee = UserID("followee@mail.com")

func TestWhenFollowUserThenUserFollowed(t *testing.T) {
	tr := stub.NewTransaction()

	Follow(tr, follower, followee)

	if len(tr.Events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(tr.Events))
	}
	if evt, ok := tr.Events[0].(UserFollowed); !ok {
		t.Errorf("expected %T, got %T", UserFollowed{}, tr.Events[0])
	} else {
		if evt.Follower != follower {
			t.Errorf("expected %s, got %s", follower, evt.Follower)
		}
		if evt.Followee != followee {
			t.Errorf("expected %s, got %s", followee, evt.Followee)
		}
	}
}
