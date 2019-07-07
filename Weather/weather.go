package main

import (
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//API Key ce2399105c47a372ccad016ed6de15d3

var apikey string = "&APPID=ce2399105c47a372ccad016ed6de15d3"

var getWeather string = "http://api.openweathermap.org/data/2.5/weather?q="

func main() {

	//Flags for the CLI
	wflag := []cli.Flag{
		cli.StringFlag{
			Name:  "city",
			Value: "London",
		},
	}

	//command line interface setup
	app := cli.NewApp()
	app.Name = "Shubs Weather tool"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:  "s",
			Usage: "Get summary of weather by city",
			Flags: wflag,
			Action: func(c *cli.Context) error {
				response, err := http.Get(getWeather + c.String("city") + apikey)

				if err != nil {
					fmt.Println(err)
				}

				data, _ := ioutil.ReadAll(response.Body)

				fmt.Println(string(data))

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
