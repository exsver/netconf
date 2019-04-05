package comware

import (
	"bytes"
	"fmt"
	"github.com/exsver/netconf"
)

/* from ACL and QoS Configuration Guide:
ACL Categories                ACL number          Match criteria
Basic ACLs                    2000 to 2999        Source IP address.
Advanced ACLs                 3000 to 3999        Source IP address, destination IP address, packet priority, protocol number, and other Layer 3 and Layer 4 header fields
Ethernet frame header ACLs    4000 to 4999        Layer 2 header fields, such as source and destination MAC addresses, 802.1p priority, and link layer protocol type
*/

func (targetDevice *TargetDevice) GetDataACL() (*ACL, error) {
	request := netconf.RPCMessage{InnerXML: []byte(`
      <get>
        <filter type="subtree">
          <top xmlns="http://www.hp.com/netconf/data:1.0"><ACL/></top>
        </filter>
      </get>`),
		Xmlns: []string{netconf.BaseURI},
	}
	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}
	return data.Top.ACL, nil
}

//Filters examples:
// all items -                        nil
// ACL with name testACL -            []string{`<GroupIndex>testACL</GroupIndex>`}
// ACLs with Description "aclDescr" - []string{`<Description>aclDescr</Description>`}
// All IPv4 ACLs -                    []string{`<GroupType>1</GroupType>`}
// All IPv6 ACLs -                    []string{`<GroupType>2</GroupType>`}
// All advanced ACLs -                []string{`<GroupCategory>2</GroupCategory>`}
// All advanced IPv4 ACLs -           []string{`<GroupCategory>2</GroupCategory>`, `<GroupType>1</GroupType>`}
// All ACLs with rule number 1 -      []string{`<RuleNum>1</RuleNum>`}
func (targetDevice *TargetDevice) GetListOfNamedACL(filters []string) ([]NamedGroup, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><ACL><NamedGroups/></ACL></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	if filters != nil {
		request.InnerXML = []byte(`
      <get>
        <filter type="subtree">
          <top xmlns="http://www.hp.com/netconf/data:1.0">
            <ACL>
              <NamedGroups>
                <Group>
                  <GroupType/>
                  <GroupCategory/>
                  <GroupIndex/>
                  <MatchOrder/>
                  <Step/>
                  <Description/>
                  <RuleNum/>
                </Group>
              </NamedGroups>
            </ACL>
          </top>
        </filter>
      </get>`)
	}
	for _, filter := range filters {
		request.InnerXML = bytes.Replace(request.InnerXML, convertToEmptyTag([]byte(filter)), []byte(filter), 1)
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	if data.Top == nil {
		return nil, fmt.Errorf("no ACLs are found.")
	}
	return data.Top.ACL.NamedGroups.Groups, nil
}

func (targetDevice *TargetDevice) GetListOfIPv4NamedAdvanceRules() ([]IPv4NamedAdvanceRule, error) {
	request := netconf.RPCMessage{InnerXML: []byte(`
      <get>
        <filter type="subtree">
          <top xmlns="http://www.hp.com/netconf/data:1.0">
            <ACL>
              <IPv4NamedAdvanceRules>
                <Rule>
                  <GroupIndex>testACL</GroupIndex>
                </Rule>
              </IPv4NamedAdvanceRules>
            </ACL>
		  </top>
        </filter>
      </get>`),
		Xmlns: []string{netconf.BaseURI},
	}
	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}
	return data.Top.ACL.IPv4NamedAdvanceRules.IPv4NamedAdvanceRules, nil
}

func (targetDevice *TargetDevice) PfilterApply(pfilter *Pfilter) error {
	return targetDevice.Configure(*pfilter.ConvertToTop(), "merge")
}

func (targetDevice *TargetDevice) PfilterRemove(pfilter *Pfilter) error {
	return targetDevice.Configure(*pfilter.ConvertToTop(), "remove")
}

func (namedGroup *NamedGroup) NewPfilter() *Pfilter {
	return &Pfilter{
		AppObjType:   1,
		AppACLType:   namedGroup.GroupType,
		AppACLGroup:  namedGroup.GroupIndex,
		AppDirection: 1,
		HardCount:    2,
	}
}

func (targetDevice *TargetDevice) ACLCreate(acl *NamedGroup) error {
	return targetDevice.Configure(*acl.ConvertToTop(), "create")
}

func (targetDevice *TargetDevice) ACLRemove(acl *NamedGroup) error {
	return targetDevice.Configure(*acl.ConvertToTop(), "remove")
}

func (targetDevice *TargetDevice) GetACLConfig() (*ACL, error) {
	request := netconf.RPCMessage{InnerXML: []byte(`
      <get-config>
        <source>
          <running/>
        </source>
        <filter type="subtree">
          <top xmlns="http://www.hp.com/netconf/config:1.0">
            <ACL>
            </ACL>
          </top>
        </filter>
      </get-config>`),
		Xmlns: []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}
	return data.Top.ACL, nil
}
