package main

import (
	"net/http"
	"time"
	"log"
	"flag"
)

// There are some golang RESTful libraries and mux libraries but i use the simplest to test.
func main() {


	addr := flag.String("addr", ":8080", "bind address, default :8080")
	sleep := flag.Int("sleep", 0, "sleep milliseconds, default 0")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if *sleep > 0 {
			time.Sleep(time.Duration(*sleep) * time.Millisecond)
		}
		w.Write([]byte("Hello world"))
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
