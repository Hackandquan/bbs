package api

import "github.com/kataras/iris/v12"

type TestController struct {
	Ctx iris.Context
}

type RespTestJson struct {
	Title string
}

func (c *TestController) Get() []RespTestJson {
	return []RespTestJson{
		{ "123" },
	}
}