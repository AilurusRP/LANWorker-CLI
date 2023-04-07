package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/antage/eventsource.v1"
)

type ReceivedMsg struct {
	Msg string `json:"msg"`
}

type WebMsgResponse struct {
	Code int `json:"code"`
}

func startServer() {
	receiveMsg()

	es := eventsource.New(nil, nil)
	defer es.Close()

	http.Handle("/events", es)
	go func() {
		id := 0
		for {
			msg := msgInput()
			id++
			es.SendEventMessage(msg, "message", strconv.Itoa(id))
		}
	}()

	log.Fatal(http.ListenAndServe(PORT, nil))
}

func receiveMsg() {
	http.HandleFunc("/webmsg", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("read body err, %v\n", err)
			return
		}
		var receivedMsg ReceivedMsg
		if err := json.Unmarshal([]byte(body), &receivedMsg); err != nil {
			fmt.Printf("parse json err, %v\n", err)
			return
		}

		response, err := json.Marshal(WebMsgResponse{Code: 200})
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Write(response)

		fmt.Println("----------------------------------------------------")
		fmt.Println("Received a message from LANWorker-web:")
		fmt.Println("-------------MESSAGE FROM LANWorker-web-------------")
		fmt.Println(receivedMsg.Msg)
		fmt.Println("----------------------------------------------------")
	})
}

func serveWebpage(webPath string) {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(webPath+"static"))))
	http.HandleFunc("/web", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, webPath+"index.html")
	})

	fmt.Println("The LANWorker webpage is now running,")
	fmt.Printf("visit http://%s%s/web to use LANWorker-web. \n", IP, PORT)
	fmt.Print("Make sure your device is in the same LAN with LANWorker-desktop.\n\n")
}
