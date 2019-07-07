package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Shubs Website Lookup"
	app.Usage = "Get information on Websites, IPS, MX Records etc"

	Flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{

		{
			Name:  "ns",
			Usage: "Looks up name servers for host",
			Flags: Flags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)

				}

				return nil
			},
		},

		{
			Name:  "ip",
			Usage: "Looks up IP for host",
			Flags: Flags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ip); i++ {

					fmt.Println(ip[i])

				}
				return nil
			},
		},

		{
			Name:  "cname",
			Usage: "Look up cname of host",
			Flags: Flags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {

					fmt.Println(err)
					return err

				}

				fmt.Println(cname)
				return nil
			},
		},

		{
			Name:  "mx",
			Usage: "Show MX records for host",
			Flags: Flags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err

				}

				for i := 0; i < len(mx); i++ {

					fmt.Println(mx[i].Host, mx[i].Pref)

				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}
