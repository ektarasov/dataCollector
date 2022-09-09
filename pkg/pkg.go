package pkg

import (
	"strconv"
)

func GetCountriesList() []string {
	return []string{"RU", "US", "GB", "FR", "BL", "AT", "BG", "DK", "CA", "ES", "CH", "TR", "PE", "NZ", "MC"}
}

func GetProvidersList() []string {
	return []string{"Topolo", "Rond", "Kildy"}
}

func CheckSmsMmsForCorrupt(s [4]string) bool {
	corr := false

	if s[0] == "" || s[1] == "" || s[2] == "" || s[3] == "" {
		corr = true
		return corr
	}

	countriesList := GetCountriesList()

	for i := 0; i < len(countriesList); i++ {
		if countriesList[i] == s[0] {
			break
		} else if i == len(countriesList)-1 {
			corr = true
			return corr
		}
	}

	providersList := GetProvidersList()

	for i := 0; i < len(providersList); i++ {
		if providersList[i] == s[1] {
			break
		} else if i == len(providersList)-1 {
			corr = true
			return corr
		}
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

func GetVoiceProvidersList() []string {
	return []string{"TransparentCalls", "E-Voice", "JustPhone"}
}

func GetEmailProvidersList() []string {
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

func CheckVoiceForCorrupt(s []string) bool {
	corr := false

	if s[0] == "" || s[1] == "" || s[2] == "" || s[3] == "" || s[4] == "" || s[5] == "" || s[6] == "" || s[7] == "" {
		return corr
	}

	countriesList := GetCountriesList()

	for i := 0; i < len(countriesList); i++ {
		if countriesList[i] == s[0] {
			break
		} else if i == len(countriesList)-1 {
			corr = true
			return corr
		}
	}

	providersList := GetVoiceProvidersList()

	for i := 0; i < len(providersList); i++ {
		if providersList[i] == s[3] {
			break
		} else if i == len(providersList)-1 {
			corr = true
			return corr
		}
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

	countryList := GetCountriesList()

	for i := 0; i < len(countryList); i++ {
		if s[0] == countryList[i] {
			break
		} else if i == len(countryList)-1 {
			corr = true
			return corr
		}
	}

	emailList := GetEmailProvidersList()

	for i := 0; i < len(emailList); i++ {
		if s[1] == emailList[i] {
			break
		} else if i == len(emailList)-1 {
			corr = true
			return corr
		}
	}

	deliveryTime, err := strconv.Atoi(s[2])
	if err != nil || deliveryTime < 0 {
		corr = true
		return corr
	}

	return corr
}
