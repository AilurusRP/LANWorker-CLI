package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func receiveMobileMsg() {
	http.HandleFunc("/mobilemsg", func(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println("Received a message from LANWorker-mobile:")
		fmt.Println("-------------MESSAGE FROM LANWorker-mobile-------------")
		fmt.Println(receivedMsg.Msg)
		fmt.Println("----------------------------------------------------")
	})
}
