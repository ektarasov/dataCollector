package pkg

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"strconv"
)

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_calls_time"`
}

func VoiceCollect() []VoiceCallData {
	curUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	var voiceTemp []VoiceCallData
	file, err := os.Open(curUser.HomeDir + "\\GolandProjects\\simulator\\voice.data")

	if err != nil {
		fmt.Println("Не удалось получить данные")
		return voiceTemp
	}

	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'

	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			//	log.Println("Cannot read line:", err)
			continue
		}
		if len(row) != 8 {
			continue
		}

		corr := CheckVoiceForCorrupt(row)

		if corr == true {
			continue
		}

		connStab, _ := strconv.ParseFloat(row[4], 32)
		connStab32 := float32(connStab)
		ttfb, _ := strconv.Atoi(row[5])
		voicePurity, _ := strconv.Atoi(row[6])
		medianOfCallsTime, _ := strconv.Atoi(row[7])

		voiceTemps := VoiceCallData{
			Country:             row[0],
			Bandwidth:           row[1],
			ResponseTime:        row[2],
			Provider:            row[3],
			ConnectionStability: connStab32,
			TTFB:                ttfb,
			VoicePurity:         voicePurity,
			MedianOfCallsTime:   medianOfCallsTime,
		}
		voiceTemp = append(voiceTemp, voiceTemps)
	}
	return voiceTemp
}
