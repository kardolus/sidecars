package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.SetFlags(log.Lshortfile)

	file := filepath.Join("/home/vcap/app/sidecar-sample.sock")
	message := "Hello,"

	conn, err := net.Dial("unix", file)
	if err != nil {
		log.Printf("error: %v\n", err)
		return
	}

	defer conn.Close()
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Printf("error: %v\n", err)
		return
	}

	log.Printf("[main-app] send to Sidecar: %s\n", message)

	err = conn.(*net.UnixConn).CloseWrite()
	if err != nil {
		log.Printf("error: %v\n", err)
		return
	}

	for {
		buf := make([]byte, 11)
		nr, err := conn.Read(buf)
		if err != nil {
			break
		}

		buf = buf[:nr]
		log.Printf("[main-app] receive from Sidecar: %s\n", string(buf))
		fmt.Fprintf(w, string(buf))
	}
}

func main() {
	http.HandleFunc("/", handler)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)
}
