package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/exsver/netconf/junos"
)

//Struct for cli Commands: show configuration *
type ConfigurationInformation struct {
	XMLName  xml.Name `xml:"configuration-information"`
	Response []byte   `xml:"configuration-output"`
}

func main() {

	sw, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	reply, _ := sw.RunCLICommand("show configuration vlans")
	var configurationInformation ConfigurationInformation
	xml.Unmarshal(reply, &configurationInformation)
	fmt.Printf("%s", configurationInformation.Response)

}
