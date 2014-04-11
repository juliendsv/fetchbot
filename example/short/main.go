package main

import (
	"fmt"
	"net/http"

	"github.com/juliendsv/fetchbot"
)

func main() {
	f := fetchbot.New(fetchbot.HandlerFunc(handler))
	queue := f.Start()
	queue.SendStringGet("http://www.google.com")
	queue.Close()
}

func handler(ctx *fetchbot.Context, res *http.Response, err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Printf("[%d] %s %s\n", res.StatusCode, ctx.Cmd.Method(), ctx.Cmd.URL())
}
