package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	getFlags()
	clientJson, err := json.Marshal(order)
	if err != nil {
		fmt.Printf("Bad arguments")
		os.Exit(1)
	}
	client := getClient()
	bodyPost := bytes.NewReader(clientJson)
	resp, err := client.Post("https://localhost:3333/buy_candy", "application/json; charset=UTF-8", bodyPost)
	must(err)

	defer resp.Body.Close()
	must(err)

	if resp.StatusCode == 201 {
		var success InlineResponse201
		json.NewDecoder(resp.Body).Decode(&success)
		fmt.Println(resp.StatusCode, success.Change, success.Thanks)
	} else if resp.StatusCode == 400 {
		var fail InlineResponse400
		json.NewDecoder(resp.Body).Decode(&fail)
		fmt.Println(resp.StatusCode, fail.Error_)
	} else if resp.StatusCode == 402 {
		var fail InlineResponse402
		json.NewDecoder(resp.Body).Decode(&fail)
		fmt.Println(resp.StatusCode, fail.Error_)
	}
}

var order Order

func getFlags() {
	flag.StringVar(&order.CandyType, "k", "", "type of candy")
	flag.IntVar(&order.CandyCount, "c", 0, "type of candy")
	flag.IntVar(&order.Money, "m", 0, "type of candy")
	flag.Parse()
}

func getClient() *http.Client {
	cp := x509.NewCertPool()
	data, _ := ioutil.ReadFile("../minica.pem")
	cp.AppendCertsFromPEM(data)

	config := &tls.Config{
		RootCAs:               cp,
		GetClientCertificate:  ClientCertReqFunc("cert.pem", "key.pem"),
		VerifyPeerCertificate: CertificateChains,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}
	return client
}

func must(err error) {
	if err != nil {
		fmt.Printf("Client error: %v\n", err)
		os.Exit(1)
	}
}
