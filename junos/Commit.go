package junos

import (
	"encoding/xml"

	"github.com/exsver/netconf/netconf"
)

type CommitHistory struct {
	Name               xml.Name            `xml:"commit-information"`
	CommitHistoryItems []CommitHistoryItem `xml:"commit-history"`
}

type CommitHistoryItem struct {
	SequenceNumber int    `xml:"sequence-number"`
	User           string `xml:"user"`
	Client         string `xml:"client"` // cli|netconf|other
	DateTime       string `xml:"date-time"`
	Comment        string `xml:"comment"`
}

//CLI equivalent: show system commit
func (targetDevice *TargetDevice) GetSystemCommit() ([]CommitHistoryItem, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get-commit-information/>`),
	}

	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return []CommitHistoryItem{}, err
	}

	if rpcReply.Error() != nil {
		return []CommitHistoryItem{}, rpcReply.Error()
	}

	var history CommitHistory
	err = xml.Unmarshal(rpcReply.Content, &history)

	return history.CommitHistoryItems, err
}

type CommitResult struct {
	XMLName       xml.Name      `xml:"outer-tag"`
	CommitResults CommitResults `xml:"commit-results"`
	OK            bool          `xml:"ok"`
}

type CommitResults struct {
	Errors []netconf.RPCError `xml:"rpc-error"`
	REs    []REReport         `xml:"routing-engine"`
}

type REReport struct {
	Name               string `xml:"name"`
	CommitCheckSuccess bool   `xml:"commit-check-success"`
	CommitSuccess      bool   `xml:"commit-success"`
}

//CLI equivalent: commit check
func (targetDevice *TargetDevice) CommitCheck() (CommitResult, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<validate><source><candidate/></source></validate>`),
	}

	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return CommitResult{}, err
	}

	if rpcReply.Error() != nil {
		return CommitResult{}, rpcReply.Error()
	}

	rpcReply.Content = netconf.ConvertToPairedTags(rpcReply.Content)
	rpcReply.Content = append([]byte("<outer-tag>"), append(rpcReply.Content, []byte("</outer-tag>")...)...)

	var commitResult CommitResult

	err = xml.Unmarshal(rpcReply.Content, &commitResult)

	return commitResult, err
}

//CLI equivalent: commit
func (targetDevice *TargetDevice) Commit() (CommitResult, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<commit/>`),
	}

	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return CommitResult{}, err
	}

	if rpcReply.Error() != nil {
		return CommitResult{}, rpcReply.Error()
	}

	rpcReply.Content = netconf.ConvertToPairedTags(rpcReply.Content)
	rpcReply.Content = append([]byte("<outer-tag>"), append(rpcReply.Content, []byte("</outer-tag>")...)...)

	var commitResult CommitResult

	err = xml.Unmarshal(rpcReply.Content, &commitResult)

	return commitResult, err
}
