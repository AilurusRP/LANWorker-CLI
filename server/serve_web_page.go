package server

import (
	"LANWorker-desktop/info"
	"fmt"
	"net/http"
	"strings"
)

func ServeWebpage(webPath string) {
	if !strings.HasSuffix(webPath, "/") {
		webPath = webPath + "/"
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(webPath+"static"))))
	http.HandleFunc("/web", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, webPath+"index.html")
	})

	fmt.Println("The LANWorker webpage is now running,")
	fmt.Printf("visit http://%s%s/web to use LANWorker-web. \n", info.IP, info.PORT)
	fmt.Print("Make sure your device is in the same LAN with LANWorker-desktop.\n\n")
}
