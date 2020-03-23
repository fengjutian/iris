package main

import (
	"github.com/kataras/iris/v12"
)

type Company struct {
	Name  string
	City  string
	Other string
}

func MyHandler(ctx iris.Context) {
	var c Company

	if err := ctx.ReadJSON(&c); 
	err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	ctx.Writef("Received: %#+v\n", c)
}

type Person struct {
	Name string `json."name"`
	Age int `json:"age"`
}

func MyHandler2(ctx iris.Context) {
	var persons []Person
	err := ctx.ReadJSON(&persons)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	ctx.Writef("Received: %#+v\n", persons)
}

func main() {
	app := iris.New()

	// use Postman or whatever to do a POST request
	// to the http://localhost:8080 with RAW BODY:
	/*
		{
			"Name": "iris-Go",
			"City": "New York",
			"Other": "Something here"
		}
	*/

	app.Post("/", MyHandler)
	app.Post("/slice", MyHandler2)

	app.Run(iris.Addr(":8081"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
}
