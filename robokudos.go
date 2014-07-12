package main

import (
	"encoding/json"
	"net/http"
)

func main() {

}

func init() {
	http.HandleFunc("/api", answerSlack)
}

func answerSlack(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	answer := "You said: " + r.FormValue("text")

	slackResponse := map[string]string{"text": answer}
	enc.Encode(slackResponse)
}
