package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func BuyCandy(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var order Order
	err := decoder.Decode(&order)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		fmt.Fprintln(w, "here")
		var response InlineResponse400
		response.Error_ = "wrong fields or types"
		convertJson, err := json.MarshalIndent(response, "", "    ")
		if err != nil {
			fmt.Println("Error MarshalIndent:", err)
		}
		fmt.Fprintf(w, string(convertJson))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	candy, err := findCandyByName(order.CandyType)
	if err != nil {
		var response InlineResponse400
		response.Error_ = "candy type " + order.CandyType + " not found"
		convertJson, err := json.MarshalIndent(response, "", "    ")
		if err != nil {
			fmt.Println("Error MarshalIndent:", err)
		}
		fmt.Fprintf(w, string(convertJson))
		w.WriteHeader(http.StatusBadRequest)
	} else if order.CandyCount <= 0 {
		var response InlineResponse400
		response.Error_ = "candy count " + strconv.Itoa(order.CandyCount) + " less than or equal 0"
		convertJson, err := json.MarshalIndent(response, "", "    ")
		if err != nil {
			fmt.Println("Error MarshalIndent:", err)
		}
		fmt.Fprintf(w, string(convertJson))
		w.WriteHeader(http.StatusBadRequest)
	} else if order.Money <= 0 {
		var response InlineResponse400
		response.Error_ = "money " + strconv.Itoa(order.Money) + " less than or equal 0"
		convertJson, err := json.MarshalIndent(response, "", "    ")
		if err != nil {
			fmt.Println("Error MarshalIndent:", err)
		}
		fmt.Fprintf(w, string(convertJson))
		w.WriteHeader(http.StatusBadRequest)
	} else {
		change := order.Money - candy.Price*order.CandyCount
		if change >= 0 {
			var response InlineResponse201

			response.Change = change
			response.Thanks = "Thank you!"
			convertJson, err := json.MarshalIndent(response, "", "    ")
			if err != nil {
				fmt.Println("Error MarshalIndent:", err)
			}
			fmt.Fprintf(w, string(convertJson))
			w.WriteHeader(http.StatusCreated)

		} else {
			var response InlineResponse402
			response.Error_ = "You need " + strconv.Itoa(-change) + " more money!"
			fmt.Println(response.Error_)
			convertJson, err := json.MarshalIndent(response, "", "    ")
			if err != nil {
				fmt.Println("Error MarshalIndent:", err)
			}
			fmt.Fprintf(w, string(convertJson))
			w.WriteHeader(http.StatusPaymentRequired)
		}
	}

}
