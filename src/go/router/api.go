package router

import (
	"../unit"
	"../controllers"
)

func apiInit() {
	App.Get(unit.CreateApi("/blogs"), controllers.GetBlogList)
}