package comware

import "encoding/xml"

type Device struct {
	/* top level
	Device
	  Base
	    TimeZone
	  Boards						 ***ReadOnly***
	    []Board                      ***ReadOnly***
	  CPUs                           ***ReadOnly***
	    []CPU                        ***ReadOnly***
	  SummerTime
	    WeekBased
	  FanDirections
	    []Fan
	  PhysicalEntities               ***ReadOnly***
	    []Entity                     ***ReadOnly***
	  ExtPhysicalEntities	         ***ReadOnly***
	    []ExtPhysicalEntity			 ***ReadOnly***
	  TemperatureSensors			 ***ReadOnly***
	    []Sensor					 ***ReadOnly***
	  Transceivers                   ***ReadOnly***
	    []Interface					 ***ReadOnly***
	  TransceiversChannels			 ***ReadOnly***   use for FortyGigE interfaces
	    []Interface					 ***ReadOnly***
	*/
	Base                 *Base                 `xml:"Base"`
	Boards               *Boards               `xml:"Boards"`
	CPUs                 *CPUs                 `xml:"CPUs"`
	PhysicalEntities     *PhysicalEntities     `xml:"PhysicalEntities"`
	ExtPhysicalEntities  *ExtPhysicalEntities  `xml:"ExtPhysicalEntities"`
	FanDirections        *FanDirections        `xml:"FanDirections"`
	TemperatureSensors   *TemperatureSensors   `xml:"TemperatureSensors"`
	Transceivers         *Transceivers         `xml:"Transceivers"`
	TransceiversChannels *TransceiversChannels `xml:"TransceiversChannels"`
	SummerTime           *SummerTime           `xml:"SummerTime"`
}

type Base struct {
	XMLName         xml.Name `xml:"Base"`
	Uptime          uint64   `xml:"Uptime,omitempty"`          //Time that elapsed after the device started up(units:second)
	HostName        string   `xml:"HostName,omitempty"`        //Device name. String length constraints must be in range(0..64). String must not begin and end with space character
	HostDescription string   `xml:"HostDescription,omitempty"` //Device description.String length constraints must be in range(0..255).
	LocalTime       string   `xml:"LocalTime,omitempty"`
	BridgeMAC       string   `xml:"BridgeMAC,omitempty"`
	LocalBridgeMAC  string   `xml:"LocalBridgeMAC,omitempty"`
	MinChassisNum   int      `xml:"MinChassisNum,omitempty"`
	MaxChassisNum   int      `xml:"MaxChassisNum,omitempty"`
	MinSlotNum      int      `xml:"MinSlotNum,omitempty"`
	MaxSlotNum      int      `xml:"MaxSlotNum,omitempty"`
	MinCPUIDNum     int      `xml:"MinCPUIDNum,omitempty"`
	MaxCPUIDNum     int      `xml:"MaxCPUIDNum,omitempty"`
}

// Boards table contains board information.
type Boards struct {
	XMLName xml.Name `xml:"Boards"`
	Boards  []Board  `xml:"Board"`
}

type Board struct {
	XMLName       xml.Name    `xml:"Board"`
	DeviceNode    *DeviceNode `xml:"DeviceNode"`
	PhysicalIndex int         `xml:"PhysicalIndex"`
	Status        int         `xml:"Status"` //Status of the device node: 1 Absent, 2 Normal, 3 Fault
	Role          int         `xml:"Role"`   //Role of the device node: 1 - Unknown, 2 - ActiveMPU, 3 - StandbyMPU, 4 - LPU. Role is Unknown when status is Absent or Fault.
}

type CPUs struct {
	XMLName xml.Name `xml:"CPUs"`
	CPUs    []CPU    `xml:"CPU"`
}

type CPU struct {
	XMLName       xml.Name `xml:"CPU"`
	PhysicalIndex int      `xml:"PhysicalIndex"`
	Chassis       int      `xml:"Chassis"`
	Slot          int      `xml:"Slot"`
	CPUID         int      `xml:"CPUID"`
	CPUUsage      int      `xml:"CPUUsage"`
}

// PhysicalEntities table contains basic information about an entity.
// ReadOnly struct
type PhysicalEntities struct {
	XMLName          xml.Name         `xml:"PhysicalEntities"`
	PhysicalEntities []PhysicalEntity `xml:"Entity"`
}

//ReadOnly struct
type PhysicalEntity struct {
	XMLName       xml.Name `xml:"Entity"`
	PhysicalIndex int      `xml:"PhysicalIndex"`
	Chassis       int      `xml:"Chassis"`
	Slot          int      `xml:"Slot"`
	SubSlot       int      `xml:"SubSlot"`
	Description   string   `xml:"Description"`
	VendorType    string   `xml:"VendorType"`
	ContainedIn   int      `xml:"ContainedIn"`
	Class         int      `xml:"Class"` //3 - physical node (slot),  6 - PSU,  7 - FAN UNIT, 8 - Sensor, 9 - Board, SubBoard, 10 - Interfaces
	ParentRelPos  int      `xml:"ParentRelPos"`
	Name          string   `xml:"Name"`
	HardwareRev   string   `xml:"HardwareRev"`
	FirmwareRev   string   `xml:"FirmwareRev"`
	SoftwareRev   string   `xml:"SoftwareRev"`
	SerialNumber  string   `xml:"SerialNumber"`
	MfgName       string   `xml:"MfgName"`
	Model         string   `xml:"Model"`
	Alias         string   `xml:"Alias"`
	AssetID       string   `xml:"AssetID"`
	FRU           bool     `xml:"FRU"`
	MfgDate       string   `xml:"MfgDate"`
	Uris          string   `xml:"Uris"`
}

// ExtPhysicalEntities table contains extended information about an entity.
type ExtPhysicalEntities struct {
	XMLName             xml.Name            `xml:"ExtPhysicalEntities"`
	ExtPhysicalEntities []ExtPhysicalEntity `xml:"Entity"`
}

type ExtPhysicalEntity struct {
	XMLName                      xml.Name `xml:"Entity"`
	PhysicalIndex                int      `xml:"PhysicalIndex"`
	Uptime                       int      `xml:"Uptime"`                       //Time that elapsed after the entity started up (units:seconds)
	AdminState                   int      `xml:"AdminState"`                   //1 - Unsupported, 2 - Disabled, 3 - Shut down, 4 - Enabled
	OperState                    int      `xml:"OperState"`                    //1 - Unsupported, 2 - Unusable, 3 - Usable, 4 - Dangerous, cannot be used
	StandbyState                 int      `xml:"StandbyState"`                 //1 - Unsupported, 2 - Hot standby, 3 - Cold standby, 4 - In service
	AlarmLight                   int      `xml:"AlarmLight"`                   //Alarm LED status
	CPUUsage                     int      `xml:"CpuUsage"`                     //CPU use ratio (percentage)
	CPUMaxUsage                  int      `xml:"CpuMaxUsage"`                  //Maximum CPU use ratio (percentage)
	CPUAvgUsage                  int      `xml:"CpuAvgUsage"`                  //Average CPU usage (percentage)
	CPUUsageThreshold            int      `xml:"CpuUsageThreshold"`            //CPU use ratio threshold (percentage)
	MemUsage                     int      `xml:"MemUsage"`                     //Memory use ratio (percentage)
	MemAvgUsage                  int      `xml:"MemAvgUsage"`                  //Average memory usage (percentage)
	MemUsageThreshold            int      `xml:"MemUsageThreshold"`            //Memory use ratio threshold (percentage)
	MemSize                      int      `xml:"MemSize"`                      //Memory size (units: bytes)
	PhyMemSize                   int      `xml:"PhyMemSize"`                   //Physical memory size (units: bytes)
	Temperature                  int      `xml:"Temperature"`                  //Temperature of the entity (units: C)
	TemperatureThreshold         int      `xml:"TemperatureThreshold"`         //High-temperature threshold (units: C)
	TemperatureCriticalThreshold int      `xml:"TemperatureCriticalThreshold"` //Critical high-voltage threshold (units: C)
	TemperatureLowThreshold      int      `xml:"TemperatureLowThreshold"`      //Low temperature threshold (units: C)
	TemperatureShutdownThreshold int      `xml:"TemperatureShutdownThreshold"` //Shutdown temperature threshold (units: C)
	ErrorStatus                  int      `xml:"ErrorStatus"`                  //Error status of the entity. Numerial value must be more than 0
	MAC                          string   `xml:"MAC"`
}

type FanDirections struct {
	XMLName xml.Name `xml:"FanDirections"`
	Fans    []Fan    `xml:"Fan"`
}

type Fan struct {
	XMLName          xml.Name `xml:"Fan"`
	Chassis          int      `xml:"Chassis,omitempty"`
	Slot             int      `xml:"Slot,omitempty"`
	CPUID            int      `xml:"CPUID,omitempty"`
	Direction        int      `xml:"Direction,omitempty"`
	DefaultDirection int      `xml:"DefaultDirection,omitempty"`
}

// TemperatureSensors table contains the temperature sensor information.
type TemperatureSensors struct {
	XMLName xml.Name `xml:"TemperatureSensors"`
	Sensors []Sensor `xml:"Sensor"`
}

type Sensor struct {
	XMLName      xml.Name `xml:"Sensor"`
	Chassis      int      `xml:"Chassis"`
	Slot         int      `xml:"Slot"`
	CPUID        int      `xml:"CPUID"`
	SensorType   int      `xml:"SensorType"`
	SensorIndex  int      `xml:"SensorIndex"`
	LowerLimit   int      `xml:"LowerLimit"`
	WarningLimit int      `xml:"WarningLimit"`
	AlarmLimit   int      `xml:"AlarmLimit"`
}

type Transceivers struct {
	XMLName    xml.Name               `xml:"Transceivers"`
	Interfaces []TransceiverInterface `xml:"Interface"`
}

type TransceiverInterface struct {
	XMLName            xml.Name `xml:"Interface"`
	IfIndex            int      `xml:"IfIndex"`
	Name               string   `xml:"Name"`
	HardwareType       string   `xml:"HardwareType"`
	TransceiverType    string   `xml:"TransceiverType"`
	RevisionNumber     string   `xml:"RevisionNumber"`
	TransceiverErrors  string   `xml:"TransceiverErrors"`
	WaveLength         string   `xml:"WaveLength"`
	VendorName         string   `xml:"VendorName"`
	SerialNumber       string   `xml:"SerialNumber"`
	FiberDiameterType  int      `xml:"FiberDiameterType"`
	TransferDistance   int      `xml:"TransferDistance"`
	Diagnostic         int      `xml:"Diagnostic"`
	CurRxPower         string   `xml:"CurRxPower"`
	MaxRxPower         string   `xml:"MaxRxPower"`
	MinRxPower         string   `xml:"MinRxPower"`
	CurTxPower         string   `xml:"CurTxPower"`
	MaxTxPower         string   `xml:"MaxTxPower"`
	MinTxPower         string   `xml:"MinTxPower"`
	Temperature        string   `xml:"Temperature"`
	Voltage            string   `xml:"Voltage"`
	BiasCurrent        string   `xml:"BiasCurrent"`
	TemperatureHiAlarm string   `xml:"TemperatureHiAlarm"`
	TemperatureLoAlarm string   `xml:"TemperatureLoAlarm"`
	TemperatureHiWarn  string   `xml:"TemperatureHiWarn"`
	TemperatureLoWarn  string   `xml:"TemperatureLoWarn"`
	VccHiAlarm         string   `xml:"VccHiAlarm"`
	VccLoAlarm         string   `xml:"VccLoAlarm"`
	VccHiWarn          string   `xml:"VccHiWarn"`
	VccLoWarn          string   `xml:"VccLoWarn"`
	BiasHiAlarm        string   `xml:"BiasHiAlarm"`
	BiasLoAlarm        string   `xml:"BiasLoAlarm"`
	BiasHiWarn         string   `xml:"BiasHiWarn"`
	BiasLoWarn         string   `xml:"BiasLoWarn"`
	PwrOutHiAlarm      string   `xml:"PwrOutHiAlarm"`
	PwrOutLoAlarm      string   `xml:"PwrOutLoAlarm"`
	PwrOutHiWarn       string   `xml:"PwrOutHiWarn"`
	PwrOutLoWarn       string   `xml:"PwrOutLoWarn"`
	RcvPwrHiAlarm      string   `xml:"RcvPwrHiAlarm"`
	RcvPwrLoAlarm      string   `xml:"RcvPwrLoAlarm"`
	RcvPwrHiWarn       string   `xml:"RcvPwrHiWarn"`
	RcvPwrLoWarn       string   `xml:"RcvPwrLoWarn"`
	VendorOUI          string   `xml:"VendorOUI"`
	Frequency          string   `xml:"Frequency"`
	ActiveITUChannel   string   `xml:"ActiveITUChannel"`
	CurWaveErr         string   `xml:"CurWaveErr"`
	WaveErrHiAlarm     string   `xml:"WaveErrHiAlarm"`
	WaveErrLoAlarm     string   `xml:"WaveErrLoAlarm"`
	CurFreqErr         string   `xml:"CurFreqErr"`
	FreqErrHiAlarm     string   `xml:"FreqErrHiAlarm"`
	FreqErrLoAlarm     string   `xml:"FreqErrLoAlarm"`
}

type TransceiversChannels struct {
	XMLName    xml.Name                         `xml:"TransceiversChannels"`
	Interfaces []TransceiversChannelsInterfaces `xml:"Interface"`
}

type TransceiversChannelsInterfaces struct {
	XMLName            xml.Name `xml:"Interface"`
	IfIndex            int      `xml:"IfIndex"`
	ChannelIndex       int      `xml:"ChannelIndex"`
	ChannelCurTXPower  string   `xml:"ChannelCurTXPower"`
	ChannelCurRXPower  string   `xml:"ChannelCurRXPower"`
	ChannelTemperature string   `xml:"ChannelTemperature"`
	ChannelBiasCurrent string   `xml:"ChannelBiasCurrent"`
}

type SummerTime struct {
	XMLName   xml.Name   `xml:"SummerTime"`
	Name      string     `xml:"Name,omitempty"`
	AddTime   string     `xml:"AddTime,omitempty"`
	WeekBased *WeekBased `xml:"WeekBased"`
}

type WeekBased struct {
	XMLName      xml.Name `xml:"WeekBased"`
	BeginMonth   int      `xml:"BeginMonth,omitempty"`
	BeginWeek    int      `xml:"BeginWeek,omitempty"`
	BeginWeekDay int      `xml:"BeginWeekDay,omitempty"`
	BeginHour    int      `xml:"BeginHour,omitempty"`
	BeginMinute  int      `xml:"BeginMinute,omitempty"`
	BeginSecond  int      `xml:"BeginSecond,omitempty"`
	EndMonth     int      `xml:"EndMonth,omitempty"`
	EndWeek      int      `xml:"EndWeek,omitempty"`
	EndWeekDay   int      `xml:"EndWeekDay,omitempty"`
	EndHour      int      `xml:"EndHour,omitempty"`
	EndMinute    int      `xml:"EndMinute,omitempty"`
	EndSecond    int      `xml:"EndSecond,omitempty"`
}
