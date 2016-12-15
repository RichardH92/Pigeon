package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/urfave/negroni"
	"github.com/zbindenren/negroni-prometheus"
	"log"
	"net/http"
	"os"
)

type Configuration struct {
	Port         string
	Emails       []Email_Type
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

	SetEmailQueries(configuration.Emails)
	SetSourceAddr(configuration.Source_Email)
	SetRegion(configuration.Region)

	r := mux.NewRouter()

	r.Handle("/metrics", prometheus.Handler())
	r.HandleFunc("/email", HandleGetSendEmail).Methods("GET")

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())
	n.Use(negroniprometheus.NewMiddleware("Pigeon"))
	n.UseHandler(r)

	log.Fatal(http.ListenAndServe(configuration.Port, n))
}
