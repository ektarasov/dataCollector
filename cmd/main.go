package main

import (
	"diplom/pkg"
	"fmt"
)

func main() {

	var smsMap []pkg.SmsData
	smsMap = pkg.SmsCollect()
	fmt.Println("Sms данные")
	for i := 0; i < len(smsMap); i++ {
		fmt.Println(smsMap[i])
	}

	var mmsMap []pkg.MMSData
	mmsMap = pkg.MmsCollect()
	fmt.Println("Mms данные")
	for i := 0; i < len(mmsMap); i++ {
		fmt.Println(mmsMap[i])
	}

	var voiceMap []pkg.VoiceCallData
	voiceMap = pkg.VoiceCollect()
	fmt.Println("Voice данные")
	for i := 0; i < len(voiceMap); i++ {
		fmt.Println(voiceMap[i])
	}

	var emailMap []pkg.EmailData
	emailMap = pkg.EmailCollect()
	fmt.Println("email Данные")
	for i := 0; i < len(emailMap); i++ {
		fmt.Println(emailMap[i])
	}

}
