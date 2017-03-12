package html

import (
	"gopkg.in/kataras/iris.v6"
)

func Index(ctx *iris.Context) {
	ctx.Render(
		// the file name of the template relative to the './templates'.
		"index.html",
		iris.Map{"Name": "Iris"},
		// the .Name inside the ./templates/hi.html,
		// you can use any custom struct that you want to bind to the requested template.
		iris.Map{"gzip": false}, // set to true to enable gzip compression.
	)
}
