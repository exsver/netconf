package comware

import "github.com/exsver/netconf"

func (targetDevice *TargetDevice) GetDataMGROUP() (*MGROUP, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><MGROUP/></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}
	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}
	return data.Top.MGROUP, nil
}

func (targetDevice *TargetDevice) NewMirroringGroup(id, groupType int) error {
	mirrorGroup := MirrorGroup{
		ID:   id,
		Type: groupType,
	}

	return targetDevice.Configure(*mirrorGroup.ConvertToTop(), "create")
}

func (targetDevice *TargetDevice) RemoveMirroringGroup(id int) error {
	mirrorGroup := MirrorGroup{
		ID:   id,
	}

	return targetDevice.Configure(*mirrorGroup.ConvertToTop(), "remove")
}