package app

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func Start() *cli.App {
	defer fmt.Println("Application completed!")
	app := cli.NewApp()
	app.Name = "Application of the command line"
	app.Usage = "Get information the URL"
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Get IP's address of URL",
			Flags:  flags,
			Action: getIp,
		},
		{
			Name:   "server",
			Usage:  "Get the server name of URL",
			Flags:  flags,
			Action: getServer,
		},
	}
	if error := app.Run(os.Args); error != nil {
		log.Fatal(error)
	}
	return app
}

func getIp(c *cli.Context) {
	host := c.String("host")
	fmt.Printf("Get IP in URL: %s\n", host)

	ips, error := net.LookupHost(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Printf("IP: %s\n", ip)
	}
}

func getServer(c *cli.Context) {
	host := c.String("host")
	fmt.Printf("Get server in URL: %s\n", host)

	servers, error := net.LookupNS(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Printf("Server: %s\n", server.Host)
	}
}
