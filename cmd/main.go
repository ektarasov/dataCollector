package main

import (
	"diplom/pkg"
	"fmt"
)

func main() {

	var smsMap []pkg.SmsData
	smsMap = pkg.SmsCollect()
	fmt.Println("\nSms данные")
	for i := 0; i < len(smsMap); i++ {
		fmt.Println(smsMap[i])
	}

	var mmsMap []pkg.MMSData
	mmsMap = pkg.MmsCollect()
	fmt.Println("\nMms данные")
	for i := 0; i < len(mmsMap); i++ {
		fmt.Println(mmsMap[i])
	}

	var voiceMap []pkg.VoiceCallData
	voiceMap = pkg.VoiceCollect()
	fmt.Println("\nVoice данные")
	for i := 0; i < len(voiceMap); i++ {
		fmt.Println(voiceMap[i])
	}

	var emailMap []pkg.EmailData
	emailMap = pkg.EmailCollect()
	fmt.Println("\nEmail данные")
	for i := 0; i < len(emailMap); i++ {
		fmt.Println(emailMap[i])
	}

	var billingMap []pkg.BillingData
	billingMap = pkg.BillingCollect()
	fmt.Println("\nBilling данные")
	for i := 0; i < len(billingMap); i++ {
		fmt.Println(billingMap[i])
	}

	var supportMap []pkg.SupportData
	supportMap = pkg.SupportCollect()
	fmt.Println("\nSupport данные")
	for i := 0; i < len(supportMap); i++ {
		fmt.Println(supportMap[i])
	}

	var incidentMap []pkg.IncidentData
	incidentMap = pkg.IncidentCollect()
	fmt.Println("\nIncident данные")
	for i := 0; i < len(incidentMap); i++ {
		fmt.Println(incidentMap[i])
	}

	pkg.ListenAndServeHTTP()
}
