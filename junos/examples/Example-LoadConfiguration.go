package main

import (
	"fmt"
	"github.com/exsver/netconf/netconf"
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
	netconf.LogLevel.Messages()

	device, err := junos.NewTargetDevice("172.21.1.250", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	LoadConfiguarationSet(device)
	LoadConfiguarationFormatXML(device)
	LoadConfiguarationFormatText(device)
}

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
    }
}
*/
func LoadConfiguarationSet(device *junos.TargetDevice) {
	configSet := `
set interfaces ae1 description testIface1
set interfaces ae1 vlan-tagging
set interfaces ae1 unit 101 vlan-id 101
set interfaces ae1 unit 101 family inet address 10.100.1.1/24
`

	loadConfigurationResults, err := device.LoadConfiguration("text", configSet, "set")
	if err != nil {
		log.Fatalf("%s", err)
	}

	if loadConfigurationResults.GetErrors() != nil {
		log.Fatalf("%s", loadConfigurationResults.GetErrors().Error())
	}

	fmt.Println("Format Set: load configuration successfully")
}

/*
interfaces {
    ae1 {
        unit 102 {
            vlan-id 102;
            family inet {
                address 10.100.2.1/24;
            }
        }
    }
}
*/
func LoadConfiguarationFormatXML(device *junos.TargetDevice) {
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

	loadConfigurationResults, err := device.LoadConfiguration("xml", configXML, "merge")
	if err != nil {
		log.Fatalf("%s", err)
	}

	if loadConfigurationResults.GetErrors() != nil {
		log.Fatalf("%s", loadConfigurationResults.GetErrors().Error())
	}

	fmt.Println("Format XML: load configuration successfully")
}

/*
interfaces {
    ae1 {
        unit 103 {
            vlan-id 103;
            family inet {
                address 10.100.3.1/24;
            }
        }
    }
}
*/
func LoadConfiguarationFormatText(device *junos.TargetDevice) {
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
	loadConfigurationResults, err := device.LoadConfiguration("text", configText, "merge")
	if err != nil {
		log.Fatalf("%s", err)
	}

	if loadConfigurationResults.GetErrors() != nil {
		log.Fatalf("%s", loadConfigurationResults.GetErrors().Error())
	}

	fmt.Println("Format Text: load configuration successfully")
}
