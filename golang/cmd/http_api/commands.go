package main

import (
	"quacker"
)

type QuackMessage struct {
	AuthorID quacker.UserID
	Content  string
}
