package client

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	pkg "github.com/njern/gonexmo"
)

func main() {
	DEBUG := true
	f, err := os.Open("key-secret.txt")
	r := bufio.NewReader(f)

	var key, secret string
	fmt.Fscanf(r, "%s %s", &key, &secret)
	nexmo, _ := pkg.NewClientFromAPI(key, secret)

	// Test if it works by retrieving your account balance
	balance, err := nexmo.Account.GetBalance()
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	fmt.Println("balance", balance)

	// Send an SMS
	// See https://docs.nexmo.com/index.php/sms-api/send-message for details.
	message := &pkg.SMSMessage{
		From:            "12013514482",
		To:              "12033924393",
		Type:            pkg.Text,
		Text:            "Wow, spring " + time.Now().String(),
		ClientReference: "gonexmo-test " + strconv.FormatInt(time.Now().Unix(), 10),
		Class:           pkg.Standard,
	}

	if DEBUG == false {
		messageResponse, _ := nexmo.SMS.Send(message)
		fmt.Println("message response", messageResponse)
	}
}