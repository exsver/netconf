package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

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
      <InterfaceCapabilities>
        <Interface>
          <IfIndex/>
          <Speed/>
        </Interface>
      </InterfaceCapabilities>
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

	type ifCapabilities struct {
		Speed int
	}

	capabilities := make(map[int]ifCapabilities)
	for _, ifCap := range data.Top.Ifmgr.InterfaceCapabilities.Interfaces {
		capabilities[ifCap.IfIndex] = ifCapabilities{
			Speed: ifCap.Speed,
		}
	}

	for _, port := range data.Top.Ifmgr.Ports.Ports {
		fmt.Printf("IfIndex: %v IfName: '%s', IfAbbreviatedName: '%s', Description: '%s' Speed: %s\n",
			port.IfIndex,
			port.Name,
			info[port.IfIndex].AbbreviatedName,
			info[port.IfIndex].Description,
			IfSpeedString(capabilities[port.IfIndex].Speed),
		)
	}
}

func IfSpeedString(in int) string {
	sp := make([]int, 0)

	if (in & 65536) != 0 {
		sp = append(sp, 25000)
	}

	if (in & 32768) != 0 {
		sp = append(sp, 5000)
	}

	if (in & 16384) != 0 {
		sp = append(sp, 100000)
	}

	if (in & 8192) != 0 {
		sp = append(sp, 40000)
	}

	if (in & 1024) != 0 {
		sp = append(sp, 10000)
	}

	if (in & 128) != 0 {
		sp = append(sp, 2500)
	}

	if (in & 32) != 0 {
		sp = append(sp, 1000)
	}

	if (in & 4) != 0 {
		sp = append(sp, 100)
	}

	if (in & 2) != 0 {
		sp = append(sp, 10)
	}

	sort.Ints(sp)

	out := ""
	for _, speed := range sp {
		out = fmt.Sprintf("%s%v/", out, speed)
	}

	return strings.TrimSuffix(out, "/")
}
