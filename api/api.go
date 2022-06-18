package api

import (
	"booking-service/models"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const cinemaServer = "http://cinema-service.com"

// Fetch attempts to retrieve all details for a given show, including available seats
func Fetch(w http.ResponseWriter, r *http.Request) {
	var show models.Show

	if err := json.NewDecoder(r.Body).Decode(&show); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Could not read req body")
		return
	}
	defer r.Body.Close()

	log.Println("show ID", show.ID)

	resp, err := http.Get(fmt.Sprintf("%s/%d/%d", cinemaServer, show.CinemaID, show.ID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Could not find", show.ID, "at", show.CinemaID)
		return
	}

	if err := json.NewDecoder(resp.Body).Decode(&show); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Could not read resp body")
		return
	}
	defer r.Body.Close()

	gob.NewEncoder(w).Encode(show)
}

// Book sends a booking request to the cinema-server for particular user selected seats
func Book(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	gob.NewDecoder(r.Body).Decode(&booking)

	bookingB, err := json.Marshal(booking)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Invalid booking request")
		return
	}

	resp, err := http.Post(fmt.Sprintf("%s/%d/%d", cinemaServer, booking.CinemaID, booking.ShowID),
		"application/json",
		bytes.NewReader(bookingB),
	)
	if err != nil || resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Could not get succesful booking")
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&booking)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Invalid booking request")
		return
	}

	log.Println("A booking has been made successfuly, ID", booking.ID)
	err = gob.NewEncoder(w).Encode(booking)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Invalid booking request")
	}
}
