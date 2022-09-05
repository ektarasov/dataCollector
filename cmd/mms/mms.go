package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type MMSData struct {
	Country string `json:"country"`

	Provider string `json:"provider"`

	Bandwidth string `json:"bandwidth"`

	ResponseTime string `json:"response_time"`
}

func main() {

	resp, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(resp.Body)

	var mmsMap []MMSData

	err = json.Unmarshal(content, &mmsMap)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(mmsMap); i++ {
		fmt.Println(mmsMap[i])
	}

	fmt.Println(resp.Status)

}
