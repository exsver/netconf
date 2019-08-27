package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/exsver/netconf/helpers/comware"
	"log"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}
	data, err := sw.GetDataMGROUP()
	if err != nil {
		log.Fatalf("%s", err)
	}
	spew.Dump(data)
}
