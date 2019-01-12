package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	UserID int    `json:"UserId,omitempty"`
	Action string `json:"Action,omitempty"`
	Value  string `json:"Value,omitempty"`
}

type User struct {
	UserID int
}

var Users []User

func RunSocketServer() {
	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		if err != nil {
			fmt.Println("WebSocket connection error: " + err.Error())
			return
		}
		User := User{len(Users)}
		userId := strconv.Itoa(User.UserID)
		if userId == "" {
			fmt.Println("Got problem with convert userId")
			return
		}
		message := Message{-1, "InitializeClient", userId}
		data, err := json.Marshal(message)
		if err != nil {
			fmt.Println("WebSocket can't marshal action " + message.Action + " with value: " + message.Value)
			return
		}
		err = conn.WriteMessage(1, data)
		Users = append(Users, User)

		for {
			// Read message from browser
			_, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			data := Message{}
			fmt.Println(string(msg))
			err = json.Unmarshal(msg, &data)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			handleMessage(data)

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			/*// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}*/
		}
	})
}

func handleMessage(message Message) {
	switch message.Action {
	case "StartTimer":
	case "StopTimer":
	}
}
