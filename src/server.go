package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"os"
	"io/ioutil"
	"fmt"
)

// global variables 

// will be used to send messages through websocket
var message_channel = make(chan User)

var clients = make(map[*websocket.Conn]bool)
var upgrader = websocket.Upgrader {
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}


func determinePort() (string, error){
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set") 
	} 

	total_port := ":" + port
	return total_port, nil
}


// start server functions
func main() {
	port, err := determinePort()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(port)

	createJson()

	file_server_remote := http.FileServer(http.Dir("public/"))
	http.Handle("/", file_server_remote)

	http.HandleFunc("/ws", handleConnections)
	http.HandleFunc("/getJson", getJson)

	go sendMessages()

	//err := http.ListenAndServe(":8080", nil)
	httpErr := http.ListenAndServe(port, nil)
	if httpErr != nil {
		log.Fatal("Error cannot host server: ListenAndServe err")
	}

	log.Println("Successful server start")
}

// receives messages and connections from clients
func handleConnections(writer http.ResponseWriter, request *http.Request) {
	websocket, err:= upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Printf("Error handling client connection: %v", err)
	}

	defer websocket.Close()
	clients[websocket] = true

	// everytime the server receives a message from a client
	// it'll append to the messages.json file
	for {
		var message User
		err := websocket.ReadJSON(&message)
		if err != nil {
			log.Printf("Error receiving message from client: %v", err)
			delete(clients, websocket)
			break
		}
		log.Printf("message: %v", message)

		message_channel <- message

		file, err := os.Open("resources/messages.json")
		if err != nil {
			log.Printf("Error cannot open/read messages.json: %v", err)
		}

		byte_value, _ := ioutil.ReadAll(file)
		var users []User
		json.Unmarshal(byte_value, &users)
		users = append(users, message)

		result, err := json.MarshalIndent(users, "", "\t")
		if err != nil {
			log.Fatal(err)
		}

		writeJson(result)
	}
}


func sendMessages() {
	for {
		message := <- message_channel
		for client := range clients {
			// data to return to client if failed
			err := client.WriteJSON(message)

			if err != nil {
				log.Printf("error-- %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}