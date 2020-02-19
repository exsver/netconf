package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/junos"
)

/*
interfaces {
    ae1 {
        description testIface1;
        vlan-tagging;
        unit 101 {
            vlan-id 101;
            family inet {
                address 10.100.1.1/24;
            }
        }
        unit 102 {
            vlan-id 102;
            family inet {
                address 10.100.2.1/24;
            }
        }
        unit 103 {
            vlan-id 103;
            family inet {
                address 10.100.3.1/24;
            }
        }
    }
}
*/

func main() {
	sw, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Format Set
	configSet := `
set interfaces ae1 description testIface1
set interfaces ae1 vlan-tagging
set interfaces ae1 unit 101 vlan-id 101
set interfaces ae1 unit 101 family inet address 10.100.1.1/24
`

	loadConfigurationResults, err := sw.LoadConfiguration("text", configSet, "set")
	if err != nil {
		log.Fatalf("%s", err)
	}

	if loadConfigurationResults.Error.Error() != nil {
		log.Fatalf("%s", loadConfigurationResults.Error.Error())
	}

	if !loadConfigurationResults.OK {
		log.Fatalf("Format Set: Unknown status")
	}

	fmt.Println("Format Set: load configuration successfully")

	// Format XML
	configXML := `
<interfaces>
	<interface>
		<name>ae1</name>
		<unit>
			<name>102</name>
			<vlan-id>102</vlan-id>
			<family>
				<inet>
					<address>
						<name>10.100.2.1/24</name>
					</address>
				</inet>
			</family>
		</unit>
	</interface>
</interfaces>`

	loadConfigurationResults, err = sw.LoadConfiguration("xml", configXML, "merge")
	if err != nil {
		log.Fatalf("%s", err)
	}

	if loadConfigurationResults.Error.Error() != nil {
		log.Fatalf("%s", loadConfigurationResults.Error.Error())
	}

	if !loadConfigurationResults.OK {
		log.Fatalf("Format XML: Unknown status")
	}

	fmt.Println("Format XML: load configuration successfully")


	// Format Text
	configText := `
interfaces {
    ae1 {
        unit 103 {
            vlan-id 103;
            family inet {
                address 10.100.3.1/24;
            }
        }
    }
}`
	loadConfigurationResults, err = sw.LoadConfiguration("text", configText, "merge")
	if err != nil {
		log.Fatalf("%s", err)
	}

	if loadConfigurationResults.Error.Error() != nil {
		log.Fatalf("%s", loadConfigurationResults.Error.Error())
	}

	if !loadConfigurationResults.OK {
		log.Fatalf("Format Text: Unknown status")
	}

	fmt.Println("Format Text: load configuration successfully")
}
