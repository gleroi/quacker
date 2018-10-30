package quacker

// FolloweeSet is a read model of whom is followed
type FolloweeSet struct {
	followees map[UserID]bool
}

func (set FolloweeSet) Contains(user UserID) bool {
	active, ok := set.followees[user]
	if !ok {
		return false
	}
	return active
}

func apply(followees map[SubscriptionID]UserID, follower UserID, evt Event) {
	switch e := evt.(type) {
	case UserFollowed:
		if e.Follower != follower {
			return
		}
		followees[e.ID] = e.Followee
	case UserUnfollowed:
		delete(followees, e.ID)
	}

}

func GetFolloweeList(store EventStore, follower UserID) FolloweeSet {
	l := FolloweeSet{}
	followees := make(map[SubscriptionID]UserID)
	for _, evt := range store.All() {
		apply(followees, follower, evt)
	}
	l.followees = make(map[UserID]bool, len(followees))
	for _, followee := range followees {
		l.followees[followee] = true
	}
	return l
}
