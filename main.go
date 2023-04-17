package main

import (
	"flag"
	"fmt"
)

var IP string

const PORT string = ":7684"

func main() {
	ip, err := getLocalIP()
	if err != nil {
		fmt.Println(err)
		return
	}
	IP = ip
	fmt.Printf("Your current desktop local IP address: %s\n\n", ip)

	qrString, err := generateQRCode(ip)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(qrString)

	// If you need to run LANWorker-web,
	// the path to the web resources should be passed from the command line.
	// Example:
	//     $ LANWorker-desktop --path /var/www/LANWorker-web
	webPath := flag.String("path", "", "the path of the LANWorker-web directory")
	flag.Parse()
	if *webPath != "" {
		serveWebpage(*webPath)
	}
	startServer()
}
