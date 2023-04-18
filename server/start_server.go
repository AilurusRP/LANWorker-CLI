package server

import (
	"log"
	"net/http"
	"strconv"

	"gopkg.in/antage/eventsource.v1"

	"LANWorker-desktop/info"
	"LANWorker-desktop/input"
)

type ReceivedMsg struct {
	Msg string `json:"msg"`
}

type WebMsgResponse struct {
	Code int `json:"code"`
}

func StartServer() {
	receiveMsg()

	es := eventsource.New(nil, nil)
	defer es.Close()

	http.Handle("/events", es)
	go func() {
		id := 0
		for {
			msg := input.MsgInput()
			id++
			es.SendEventMessage(msg, "message", strconv.Itoa(id))
		}
	}()

	log.Fatal(http.ListenAndServe(info.PORT, nil))
}

func receiveMsg() {
	receiveWebMsg()
	receiveMobileMsg()
}
