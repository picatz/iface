package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "iface"
	app.Version = "1.0.1"
	app.Usage = "network interface command-line utility"

	app.Commands = []cli.Command{
		{
			Name:  "names",
			Usage: "all network interface names",
			Action: func(c *cli.Context) error {
				ifaces, err := net.Interfaces()
				if err != nil {
					log.Fatal(err)
				}
				for _, iface := range ifaces {
					fmt.Println(iface.Name)
				}
				return nil
			},
		},
		{
			Name:  "macs",
			Usage: "network interfaces with their mac addresses",
			Action: func(c *cli.Context) error {
				ifaces, err := net.Interfaces()
				if err != nil {
					log.Fatal(err)
				}
				for _, iface := range ifaces {
					if iface.HardwareAddr != nil {
						fmt.Println(iface.HardwareAddr, iface.Name)
					}
				}
				return nil
			},
		},
		{
			Name:  "mtus",
			Usage: "network interfaces with their maximum transmission unit (mtu)s",
			Action: func(c *cli.Context) error {
				ifaces, err := net.Interfaces()
				if err != nil {
					log.Fatal(err)
				}
				for _, iface := range ifaces {
					fmt.Println(iface.MTU, iface.Name)
				}
				return nil
			},
		},
		{
			Name:  "ips",
			Usage: "network interfaces with their (local) ip addresses",
			Action: func(c *cli.Context) error {
				ifaces, err := net.Interfaces()
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}
				for _, iface := range ifaces {
					addrs, err := iface.Addrs()
					if err != nil {
						log.Fatal(err)
						os.Exit(1)
					}
					for _, addr := range addrs {
						fmt.Println(iface.Name, addr)
					}
				}
				return nil
			},
		},
		{
			Name:  "default",
			Usage: "the first non-loopback network interface",
			Action: func(c *cli.Context) error {
				ifaces, err := net.Interfaces()
				if err != nil {
					log.Fatal(err)
				}
				for _, iface := range ifaces {
					if iface.HardwareAddr != nil {
						fmt.Println(iface.Name)
						break
					}
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
