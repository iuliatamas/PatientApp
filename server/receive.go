package server

import (
	"log"
	"net/http"
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

func main() {
	messages := make(chan *pkg.RecvdMessage, 1)
	h := pkg.NewMessageHandler(messages)

	go func() {
		for {
			msg := <-messages
			log.Printf("%v\n", msg)

			// mesg.MSIDN is the sender ID (phone number)
			// messg.Text is text body
			// should add message to databse with primary key sender ID
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
