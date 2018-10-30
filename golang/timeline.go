package quacker

// Timeline is a read model of the messages received by a user
type Timeline struct {
	UserID   UserID
	Messages []TimelineMessage
}

func (tl *Timeline) add(msg TimelineMessage) {
	tl.Messages = append(tl.Messages, msg)
}

// TimelineMessage is a message in the timeline
type TimelineMessage struct {
	MessageID string
	AuthorID  UserID
	Content   string
}

func (tl *Timeline) apply(evt Event) {
	switch e := evt.(type) {
	case MessageQuacked:
		if e.Author != tl.UserID {
			return
		}
		tl.add(TimelineMessage{
			MessageID: e.ID.String(),
			AuthorID:  e.Author,
			Content:   e.Content,
		})
	}
}

func GetTimeline(store EventStore, userID UserID) Timeline {
	tl := Timeline{
		UserID:   userID,
		Messages: make([]TimelineMessage, 0),
	}
	for _, evt := range store.All() {
		tl.apply(evt)
	}
	return tl
}
