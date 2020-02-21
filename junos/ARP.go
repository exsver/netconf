package junos

import (
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/exsver/netconf/netconf"
)

type ARPTableInformation struct {
	XMLName         xml.Name        `xml:"arp-table-information"`
	ARPEntryCount   int             `xml:"arp-entry-count"`
	ARPTableEntries []ARPTableEntry `xml:"arp-table-entry"`
}
type ARPTableEntry struct {
	MACAddress         string `xml:"mac-address"`
	IPAddress          string `xml:"ip-address"`
	InterfaceName      string `xml:"interface-name"`
	ARPTableEntryFlags flags  `xml:"arp-table-entry-flags"`
}

type flags struct {
	ARPTableEntryFlags []byte `xml:",innerxml"`
}

func (targetDevice *TargetDevice) GetARPTableInformation() (ARPTableInformation, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get-arp-table-information><no-resolve/></get-arp-table-information>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return ARPTableInformation{}, err
	}

	if len(rpcReply.Errors) != 0 {
		errorString := ""
		for _, v := range rpcReply.Errors {
			errorString = fmt.Sprintf("%s, error-type: %s, error-message: %s error-info: %s\n", errorString, v.ErrorType, v.ErrorMessage, v.ErrorInfo)
		}

		return ARPTableInformation{}, errors.New(errorString)
	}

	rpcReply.Content = netconf.Normalize(rpcReply.Content)

	var arpTable ARPTableInformation

	err = xml.Unmarshal(rpcReply.Content, &arpTable)
	if err != nil {
		return arpTable, err
	}

	return arpTable, nil
}

type ClearARPTableResults struct {
	XMLName xml.Name              `xml:"clear-arp-table-results"`
	Results []ClearARPTableResult `xml:"clear-arp-table-result"`
}

type ClearARPTableResult struct {
	XMLName      xml.Name `xml:"clear-arp-table-result"`
	IPAddress    string   `xml:"ip-address"`
	ClearSuccess bool     `xml:"clear-success"`
}

//
func (targetDevice *TargetDevice) ClearARPHostname(hostname string) (ClearARPTableResults, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(fmt.Sprintf("<clear-arp-table><hostname>%s</hostname></clear-arp-table>", hostname)),
		Xmlns:    []string{netconf.BaseURI},
	}

	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return ClearARPTableResults{}, err
	}

	rpcReply.Content = netconf.Normalize(rpcReply.Content)
	rpcReply.Content = netconf.ConvertToPairedTags(rpcReply.Content)

	var clearResults ClearARPTableResults
	err = xml.Unmarshal(rpcReply.Content, &clearResults)

	return clearResults, err
}
