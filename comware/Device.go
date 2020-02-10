package comware

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/exsver/netconf/netconf"
)

func (targetDevice *TargetDevice) GetDataDevice() (*Device, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><Device/></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.Device, nil
}

func (targetDevice *TargetDevice) GetExtPhysicalEntities(physicalIndexes []int) ([]ExtPhysicalEntity, error) {
	entities := ""

	for _, v := range physicalIndexes {
		if v == 0 {
			return nil, fmt.Errorf("invalid PhysicalIndex: %v", v)
		}

		entities = fmt.Sprintf("%s<Entity><PhysicalIndex>%v</PhysicalIndex></Entity>", entities, v)
	}

	request := netconf.RPCMessage{
		InnerXML: []byte(`
		<get>
		  <filter type="subtree">
            <top xmlns="http://www.hp.com/netconf/data:1.0">
              <Device>
                <ExtPhysicalEntities>change_data</ExtPhysicalEntities>
              </Device>
            </top>
          </filter>
        </get>`),
		Xmlns: []string{netconf.BaseURI},
	}

	request.InnerXML = bytes.Replace(request.InnerXML, []byte("change_data"), []byte(entities), 1)

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.Device.ExtPhysicalEntities.ExtPhysicalEntities, nil
}

func (targetDevice *TargetDevice) GetSlotsInfo() ([]PhysicalEntity, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`
<get>
  <filter type="subtree">
    <top xmlns="http://www.hp.com/netconf/data:1.0">
      <Device>
        <PhysicalEntities>
          <Entity>
            <PhysicalIndex/><Chassis/><Slot/><SubSlot/><Description/><VendorType/><ContainedIn/>
            <Class>3</Class>
            <ParentRelPos/><Name/><HardwareRev/><FirmwareRev/><SoftwareRev/><SerialNumber/><MfgName/><Model/><Alias/><AssetID/><FRU/><MfgDate/><Uris/>
          </Entity>
        </PhysicalEntities>
      </Device>
    </top>
  </filter>
</get>`),
		Xmlns: []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.Device.PhysicalEntities.PhysicalEntities, nil
}

func (targetDevice *TargetDevice) GetIndexBoards() (index []int, err error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`
<get>
  <filter type="subtree">
    <top xmlns="http://www.hp.com/netconf/data:1.0">
      <Device>
        <Boards>
          <Board>
            <Status>2</Status>
            <PhysicalIndex/>
          </Board>
        </Boards>
      </Device>
    </top>
  </filter>
</get>`),
		Xmlns: []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	for _, board := range data.Top.Device.Boards.Boards {
		index = append(index, board.PhysicalIndex)
	}

	return index, nil
}

//Class:
// 0 - all classes
// 3 - physical switch (slot)
// 5 - CONTAINER ???
// 6 - PSU
// 7 - FAN UNIT
// 8 - Sensor
// 9 - Board, SubBoard
// 10 - Interfaces
// 11 - top container ???
func (targetDevice *TargetDevice) GetPhysicalIndex(class int) (index []int, err error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`
<get>
  <filter type="subtree">
    <top xmlns="http://www.hp.com/netconf/data:1.0">
      <Device>
        <PhysicalEntities>
          <Entity>
            <Class/>
            <PhysicalIndex/>
          </Entity>
        </PhysicalEntities>
      </Device>
    </top>
  </filter>
</get>`),
		Xmlns: []string{netconf.BaseURI},
	}

	if class != 0 {
		request.InnerXML = bytes.Replace(request.InnerXML, []byte("<Class/>"), []byte(fmt.Sprintf("<Class>%s</Class>", strconv.Itoa(class))), 1)
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	for _, sensor := range data.Top.Device.PhysicalEntities.PhysicalEntities {
		index = append(index, sensor.PhysicalIndex)
	}

	return index, nil
}

func (targetDevice *TargetDevice) GetPhysicalIndexSensors() (index []int, err error) {
	return targetDevice.GetPhysicalIndex(8)
}

func (targetDevice *TargetDevice) GetPhysicalIndexPSU() (index []int, err error) {
	return targetDevice.GetPhysicalIndex(6)
}

func (targetDevice *TargetDevice) GetPhysicalIndexFan() (index []int, err error) {
	return targetDevice.GetPhysicalIndex(7)
}
