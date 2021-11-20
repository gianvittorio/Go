package main

import (
	_ "bytes"
	"encoding/json"
	_ "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	_"time"
)

func main() {
	// url := "https://enc709vwym8s0vs.m.pipedream.net"
	// resp, err := http.Get(url)
	// inspectResponse(resp, err)
	// type s struct {
	// 	X int
	// 	Y float32
	// }
	// v := s{4, 3.8}

	// data, err := json.Marshal(v)
	// if err != nil {
	// 	log.Fatal("Error occurred while marshaling json ", err)
	// }

	// resp, err = http.Post(url, "application/json", bytes.NewReader(data))
	// inspectResponse(resp, err)

	// client := http.Client{
	// 	Timeout: 3 * time.Second,
	// }
	// client.Get(url)

	// req, err := http.NewRequest(http.MethodPut, url, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// req.Header.Add("x-testheader", "learning go header")
	// req.Header.Set("User-Agent", "Go learning HTTP/1.1")
	// resp, err := client.Do(req)
	// inspectResponse(resp, err)

	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	v := struct {
		IP string `json: "ip"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(v.IP)
}

func inspectResponse(resp *http.Response, err error) {
	if err != nil {
		log.Fatal("Error occurred while marshaling json ", err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error occurred while trying to read response body ", err)
	}
	log.Println(string(b))
}