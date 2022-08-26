package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "t",
	Short: "Trace the IP",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				showData(ip)
			}
		} else {
			fmt.Println("Please provide IP to trace.")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

// json response from ipinfo
// {
// 	"ip": "8.8.8.8",
// 	"hostname": "dns.google",
// 	"anycast": true,
// 	"city": "Mountain View",
// 	"region": "California",
// 	"country": "US",
// 	"loc": "37.4056,-122.0775",
// 	"org": "AS15169 Google LLC",
// 	"postal": "94043",
// 	"timezone": "America/Los_Angeles",
// 	"readme": "https://ipinfo.io/missingauth"
//   }

type Ip struct {
	IP       string `json:"ip"`
	HostName string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Timezone string `json:"timezone"`
	Postal   string `json:"postal"`
}

func showData(ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"
	responseByte := getData(url)

	data := Ip{}

	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Println("Unable to unmarshal the response")
	}

	c := color.New(color.FgRed).Add(color.Underline).Add(color.Bold)
	c.Println("DATA FOUND :")

	fmt.Printf("IP :%s\nCITY :%s\nHOSTNAME :%s\nREGION :%s\nCOUNTRY :%s\nLOCATION :%s\nTIMEZONE:%s\nPOSTAL :%s\n", data.IP, data.City, data.HostName, data.Region, data.Country, data.Loc, data.Timezone, data.Postal)

}

func getData(url string) []byte {

	response, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get the response")
	}

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Unable to read the response")
	}

	return responseByte
}
