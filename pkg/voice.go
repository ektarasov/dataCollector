package pkg

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"os/user"
	"strconv"
)

type VoiceCallData struct {
	Country             string
	Bandwidth           string
	ResponseTime        string
	Provider            string
	ConnectionStability float32
	TTFB                int
	VoicePurity         int
	MedianOfCallsTime   int
}

func VoiceCollect() []VoiceCallData {
	curUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(curUser.HomeDir + "\\GolandProjects\\simulator\\voice.data")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'

	var voiceTemp []VoiceCallData
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
