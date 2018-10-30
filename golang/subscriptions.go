package quacker

import "github.com/google/uuid"

// SubscriptionID is a unique key identifying a subscription
type SubscriptionID uuid.UUID

// Subscription is the relationship between a follower and followee.
type Subscription struct {
	id       SubscriptionID
	followee UserID
	follower UserID
}

type UserFollowed struct {
	ID       SubscriptionID
	Followee UserID
	Follower UserID
}

func Follow(tr Transaction, follower UserID, followee UserID) {
}
