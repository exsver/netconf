package junos

import "encoding/xml"

type System struct {
	XMLName            xml.Name            `xml:"system"`
	HostName           string              `xml:"host-name,omitempty"`
	TimeZone           string              `xml:"time-zone,omitempty"`
	ARP                *ARP                `xml:"arp"`
	RootAuthentication *RootAuthentication `xml:"root-authentication"`
	NameServer         []NameServer        `xml:"name-server,omitempty"`
	Login              *Login              `xml:"login"`
	Services           *Services           `xml:"services"`
	Syslog             *Syslog             `xml:"syslog"`
	NTP                *NTP                `xml:"ntp"`
}

type ARP struct {
	XMLName    xml.Name `xml:"arp"`
	AgingTimer int      `xml:"aging-timer"`
}

type RootAuthentication struct {
	XMLName           xml.Name `xml:"root-authentication"`
	EncryptedPassword string   `xml:"encrypted-password,omitempty"`
}

type NameServer struct {
	XMLName xml.Name `xml:"name-server"`
	Name    string   `xml:"name"`
}

type Login struct {
	XMLName     xml.Name    `xml:"login"`
	UserClasses []UserClass `xml:"class,omitempty"`
	Users       []User      `xml:"user,omitempty"`
}

type UserClass struct {
	XMLName     xml.Name `xml:"class"`
	Name        string   `xml:"name,omitempty"`
	Permissions []string `xml:"permissions,omitempty"`
}

type User struct {
	XMLName        xml.Name        `xml:"user"`
	Name           string          `xml:"name,omitempty"`
	UID            int             `xml:"uid,omitempty"`
	Class          string          `xml:"class,omitempty"`
	Authentication *Authentication `xml:"authentication,omitempty"`
}

type Authentication struct {
	XMLName           xml.Name `xml:"authentication"`
	EncryptedPassword string   `xml:"encrypted-password,omitempty"`
}

type Services struct {
	XMLName xml.Name `xml:"services"`
}

type Syslog struct {
	XMLName xml.Name `xml:"syslog"`
}

type NTP struct {
	XMLName xml.Name `xml:"ntp"`
}
