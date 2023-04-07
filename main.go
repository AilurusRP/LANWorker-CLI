package main

import (
	"fmt"
	"os"
)

var IP string
var PORT string = ":7684"

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
	//     $ LANWorker-desktop /var/www/LANWorker-web
	if len(os.Args) >= 2 {
		serveWebpage(os.Args[1])
	}
	startServer()
}
