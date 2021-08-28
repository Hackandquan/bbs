package model

import (
	"bbs/utils"
)

func (u *User) IsForbidden() bool {
	return u.ForbiddenEndTime >= utils.TimestampNow()
}
