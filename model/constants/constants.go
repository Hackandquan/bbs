package constants

type TopicType int

const (
	StatusOk      = 0
	StatusDeleted = 1
	StatusPending = 2
)

const (
	EntityArticle = "article"
	EntityTopic   = "topic"
	EntityComment = "comment"
	EntityUser    = "user"
	EntityCheckIn = "checkIn"
)
