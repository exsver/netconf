package netconf

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCorrectBackspaces(t *testing.T) {
	cases := []struct {
		caseDescription string
		in              []byte
		out             []byte
	}{
		{ // 2 backspaces
			caseDescription: "1",
			in:              []byte{0x53, 0x53, 0x8, 0x53, 0x53, 0x53, 0x53, 0x53, 0x8, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53},
			out:             []byte{0x53, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53},
		},
		{
			caseDescription: "2",
			in:              []byte{0x50, 0x51, 0x52, 0x8, 0x8, 0x53, 0x54, 0x55},
			out:             []byte{0x50, 0x53, 0x54, 0x55},
		},
		{
			caseDescription: "3",
			in:              []byte{0x50, 0x51, 0x52, 0x53, 0x54, 0x8, 0x8, 0x8, 0x8, 0x8, 0x55, 0x56},
			out:             []byte{0x55, 0x56},
		},
		{
			caseDescription: "4",
			in:              []byte{0x50, 0x51, 0x52, 0x8},
			out:             []byte{0x50, 0x51},
		},
		{
			caseDescription: "5",
			in:              []byte{0x50, 0x51, 0x52, 0x8, 0x8},
			out:             []byte{0x50},
		},
		{
			caseDescription: "6",
			in:              []byte{0x50, 0x51, 0x52, 0x8, 0x8, 0x8},
			out:             []byte{},
		},
		{
			caseDescription: "7",
			in:              []byte{0x8, 0x50, 0x51},
			out:             []byte{0x50, 0x51},
		},
		{
			caseDescription: "8",
			in:              []byte{0x8, 0x8, 0x50, 0x51},
			out:             []byte{0x50, 0x51},
		},
		{
			caseDescription: "9",
			in:              []byte{0x8, 0x8, 0x50, 0x51, 0x8, 0x8},
			out:             []byte{},
		},
		{
			caseDescription: "10",
			in:              []byte{0x8, 0x8, 0x50, 0x51, 0x52, 0x8, 0x8},
			out:             []byte{0x50},
		},
		{
			caseDescription: "11",
			in:              []byte{0x8, 0x8},
			out:             []byte{},
		},
		{
			caseDescription: "12",
			in:              []byte{0x50, 0x51, 0x8, 0x8, 0x8, 0x8},
			out:             []byte{},
		},
		{
			caseDescription: "13",
			in:              []byte{0x8, 0x8, 0x50},
			out:             []byte{0x50},
		},
		{
			caseDescription: "14",
			in:              []byte{0x41, 0x42, 0x8, 0x43, 0x44, 0x45, 0x8, 0x46, 0x47, 0x48, 0x49, 0x50, 0x51, 0x52, 0x53, 0x54, 0x8, 0x8, 0x8, 0x55, 0x56, 0x57, 0x8, 0x58, 0x59, 0x8},
			out:             []byte{0x41, 0x43, 0x44, 0x46, 0x47, 0x48, 0x49, 0x50, 0x51, 0x55, 0x56, 0x58},
		},
		{
			caseDescription: "15",
			in:              []byte{0x50, 0x51, 0x52},
			out:             []byte{0x50, 0x51, 0x52},
		},
	}
	for _, testCase := range cases {
		out := CorrectBackspaces(testCase.in)
		if !reflect.DeepEqual(out, testCase.out) {
			t.Errorf("CaseDescription: %s\n Got :%s\n Want:%s", testCase.caseDescription, out, testCase.out)
		}
	}
}

func TestNormalize(t *testing.T) {
	cases := []struct {
		caseDescription string
		in              []byte
		out             []byte
	}{
		{
			caseDescription: "1",
			in: []byte(`<rpc>
                          <validate>
                            <source>
                              <candidate/>
                            </source>
                          </validate>
                        </rpc>`),
			out: []byte(`<rpc><validate><source><candidate/></source></validate></rpc>`),
		},
		{
			caseDescription: "2",
			in: []byte(`
                       <rpc message-id="102" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                          <get-sessions/>
                       </rpc>`),
			out: []byte(`<rpc message-id="102" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0"><get-sessions/></rpc>`),
		},
		{
			caseDescription: "3",
			in: []byte(`
<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X53/junos">
<ping-results xmlns="http://xml.juniper.net/junos/15.1X53/junos-probe-tests">
<target-host>
8.8.8.8
</target-host>
<target-ip>
8.8.8.8
</target-ip>
<packet-size>
56
</packet-size>
<probe-result date-determined="1527113979">
<probe-index>
1
</probe-index>
<probe-success/>
<sequence-number>
0
</sequence-number>
<ip-address>
8.8.8.8
</ip-address>
<time-to-live>
48
</time-to-live>
<response-size>
64
</response-size>
<rtt>
33099
</rtt>
</probe-result>
<probe-results-summary>
<probes-sent>
1
</probes-sent>
<responses-received>
1
</responses-received>
<packet-loss>
0
</packet-loss>
<rtt-minimum>
33099
</rtt-minimum>
<rtt-maximum>
33099
</rtt-maximum>
<rtt-average>
33099
</rtt-average>
<rtt-stddev>
0
</rtt-stddev>
</probe-results-summary>
<ping-success/>
</ping-results>
</rpc-reply>`),
			out: []byte(`<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X53/junos"><ping-results xmlns="http://xml.juniper.net/junos/15.1X53/junos-probe-tests"><target-host>8.8.8.8</target-host><target-ip>8.8.8.8</target-ip><packet-size>56</packet-size><probe-result date-determined="1527113979"><probe-index>1</probe-index><probe-success/><sequence-number>0</sequence-number><ip-address>8.8.8.8</ip-address><time-to-live>48</time-to-live><response-size>64</response-size><rtt>33099</rtt></probe-result><probe-results-summary><probes-sent>1</probes-sent><responses-received>1</responses-received><packet-loss>0</packet-loss><rtt-minimum>33099</rtt-minimum><rtt-maximum>33099</rtt-maximum><rtt-average>33099</rtt-average><rtt-stddev>0</rtt-stddev></probe-results-summary><ping-success/></ping-results></rpc-reply>`),
		},
	}

	for _, testCase := range cases {
		out := Normalize(testCase.in)
		if !bytes.Equal(out, testCase.out) {
			t.Errorf("CaseDescription: %s\n Got :%s\n Want:%s", testCase.caseDescription, out, testCase.out)
		}
	}
}

func TestConvertToPairedTags(t *testing.T) {
	cases := []struct {
		caseDescription string
		in              []byte
		out             []byte
	}{
		{
			caseDescription: "1 - simple changes",
			in:              []byte("<if-device-flags><ifdf-present/><ifdf-running/></if-device-flags><if-config-flags><iff-snmp-traps/></if-config-flags>"),
			out:             []byte("<if-device-flags><ifdf-present>true</ifdf-present><ifdf-running>true</ifdf-running></if-device-flags><if-config-flags><iff-snmp-traps>true</iff-snmp-traps></if-config-flags>"),
		},
		{
			caseDescription: "2 - simple changes",
			in: []byte(`
						<arp-table-entry>
    						<mac-address>d8:50:e6:ae:67:28</mac-address>
    						<ip-address>172.21.1.128</ip-address>
							<hostname>172.21.1.128</hostname>
							<interface-name>em0.0</interface-name>
    						<arp-table-entry-flags>
        						<none/>
    						</arp-table-entry-flags>
						</arp-table-entry>`),
			out: []byte(`
						<arp-table-entry>
    						<mac-address>d8:50:e6:ae:67:28</mac-address>
    						<ip-address>172.21.1.128</ip-address>
							<hostname>172.21.1.128</hostname>
							<interface-name>em0.0</interface-name>
    						<arp-table-entry-flags>
        						<none>true</none>
    						</arp-table-entry-flags>
						</arp-table-entry>`),
		},
		{
			caseDescription: "3 - No change",
			in: []byte(`<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X53/junos">
                                       <system-information>
                                         <hardware-model>qfx5200-32c-32q</hardware-model>
                                         <os-name>junos-qfx</os-name>
                                         <os-version>15.1X53-D30.5</os-version>
                                         <serial-number>WH0000000000</serial-number>
                                         <host-name>test-qfx5200</host-name>
                                       </system-information>
                                     </rpc-reply>`),
			out: []byte(`<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X53/junos">
                                       <system-information>
                                         <hardware-model>qfx5200-32c-32q</hardware-model>
                                         <os-name>junos-qfx</os-name>
                                         <os-version>15.1X53-D30.5</os-version>
                                         <serial-number>WH0000000000</serial-number>
                                         <host-name>test-qfx5200</host-name>
                                       </system-information>
                                     </rpc-reply>`),
		},
		{
			caseDescription: "4 - Multilevel",
			in: []byte(`<level1>
                                       <tagA/>
                                       <tagB/>
                                       <tagC/>
                                       <level2>
                                         <tagAA/>
                                         <tagBB/>
                                         <tagCC/>
                                         <level3>
                                           <tagAAA/>
                                           <tagBBB/>
                                           <tagCCC/>
                                         </level3>
                                       </level2>
                                     </level1>`),
			out: []byte(`<level1>
                                       <tagA>true</tagA>
                                       <tagB>true</tagB>
                                       <tagC>true</tagC>
                                       <level2>
                                         <tagAA>true</tagAA>
                                         <tagBB>true</tagBB>
                                         <tagCC>true</tagCC>
                                         <level3>
                                           <tagAAA>true</tagAAA>
                                           <tagBBB>true</tagBBB>
                                           <tagCCC>true</tagCCC>
                                         </level3>
                                       </level2>
                                     </level1>`),
		},
	}

	for _, testCase := range cases {
		out := ConvertToPairedTags(testCase.in)
		if !bytes.Equal(out, testCase.out) {
			t.Errorf("CaseDescription: %s\n Got :%s\n Want:%s", testCase.caseDescription, out, testCase.out)
		}
	}
}

func TestConvertToXML(t *testing.T) {
	cases := []struct {
		caseDescription string
		in              []byte
		out             []byte
	}{
		{
			caseDescription: "1",
			in:              []byte(`level1`),
			out:             []byte(`<level1/>`),
		},
		{
			caseDescription: "2",
			in:              []byte(`level1/level2`),
			out:             []byte(`<level1><level2/></level1>`),
		},
		{
			caseDescription: "3",
			in:              []byte(`level1/level2/level3`),
			out:             []byte(`<level1><level2><level3/></level2></level1>`),
		},
	}

	for _, testCase := range cases {
		out := ConvertToXML(testCase.in)
		if !bytes.Equal(out, testCase.out) {
			t.Errorf("CaseDescription: %s\n Got :%s\n Want:%s", testCase.caseDescription, out, testCase.out)
		}
	}
}

func TestConvertToSelfClosingTag(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		out  []byte
	}{
		{
			in:  []byte("<flags><none>true</none></flags>"),
			out: []byte("<flags><none/></flags>"),
		},
		{
			in:  []byte("<flags><none>true</none><enabled>true</enabled><disabled>true</disabled></flags>"),
			out: []byte("<flags><none/><enabled/><disabled/></flags>"),
		},
		{
			in:  []byte("<flags><none>true</none></flags><options><none>true</none></options>"),
			out: []byte("<flags><none/></flags><options><none/></options>"),
		},
		{
			in:  []byte("<flags><none>true</none><options><none>true</none></options></flags>"),
			out: []byte("<flags><none/><options><none/></options></flags>"),
		},
	}

	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToSelfClosingTag(tt.in); !reflect.DeepEqual(got, tt.out) {
				t.Errorf("ConvertToSelfClosingTag() = %v, want %v", got, tt.out)
			}
		})
	}
}
