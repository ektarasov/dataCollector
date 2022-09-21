package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

func IncidentCollect() []IncidentData {
	var incidentTemp []IncidentData
	resp, err := http.Get("http://127.0.0.1:8383/accendent")
	if err != nil {
		fmt.Println("Не удалось получить данные")
		return incidentTemp
	}
	if resp.Status == "200 OK" {

		content, err := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal(content, &incidentTemp)
		if err != nil {
			log.Fatal(err)
			return incidentTemp
		}
	}

	return incidentTemp
}
