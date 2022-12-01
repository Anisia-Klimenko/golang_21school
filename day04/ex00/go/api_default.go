package swagger

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func BuyCandy(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var order Order
	err := decoder.Decode(&order)
	if err != nil {
		panic(err)
	}
	fmt.Println("here")
	fmt.Fprintf(w, "%s", order.toString())
	log.Printf("%v", order)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(order.toString()))
	if err != nil {
		return
	}

}
