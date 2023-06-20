package comware

import "encoding/xml"

type Diagnostic struct {
	/* top level
	Diagnostic
	  CPUAlert
	    []Alert
	  MemoryAlert
	    []Alert
	  Memories
	    []Memory
	  CPUHistory
	    []CPU
	  ProcessesCPUHistory
	    []Processes
	*/
	CPUAlert            *DiagnosticCPUAlerts    `xml:"CPUAlert"`
	MemoryAlert         *DiagnosticMemoryAlerts `xml:"MemoryAlert"`
	Memories            *DiagnosticMemories     `xml:"Memories"`
	CPUHistories        *DiagnosticCPUHistories `xml:"CPUHistory"`
	ProcessesCPUHistory *ProcessesCPUHistory    `xml:"ProcessesCPUHistory"`
}

// DiagnosticCPUAlerts table contains warning threshold information about the logical CPUs.
type DiagnosticCPUAlerts struct {
	XMLName xml.Name             `xml:"CPUAlert"`
	Alerts  []DiagnosticCPUAlert `xml:"Alert"`
}

type DiagnosticCPUAlert struct {
	XMLName           xml.Name `xml:"Alert"`
	Chassis           int      `xml:"Chassis"`
	Slot              int      `xml:"Slot"`
	CPUID             int      `xml:"CPUID"`
	MinorThreshold    int      `xml:"MinorThreshold"`
	WarningThreshold  int      `xml:"WarningThreshold"`
	RecoveryThreshold int      `xml:"RecoveryThreshold"`
	// CurrentState - Current alert state:
	//   Normal
	//   Minor
	//   Severe
	CurrentState string `xml:"CurrentState"`
}

type DiagnosticMemoryAlerts struct {
	XMLName xml.Name                `xml:"MemoryAlert"`
	Alerts  []DiagnosticMemoryAlert `xml:"Alert"`
}

type DiagnosticMemoryAlert struct {
	XMLName               xml.Name `xml:"Alert"`
	Chassis               int      `xml:"Chassis"`
	Slot                  int      `xml:"Slot"`
	CPUID                 int      `xml:"CPUID"`
	MinorThreshold        int      `xml:"MinorThreshold"`
	SevereThreshold       int      `xml:"SevereThreshold"`
	CriticalThreshold     int      `xml:"CriticalThreshold"`
	NormalThreshold       int      `xml:"NormalThreshold"`
	EarlyWarningThreshold int      `xml:"EarlyWarningThreshold"`
	SecureThreshold       int      `xml:"SecureThreshold"`
	// Unit for memory thresholds:
	//   MB
	//   Percentage
	Unit string `xml:"Unit"`
	// CurrentState - Current alert state:
	//   Minor
	//   Severe
	//   Critical
	//   Normal
	//   EarlyWarning
	//   Secure
	CurrentState string `xml:"CurrentState"`
}

// DiagnosticMemories table contains memory information
type DiagnosticMemories struct {
	XMLName  xml.Name           `xml:"Memories"`
	Memories []DiagnosticMemory `xml:"Memory"`
}

type DiagnosticMemory struct {
	XMLName xml.Name `xml:"Memory"`
	Chassis int      `xml:"Chassis"`
	Slot    int      `xml:"Slot"`
	CPUID   int      `xml:"CPUID"`
	Total   int      `xml:"Total"`
	Used    int      `xml:"Used"`
	Free    int      `xml:"Free"`
	// FreeRatio - free memory ratio (%).
	FreeRatio string `xml:"FreeRatio"`
	Buffers   int    `xml:"Buffers"`
	Cached    int    `xml:"Cached"`
	// CommitLimit - memory space currently available for allocation (KB).
	CommitLimit int `xml:"CommitLimit"`
	// CommittedAS - memory space allocated by processes (KB).
	CommittedAS int `xml:"CommittedAS"`
	// VmallocTotal - total allocated virtual memory space (KB)
	VmallocTotal int `xml:"VmallocTotal"`
	// VmallocUsed - used virtual memory space (KB).
	VmallocUsed int `xml:"VmallocUsed"`
	// VmallocChunk - free virtual memory space (KB).
	VmallocChunk int `xml:"VmallocChunk"`
	// SwapTotal  - total size of the memory space used for swapping (KB).
	SwapTotal int `xml:"SwapTotal"`
	// SwapFree - size of the free memory space used for swapping (KB).
	SwapFree  int `xml:"SwapFree"`
	LowTotal  int `xml:"LowTotal"`
	LowFree   int `xml:"LowFree"`
	HighTotal int `xml:"HighTotal"`
	HighFree  int `xml:"HighFree"`
}

// DiagnosticCPUHistories table contains usage information about the logical CPUs.
type DiagnosticCPUHistories struct {
	XMLName xml.Name               `xml:"CPUHistory"`
	CPU     []DiagnosticCPUHistory `xml:"CPU"`
}

type DiagnosticCPUHistory struct {
	XMLName              xml.Name `xml:"CPU"`
	Chassis              int      `xml:"Chassis"`
	Slot                 int      `xml:"Slot"`
	CPUID                int      `xml:"CPUID"`
	CPUUsage             int      `xml:"CPUUsage"`
	Last1mUsage          int      `xml:"Last1mUsage"`
	Last5mUsage          int      `xml:"Last5mUsage"`
	MoniterEnable        int      `xml:"MoniterEnable"`
	HistroyCycleInSecond int      `xml:"HistroyCycleInSecond"`
	CPUUsageLimit        int      `xml:"CPUUsageLimit"`
	CPUHistory           string   `xml:"CPUHistory"`
}

// ProcessesCPUHistory table contains historical CPU usage statistics by process ID.
type ProcessesCPUHistory struct {
	XMLName   xml.Name            `xml:"ProcessesCPUHistory"`
	Processes []ProcessCPUHistory `xml:"Process"`
}

type ProcessCPUHistory struct {
	XMLName xml.Name `xml:"Process"`
	Chassis int      `xml:"Chassis"`
	Slot    int      `xml:"Slot"`
	CPUID   int      `xml:"CPUID"`
	// ProcessJobID - Process ID
	ProcessJobID      int    `xml:"ProcessJobID"`
	ProcessName       string `xml:"ProcessName"`
	CPUHistoryProcess string `xml:"CPUHistoryProcess"`
}
