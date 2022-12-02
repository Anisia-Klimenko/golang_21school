package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Candy struct {
	Name  string
	Price int
}

func initCandy() []Candy {
	var menu []Candy
	menu = append(menu, Candy{"CE", 10})
	menu = append(menu, Candy{"AA", 15})
	menu = append(menu, Candy{"NT", 17})
	menu = append(menu, Candy{"DE", 21})
	menu = append(menu, Candy{"YR", 23})
	return menu
}

var menu = initCandy()
var ErrorNotFound = errors.New("candy not found")

func findCandyByName(name string) (Candy, error) {
	for _, candy := range menu {
		if candy.Name == name {
			return candy, nil
		}
	}
	return Candy{}, ErrorNotFound
}

func main() {
	log.Printf("Server started")
	server := getServer()
	http.HandleFunc("/buy_candy", BuyCandy)
	must(server.ListenAndServeTLS("", ""))
}

//func myHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("Handling request")
//	w.Write([]byte("Hey GopherCon!"))
//}

func getServer() *http.Server {
	cp := x509.NewCertPool()
	data, _ := ioutil.ReadFile("../minica.pem")
	cp.AppendCertsFromPEM(data)

	// c, _ := tls.LoadX509KeyPair("cert.pem", "key.pem")

	tls := &tls.Config{
		// Certificates:          []tls.Certificate{c},
		ClientCAs:             cp,
		ClientAuth:            tls.RequireAndVerifyClientCert,
		GetCertificate:        CertReqFunc("cert.pem", "key.pem"),
		VerifyPeerCertificate: CertificateChains,
	}

	server := &http.Server{
		Addr:      ":3333",
		TLSConfig: tls,
	}
	return server
}

func must(err error) {
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
		os.Exit(1)
	}
}

// cert := "cert"
// fmt.Println("My certificate:")
// must(utils.OutputPEMFile(cert))
// c, _ = tls.LoadX509KeyPair(cert, "key")

// fmt.Println("Certificate authority:")
// must(utils.OutputPEMFile("../ca/cert"))
// cp, _ := x509.SystemCertPool()
// data, _ := ioutil.ReadFile("../ca/cert")
// cp.AppendCertsFromPEM(data)

// NoClientCert ClientAuthType = iota
// RequestClientCert
// RequireAnyClientCert
// VerifyClientCertIfGiven
// RequireAndVerifyClientCert

// RootCAs:               cp,
// ClientCAs:             cp,
// VerifyPeerCertificate: utils.CertificateChains,
// GetCertificate:        getCert,
// GetClientCertificate:  getClientCert,
