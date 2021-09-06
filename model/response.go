package model

import "bbs/model/constants"

type TopicResponse struct {
	TopicId         int64               `json:"topicId"`
	Type            constants.TopicType `json:"type"`
	User            *UserInfo           `json:"user"` // 关联topic的用户信息 1:1
	Node            *NodeResponse       `json:"node"` // 关联topic节点信息 1:1
	Tags            *[]TagResponse      `json:"tags"` // 关联topic的标签信息 1:n
	Title           string              `json:"title"`
	Summary         string              `json:"summary"`
	Content         string              `json:"content"`
	ImageList       []ImageInfo         `json:"imageList"` // 关联？
	LastCommentTime int64               `json:"lastCommentTime"`
	ViewCount       int64               `json:"viewCount"`
	CommentCount    int64               `json:"commentCount"`
	LikeCount       int64               `json:"likeCount"`
	Liked           bool                `json:"liked"` // 标记当前用户是否点了喜欢
	CreateTime      int64               `json:"createTime"`
	Recommend       bool                `json:"recommend"`
	RecommendTime   int64               `json:"recommendTime"`
}

type NodeResponse struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	SortNo      int64  `json:"sortNo"`
	CreateTime  int64  `json:"createTime"`
}

type ImageInfo struct{}

type UserInfo struct {
	Id                   int64    `json:"id"`
	Username             string   `json:"username"`
	Email                string   `json:"email"`
	EmailVerified        bool     `json:"emailVerified"`
	Nickname             string   `json:"nickname"`
	Avatar               string   `json:"avatar"`
	SmallAvatar          string   `json:"smallAvatar"`
	BackgroundImage      string   `json:"backgroundImage"`
	SmallBackgroundImage string   `json:"smallBackgroundImage"`
	Type                 int      `json:"type"`
	Roles                []string `json:"roles"`
	HomePage             string   `json:"homePage"`
	Description          string   `json:"description"`
	Score                int      `json:"score"`        // 积分
	TopicCount           int      `json:"topicCount"`   // 话题数量
	CommentCount         int      `json:"commentCount"` // 跟帖数量
	PasswordSet          bool     `json:"passwordSet"`  // 密码已设置
	Forbidden            bool     `json:"forbidden"`    // 是否禁言
	Status               int      `json:"status"`
	CreateTime           int64    `json:"createTime"`
}

// select *, If(forbidden_end_time<timestamp(now()),0,1) isForbidden from t_user where type='1' And status='1'

type TagResponse struct {
	TagId   int64  `json:"tagId"`
	TagName string `json:"tagName"`
}
