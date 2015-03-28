package client

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	pkg "github.com/njern/gonexmo"
)

/*
	FROM: app
	TO: patients
	WHAT: check-ins (questions)
*/

type Server struct {
	PhoneNumber string

	// queue of actions
	ActLock sync.Mutex
	Actions []*Action
}

func NewServer(pn string) *Server {
	return Server{PhoneNumber: pn, Actions: make([]Action, 0)}
}

var S Server = NewServer("12013514482")

// Send an SMS, with Text to Patient
// See https://docs.nexmo.com/index.php/sms-api/send-message for details.
func SendSMS(p common.Patient, t Text) {
	message := &pkg.SMSMessage{
		From:            OurPhoneNumber,
		To:              p.Phone,
		Type:            pkg.Text,
		Text:            t,
		ClientReference: "gonexmo-test " + strconv.FormatInt(time.Now().Unix(), 10),
		Class:           pkg.Standard,
	}

}

func (s *Server) CheckOnPatients() {
	// for p := range common.SamplePatients {
	// 	t := "Hi,", p.name, ", how are you feeling this morning?"
	// 	SendSMS(p, t)
	// }

	for _, act := range s.Actions {
		s.ActOnAction(act)
	}
}

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

	S.CheckOnPatients()

	if DEBUG == false {
		messageResponse, _ := nexmo.SMS.Send(message)
		fmt.Println("message response", messageResponse)
	}
}
