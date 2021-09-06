package api

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

type GithubLoginController struct {
	Ctx iris.Context
}

type callbackQuery struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

// 接收 github 登录的 code
func (c *GithubLoginController) GetCallback() {
	query := callbackQuery{}
	c.Ctx.ReadQuery(&query)

	fmt.Println(query.Code, query.State)

	// 查询是否已经绑定过本站用户，到 t_third_account 表中查找；如果已经有了就登录，没有就让用户注册绑定
	
}
