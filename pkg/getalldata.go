package pkg

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type ResultSetT struct {
	SMS       [][]SmsData              `json:"sms"`
	MMS       [][]MMSData              `json:"mms"`
	VoiceCall []VoiceCallData          `json:"voice_call"`
	Email     map[string][][]EmailData `json:"email"`
	Billing   []BillingData            `json:"billing"`
	Support   [2]int                   `json:"support"`
	Incident  []IncidentData           `json:"incident"`
}

type ResultT struct {
	Status bool       `json:"status"` // true, если все этапы сбора данных прошли успешно, false во всех остальных случаях
	Data   ResultSetT `json:"data"`   // заполнен, если все этапы сбора данных прошли успешно, nil во всех остальных случаях
	Error  string     `json:"error"`  // пустая строка если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки (детали ниже)

}

func mapRenameCountry() map[string]string {
	curUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(filepath.Join(curUser.HomeDir, "GolandProjects", "diplom", "iso3166.data"))
	if err != nil {
		fmt.Println("Не удалось получить данные")
	}

	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'

	isoMap := make(map[string]string)

	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}

		isoMap[row[0]] = row[1]
	}
	return isoMap
}

func smsResult() [][]SmsData {
	var smsSortByCountry []SmsData
	smsSortByCountry = SmsCollect()
	var smsTemp [][]SmsData

	if smsSortByCountry == nil {
		return smsTemp
	}

	iso3166Map := mapRenameCountry()

	for i := 0; i < len(smsSortByCountry); i++ {

		countryName, ok := iso3166Map[smsSortByCountry[i].Country]
		if !ok {
			continue
		}
		smsSortByCountry[i].Country = countryName

	}

	smsSortByProvider := make([]SmsData, len(smsSortByCountry))
	copy(smsSortByProvider, smsSortByCountry)

	for i := 0; i < len(smsSortByProvider); i++ {
		min := i
		for j := i; j < len(smsSortByProvider); j++ {
			if smsSortByProvider[min].Provider > smsSortByProvider[j].Provider {
				min = j
			}
		}
		smsSortByProvider[i], smsSortByProvider[min] = smsSortByProvider[min], smsSortByProvider[i]
	}

	for i := range smsSortByCountry {
		min := i
		for j := i; j < len(smsSortByCountry); j++ {
			if smsSortByCountry[min].Country > smsSortByCountry[j].Country {
				min = j
			}
		}
		smsSortByCountry[i], smsSortByCountry[min] = smsSortByCountry[min], smsSortByCountry[i]
	}
	smsTemp = make([][]SmsData, 2)
	for i := range smsTemp {
		smsTemp[i] = make([]SmsData, len(smsSortByCountry))
	}

	for i := range smsTemp {
		for j := range smsTemp[i] {
			if i == 0 {
				smsTemp[i][j] = smsSortByProvider[j]
			} else {
				smsTemp[i][j] = smsSortByCountry[j]
			}
		}
	}
	return smsTemp
}

func mmsResult() [][]MMSData {
	var mmsSortByCountry []MMSData
	mmsSortByCountry = MmsCollect()
	var mmsTemp [][]MMSData

	if mmsSortByCountry == nil {
		return mmsTemp
	}

	iso3166Map := mapRenameCountry()
	for i := 0; i < len(mmsSortByCountry); i++ {

		countryName, ok := iso3166Map[mmsSortByCountry[i].Country]
		if !ok {
			continue
		}
		mmsSortByCountry[i].Country = countryName

	}

	mmsSortByProvider := make([]MMSData, len(mmsSortByCountry))
	copy(mmsSortByProvider, mmsSortByCountry)

	for i := 0; i < len(mmsSortByProvider); i++ {
		min := i
		for j := i; j < len(mmsSortByProvider); j++ {
			if mmsSortByProvider[min].Provider > mmsSortByProvider[j].Provider {
				min = j
			}
		}
		mmsSortByProvider[i], mmsSortByProvider[min] = mmsSortByProvider[min], mmsSortByProvider[i]
	}

	for i := range mmsSortByCountry {
		min := i
		for j := i; j < len(mmsSortByCountry); j++ {
			if mmsSortByCountry[min].Country > mmsSortByCountry[j].Country {
				min = j
			}
		}
		mmsSortByCountry[i], mmsSortByCountry[min] = mmsSortByCountry[min], mmsSortByCountry[i]
	}

	mmsTemp = make([][]MMSData, 2)
	for i := range mmsTemp {
		mmsTemp[i] = make([]MMSData, len(mmsSortByCountry))
	}

	for i := range mmsTemp {
		for j := range mmsTemp[i] {
			if i == 0 {
				mmsTemp[i][j] = mmsSortByProvider[j]
			} else {
				mmsTemp[i][j] = mmsSortByCountry[j]
			}
		}
	}
	return mmsTemp
}

func mailResult() map[string][][]EmailData {

	var emailTemp []EmailData
	emailTemp = EmailCollect()
	var emailResult map[string][][]EmailData

	if emailTemp == nil {
		return emailResult
	}

	emailResult = make(map[string][][]EmailData, 3)
	for i := 0; i < len(emailTemp); i++ {
		min := i
		for j := i; j < len(emailTemp); j++ {
			if emailTemp[min].Country > emailTemp[j].Country {
				min = j
			}
		}
		emailTemp[i], emailTemp[min] = emailTemp[min], emailTemp[i]
	}

	ind := make([]int, 1)

	for i := 0; i < len(emailTemp); i++ {
		min := i
		for j := i; j < len(emailTemp); j++ {
			if emailTemp[i].Country == emailTemp[j].Country {
				if emailTemp[min].DeliveryTime > emailTemp[j].DeliveryTime {
					min = j
				}
			} else {
				if ind[len(ind)-1] != j {
					ind = append(ind, j)
				}
				break
			}
		}
		emailTemp[i], emailTemp[min] = emailTemp[min], emailTemp[i]
	}

	var emailDeliveryTimeBest []EmailData
	var emailDeliveryTimeWorst []EmailData

	for i := 1; i < len(ind); i++ {

		emailDeliveryTimeBest = append(emailTemp[ind[i-1] : ind[i-1]+3])

		emailDeliveryTimeWorst = append(emailTemp[ind[i]-3 : ind[i]])

		emailMid := make([][]EmailData, 2)
		for l := range emailMid {
			emailMid[l] = make([]EmailData, 3)
		}

		for m := range emailMid {
			for j := range emailMid[m] {
				if m == 0 {
					emailMid[m][j] = emailDeliveryTimeBest[j]
				} else {
					emailMid[m][j] = emailDeliveryTimeWorst[j]
				}
			}
		}
		emailResult[emailMid[0][0].Country] = emailMid
	}

	return emailResult
}

func supportResult() [2]int {
	var supportTemp []SupportData
	supportTemp = SupportCollect()

	var sumTicket int

	var supportRes [2]int

	if supportTemp == nil {
		return supportRes
	}

	for i := range supportTemp {
		sumTicket += supportTemp[i].ActiveTickets
	}

	if sumTicket < 9 {
		supportRes[0] = 1
	} else if sumTicket > 9 && sumTicket < 16 {
		supportRes[0] = 2
	} else {
		supportRes[0] = 3
	}

	avgTimeForOneTicket := 60 / 18
	supportRes[1] = avgTimeForOneTicket * sumTicket

	return supportRes
}

func incidentResult() []IncidentData {
	var incidentTemp []IncidentData
	incidentTemp = IncidentCollect()

	if incidentTemp == nil {
		return incidentTemp
	}

	for i, v := range incidentTemp {

		if v.Status == "closed" {
			for j := i + 1; j < len(incidentTemp); j++ {
				if incidentTemp[j].Status == "active" {
					incidentTemp[j].Status, incidentTemp[i].Status = incidentTemp[i].Status, incidentTemp[j].Status
					break
				}
			}
		}
	}

	return incidentTemp
}

func GetResultData() ResultSetT {

	var resultSet ResultSetT

	resultSet.SMS = smsResult()

	resultSet.MMS = mmsResult()

	resultSet.VoiceCall = VoiceCollect()

	resultSet.Email = mailResult()

	resultSet.Billing = BillingCollect()

	resultSet.Support = supportResult()

	resultSet.Incident = incidentResult()

	return resultSet

}
