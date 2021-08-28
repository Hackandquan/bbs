package render

import "bbs/model"

func BuildNode(node *model.TopicNode) *model.NodeResponse {
	if node == nil {
		return nil
	}
	return &model.NodeResponse{
		Id:          node.Id,
		Name:        node.Name,
		Logo:        node.Logo,
		Description: node.Description,
	}
}
