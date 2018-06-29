package main

import (
	"net/http"
	"meander"
	"encoding/json"
)


func main() {
	//meander.APIKey = "TODO"
	http.HandleFunc("/journeys", func(w http.ResponseWriter,
	r *http.Request) {
		respond(w, r, meander.Journeys)
	})
	http.ListenAndServe(":8080", http.DefaultServeMux)
}


func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	publicData := make([]interface{}, len(data))
	for i, d := range data {
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}
