package main

/*
#include <stdlib.h>
#include "cow.h"
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"unsafe"
)

func BuyCandy(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var order Order
	err := decoder.Decode(&order)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var response InlineResponse400
		response.Error_ = "wrong fields or types"
		convertJson, err := json.MarshalIndent(response, "", "    ")
		if err != nil {
			fmt.Println("Error MarshalIndent:", err)
		}
		fmt.Fprintf(w, string(convertJson))
		return
	}

	candy, err := findCandyByName(order.CandyType)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var response InlineResponse400
		response.Error_ = "candy type " + order.CandyType + " not found"
		convertJson, err := json.MarshalIndent(response, "", "    ")
		if err != nil {
			fmt.Println("Error MarshalIndent:", err)
		}
		fmt.Fprintf(w, string(convertJson))
	} else if order.CandyCount <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		var response InlineResponse400
		response.Error_ = "candy count " + strconv.Itoa(order.CandyCount) + " less than or equal 0"
		convertJson, err := json.MarshalIndent(response, "", "    ")
		if err != nil {
			fmt.Println("Error MarshalIndent:", err)
		}
		fmt.Fprintf(w, string(convertJson))
	} else if order.Money <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		var response InlineResponse400
		response.Error_ = "money " + strconv.Itoa(order.Money) + " less than or equal 0"
		convertJson, err := json.MarshalIndent(response, "", "    ")
		if err != nil {
			fmt.Println("Error MarshalIndent:", err)
		}
		fmt.Fprintf(w, string(convertJson))
	} else {
		change := order.Money - candy.Price*order.CandyCount
		if change >= 0 {
			w.WriteHeader(http.StatusCreated)
			var response InlineResponse201

			cs := C.CString("Thank you!")
			defer C.free(unsafe.Pointer(cs))
			fooRes := C.ask_cow(cs)
			result := C.GoString(fooRes)

			response.Change = change
			response.Thanks = result
			//success := InlineResponse201{Change: change, Thanks: result}
			//json.NewEncoder(w).Encode(success)
			convertJson, err := json.MarshalIndent(response, "", "    ")
			if err != nil {
				fmt.Println("Error MarshalIndent:", err)
			}
			fmt.Fprintf(w, string(convertJson))

		} else {
			w.WriteHeader(http.StatusPaymentRequired)
			var response InlineResponse402
			response.Error_ = "You need " + strconv.Itoa(-change) + " more money!"
			fmt.Println(response.Error_)
			convertJson, err := json.MarshalIndent(response, "", "    ")
			if err != nil {
				fmt.Println("Error MarshalIndent:", err)
			}
			fmt.Fprintf(w, string(convertJson))
		}
	}

}
