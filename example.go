package main

import (
	"github.com/kataras/iris/v12"
)

// func main() {
// 	app := iris.New()
// 	app.Logger().SetLevel("debug")

// 	app.Use(recover.New())
// 	app.Use(logger.New())

// 	app.Handle("GET", "/", func(ctx iris.Context) {
// 		ctx.HTML("<h1>Welcome</h1>")
// 	})

// 	app.Get("/ping", func(ctx iris.Context) {
// 		ctx.WriteString("pong")
// 	})

// 	app.Get("/hello", func(ctx iris.Context) {
// 		ctx.JSON(iris.Map{"message": "Hello Iris!"})
// 	})

// 	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
// }

func main() {
	// Creates an application with default middleware:
	// logger and recovery (crash-free) middleware.
	app := iris.Default()

	// app.Post("/user/{name:string}/{action:path}", func(ctx iris.Context) {
	// 	name := ctx.Params().Get("name")
	// 	action := ctx.Params().Get("action")
	// 	message := name + " is " + action
	// 	ctx.WriteString(message)
	// })

	app.Get("/someGet", getting)
	app.Post("/user/{name:string}/{action:path}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		action := ctx.Params().Get("action")
		message := name + " is " + action
		ctx.WriteString(message)
	})

	// app.Put("/somePut", putting)
	// app.Delete("/someDelete", deleting)
	// app.Patch("/somePatch", patching)
	// app.Head("/someHead", head)
	// app.Options("/someOptions", options)

	app.Run(iris.Addr(":8080"))
}

func getting(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "Hello Iris!"})
}
