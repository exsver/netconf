package comware

import (
	"errors"
	"reflect"
	"testing"
)

func TestVlanListToIntSlice(t *testing.T) {
	cases := []struct {
		caseDescription string
		in              string
		out             []int //result
		err             error //result
	}{
		{
			caseDescription: "1",
			in:              "99,203-206,213,240",
			out:             []int{99, 203, 204, 205, 206, 213, 240},
			err:             nil,
		},
		{
			caseDescription: "2",
			in:              "10-12,20-25,30-33",
			out:             []int{10, 11, 12, 20, 21, 22, 23, 24, 25, 30, 31, 32, 33},
			err:             nil,
		},
		{
			caseDescription: "3",
			in:              "99-110",
			out:             []int{99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110},
			err:             nil,
		},
		{
			caseDescription: "4",
			in:              "10",
			out:             []int{10},
			err:             nil,
		},
		{
			caseDescription: "5",
			in:              "",
			out:             nil,
			err:             nil,
		},
		{
			caseDescription: "6",
			in:              "10-9",
			out:             nil,
			err:             errors.New("error while trying to convert a string into a int slice. min>max: 10-9"),
		},
		{
			caseDescription: "7",
			in:              "1,5-3,10",
			out:             nil,
			err:             errors.New("error while trying to convert a string into a int slice. min>max: 5-3"),
		},
		{
			caseDescription: "8",
			in:              "ttttt",
			out:             nil,
			err:             errors.New(`error while trying to convert a string into a int slice: strconv.Atoi: parsing "ttttt": invalid syntax`),
		},
		{
			caseDescription: "9",
			in:              "10,20,30,ttt",
			out:             []int{10, 20, 30},
			err:             errors.New(`error while trying to convert a string into a int slice: strconv.Atoi: parsing "ttt": invalid syntax`),
		},
	}
	for _, testCase := range cases {
		out, err := VlanListToIntSlice(testCase.in)
		if !reflect.DeepEqual(out, testCase.out) || !reflect.DeepEqual(err, testCase.err) {
			t.Errorf("CaseDescription: %s\n Got :%v %t\n Want:%v %t\n", testCase.caseDescription, out, err, testCase.out, testCase.err)
		}
	}
}

func TestParseVlansFromConfigString(t *testing.T) {
	cases := []struct {
		caseDescription string
		in              string
		out             []int
	}{
		{
			caseDescription: "1",
			in:              `port trunk permit vlan 105 to 110 120`,
			out:             []int{105, 106, 107, 108, 109, 110, 120},
		},
		{
			caseDescription: "2",
			in:              `port access vlan 99`,
			out:             []int{99},
		},
		{
			caseDescription: "3",
			in:              `port hybrid vlan 140 untagged`,
			out:             []int{140},
		},
		{
			caseDescription: "4",
			in:              `port hybrid vlan 140 tagged`,
			out:             []int{140},
		},
		{
			caseDescription: "5",
			in:              "port link-type trunk\nundo port trunk permit vlan 1\nport trunk permit vlan 105 to 110 120",
			out:             []int{105, 106, 107, 108, 109, 110, 120},
		},
		{
			caseDescription: "6",
			in:              "interface GigabitEthernet1/0/1\n port access vlan 99\n",
			out:             []int{99},
		},
		{
			caseDescription: "7",
			in:              "#\ninterface GigabitEthernet1/0/1\n port access vlan 99\n#\n",
			out:             []int{99},
		},
		{
			caseDescription: "8",
			in:              "#\n\n\ninterface GigabitEthernet1/0/1\n\n\n port access vlan 100\n#\n\n\n",
			out:             []int{100},
		},
		{
			caseDescription: "9",
			in:              "#\ninterface GigabitEthernet1/0/1\n port access vlan 99\n#\nreturn\n",
			out:             []int{99},
		},
		{
			caseDescription: "10",
			in: `#
interface GE1/0/10
 default
 description PXEbootPort
 port link-type trunk
 undo port trunk permit vlan 1
 port trunk permit vlan 302
 port trunk pvid vlan 302
 broadcast-suppression pps 100
 multicast-suppression pps 100
 unicast-suppression pps 2000
 speed auto
 bpdu-drop any
 ipv6 nd raguard role host
 loopback-detection enable vlan all
 loopback-detection action shutdown
#`,
			out: []int{302},
		},
	}
	for _, testCase := range cases {
		out := ParseVlansFromConfigString(testCase.in)
		if !reflect.DeepEqual(out, testCase.out) {
			t.Errorf("CaseDescription: %s\n Got :%v\n Want:%v\n", testCase.caseDescription, out, testCase.out)
		}
	}
}
