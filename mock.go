package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/h2non/gock.v1"
)

type MockType struct {
	Request struct {
		Method        string
		URL           string
		Host          string
		Type          string
		Path          string
		Persist       bool
		Times         int
		JSON          interface{}
		XML           interface{}
		HeaderPresent []string
		ParamPresent  []string
		MatchHeaders  map[string]string
		MatchParams   map[string]string
		PathParams    map[string]string
	} `json:"Request"`
	Response struct {
		Status  int
		Body    string
		Headers map[string]string
		Type    string
		Delay   string
		JSON    interface{}
		XML     interface{}
	}
}
type MockTypeList []MockType

func prepareMockFromFile(fname string) {
	body, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	prepareMockFromBuffer(body)
}

func prepareMockFromBuffer(body []byte) {
	var f MockTypeList
	err := json.Unmarshal(body, &f)
	if err != nil {
		panic(err)
	}
	gock.Flush()
	for k, v := range f {
		req := v.Request
		rsp := v.Response
		if k == 0 {
			DefaultHost = req.Host
			log.Println("Setting default host as " + DefaultHost)
		}
		log.Printf("[%d] Setting %s %s\n", k, req.Method, req.Host+req.URL)
		g := gock.New(v.Request.Host)
		switch req.Method {
		case "GET":
			g.Get(req.URL)
			break
		case "POST":
			g.Post(req.URL)
			break
		case "HEAD":
			g.Head(req.URL)
			break
		case "DELETE":
			g.Delete(req.URL)
			break
		}
		if req.Type != "" {
			g.MatchType(req.Type)
		}
		if req.Path != "" {
			g.Path(req.Path)
		}
		if req.Persist {
			g.Persist()
		}
		if req.Times != 0 {
			g.Times(req.Times)
		}
		if &req.JSON == nil {
			g.JSON(req.JSON)
		}
		if &req.XML == nil {
			g.XML(req.XML)
		}
		for _, b := range req.HeaderPresent {
			g.HeaderPresent(b)
		}
		for _, b := range req.ParamPresent {
			g.ParamPresent(b)
		}
		for a, b := range req.MatchHeaders {
			g.MatchHeader(a, b)
		}
		for a, b := range req.MatchParams {
			g.MatchParam(a, b)
		}
		for a, b := range req.PathParams {
			g.PathParam(a, b)
		}
		r := g.Reply(rsp.Status)
		if rsp.Body != "" {
			r.BodyString(rsp.Body)
		}
		if rsp.Type != "" {
			r.Type(rsp.Type)
		}
		if &rsp.JSON == nil {
			r.JSON(rsp.JSON)
		}
		if &rsp.XML == nil {
			r.XML(rsp.XML)
		}
		if rsp.Delay != "" {
			dur, err := time.ParseDuration(rsp.Delay)
			if err == nil {
				r.Delay(dur)
			} else {
				log.Println("Error:", err)
			}
		}
		for a, b := range rsp.Headers {
			r.AddHeader(a, b)
		}
	}
}
