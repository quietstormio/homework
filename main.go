package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Make custom client and check Redirects
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
			// return nil
			// The return nil will allow the redirect
			// and I make request with Location value below
		},
	}

	resp, err := client.Get("https://localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// How to get redirect location
	// This doesn't return the error
	dione, _ := resp.Location()
	fmt.Println(dione)
}

/*********************************

How to add cert to TLS Client

func main() {
	certFile := flag.String("cert", "client.crt", "Client certificate file")
	keyFile := flag.String("key", "client.key", "Client private key file")
	caFile := flag.String("ca", "ca.crt", "CA certificate file")
	flag.Parse()

	// Load client certificate and key
	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Load CA certificate
	caCert, err := ioutil.ReadFile(*caFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Configure TLS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	// Create HTTP client with TLS configuration
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// Make a request
	resp, err := client.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Handle the response
	// ...
}

****************/
