package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const data_json = "https://vl-testovac.s3.amazonaws.com/challenges/childspeak_v2/test.in.json"

func main() {
	// Get data from remote JSON.
	res, err := http.Get(data_json)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	// Read data from JSON.
	var arr []string
	data, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal([]byte(data), &arr)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	} else {
		log.Println("Reading data - [Successfully!]")
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
