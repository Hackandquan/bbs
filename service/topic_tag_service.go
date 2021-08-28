package service

import (
	"bbs/model"
	"fmt"

	"github.com/mlogclub/simple"
)

var TopicTagService = newTopicTagService()

func newTopicTagService() *topicTagService {
	return &topicTagService{}
}

type topicTagService struct{}

func (s *topicTagService) GetTagsByTopicId(topicId int64, status int64) *[]model.TopicTag {
	sql := `SELECT * FROM t_topic_tag where topic_id=? AND status=1` // status=0为11删除
	ret := &[]model.TopicTag{}
	if err := simple.DB().Exec(sql, topicId, status).Find(&ret).Error; err != nil {
		fmt.Println("service.GteTagsByTopicId Error: ", err)
	}
	return ret
}
