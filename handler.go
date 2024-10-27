package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// handler function do?
// 1. doing the work required
//  * querying a database
//  * getting data from file
// 2. returning the result
//  * returning result or an error

func addReqId(context *fiber.Ctx) error {
    uId := "xnxx1234"
    context.Request().Header.Add("request-id", uId)
    return context.Next()
}

func requestLogger(context *fiber.Ctx) error {
    reqId := context.Request().Header.Peek("request-id")
	slog.Info("got request", "method", context.Method(), "path", context.Path(), "param", context.Params("id"), "reqid", reqId)
	// if we return that means request is completed and fiber will send response to the client
	// return nil
    // return next to invoke the next handler function
    return context.Next()
}

// handler function takes a pointer of fiber.Ctx (fiber context)
// fiber context is designed for receiving request at the moment
// fiber will reuse context for better performance
// don't store fiber context in global variables but only use fiber context in the handler function
// don't store reference value in fiber context
// if we want to store any reference to fiber context only use the fiber context in the handler function
func getAllItems(context *fiber.Ctx) error {
    return context.JSON(items)
}

func getItem(context *fiber.Ctx) error {
    id, err := context.ParamsInt("id")
    if err!= nil {
        // sending an error with message
        return fiber.NewError(http.StatusBadRequest, "id is invalid")
    }
    for _, item := range items {
        if item.ID == id {
            // sending json as a response
            return context.JSON(item)
        }
    }
    return fiber.NewError(http.StatusNotFound, "can not find the item")
}

func createtItem(context *fiber.Ctx) error {
    var item Item
    // bind the request body to struct
    err := context.BodyParser(&item); 
    if err!= nil {
        return context.Status(fiber.StatusBadRequest).SendString("Bad request")
    }
    item.ID = len(items)+1
    items = append(items, item)
    return context.JSON(item)
}

func getHandlerId() string {
    c := grID[idx%26]
    idx++
    return fmt.Sprintf("grID-%v-%c", idx, c)
}

func getName(context *fiber.Ctx) error {
    ccId := getHandlerId()
    n := context.Params("name")
    // using goroutines to simulate that fiber context will be reused and must not use outside the handler function
    // before go runtime reach to return, fiber context is mutable because fiber context will be reused
    go func() {
        slog.Info("starting handler","ccId",ccId,"name",n)
        t := time.After(10*time.Second)
        for {
            select {
            case <-t:
                slog.Info("handler done","ccId",ccId,"name",n)
                return
            default:
                slog.Info("handler still running","ccId",ccId,"name",n)
                time.Sleep(1*time.Second)
            }
        }
    
    }()
    slog.Info("request received","name",n)
    return nil
}

func getAllBooksLC(context *fiber.Ctx) error {
    slog.Info("get all books - lowercase")
    return nil
}

func getAllBooksUC(context *fiber.Ctx) error {
    slog.Info("get all books - uppercase")
    return nil
}

// func getBook(context *fiber.Ctx) error {
//     bookId := context.Params("id")
//     slog.Info("get book by id", "bookId", bookId)
//     return nil
// }

func getAuthor(context *fiber.Ctx) error {
    authorId := context.Params("id")
    if authorId == "" {
        slog.Info("get all author")
        return nil
    }
    slog.Info("get author by id", "authorId", authorId)
    return nil
}

func getBook(context *fiber.Ctx) error {
    itemPath := context.Params("*")
    if itemPath == "" {
        slog.Info("request all items")
        return nil
    }
    slog.Info("request item with sub-path")
    // get all sub-paths from the item path
    subPaths := strings.Split(itemPath, "/")
    for _, subPath := range subPaths {
        slog.Info("sub-path", "subPath", subPath)
    }
    return nil
}