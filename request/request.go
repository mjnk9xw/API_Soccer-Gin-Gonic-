package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// OpenFile doc file co chua chuoi json
func OpenFile(filename string) {
	counts := make(map[string]int)
	data, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(string(data), "\n") {
		counts[line]++
	}
	if filename == "tour.txt" {
		for line, n := range counts {
			if n >= 1 {
				Req("http://localhost:8080/api/v1/Tour/", line)
			}
		}
	}
	if filename == "round.txt" {
		for line, n := range counts {
			if n >= 1 {
				Req("http://localhost:8080/api/v1/Tour/1/Round/", line)
			}
		}
	}
	if filename == "table.txt" {
		for line, n := range counts {
			if n >= 1 {
				Req("http://localhost:8080/api/v1/Tour/1/Round/1/Table", line)
			}
		}
	}
	if filename == "player.txt" {
		for line, n := range counts {
			if n >= 1 {
				Req("http://localhost:8080/api/v1/Tour/1/Round/1/Table/1/Player", line)
			}
		}
	}
	if filename == "match.txt" {
		for line, n := range counts {
			if n >= 1 {
				Req("http://localhost:8080/api/v1/Tour/1/Round/1/Table/1/Match", line)
			}
		}
	}
}

// Req Thuc hien Post cac chuoi JSON
func Req(url string, js string) {
	var jsonStr = []byte(js)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
}
func main() {
	OpenFile("tour.txt")
	OpenFile("round.txt")
	OpenFile("table.txt")
	OpenFile("player.txt")
	OpenFile("match.txt")
}
