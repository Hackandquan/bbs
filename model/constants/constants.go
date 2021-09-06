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

// 空值
const NilInt64 int64 = -1
const NilInt int = -1
const NilString string = ""
const NilFloat float32 = -1.0
const NilFloat64 float64 = -1.0
