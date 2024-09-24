package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/adwait-upadhyaya/cli-chat-app/internal/database"
	socketio "github.com/googollee/go-socket.io"
)

func InitServer() {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {

		s.SetContext("")
		fmt.Println("connected:", s.ID())
		s.Join("chat_room")
		return nil
	})

	server.OnEvent("/", "message", func(s socketio.Conn, msg string) {
		fmt.Printf("Message recieved from %s:  %s \n", s.ID(), msg)
		intId, _ := strconv.Atoi(s.ID())
		database.LogMessage(msg, intId)
		server.BroadcastToRoom("/", "chat_room", "message", msg)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Printf("User %s disconnected: %s\n", s.ID(), reason)
		s.Leave("chat_room")
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
