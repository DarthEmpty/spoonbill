package main

import (
	"encoding/json/v2"
	"net/http"
)

type Ticket struct {
	ID   int
	Inc  bool
	Desc string
}

// To be replaced with db connection
var tickets = [5]Ticket{
	{1, true, "PC not working, URGENT"},
	{ID: 2, Desc: "Asking for Bitlocker recovery key"},
	{3, true, "Laptop needs reimaging"},
	{4, false, "New User starting tomorrow - accounts needed urgently"},
	{5, true, "Toilet Clogged"},
}

func getAllTickets(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(tickets)

	if err != nil {
		panic("Can't retrieve all tickets")
	}

	w.Write(bytes)
}
