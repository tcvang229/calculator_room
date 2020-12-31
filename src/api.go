package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"io/ioutil"
)


// will be used as a struct to send and receive messages from clients
type User struct {
	Username string `json:"username"`
	Message string `json:"message"`
}

// creates messages.json if not already created
func createJson() {
	file, err := os.OpenFile("resources/messages.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Println("Error cannot open or create messages.json. Throwing Fatal()")
		log.Fatal(err)
	}

	if err := file.Close(); err != nil {
		log.Println("Error cannot close messages.json. Throwing Fatal()")
		log.Fatal(err)
	}
}

func getJson(writer http.ResponseWriter, request *http.Request) {
	file, err := os.Open("resources/messages.json")
	if err != nil {
		log.Printf("Error cannot read messages.json: %v", err)
	}

	byte_value, _ := ioutil.ReadAll(file)
	var users []User

	// unmarshal into users []object
	json.Unmarshal(byte_value, &users)

	// returns json format
	json.NewEncoder(writer).Encode(users)
}

// messages.json structure is a slice of User -> User[]
func writeJson(json_data []byte) {
	file, err := os.OpenFile("resources/messages.json", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write(json_data)
	if err != nil {
		log.Fatal(err)
	}

	file.Close()
}