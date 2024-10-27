package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Item struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Price float64 `json:"price"`
}

var items []Item = []Item{{ID: 1,Title:"item1",Price: 10}, {2,"item2",5}, {3,"item3",7}}

const grID = "abcdefghijklmnopqrstuvwxyz"
var idx = 0

func main() {
    // fiber.config is used to configure application
    
    appConfig := fiber.Config{
        // prefork is an ability to spin up multiple processes essentially go routines
        // to listen the same port
        Prefork: false,
        // set up application name
        AppName:"My awesome app beta v1.0",
        // show list of routes and their handler in console
        EnablePrintRoutes: true,
        // set the server header name and it will show to the response header
        ServerHeader: "My awesome app",
        // prevent the fiber context from reusable(mutable), so it will create a copied value every time
        Immutable: true,
        // [NOT PRACTICAL] case sensitive for url path
        CaseSensitive: true,
    }
	app := fiber.New(appConfig)

    app.Get("/item", getAllItems)
    app.Get("/item/:id", addReqId,requestLogger, getItem)
    app.Post("/item", createtItem)
    // app.Get("/:name", getName)

    // multiple paths with the same name except for different cases
    app.Get("/books",getAllBooksLC)
    app.Get("/Books",getAllBooksUC)

    // app.Get("/books/:id",getBook)
    // make params to be optional
    app.Get("/author/:id?",getAuthor)
    // * is a modifier which is to be additional information to determine which set of items can be returned
    // don't use * unless you really want to accept just about anything from user
    app.Get("/book/*",getBook)

    // start server
    err := app.Listen(":8000")
    if err!= nil {
        log.Fatal(err)
    }
}