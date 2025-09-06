package main

import (
	"LANWorker-CLI/info"
	"LANWorker-CLI/server"
	"flag"
	"fmt"
	"os/user"
)

func main() {
	info.GetLocalIP()
	fmt.Printf("Your current desktop local IP address: %s\n\n", info.IP)

	qrString, err := generateQRCode(info.IP)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(qrString)

	// If you need to run LANWorker-web,
	// the path to the web resources should be passed from the command line.
	// Example:
	//     $ LANWorker-CLI --path /var/www/LANWorker-web
	webPath := flag.String("path", "", "the path of the LANWorker-web directory")
	flag.Parse()
	if *webPath != "" {
		server.ServeWebpage(*webPath)
	} else {
		currentUser, err := user.Current()
		if err != nil {
			fmt.Println(err)
			return
		}
		server.ServeWebpage(currentUser.HomeDir + "/LANWorker-web")
	}
	server.StartServer()
}
