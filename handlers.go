package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func DrmShowsIndexAction(w http.ResponseWriter, r *http.Request) {
	var errResp ErrorResponse
	errResp.ErrorMesg = "Could not decode request: JSON parsing failed"

	var clireq ClientRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Println("Error reading response :  ")
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Println("Error in request body :  ")
		return
	}
	if err := json.Unmarshal(body, &clireq); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(400) // Bad request
		if err := json.NewEncoder(w).Encode(errResp); err != nil {
			log.Println("Error in sending ErrResp :  ")
		}
		return
	}

	sResp := FindDrmShows(clireq.Payload)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(sResp); err != nil {
		log.Println("Error encoding response JSON :  ", sResp)
	}

	return
}
