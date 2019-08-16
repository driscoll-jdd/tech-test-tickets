package Structures

import (
	"errors"
	"strings"
	"sync"
	"time"
)

	// Get a new instance of the ticket issuer with existing tickets loaded from a data storage medium
	func CreateIssuer(tickets []Ticket) Issuer {

		issuer := Issuer{ tickets: tickets }

		// Kick start a thread for the releaser
		go issuer.thread_releaser()

		// Send this back
		return issuer
	}

	type Issuer struct {

		lock sync.RWMutex
		tickets []Ticket
		unpaidTickets []int
	}

	func (i *Issuer) Reserve(name string) (Ticket, error) {

		i.lock.Lock()
		defer i.lock.Unlock()

		// Preformat the name
		name = strings.ToLower(strings.TrimSpace(name))

		// Does this person already have a ticket?
		// Using a map allows _, exists := mapname[value] but maps are expensive, so maybe a slice is better
		for _, ticket := range i.tickets {

			if(ticket.Holder == name) {

				// If this person already has a ticket, send back their actual ticket so the remaining time for them to pay to confirm it can be shown
				return ticket, errors.New("This person already holds a ticket")
			}
		}

		// Also our ticket supply is limited to five tickets
		if(len(i.tickets) > 4) {

			// We are over limit
			return Ticket{}, errors.New("No more tickets available")
		}

		// OK so, issue this ticket
		myTicket := Ticket{ Serial: len(i.tickets), Holder: name, Paid: false, ReleasePoint: time.Now().Add(time.Minute * 5) }
		i.tickets = append(i.tickets, myTicket)

		// They haven't paid; put them on the list
		i.unpaidTickets = append(i.unpaidTickets, myTicket.Serial)

		// Give this to the customer
		return myTicket, nil
	}

	// Purchase a ticket, if the ticket hasn't expired already
	func (i *Issuer) Purchase(ticket Ticket) bool {

		// First, gain exclusivity over the ticket stock
		i.lock.Lock()
		defer i.lock.Unlock()

		// Does this ticket still exist in our catalogue (hasn't timed out)?
		stillValid := false
		var unpaidID int
		var ticketID int
		for unpaidID, ticketID = range i.unpaidTickets {

			// At this point I'm thinking a map and _, exists := myMapVariable[key] would be easier but there is an additional slowness cost with maps. Which is better?
			if(ticketID == ticket.Serial) {

				stillValid = true
			}
		}

		if(!stillValid) {

			// Shame, too bad
			return false
		}

		// This is OK, purchase the ticket
		i.tickets[ticket.Serial].Paid = true
		i.deleteUnpaidRecord(unpaidID)
		return true
	}

	// Release tickets back into availability if someone does not complete the purchase
	func (i *Issuer) thread_releaser() {

		for {

			// Very briefly grab exclusivity
			i.lock.RLock()

				unpaid := i.unpaidTickets

			i.lock.RUnlock()

			// Go through these releasing them
			for unpaidID, ticket := range unpaid {

				if(time.Now().After(i.tickets[ticket].ReleasePoint)) {

					i.lock.Lock()

						// Dump this ticket
						i.deleteTicket(ticket)

						// Dump the 'unpaid' flag
						i.deleteUnpaidRecord(unpaidID)

					i.lock.Unlock()
				}
			}

			// Keep polling and releasing tickets as soon as possible. Demand is high.
			time.Sleep(time.Second * time.Duration(1))
		}
	}

	// delete a ticket entry from our slice - unsafe function - should only be called from within a Lock() / Unlock()
	func (i *Issuer) deleteTicket(ticketID int) {

		if(len(i.tickets) > ticketID) {

			i.tickets = append(i.tickets[:ticketID], i.tickets[ticketID + 1:]...)

		} else {

			i.tickets = append(i.tickets[:ticketID])
		}
	}

	// Delete a record of a ticket being unpaid - unsafe function - should only be called from within a Lock() / Unlock()
	func (i *Issuer) deleteUnpaidRecord(unpaidID int) {

		if(len(i.unpaidTickets) > unpaidID) {

			i.unpaidTickets = append(i.unpaidTickets[:unpaidID], i.unpaidTickets[unpaidID + 1:]...)

		} else {

			i.unpaidTickets = append(i.unpaidTickets[:unpaidID])
		}
	}

	type Ticket struct {

		Serial int
		Holder string
		Paid bool
		ReleasePoint time.Time
	}
