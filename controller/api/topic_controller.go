package api

import (
	"bbs/controller/render"
	"bbs/model"
	"bbs/service"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
)

type TopicController struct {
	Ctx iris.Context
}

type TopicListResp struct {
	model.TopicListCondition
	Id int
}

// 增
// 发表topic
func (c *TopicController) PostCreate() *simple.JsonResult {
	// 判断是否登录
	// 判断用户status：是否已经被删
	// 判断用户是否在禁言状态：isForbidden
	// 操作
}

// 删
// 删除topics
func (c *TopicController) DeleteDelete() *simple.JsonResult {}

// 查
// fmt.Println(time.Now().Format("2006-01-02"))
// 根据发表时间createTime、是否推荐bool、类型查询nodeId 分页
// 如果某个字段类型为int64，在json 中是字符串不会自动转为数值，值为0
func (c *TopicController) GetList() *simple.JsonResult {
	body := model.TopicListCondition{}
	err := c.Ctx.ReadJSON(&body)
	if err != nil {
		println(err.Error())
		return simple.JsonErrorMsg(err.Error())
	}
	topics, cursor, hasMore := service.TopicService.GetTopics(body.NodeId, body.Cursor, body.Recommend)
	// 获取当前登录用户
	var user *model.User = nil
	return simple.JsonCursorData(render.BuldSimpleTopics(topics, user), strconv.FormatInt(cursor, 10), hasMore)
}

// 改
// 修改topic
func (c *TopicController) PutUpdate() *simple.JsonResult {}
