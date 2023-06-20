package comware

import "encoding/xml"

type Syslog struct {
	/* top level
	   Syslog
	     Configuration
	     LogBuffer
	     LogHosts
	       []Host
	     Logs                  ***ReadOnly***
	       []Log               ***ReadOnly***
	     OutputRules
	       []OutputRule
	*/
	Configuration *SyslogConfiguration `xml:"Configuration"`
	LogBuffer     *LogBuffer           `xml:"LogBuffer"`
	LogHosts      *LogHosts            `xml:"LogHosts"`
	Logs          *Logs                `xml:"Logs"`
}

// SyslogConfiguration contains information about syslog (the information center).
type SyslogConfiguration struct {
	XMLName xml.Name `xml:"Configuration"`
	// Status of syslog (enable or disable)
	State string `xml:"State,omitempty"`
	// Status of the duplicate log suppression feature
	DuplicateLogSuppression string `xml:"DuplicateLogSuppression,omitempty"`
}

type LogBuffer struct {
	XMLName xml.Name `xml:"LogBuffer"`
	State   string   `xml:"State,omitempty"`
	// BufferSize - Maximum log buffer size configured by the user.
	BufferSize int `xml:"BufferSize,omitempty"`
	// BufferSizeLimit - Maximum log buffer size supported by the device.
	BufferSizeLimit int `xml:"BufferSizeLimit,omitempty"` //***ReadOnly***
	// LogsCount - Number of logs stored in the log buffer.
	LogsCount int `xml:"LogsCount,omitempty"` //***ReadOnly***
	// DroppedLogsCount - Number of dropped logs.
	DroppedLogsCount int `xml:"DroppedLogsCount,omitempty"` //***ReadOnly***
	// OverwrittenLogsCount - Number of overwritten logs.
	OverwrittenLogsCount int                   `xml:"OverwrittenLogsCount,omitempty"` //***ReadOnly***
	LogsCountPerSeverity *LogsCountPerSeverity `xml:"LogsCountPerSeverity"`           //***ReadOnly***
}

//ReadOnly struct
type LogsCountPerSeverity struct {
	XMLName       xml.Name `xml:"LogsCountPerSeverity"`
	Emergency     int      `xml:"Emergency"`
	Alert         int      `xml:"Alert"`
	Critical      int      `xml:"Critical"`
	Error         int      `xml:"Error"`
	Warning       int      `xml:"Warning"`
	Notice        int      `xml:"Notice"`
	Informational int      `xml:"Informational"`
	Debug         int      `xml:"Debug"`
}

// LogHosts table contains log hosts information.
type LogHosts struct {
	XMLName xml.Name  `xml:"LogHosts"`
	Hosts   []LogHost `xml:"Hosts"`
}

type LogHost struct {
	XMLName xml.Name `xml:"Host"`
	Address string   `xml:"Address,omitempty"`
	VRF     string   `xml:"VRF,omitempty"`
	Port    int      `xml:"Port,omitempty"`
	//Logging facility used by the log host, as follows:
	//128 - local0, 136 - local1, 144 - local2,  152 - local3, 160 - local4, 168 - local5, 176 - local6, 184 - local7.
	Facility int `xml:"Facility,omitempty"`
}

// Logs table contains information about the logs in the log buffer (syslog messages).
// ReadOnly struct
type Logs struct {
	XMLName xml.Name `xml:"Logs"`
	Logs    []Log    `xml:"Log"`
}

// ReadOnly struct
type Log struct {
	XMLName xml.Name `xml:"Log"`
	// Index - number of the log as table index.
	Index int `xml:"Index"`
	// Time when the log was generated.
	Time string `xml:"Time"`
	// Group - module that generated the log. String length constraints must be in range(1..8).
	Group string `xml:"Group"`
	// Digest - Brief description of the log, String length constraints must be in range(1..8).
	Digest string `xml:"Digest"`
	// Severity level of the log:
	// 0(Emergency); 1(Alert); 2(Critical); 3(Error); 4(Warning); 5(Notification); 6(Informational); 7(Debugging)
	Severity int `xml:"Severity"`
	// Content of the log. String length constraints must be in range(0..1023).
	Content string `xml:"Content"`
}
