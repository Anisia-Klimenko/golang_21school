package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func BuyCandy(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var convertJson []byte
	var order Order
	err := decoder.Decode(&order)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := InlineResponse400{Error_: "wrong fields or types"}
		convertJson, err = json.MarshalIndent(response, "", "    ")
		if err != nil {
			fmt.Println("Error MarshalIndent:", err)
		}
		fmt.Fprintf(w, string(convertJson))
		return
	}

	candy, err := findCandyByName(order.CandyType)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := InlineResponse400{Error_: "candy type " + order.CandyType + " not found"}
		convertJson, err = json.MarshalIndent(response, "", "    ")
	} else if order.CandyCount <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		response := InlineResponse400{Error_: "candy count " + strconv.Itoa(order.CandyCount) + " less than or equal 0"}
		convertJson, err = json.MarshalIndent(response, "", "    ")
	} else if order.Money <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		response := InlineResponse400{Error_: "money " + strconv.Itoa(order.Money) + " less than or equal 0"}
		convertJson, err = json.MarshalIndent(response, "", "    ")
	} else {
		change := order.Money - candy.Price*order.CandyCount
		if change >= 0 {
			w.WriteHeader(http.StatusCreated)
			response := InlineResponse201{Change: change, Thanks: "Thank you!"}
			convertJson, err = json.MarshalIndent(response, "", "    ")
		} else {
			w.WriteHeader(http.StatusPaymentRequired)
			response := InlineResponse402{Error_: "You need " + strconv.Itoa(-change) + " more money!"}
			convertJson, err = json.MarshalIndent(response, "", "    ")
		}
	}
	if err != nil {
		fmt.Println("Error MarshalIndent:", err)
	}
	fmt.Fprintf(w, string(convertJson))
}
