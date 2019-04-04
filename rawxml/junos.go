package rawxml

var XMLMessagesJunOS = map[string]string{
	"GetChassisInventory":      `<rpc><get-chassis-inventory/></rpc>`,
	"GetSoftwareInformation":   `<rpc><get-software-information/></rpc>`,
	"GetSystemInformation":     `<rpc><get-system-information/></rpc>`,
	"GetChassisMacAddresses":   `<rpc><get-chassis-mac-addresses/></rpc>`,
	"GetInterfacesInformation": `<rpc><get-interface-information><terse/></get-interface-information></rpc>`,
	"GetVRRPInformation":       `<rpc><get-vrrp-information/></rpc>`,
	"GetARP":                   `<rpc><get-arp-table-information/></rpc>`,
	"Commit":                   `<rpc><commit/></rpc>`,
	"Reboot":                   `<rpc><request-reboot/></rpc>`,
}
