package router

import (
	"../html"
)

func htmlInit() {
	App.Get("/", html.Index)
}