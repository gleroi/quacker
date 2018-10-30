package quacker

import (
	"fmt"

	"github.com/google/uuid"
)

// SubscriptionID is a unique key identifying a subscription
type SubscriptionID uuid.UUID

func NewSubscriptionID() SubscriptionID {
	return SubscriptionID(uuid.New())
}

// Subscription is the relationship between a follower and followee.
type Subscription struct {
	id       SubscriptionID
	followee UserID
	follower UserID
}

func (s *Subscription) Apply(evt Event) error {
	switch e := evt.(type) {
	case UserFollowed:
		s.id = e.ID
		s.followee = e.Followee
		s.follower = e.Follower
	default:
		return fmt.Errorf("unexpected event of type %T", evt)
	}
	return nil
}

type UserFollowed struct {
	ID       SubscriptionID
	Followee UserID
	Follower UserID
}

func (e UserFollowed) AggregateID() uuid.UUID {
	return uuid.UUID(e.ID)
}

func Follow(tr Transaction, follower UserID, followee UserID) {
	tr.Append(UserFollowed{
		ID:       NewSubscriptionID(),
		Follower: follower,
		Followee: followee,
	})
}

type UserUnfollowed struct {
	ID SubscriptionID
}

func (e UserUnfollowed) AggregateID() uuid.UUID {
	return uuid.UUID(e.ID)
}

func (s Subscription) Unfollow(tr Transaction) {
	tr.Append(UserUnfollowed{
		ID: s.id,
	})
}
