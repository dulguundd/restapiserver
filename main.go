package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"strings"
)

type NewUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ValidationError struct {
	ValidationError NewUserValidationError `json:"validationError"`
}

type NewUserValidationError struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/home", home).Methods(http.MethodGet)
	router.HandleFunc("/api/1.0/users", getBody).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}

func home(w http.ResponseWriter, r *http.Request) {
	ip, _ := getIP(r)
	headers := r.Header
	headersAsString := headersToString(headers)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	log.Println("handler routing /home get request" + ip + ", " + headersAsString)
	responseBody := "Hello world! /home IP Address: " + ip
	if err := json.NewEncoder(w).Encode(responseBody); err != nil {
		panic(err)
	}
}

func getBody(w http.ResponseWriter, r *http.Request) {
	var req NewUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Request body parsing error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		incorrectRequestBool, newUserValidationError := newUserRequestValidation(req)
		if incorrectRequestBool != false {
			log.Println("Request data is incorrect")
			log.Println(newUserValidationError.Password)
			log.Println(newUserValidationError.Username)
			log.Println(newUserValidationError.Email)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(200)
			var res ValidationError
			res.ValidationError = *newUserValidationError
			if err := json.NewEncoder(w).Encode(res); err != nil {
				panic(err)
			}
		} else {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(200)
			log.Println("User created")
			if err := json.NewEncoder(w).Encode("user created "); err != nil {
				panic(err)
			}
		}
	}
}

func newUserRequestValidation(req NewUserRequest) (bool, *NewUserValidationError) {
	var incorrectRequestBool bool = false
	var newUserValidationError NewUserValidationError
	if len([]rune(req.Username)) < 6 {
		incorrectRequestBool = true
		newUserValidationError.Username = "username too short"
	}
	if strings.Index(req.Email, "@") == -1 {
		incorrectRequestBool = true
		newUserValidationError.Email = "mail is incorrect"
	}
	if len([]rune(req.Password)) < 6 {
		incorrectRequestBool = true
		newUserValidationError.Password = "password too short"
	}
	if incorrectRequestBool == true {
		return incorrectRequestBool, &newUserValidationError
	} else {
		return incorrectRequestBool, nil
	}

}

func getIP(r *http.Request) (string, error) {
	ips := r.Header.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")

	if len(splitIps) > 0 {
		// get last IP in list since ELB prepends other user defined IPs, meaning the last one is the actual client IP.
		netIP := net.ParseIP(splitIps[len(splitIps)-1])
		if netIP != nil {
			return netIP.String(), nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	netIP := net.ParseIP(ip)
	if netIP != nil {
		ip := netIP.String()
		if ip == "::1" {
			return "127.0.0.1", nil
		}
		return ip, nil
	}

	return "", errors.New("IP not found")
}

func headersToString(headers http.Header) string {
	var headerStrings []string

	for key, values := range headers {
		for _, value := range values {
			headerStrings = append(headerStrings, fmt.Sprintf("%s: %s", key, value))
		}
	}

	return strings.Join(headerStrings, "\n")
}
