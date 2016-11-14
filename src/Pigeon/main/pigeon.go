package main

import (
	"pigeon/user_API"
	"pigeon/utilities"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Configuration struct {
	Port         string
	Emails       []user_API.Email_Type
	Source_Email string
	Region       string
}

func main() {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
		fmt.Println("error:", err)
	}

	user_API.SetEmailQueries(configuration.Emails)
	utilities.SetSourceAddr(configuration.Source_Email)
	utilities.SetRegion(configuration.Region)

	r := mux.NewRouter()

	r.HandleFunc("/email", user_API.HandleGetSendEmail).Methods("GET")

	log.Fatal(http.ListenAndServe(configuration.Port, r))
}
