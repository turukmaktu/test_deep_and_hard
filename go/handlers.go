package main

import (
	"encoding/json"
	"net/http"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	label := r.FormValue("label")

	status := resultApi{ Status: true, Message: "Success send"}

	checkInput, jsonResultCheck := checkInput(id, label)
	if !checkInput {
		status = jsonResultCheck
	}else{
		go sendToRedis( eventApi{ Id:id,Label: label } )
	}

	js, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
