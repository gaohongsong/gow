package main

import (
	"time"
	"log"
	"flag"
	"net/http"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

// There are some golang RESTful libraries and mux libraries but i use the simplest to test.
func main() {

	addr := flag.String("addr", ":8080", "bind address, default :8080")
	sleep := flag.Int("sleep", 0, "sleep milliseconds, default 0")
	flag.Parse()

	// close debug info
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.New()
	//r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		if *sleep > 0 {
			time.Sleep(time.Duration(*sleep) * time.Millisecond)
		}
		c.String(http.StatusOK, "Hello world")
	})

	err := r.Run(*addr)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
