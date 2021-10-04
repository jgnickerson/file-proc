package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var nc *nats.Conn

type PostFileRequest struct {
	Filename string `json:"filename"`
}

type PostFileResponse struct {
	UUID string `json:"uuid"`
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	//won't work for big files
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	}

	id := uuid.New()

	err = os.WriteFile(fmt.Sprintf("/tmp/%s", id.String()), data, 0644)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	}

	resp := &PostFileResponse{UUID: id.String()}
	w.WriteHeader(http.StatusAccepted)
	b, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	}
	w.Write(b)
}

//func handleGet(w http.ResponseWriter, r *http.Request) {
//
//}

func publishLocation() {
	_, err := nc.Subscribe("file.server", func(msg *nats.Msg) {
		err := nc.Publish(msg.Reply, []byte("localhost:8080"))
		if err != nil {
			log.Println(err)
		}
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var err error
	nc, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlePost(w, r)
		}
	})

	publishLocation()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
