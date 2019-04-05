package comware

import (
	"encoding/xml"
	"fmt"
	"github.com/exsver/netconf"
	"strconv"
)

//CLI equivalent "save force"
func (targetDevice *TargetDevice) SaveForce() error {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<save/>`),
		Xmlns:    []string{netconf.BaseURI},
	}
	return targetDevice.PerformAction(request)
}

//The name of the specified configuration file must end with the extension .cfg.
//The total length of the save path and file name must be no more than 191 characters.
func (targetDevice *TargetDevice) Save(fileName string) error {
	request := netconf.RPCMessage{
		InnerXML: []byte(fmt.Sprintf(`<save><file>%s</file></save>`, fileName)),
		Xmlns:    []string{netconf.BaseURI},
	}
	return targetDevice.PerformAction(request)
}

type SavePoint struct {
	XMLName           xml.Name             `xml:"save-point"`
	Commit            *Commit              `xml:"commit"`
	CommitInformation *[]CommitInformation `xml:"commit-information"`
}

type Commit struct {
	XMLName  xml.Name `xml:"commit"`
	CommitID int      `xml:"commit-id,omitempty"`
	Label    string   `xml:"label,omitempty"`
	Comment  string   `xml:"comment,omitempty"`
}

type CommitInformation struct {
	XMLName   xml.Name `xml:"commit-information"`
	CommitID  int      `xml:"commit-id"`
	TimeStamp string   `xml:"TimeStamp"`
	UserName  string   `xml:"UserName"`
	Label     string   `xml:"label"`
	Comment   string   `xml:"comment"`
}

//The confirmTimeout parameter specifies the rollback idle timeout time in the range of 1 to 65535 seconds.
func (targetDevice *TargetDevice) SavePointBegin(confirmTimeout int) (commitID int, err error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(fmt.Sprintf(`<save-point><begin><confirm-timeout>%s</confirm-timeout></begin></save-point>`, strconv.Itoa(confirmTimeout))),
		Xmlns:    []string{netconf.BaseURI},
	}
	data, err := targetDevice.RetrieveData(request)
	return data.SavePoint.Commit.CommitID, err
}

//The system supports a maximum of 50 rollback points.
func (targetDevice *TargetDevice) SavePointEnd() error {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<save-point><end/></save-point>`),
		Xmlns:    []string{netconf.BaseURI},
	}
	return targetDevice.PerformAction(request)
}

func (targetDevice *TargetDevice) SavePointGetCommits() (*SavePoint, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<save-point><get-commits/></save-point>`),
		Xmlns:    []string{netconf.BaseURI},
	}
	data, err := targetDevice.RetrieveData(request)
	return data.SavePoint, err
}

func (targetDevice *TargetDevice) SavePointCommit(label string, comment string) error {
	request := netconf.RPCMessage{
		InnerXML: []byte(fmt.Sprintf(`<save-point><commit><label>%s</label><comment>%s</comment></commit></save-point>`, label, comment)),
		Xmlns:    []string{netconf.BaseURI},
	}
	return targetDevice.PerformAction(request)
}

func (targetDevice *TargetDevice) SavePointRollback(commitID int) error {
	request := netconf.RPCMessage{
		InnerXML: []byte(fmt.Sprintf(`<save-point><rollback><commit-id>%v</commit-id></rollback></save-point>`, commitID)),
		Xmlns:    []string{netconf.BaseURI},
	}
	return targetDevice.PerformAction(request)
}

func (targetDevice *TargetDevice) SavePointRollbackByLabel(label string) error {
	request := netconf.RPCMessage{
		InnerXML: []byte(fmt.Sprintf(`<save-point><rollback><commit-label>%s</commit-label></rollback></save-point>`, label)),
		Xmlns:    []string{netconf.BaseURI},
	}
	return targetDevice.PerformAction(request)
}
