/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long:  `Testing trace sub-command`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {

				showData(ip)
			}
		} else {
			fmt.Println("Please provide IP to trace")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)

}

type Ip struct {
	IP       string `json::"ip"`
	City     string `json::"city"`
	Region   string `json::"region"`
	Country  string `json::"country"`
	Loc      string `json::"loc"`
	Timezone string `json::"timezone"`
	Postal   string `json::"postal"`
}

func showData(ip string) {
	//using postman
	url := "https://ipinfo.io/1.1.1.1/geo"
	responseByte := getData(url)
	data := Ip{}

	//Unmarshal is the contrary of marshal. It allows you to convert byte data into the original data structure.
	//In go, unmarshaling is handled by the json.Unmarshal() method.
	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Println("Unable to unmarshal the response")
	}
	fmt.Println("Data Found")
	fmt.Printf("IP : %s\nCITY :%s\nREGION :%s\nCOUNTRY :%s\nLOCATION :%s\nTIMEZONE :%s\nPOSTAL :%s\n", data.IP, data.City, data.Region, data.Country, data.Loc, data.Timezone, data.Postal)
}

func getData(url string) []byte {

	response, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get the response")
	}
	//ioutil.ReadAll is a useful io utility function for reading all data from a io.
	//Reader until EOF. It’s often used to read data such as HTTP response body, files
	// and other data sources which implement io.Reader interface. Be careful though because a
	//lot can go wrong if you don’t take care while using this small seemingly harmless function.

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Unable to read the response")
	}

	return responseByte
}
