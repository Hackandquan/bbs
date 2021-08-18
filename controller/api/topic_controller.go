package api

import (
	"bbs/model"
	"fmt"
	"github.com/kataras/iris/v12"
	"time"
)


type TopicController struct {
	Ctx iris.Context
}



type TopicListResp struct {
	model.TopicListCondition
	Id int
}

// 根据发表时间、是否推荐、类型查询 分页
func (c *TopicController) GetList() model.TopicListCondition {
	body := model.TopicListCondition{}
	err := c.Ctx.ReadJSON(&body)
	if err!= nil {
		println(err.Error())
	}
	fmt.Println(time.Now().Format("2006-01-02"))
	return body
}
