package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var DefaultHost string = "localhost"

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

type proxy struct {
}

func (p *proxy) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost && req.RequestURI == "/MOCK" {
		buf, err := ioutil.ReadAll(req.Body)
		log.Println("Mock definition request received!")
		if err == nil {
			prepareMockFromBuffer(buf)
			http.Error(wr, "Mock definition request accepted", http.StatusAccepted)
		} else {
			log.Println(err, string(buf))
			http.Error(wr, "Mock definition request failed", http.StatusBadRequest)
		}
		return
	}
	req.URL.Host = req.Header.Get("x-host")
	if req.URL.Host == "" {
		req.URL.Host = DefaultHost
	}
	req.URL.Scheme = "http"
	log.Printf("Requesting: %s %#s\n", req.Method, req.URL)

	client := &http.Client{}
	req.RequestURI = ""

	resp, err := client.Do(req)
	if err != nil {
		http.Error(wr, "Server Error", http.StatusInternalServerError)
		wr.Write([]byte(err.Error() + "\n"))
		return
	}
	defer resp.Body.Close()

	log.Println(req.RemoteAddr, " ", resp.Status)

	copyHeader(wr.Header(), resp.Header)
	wr.WriteHeader(resp.StatusCode)
	io.Copy(wr, resp.Body)
}
