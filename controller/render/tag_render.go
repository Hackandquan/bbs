package render

import "bbs/model"

func BuidlTag(tag *model.TopicTag) *model.TagResponse {
	if tag == nil {
		return nil
	}
	return &model.TagResponse{
		TagId:   tag.Id,
		TagName: tag.Name,
	}
}

func BuildTags(tags *[]model.TopicTag) *[]model.TagResponse {
	if len(tags) == 0 {
		return &[]model.TagResponse{}
	}
	responses := &[]model.TagResponse{}
	for _, tag := range tags {
		reponses = append(responses, *BuildTag(&tag))
	}
	return responses
}
