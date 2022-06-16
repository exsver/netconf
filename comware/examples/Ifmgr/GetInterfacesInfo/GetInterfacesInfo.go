package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"

	"github.com/exsver/netconf/comware"
	"github.com/exsver/netconf/netconf"
)

func main() {
	// Setting the Log Level for netconf lib.
	// One of:
	//   netconf.LogLevel.Silent()
	//   netconf.LogLevel.Default() - default
	//   netconf.LogLevel.Messages()
	//   netconf.LogLevel.Verbose()
	netconf.LogLevel.Verbose()

	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Filters examples:
	//   all items -                           nil
	//   all BAGG interfaces -                 []string{`<ifType>161</ifType>`}
	//   all ethernet Interfaces -             []string{`<ifType>6</ifType>`}
	//   all Vlan-interfaces -                 []string{`<ifType>136</ifType>`}
	//   interface with ifIndex 10 -           []string{`<IfIndex>10</IfIndex>`}
	//   interface with index 10 -             []string{`<PortIndex>10</PortIndex>`}
	//   interface with name GE1/0/10          []string{`<AbbreviatedName>GE1/0/10</AbbreviatedName>`}
	//   all Ports in Up state -               []string{`<OperStatus>1</OperStatus>`}
	//   interfaces with Description "test" -  []string{`<Description>test</Description>`}
	//   all ethernet Interfaces in UP state - []string{`<ifType>6</ifType>`, `<OperStatus>1</OperStatus>`}

	// Matching interface by IfIndex
	interfaces, err := sw.GetInterfacesInfo([]string{`<IfIndex>10</IfIndex>`})
	spew.Dump(interfaces, err)

	// Matching interface by AbbreviatedName
	interfaces, err = sw.GetInterfacesInfo([]string{`<AbbreviatedName>GE1/0/10</AbbreviatedName>`})
	spew.Dump(interfaces, err)
}
