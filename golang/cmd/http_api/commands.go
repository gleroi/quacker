package main

import (
	"quacker"
)

type QuackMessage struct {
	AuthorID quacker.UserID
	Content  string
}

type FollowUser struct {
	Follower quacker.UserID
	Followee quacker.UserID
}
