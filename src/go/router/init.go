package router

import (
	"gopkg.in/kataras/iris.v6"
)

var App *iris.Framework

func Init(app *iris.Framework) {
	App = app
	apiInit()
	htmlInit()
}