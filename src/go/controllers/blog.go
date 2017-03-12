package controllers

import (
	"gopkg.in/kataras/iris.v6"
)

func GetBlogList(ctx *iris.Context) {
	ctx.HTML(iris.StatusOK, "<h1>GetBlogList</h1>")
}