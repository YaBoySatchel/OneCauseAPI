package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type UserLogin struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	OneTimeCode int    `json:"oneTimeCode"`
}

func validateUserCredentials(userlogin *UserLogin) bool {
	// If we wanted we could check separately and return a
	// more determinate response so the end user knew if an
	// email existed or not
	return userlogin.Email == "c137@onecause.com" && userlogin.Password == "#th@nH@rm#y#r!$100%D0p#"
}

func validateUserCode(userlogin *UserLogin, currentUTCTime time.Time) bool {
	hourString := fmt.Sprint(currentUTCTime.Hour())
	minuteString := fmt.Sprint(currentUTCTime.Minute())
	concatenatedTime := hourString + minuteString
	timeCode, err := strconv.Atoi(concatenatedTime)

	return err == nil && userlogin.OneTimeCode == timeCode
}

func login(w http.ResponseWriter, r *http.Request) {
	// Wouldn't normally open up the endpoint to the world
	// but for this local integration it works
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var userLogin UserLogin
	_ = json.NewDecoder(r.Body).Decode(&userLogin)

	credentialResult := validateUserCredentials(&userLogin)
	currentUTCTime := time.Now().UTC()
	userCodeResult := validateUserCode(&userLogin, currentUTCTime)

	if credentialResult && userCodeResult {
		// log successful login attempt by user email and send response
		w.WriteHeader(http.StatusOK)
	} else {
		// log unsuccessful login attempt by user email and send response
		w.WriteHeader(http.StatusBadRequest)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "App Healthy")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", healthCheck).Methods("GET")
	router.HandleFunc("/login", login).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	handleRequests()
}
