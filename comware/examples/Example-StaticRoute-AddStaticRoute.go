package main

import (
	"log"

	"github.com/exsver/netconf/comware"
)

func main() {
	// netconf.LogLevel.Messages()
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// ip route-static 100.100.100.0 24 10.10.10.1
	IPv4Route := comware.IPv4StaticRoute{
		Ipv4Address:        "100.100.100.0",
		NexthopIpv4Address: "10.10.10.1",
		Ipv4PrefixLength:   24,
	}

	err = sw.AddIPv4StaticRoute(&IPv4Route)
	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Printf("IPv4 OK\n")

	// ipv6 route-static 2001:DB8:: 32 Vlan-interface99
	// 719 - ifIndex of Vlan-interface99
	IPv6Route := comware.IPv6StaticRoute{
		Ipv6Address:        "2001:DB8::",
		Ipv6PrefixLength:   32,
		NexthopIpv6Address: "0::0",
		IfIndex: 719,
	}

	err = sw.AddIPv6StaticRoute(&IPv6Route)
	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Printf("IPv6 OK\n")
}
