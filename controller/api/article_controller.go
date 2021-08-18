package api

import (
	"github.com/kataras/iris/v12"
	"time"
)

type ArticleController struct {
	Ctx iris.Context
}

// 查询文章列表
// 根据发表时间、是否推荐、类型查询
type ArticleListCondition struct {
	CreateTime time.Time

}
func (receiver *ArticleController) GetList()  {

}
