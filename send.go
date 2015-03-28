package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	pkg "github.com/njern/gonexmo"
)

/*
	FROM: app
	TO: patients
	WHAT: check-ins (questions)
*/
var DEBUG bool
var nexmo *pkg.Client
var S *Server = NewServer("12013514482")

// Send an SMS, with Text to a Person ( Patient, Clinician, Conact-Person)
// See https://docs.nexmo.com/index.php/sms-api/send-message for details.
func (s *Server) SendSMS(p Person, t string) {
	message := &pkg.SMSMessage{
		From:            s.PhoneNumber,
		To:              p.Phone(),
		Type:            pkg.Text,
		Text:            t,
		ClientReference: "gonexmo-test " + strconv.FormatInt(time.Now().Unix(), 10),
		Class:           pkg.Standard,
	}

	if DEBUG == false {
		messageResponse, _ := nexmo.SMS.Send(message)
		// TODO, check: This response confirms receiving
		fmt.Println("message response", messageResponse)
	}

}

func (s *Server) CheckOnPatients() {
	// for p := range common.SamplePatients {
	// 	t := "Hi,", p.name, ", how are you feeling this morning?"
	// 	SendSMS(p, t)
	// }

	// get actions from heap
}

func init() {
	// get key secret
	f, _ := os.Open("key-secret.txt")
	r := bufio.NewReader(f)
	var key, secret string
	fmt.Fscanf(r, "%s %s", &key, &secret)

	DEBUG = true
	nexmo, _ = pkg.NewClientFromAPI(key, secret)

	// getBalance()

	S.CheckOnPatients()

}

func getBalance() {
	// Test if it works by retrieving your account balance
	balance, err := nexmo.Account.GetBalance()
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	fmt.Println("balance", balance)
}