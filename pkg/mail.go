package pkg

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"os/user"
	"strconv"
)

type EmailData struct {
	Country      string
	Provider     string
	DeliveryTime int
}

func EmailCollect() []EmailData {

	curUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(curUser.HomeDir + "\\GolandProjects\\simulator\\email.data")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'

	var emailTemp []EmailData

	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			//log.Println("Cannot read line:", err)
			continue
		}
		if len(row) != 3 {
			continue
		}

		corr := CheckEmailForCorrupt(row)

		if corr == true {
			continue
		}

		deliveryTime, _ := strconv.Atoi(row[2])

		emailTemps := EmailData{
			Country:      row[0],
			Provider:     row[1],
			DeliveryTime: deliveryTime,
		}
		emailTemp = append(emailTemp, emailTemps)
	}
	return emailTemp
}
