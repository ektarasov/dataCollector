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

func getCountriesList() []string {
	return []string{"RU", "US", "GB", "FR", "BL", "AT", "BG", "DK", "CA", "ES", "CH", "TR", "PE", "NZ", "MC"}
}

func getProvidersList() []string {
	return []string{"Topolo", "Rond", "Kildy"}
}

func getVoiceProvidersList() []string {
	return []string{"TransparentCalls", "E-Voice", "JustPhone"}
}

func getEmailProvidersList() []string {
	return []string{
		"Gmail",
		"Yahoo",
		"Hotmail",
		"MSN",
		"Orange",
		"Comcast",
		"AOL",
		"Live",
		"RediffMail",
		"GMX",
		"Protonmail",
		"Yandex",
		"Mail.ru",
	}
}

func CheckSmsMmsForCorrupt(s [4]string) bool {
	corr := false

	if s[0] == "" || s[1] == "" || s[2] == "" || s[3] == "" {
		corr = true
		return corr
	}

	countriesList := getCountriesList()
	corr = cycleCheck(s[0], countriesList)
	if corr == true {
		return corr
	}

	providersList := getProvidersList()
	corr = cycleCheck(s[1], providersList)
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

	if s[0] == "" || s[1] == "" || s[2] == "" || s[3] == "" || s[4] == "" || s[5] == "" || s[6] == "" || s[7] == "" {
		corr = true
		return corr
	}

	countriesList := getCountriesList()
	corr = cycleCheck(s[0], countriesList)
	if corr == true {
		return corr
	}

	providersList := getVoiceProvidersList()
	corr = cycleCheck(s[3], providersList)
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

func cycleCheck(s string, l []string) bool {
	corr := false

	for i := 0; i < len(l); i++ {
		if s == l[i] {
			break
		} else if i == len(l)-1 {
			corr = true
			return corr
		}
	}
	return corr
}

func CheckEmailForCorrupt(s []string) bool {
	corr := false

	countriesList := getCountriesList()

	corr = cycleCheck(s[0], countriesList)
	if corr == true {
		return corr
	}

	emailList := getEmailProvidersList()
	corr = cycleCheck(s[1], emailList)
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
	//resultSet.MMS = nil

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
