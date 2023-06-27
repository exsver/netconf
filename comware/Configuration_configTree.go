package comware

import "encoding/xml"

type Configuration struct {
	Files *ConfigurationFiles `xml:"Files"`
}

// ConfigurationFiles table contains the configuration files information.
type ConfigurationFiles struct {
	XMLName xml.Name            `xml:"Files"`
	Files   []ConfigurationFile `xml:"File"`
}

type ConfigurationFile struct {
	XMLName    xml.Name `xml:"File"`
	Name       string   `xml:"Name"`
	Running    bool     `xml:"Running"`
	NextMain   bool     `xml:"NextMain"`
	NextBackup bool     `xml:"NextBackup"`
}
