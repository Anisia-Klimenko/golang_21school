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
	//resp, err := client.Get("https://localhost:3333/buy_candy")
	bodyPost := bytes.NewReader(clientJson)
	resp, err := client.Post("https://localhost:3333/buy_candy", "application/json; charset=UTF-8", bodyPost)
	must(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	must(err)

	fmt.Printf("Status: %s  Body: %s\n", resp.Status, string(body))
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

	// c, _ := tls.LoadX509KeyPair("signed-cert", "key")

	config := &tls.Config{
		// Certificates: []tls.Certificate{c},
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

// fmt.Println("Certificate authority:")
// must(utils.OutputPEMFile("../ca/cert"))
// cp, _ := x509.SystemCertPool() or
// cp := x509.NewCertPool()
// data, _ := ioutil.ReadFile("../ca/cert")
// cp.AppendCertsFromPEM(data)

// fmt.Println("My certificate:")
// must(utils.OutputPEMFile("signed-cert"))
// c, _ := tls.LoadX509KeyPair("signed-cert", "key")

// InsecureSkipVerify: true,
// RootCAs:               cp,
// Certificates:          []tls.Certificate{c},
