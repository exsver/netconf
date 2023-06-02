package comware

import (
	"bytes"
	"fmt"

	"github.com/exsver/netconf/netconf"
)

/* from ACL and QoS Configuration Guide:
ACL Categories                ACL number          Match criteria
Basic ACLs                    2000 to 2999        Source IP address.
Advanced ACLs                 3000 to 3999        Source IP address, destination IP address, packet priority, protocol number, and other Layer 3 and Layer 4 header fields
Ethernet frame header ACLs    4000 to 4999        Layer 2 header fields, such as source and destination MAC addresses, 802.1p priority, and link layer protocol type
*/

func (targetDevice *TargetDevice) GetDataACL() (*ACL, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><ACL/></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.ACL, nil
}

// Filters examples:
// all items -                        nil
// ACL with name testACL -            []string{`<GroupIndex>testACL</GroupIndex>`}
// ACLs with Description "aclDescr" - []string{`<Description>aclDescr</Description>`}
// All IPv4 ACLs -                    []string{`<GroupType>1</GroupType>`}
// All IPv6 ACLs -                    []string{`<GroupType>2</GroupType>`}
// All advanced ACLs -                []string{`<GroupCategory>2</GroupCategory>`}
// All advanced IPv4 ACLs -           []string{`<GroupCategory>2</GroupCategory>`, `<GroupType>1</GroupType>`}
// All ACLs with rule number 1 -      []string{`<RuleNum>1</RuleNum>`}
func (targetDevice *TargetDevice) ACLGetNamedGroups(filters []string) ([]NamedGroup, error) {
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
		return nil, fmt.Errorf("no ACLs are found")
	}

	return data.Top.ACL.NamedGroups.Groups, nil
}

// ACLIPv4NamedAdvanceRulesGet returns ACL rules ([]IPv4NamedAdvanceRule) and error
// Available filters are:
//   - GroupIndex
//   - RuleID
//   - Action
//   - ProtocolType
//   - Fragment
//   - Logging
//   - Counting
//   - Status
//   - SrcIPv4Addr
//   - SrcIPv4Wildcard
//   - DstIPv4Addr
//   - DstIPv4Wildcard
//   - SrcPortOp
//   - SrcPortValue1
//   - SrcPortValue2
//   - DstPortOp
//   - DstPortValue1
//   - DstPortValue2
//   - Comment
//
// Filter examples:
//   - all rules for acl named testACL                           []comware.XMLFilter{{Key: "GroupIndex", Value: "testACL", IsRegExp: false}}
//   - all rules with ProtocolType "ICMP"					    []comware.XMLFilter{{Key: "ProtocolType", Value: strconv.Itoa(comware.ProtocolICMP), IsRegExp: false}}
//   - all rules with Action "Deny"							    []comware.XMLFilter{{Key: "Action", Value: strconv.Itoa(comware.ACLRuleActionDeny), IsRegExp: false}}
//   - all rules with ProtocolType "ICMP" and Action "Deny"      []comware.XMLFilter{{Key: "ProtocolType", Value: strconv.Itoa(comware.ProtocolICMP), IsRegExp: false}, {Key: "Action", Value: strconv.Itoa(comware.ACLRuleActionDeny), IsRegExp: false}}
//   - all rules with DstIPv4Addr matches by regexp "^10.100"    []comware.XMLFilter{{Key: "DstIPv4Addr", Value: "^10.100", IsRegExp: true}}
//   - all rules with Counting                                   []comware.XMLFilter{{Key: "Counting", Value: "true", IsRegExp: false}}
func (targetDevice *TargetDevice) ACLIPv4NamedAdvanceRulesGet(filters []XMLFilter) ([]IPv4NamedAdvanceRule, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><ACL><IPv4NamedAdvanceRules/></ACL></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	if filters != nil {
		request.InnerXML = []byte(`
          <get>
            <filter type="subtree">
              <top xmlns="http://www.hp.com/netconf/data:1.0">
                <ACL>
                  <IPv4NamedAdvanceRules>
                    <Rule>
                      <GroupIndex/>
                      <RuleID/>
                      <Action/>
                      <ProtocolType/>
                      <Count/>
                      <Status/>
                      <Fragment/>
                      <Logging/>
                      <Counting/>
                      <SrcAny/>
                      <DstAny/>
                      <SrcIPv4>
                        <SrcIPv4Addr/>
                        <SrcIPv4Wildcard/> 
                      </SrcIPv4>
                      <DstIPv4>
                        <DstIPv4Addr/>
                        <DstIPv4Wildcard/> 
                      </DstIPv4>
                      <SrcPort>
                        <SrcPortOp/>
                        <SrcPortValue1/>
                        <SrcPortValue2/> 
                      </SrcPort>
                      <DstPort>
                        <DstPortOp/>
                        <DstPortValue1/>
                        <DstPortValue2/> 
                      </DstPort>
                      <Comment/>
                    </Rule>
                  </IPv4NamedAdvanceRules>
                </ACL>
              </top>
            </filter>
          </get>`)

		for _, filter := range filters {
			if !((filter.Key == "GroupIndex") ||
				(filter.Key == "RuleID") ||
				(filter.Key == "Action") ||
				(filter.Key == "ProtocolType") ||
				(filter.Key == "Count") ||
				(filter.Key == "Status") ||
				(filter.Key == "Fragment") ||
				(filter.Key == "Logging") ||
				(filter.Key == "Counting") ||
				(filter.Key == "SrcAny") ||
				(filter.Key == "DstAny") ||
				(filter.Key == "SrcIPv4Addr") ||
				(filter.Key == "SrcIPv4Wildcard ") ||
				(filter.Key == "DstIPv4Addr") ||
				(filter.Key == "DstIPv4Wildcard") ||
				(filter.Key == "SrcPortOp") ||
				(filter.Key == "SrcPortValue1") ||
				(filter.Key == "SrcPortValue2") ||
				(filter.Key == "DstPortOp") ||
				(filter.Key == "DstPortValue1") ||
				(filter.Key == "DstPortValue2") ||
				(filter.Key == "Comment")) {
				return nil, fmt.Errorf("invalid filter: %s", filter.Key)
			}

			request.InnerXML = bytes.Replace(request.InnerXML, []byte(fmt.Sprintf("<%s/>", filter.Key)), filter.ConvertToXML(), 1)
		}
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	// nothing found
	if data.Top == nil {
		return nil, nil
	}

	return data.Top.ACL.IPv4NamedAdvanceRules.IPv4NamedAdvanceRules, nil
}

func (targetDevice *TargetDevice) ACLIPv4NamedAdvanceRulesAdd(rules *IPv4NamedAdvanceRules) error {
	if rules == nil {
		return fmt.Errorf("IPv4NamedAdvanceRules is nil")
	}

	return targetDevice.Configure(*rules.ConvertToTop(), "create")
}

func (targetDevice *TargetDevice) ACLIPv4NamedAdvanceRulesRemove(rules *IPv4NamedAdvanceRules) error {
	if rules == nil {
		return fmt.Errorf("IPv4NamedAdvanceRules is nil")
	}

	if rules.IPv4NamedAdvanceRules == nil {
		return fmt.Errorf("IPv4NamedAdvanceRules slice is nil")
	}

	removeRules := IPv4NamedAdvanceRules{
		IPv4NamedAdvanceRules: make([]IPv4NamedAdvanceRule, 0, len(rules.IPv4NamedAdvanceRules)),
	}

	// When the delete or remove operation is issued, data cannot be assigned to non-index columns
	// Index fields for IPv4NamedAdvanceRule struct are:
	//  - GroupIndex
	//  - RuleID
	for _, rule := range rules.IPv4NamedAdvanceRules {
		removeRules.IPv4NamedAdvanceRules = append(removeRules.IPv4NamedAdvanceRules, IPv4NamedAdvanceRule{
			GroupIndex: rule.GroupIndex,
			RuleID:     rule.RuleID,
		})
	}

	return targetDevice.Configure(*removeRules.ConvertToTop(), "remove")
}

func (targetDevice *TargetDevice) PfilterApply(pfilter *Pfilter) error {
	if pfilter == nil {
		return fmt.Errorf("Pfilter is nil")
	}

	return targetDevice.Configure(*pfilter.ConvertToTop(), "merge")
}

func (targetDevice *TargetDevice) PfilterRemove(pfilter *Pfilter) error {
	if pfilter == nil {
		return fmt.Errorf("Pfilter is nil")
	}

	return targetDevice.Configure(*pfilter.ConvertToTop(), "remove")
}

func (namedGroup *NamedGroup) NewPfilter(applyObjectType, applyObjectIndex int, applyDirection ACLApplyDirection) *Pfilter {
	return &Pfilter{
		AppObjType:   applyObjectType,
		AppObjIndex:  applyObjectIndex,
		AppACLType:   namedGroup.GroupType,
		AppACLGroup:  namedGroup.GroupIndex,
		AppDirection: applyDirection,
	}
}

func (targetDevice *TargetDevice) ACLCreate(acl *NamedGroup) error {
	if acl == nil {
		return fmt.Errorf("NamedGroup is nil")
	}

	return targetDevice.Configure(*acl.ConvertToTop(), "create")
}

func (targetDevice *TargetDevice) ACLRemove(acl *NamedGroup) error {
	if acl == nil {
		return fmt.Errorf("NamedGroup is nil")
	}

	return targetDevice.Configure(*acl.ConvertToTop(), "remove")
}
