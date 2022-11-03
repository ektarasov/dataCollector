package pkg

import (
	"io/ioutil"
	"log"
	"os/user"
	"strconv"
)

type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

func BillingCollect() []BillingData {

	var billingTemp []BillingData
	curUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fContent, err := ioutil.ReadFile(curUser.HomeDir + "\\GolandProjects\\simulator\\billing.data")
	if err != nil {
		panic(err)
	}
	str := string(fContent)

	for len(str) >= 6 {
		str1 := str[:6]
		j, _ := strconv.ParseInt(str1, 2, 8)
		u8 := uint8(j)
		var x [6]bool
		for i := 0; i < 6; i++ {
			var nbit = u8 & 1
			if nbit == 1 {
				x[i] = true
			} else if nbit == 0 {
				x[i] = false
			}
			u8 = u8 >> 1
		}
		billingTemps := BillingData{
			CreateCustomer: x[0],
			Purchase:       x[1],
			Payout:         x[2],
			Recurring:      x[3],
			FraudControl:   x[4],
			CheckoutPage:   x[5],
		}
		billingTemp = append(billingTemp, billingTemps)
		str = str[6:]
	}

	return billingTemp
}
