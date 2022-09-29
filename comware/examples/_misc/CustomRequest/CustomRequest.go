package main

import (
	"fmt"
	"log"

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

	// Creating a new device.
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	request := netconf.RPCMessage{
		Xmlns: []string{netconf.BaseURI},
		InnerXML: []byte(`
<get><filter type="subtree">
  <top xmlns="http://www.hp.com/netconf/data:1.0">
    <Device>
      <Base>
        <Uptime/><HostName/><LocalTime/><BridgeMAC/>
      </Base>
      <PhysicalEntities>
        <Entity>
          <PhysicalIndex/><Chassis/><Slot/><Description/><Name/><HardwareRev/><FirmwareRev/><SoftwareRev/><SerialNumber/><MfgName/><Model/><MfgDate/>
          <Class>3</Class>
        </Entity>
      </PhysicalEntities>
    </Device>
    <Ifmgr>
      <Interfaces>
        <Interface>
          <IfIndex/>
          <Name/>
          <AbbreviatedName/>
          <Description/>
        </Interface>
      </Interfaces>
      <Ports/>
    </Ifmgr>
  </top>
</filter></get>`),
	}

	data, err := sw.RetrieveData(request)
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Base info
	fmt.Printf("Hostname: '%s', Uptime: %v seconds, LocalTime: '%s', BridgeMAC: '%s'\n",
		data.Top.Device.Base.HostName,
		data.Top.Device.Base.Uptime,
		data.Top.Device.Base.LocalTime,
		data.Top.Device.Base.BridgeMAC,
	)

	// Slots info
	for _, slot := range data.Top.Device.PhysicalEntities.PhysicalEntities {
		fmt.Printf("Chassis: %v Slot: %v, Model: '%s', Name: '%s', Description: '%s', HardwareRev: '%s', FirmwareRev: '%s', SoftwareRev: '%s', SerialNumber: '%s' MfgName: '%s', MfgDate: '%s'\n",
			slot.Chassis,
			slot.Slot,
			slot.Model,
			slot.Name,
			slot.Description,
			slot.HardwareRev,
			slot.FirmwareRev,
			slot.SoftwareRev,
			slot.SerialNumber,
			slot.MfgName,
			slot.MfgDate,
		)
	}

	// Ports info
	type ifInfo struct {
		Name            string
		AbbreviatedName string
		Description     string
	}

	info := make(map[int]ifInfo)
	for _, iface := range data.Top.Ifmgr.Interfaces.Interfaces {
		info[iface.IfIndex] = ifInfo{
			Name:            iface.Name,
			AbbreviatedName: iface.AbbreviatedName,
			Description:     iface.Description,
		}
	}

	for _, port := range data.Top.Ifmgr.Ports.Ports {
		fmt.Printf("IfIndex: %v IfName: '%s', IfAbbreviatedName: '%s', Description: '%s'\n", port.IfIndex, port.Name, info[port.IfIndex].AbbreviatedName, info[port.IfIndex].Description)
	}
}
