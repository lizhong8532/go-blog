package main

import (
	"./router"
	"./config"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
)

func initApp () {
	app := iris.New(iris.Configuration{
		Gzip: false,
		Charset: "UTF-8",
	})

	//app.Adapt(iris.DevLogger())
	app.Adapt(view.HTML("./templates", ".html"))
	app.Adapt(httprouter.New())
	router.Init(app)
	app.Listen(config.HOST + ":" + config.PORT)
}

func main() {

	// Init App
	initApp()
}