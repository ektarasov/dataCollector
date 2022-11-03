package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func MmsCollect() []MMSData {

	var mmsTemp []MMSData

	resp, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		fmt.Println("Не удалось получить данные")
		return mmsTemp
	}

	if resp.Status == "200 OK" {

		content, err := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal(content, &mmsTemp)
		if err != nil {
			log.Fatal(err)
			return mmsTemp
		}

		for i := 0; i < len(mmsTemp); i++ {
			var str [4]string

			str[0] = mmsTemp[i].Country
			str[1] = mmsTemp[i].Provider
			str[2] = mmsTemp[i].Bandwidth
			str[3] = mmsTemp[i].ResponseTime

			corr := CheckSmsMmsForCorrupt(str)

			if corr != false {
				mmsTemp = append(mmsTemp[:i], mmsTemp[i+1:]...)
				i--
			}
		}
		return mmsTemp
	}
	return mmsTemp
}
