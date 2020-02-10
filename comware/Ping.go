package comware

import (
	"bytes"
	"encoding/xml"

	"github.com/exsver/netconf/netconf"
)

type PingResult struct {
	XMLName     xml.Name       `xml:"action-result"`
	RequestPath string         `xml:"request-path"`
	ResultData  PingResultData `xml:"result-data"`
}

type PingResultData struct {
	Ping Ping `xml:"Ping"`
}

type Ping struct {
	IPv4Ping IPv4Ping `xml:"IPv4Ping"`
}

type IPv4Ping struct {
	PingTest PingTest `xml:"PingTest"`
}

type PingTest struct {
	Host                string      `xml:"Host"`
	PayloadLength       int         `xml:"PayloadLength"`
	TotalTransmitPacket int         `xml:"TotalTransmitPacket"`
	TotalReceivePacket  int         `xml:"TotalReceivePacket"`
	LossRate            int         `xml:"LossRate"`
	MinReplyTime        int         `xml:"MinReplyTime"`
	MaxReplyTime        int         `xml:"MaxReplyTime"`
	AvgReplyTime        int         `xml:"AvgReplyTime"`
	StandardDeviation   int         `xml:"StandardDeviation"`
	EchoReplies         []EchoReply `xml:"EchoReply"`
}

type EchoReply struct {
	IcmpSequence int `xml:"IcmpSequence"`
	TTLValue     int `xml:"TTLValue"`
	ReplyTime    int `xml:"ReplyTime"`
}

func (targetDevice *TargetDevice) PingIPv4(host string) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`
			<action>
              <top xmlns="http://www.hp.com/netconf/action:1.0">
                <Ping>
                  <IPv4Ping>
                    <PingTest>
                      <Host>dst_IP</Host>
                      <VRF/>
                    </PingTest>
                  </IPv4Ping>
                </Ping>
              </top>
            </action>
`),
		Xmlns: []string{netconf.BaseURI},
	}

	request.InnerXML = bytes.Replace(request.InnerXML, []byte("dst_IP"), []byte(host), 1)
	targetDevice.RetrieveData(request)
}
