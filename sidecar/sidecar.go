package main

import (
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {
	log.SetFlags(log.Lshortfile)
	file := filepath.Join("/home/vcap/app/sidecar-sample.sock")
	os.Remove(file)
	listener, err := net.Listen("unix", file)
	defer listener.Close()

	if err != nil {
		log.Printf("error: %v\n", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}

		go func() {
			defer conn.Close()
			send_data := make([]byte, 0)
			for {
				buf := make([]byte, 10)

				recv_data, err := conn.Read(buf)
				if err != nil {
					if err != io.EOF {
						log.Printf("error: %v", err)
					}
					break
				}

				buf = buf[:recv_data]
				log.Printf("[sidecar] receive from Main-App: %s\n", string(buf))
				send_data = append(buf, []byte(" Super_Secret_Credentials")...)
			}
			conn.Write(send_data)
			log.Printf("[sidecar] send to Main-App: %s\n", string(send_data))
		}()
	}
}
