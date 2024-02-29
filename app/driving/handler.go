package driving

import (
	"encoding/json"
	"github.com/dulguundd/logError-lib/logger"
	"log"
	"net/http"
	"restAPIServer/app/dto"
	"restAPIServer/app/service"
	"time"
)

type Handlers struct {
	service service.Service
}

func home(w http.ResponseWriter, r *http.Request) {
	ip, _ := getIP(r)
	headers := r.Header
	headersAsString := HeadersToString(headers)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	log.Println("handler routing /home get request" + ip + ", " + headersAsString)
	responseBody := "Hello world! /home IP Address: " + ip
	if err := json.NewEncoder(w).Encode(responseBody); err != nil {
		panic(err)
	}
}

func getBody(w http.ResponseWriter, r *http.Request) {
	var req dto.NewUserRequest
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
			var res dto.ValidationError
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

func (h *Handlers) MongoTest(w http.ResponseWriter, _ *http.Request) {
	startTime := time.Now()
	_ = h.service.MongoList()
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	log.Println("Hello world! /mongotest duration: ", executionTime)
	responseBody := "Hello world! /mongotest duration: "
	if err := json.NewEncoder(w).Encode(responseBody); err != nil {
		panic(err)
	}
}

func (h *Handlers) MongoTestById(w http.ResponseWriter, _ *http.Request) {
	startTime := time.Now()
	result, err := h.service.MongoById()
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(500)
		if err := json.NewEncoder(w).Encode(err.Message); err != nil {
			panic(err)
		}
	} else {
		logger.Info("This is driving log: " + result.Id)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		//responseBody := "Hello world! /mongotest/id duration: "
		if err := json.NewEncoder(w).Encode(result); err != nil {
			panic(err)
		}
	}
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	log.Println("Hello world! /mongotest/id duration: ", executionTime)
}
