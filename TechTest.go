package main

import (
	"TechTest2/Structures"
	"TechTest2/Templates"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var issuer Structures.Issuer

	func main() {

		// Create a ticket issuer instance with a fresh book of tickets (could be loaded in from database, redis, other source of information)
		issuer = Structures.CreateIssuer(make([]Structures.Ticket, 0))

		// Set up API routing
		router := httprouter.New()

		// The main interface
		router.GET("/", Splash)

		// It would be good to make this an ajax endpoint and access this via ajax with json output
		router.GET("/reserve", Reserve)
		router.GET("/reserve/:name", Reserve)

		router.GET("/charge", Charge)

		// Bring up the server on 8080
		http.ListenAndServe(":8080",router)
	}

	func Splash(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		w.Write([]byte(Templates.Page_Splash()))
	}

	func Charge(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		// Payment for tickets here using the issuer.purchase() method and the ticket struct from issuer.Reserve()
	}

	func Reserve(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		name := p.ByName("name")
		response := Structures.AjaxReservationResponse{}

		if(len(name) < 1) {

			// Indicate the failure
			response.Faults.Name = true

			// Send this back
			responseBytes, _ := json.Marshal(response)
			w.Write(responseBytes)
			return
		}

		// Try to reserve a ticket
		ticket, err := issuer.Reserve(name)
		response.Ticket = ticket

		if(err != nil) {

			// Some kind of problem
			response.Error = err.Error()

			// Send this back
			responseBytes, _ := json.Marshal(response)
			w.Write(responseBytes)
			return
		}

		// Great, we have our ticket
		response.Outcome = true
		response.Ticket = ticket
		responseBytes, _ := json.Marshal(response)
		w.Write(responseBytes)
	}
