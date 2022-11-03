package main

import (
	"diplom/pkg"
	"encoding/json"
	"fmt"
)

type ResultSetT struct {
	SMS       [][]pkg.SmsData              `json:"sms"`
	MMS       [][]pkg.MMSData              `json:"mms"`
	VoiceCall []pkg.VoiceCallData          `json:"voice_call"`
	Email     map[string][][]pkg.EmailData `json:"email"`
	Billing   pkg.BillingData              `json:"billing"`
	Support   []int                        `json:"support"`
	Incident  []pkg.IncidentData           `json:"incident"`
}

type ResultT struct {
	Status bool       `json:"status"` // true, если все этапы сбора данных прошли успешно, false во всех остальных случаях
	Data   ResultSetT `json:"data"`   // заполнен, если все этапы сбора данных прошли успешно, nil во всех остальных случаях
	Error  string     `json:"error"`  // пустая строка если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки (детали ниже)

}

func smsres() [][]pkg.SmsData {
	var test []pkg.SmsData
	test = pkg.SmsCollect()

	v, _ := json.Marshal(test)

	var test1 []pkg.SmsData

	json.Unmarshal(v, &test1)

	for i := 0; i < len(test1); i++ {

		switch {
		case test1[i].Country == "RU":
			test1[i].Country = "Russia"

		case test1[i].Country == "US":
			test1[i].Country = "USA"

		case test1[i].Country == "GB":
			test1[i].Country = "United Kingdom"

		case test1[i].Country == "FR":
			test1[i].Country = "France"

		case test1[i].Country == "BL":
			test1[i].Country = "Saint Barthélemy"

		case test1[i].Country == "AT":
			test1[i].Country = "Austria"

		case test1[i].Country == "BG":
			test1[i].Country = "Bulgaria"

		case test1[i].Country == "DK":
			test1[i].Country = "Denmark"

		case test1[i].Country == "CA":
			test1[i].Country = "Canada"

		case test1[i].Country == "ES":
			test1[i].Country = "Spain"

		case test1[i].Country == "CH":
			test1[i].Country = "Switzerland"

		case test1[i].Country == "TR":
			test1[i].Country = "Türkiye"

		case test1[i].Country == "PE":
			test1[i].Country = "Peru"

		case test1[i].Country == "NZ":
			test1[i].Country = "New Zealand"

		case test1[i].Country == "MC":
			test1[i].Country = "Monaco"
		}

	}
	test2 := make([]pkg.SmsData, len(test1))
	copy(test2, test1)

	for i := 0; i < len(test2); i++ {
		min := i
		for j := i; j < len(test2); j++ {
			if test2[min].Provider > test2[j].Provider {
				min = j
			}
		}
		test2[i], test2[min] = test2[min], test2[i]
	}

	for i := range test1 {
		min := i
		for j := i; j < len(test1); j++ {
			if test1[min].Country > test1[j].Country {
				min = j
			}
		}
		test1[i], test1[min] = test1[min], test1[i]
	}

	ssssm := make([][]pkg.SmsData, 2)
	for i := range ssssm {
		ssssm[i] = make([]pkg.SmsData, len(test1))
	}

	for i := range ssssm {
		for j := range ssssm[i] {
			if i == 0 {
				ssssm[i][j] = test2[j]
			} else {
				ssssm[i][j] = test1[j]
			}
		}
	}
	return ssssm
}

func main() {

	var smsresult [][]pkg.SmsData
	smsresult = smsres()
	fmt.Println("")
	for i := 0; i < len(smsresult); i++ {
		fmt.Println(smsresult[i])
	}

}
