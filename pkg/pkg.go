package pkg

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var dstServerAddress = flag.String("dstServerAddress", "", "Сетевой адрес HTTP DST")
var srcServerAddress = flag.String("srcServerAddress", "", "Сетевой адрес HTTP SRC")
var mapCountriesList = map[string]bool{
	"RU": false,
	"US": false,
	"GB": false,
	"FR": false,
	"BL": false,
	"AT": false,
	"BG": false,
	"DK": false,
	"CA": false,
	"ES": false,
	"CH": false,
	"TR": false,
	"PE": false,
	"NZ": false,
	"MC": false,
}
var mapProvidersList = map[string]bool{
	"Topolo": false,
	"Rond":   false,
	"Kildy":  false,
}
var mapVoiceProvidersList = map[string]bool{
	"TransparentCalls": false,
	"E-Voice":          false,
	"JustPhone":        false,
}
var mapEmailProvidersList = map[string]bool{
	"Gmail":      false,
	"Yahoo":      false,
	"Hotmail":    false,
	"MSN":        false,
	"Orange":     false,
	"Comcast":    false,
	"AOL":        false,
	"Live":       false,
	"RediffMail": false,
	"GMX":        false,
	"Protonmail": false,
	"Yandex":     false,
	"Mail.ru":    false,
}

func mapCheck(s string, m map[string]bool) bool {
	_, ok := m[s]
	return !ok
}

func CheckSmsMmsForCorrupt(s [4]string) bool {
	corr := false

	for _, v := range s {
		if v == "" {
			corr = true
			return corr
		}
	}

	corr = mapCheck(s[0], mapCountriesList)
	if corr == true {
		return corr
	}

	corr = mapCheck(s[1], mapProvidersList)
	if corr == true {
		return corr
	}

	bandInt, err := strconv.Atoi(s[2])
	if err != nil || (bandInt < 0 || bandInt > 100) {
		corr = true
		return corr
	}
	responseInt, err := strconv.Atoi(s[3])
	if err != nil || responseInt < 0 {
		corr = true
		return corr
	}

	return corr
}

func CheckVoiceForCorrupt(s []string) bool {
	corr := false

	for _, v := range s {
		if v == "" {
			corr = true
			return corr
		}
	}

	corr = mapCheck(s[0], mapCountriesList)
	if corr == true {
		return corr
	}

	corr = mapCheck(s[3], mapVoiceProvidersList)
	if corr == true {
		return corr
	}

	bandInt, err := strconv.Atoi(s[1])
	if err != nil || (bandInt < 0 || bandInt > 100) {
		corr = true
		return corr
	}
	responseInt, err := strconv.Atoi(s[2])
	if err != nil || responseInt < 0 {
		corr = true
		return corr
	}
	connectionStability, err := strconv.ParseFloat(s[4], 32)
	if err != nil || connectionStability < 0 {
		corr = true
		return corr
	}
	TTFB, err := strconv.Atoi(s[5])
	if err != nil || TTFB < 0 {
		corr = true
		return corr
	}
	voicePurity, err := strconv.Atoi(s[6])
	if err != nil || voicePurity < 0 {
		corr = true
		return corr
	}
	medianOfCallsTime, err := strconv.Atoi(s[7])
	if err != nil || medianOfCallsTime < 0 {
		corr = true
		return corr
	}

	return corr
}

func CheckEmailForCorrupt(s []string) bool {
	corr := false

	for _, v := range s {
		if v == "" {
			corr = true
			return corr
		}
	}

	corr = mapCheck(s[0], mapCountriesList)
	if corr == true {
		return corr
	}

	corr = mapCheck(s[1], mapEmailProvidersList)
	if corr == true {
		return corr
	}

	deliveryTime, err := strconv.Atoi(s[2])
	if err != nil || deliveryTime < 0 {
		corr = true
		return corr
	}

	return corr
}

func handleConnections(w http.ResponseWriter, r *http.Request) {

	var resultSet ResultSetT
	var Result ResultT

	resultSet = GetResultData()

	if resultSet.SMS == nil || resultSet.MMS == nil || resultSet.Email == nil || resultSet.Support[0] == 0 || resultSet.Incident == nil || resultSet.Billing.Err == true || resultSet.VoiceCall == nil {

		Result.Status = false
		Result.Error = "Error on collect data"

	} else {
		Result.Data = resultSet
		Result.Status = true
	}

	jsonResult, err := json.Marshal(Result)

	if err != nil {
		w.Write([]byte("Ошибка получения данных"))
	}

	w.Write(jsonResult)
}

func ListenAndServeHTTP() {

	router := mux.NewRouter()
	router.HandleFunc("/", handleConnections)
	http.ListenAndServe(*srcServerAddress, router)

}
