package client

import (
	"bufio"
	"fmt"
	"log"
	"os"

	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

func InitClient(username string, userId int) {

	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	opts.Query["user"] = username
	opts.Query["userId"] = fmt.Sprintf("%v", userId)

	uri := "http://localhost:8000/"

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		log.Printf("NewClient error:%v\n", err)
		return
	}

	client.On("error", func() {
		log.Printf("on error\n")
	})
	client.On("connect", func() {
		log.Printf("Connected to chat server\n")
	})
	client.On("message", func(msg string) {
		log.Printf("sent message:%v\n", msg)
	})
	client.On("disconnection", func() {
		log.Printf("Disconnected from server\n")
	})

	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		command := string(data)
		client.Emit("message", command)
		// log.Printf("sent message:%v\n", command)
	}
}
