package main

import (
	"crypto/tls"
	"fmt"
)

func getCert(certfile, keyfile string) (c tls.Certificate, err error) {
	if certfile != "" && keyfile != "" {
		c, err = tls.LoadX509KeyPair(certfile, keyfile)
		if err != nil {
			fmt.Printf("Error loading key pair: %v\n", err)
		}
	} else {
		err = fmt.Errorf("I have no certificate")
	}
	return
}

// ClientCertReqFunc returns a function for tlsConfig.GetClientCertificate
func ClientCertReqFunc(certfile, keyfile string) func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
	c, err := getCert(certfile, keyfile)

	return func(certReq *tls.CertificateRequestInfo) (*tls.Certificate, error) {
		fmt.Println("Received certificate request: sending certificate")
		if err != nil || certfile == "" {
			fmt.Println("I have no certificate")
		} else {
			err := OutputPEMFile(certfile)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		}
		Wait()
		return &c, nil
	}
}
