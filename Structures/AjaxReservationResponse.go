package Structures

	type AjaxReservationResponse struct {

		Outcome bool
		Ticket Ticket
		Faults Faults
		Error string
	}

	type Faults struct {

		Name bool
	}
