package pkg

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"os/user"
)

type SmsData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

func SmsCollect() []SmsData {
	curUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(curUser.HomeDir + "\\GolandProjects\\simulator\\sms.data")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'
	var smsTemp []SmsData
	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			//	log.Println("Cannot read line:", err)
			continue
		}
		if len(row) != 4 {
			continue
		}
		var str [4]string
		str[0] = row[0]
		str[1] = row[3]
		str[2] = row[1]
		str[3] = row[2]

		corr := CheckSmsMmsForCorrupt(str)
		if corr == true {
			continue
		}

		smsTemps := SmsData{
			Country:      row[0],
			Bandwidth:    row[1],
			ResponseTime: row[2],
			Provider:     row[3],
		}
		smsTemp = append(smsTemp, smsTemps)
	}
	return smsTemp
}
