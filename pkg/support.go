package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func SupportCollect() []SupportData {
	var supportTemp []SupportData

	resp, err := http.Get(*dstServerAddress + "/support")
	if err != nil {
		fmt.Println("Не удалось получить данные")
		return supportTemp
	}
	if resp.Status == "200 OK" {

		content, err := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal(content, &supportTemp)
		if err != nil {
			log.Fatal(err)
			return supportTemp
		}
	}

	return supportTemp
}
