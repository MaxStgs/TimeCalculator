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
	UserID int    `json:"UserID,omitempty"`
	Action string `json:"Action,omitempty"`
	Value  string `json:"Value,omitempty"`
}

type User struct {
	UserID int
	Conn   *websocket.Conn
}

var Users []User

func RunSocketServer() {
	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		if err != nil {
			fmt.Println("WebSocket connection error: " + err.Error())
			return
		}
		User := User{len(Users), conn}
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
		index, err := strconv.Atoi(message.Value)
		index -= 1
		if err != nil {
			fmt.Println("Can't StartTimer : " + err.Error())
			return
		}
		handleTimerStart(index)
	case "StopTimer":
		index, err := strconv.Atoi(message.Value)
		index -= 1
		if err != nil {
			fmt.Println("Can't StopTimer : " + err.Error())
			return
		}
		handleTimerStop(index)
	}
}

func sendMessage(message Message, users []User) {
	for k, v := range users {
		message.UserID = 0
		// TODO: Client not get UserID and think that it is not his package
		data, err := json.Marshal(message)
		if err != nil {
			fmt.Println("WebSocket can't marshal sendMessage " + message.Action + " with value: " + message.Value + " to UserID(" + strconv.Itoa(v.UserID) + ")")
			continue
		}
		if err = v.Conn.WriteMessage(1, data); err != nil {
			Users = append(Users[:k], Users[k+1:]...)
			fmt.Println("Removed socket from Users with index:" + strconv.Itoa(v.UserID) + " reason:" + "socket close")
			continue
		}
	}
}
