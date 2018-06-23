package main

import (
	"time"
	"github.com/kataras/iris"
	"flag"
)

func main() {

	addr := flag.String("addr", ":8080", "bind address, default :8080")
	sleep := flag.Int("sleep", 0, "sleep milliseconds, default 0")
	flag.Parse()

	api := iris.New()
	api.Get("/", func(c iris.Context) {
		if *sleep > 0 {
			time.Sleep(time.Duration(*sleep) * time.Millisecond)
		}
		c.Writef("Hello world")
	})
	api.Run(iris.Addr(*addr))
}
