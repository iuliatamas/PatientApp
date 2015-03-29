package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/gorilla/context"

	pkg "github.com/njern/gonexmo"
)

/*
	Client receives messages on the phone
	But we have callback to listen for
*/

type Page struct {
	Title string
}

var DEMO_MODE_TEXT string = "Please send 'demo' if you would like to try out the Virtual CareTaker"

func main() {
	messages := make(chan *pkg.RecvdMessage, 1)
	h := pkg.NewMessageHandler(messages)

	go func() {
		for {
			msg := <-messages
			WantDemo := strings.ToLower(msg.Text) == "demo"
			log.Printf("%v\n", msg)

			// senderID is a phone number
			senderID := msg.MSISDN

			var p *Patient
			var exist bool
			if p, exist = Patients[senderID]; !exist || p == nil || WantDemo {
				if !DEMO {
					log.Println("Only operating in demo mode now. Exiting.")
					return
				}

				// must create new patient
				p = NewPatient(senderID)
				p.DecisionTree = NewDemoTree(p)
				Patients[senderID] = p
				fmt.Println("News patient:", p)

				// DEMO
				if !WantDemo {
					// inform of demo mode
					fmt.Println("informing of demo mode")
					S.SendSMS(p, DEMO_MODE_TEXT)
					continue
				} else {
					log.Println("Demo request from", senderID)
					S.SendSMS(p, "Thank you for the demo request")
					// initiate conversation
					S.SendSMS(p, p.DecisionTree.Action.String())

				}
			} else {
				fmt.Println("Acting on respose for patient", p)
				S.ActOnResponse(p, msg.Text)
			}
		}
	}()

	ph := func(w http.ResponseWriter, req *http.Request) {
		log.Println("Get received")
		p := &Page{Title: "Home"}
		t, err := template.ParseFiles("index.html")
		if err != nil {
			log.Fatal(err)
		}

		t.Execute(w, p)
		h(w, req)
	}
	// Set your Nexmo callback url to http://<domain or ip>:8080/get/
	http.HandleFunc("/get/", ph)
	if err := http.ListenAndServe(":8000", context.ClearHandler(http.DefaultServeMux)); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
