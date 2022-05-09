module github.com/exsver/netconf/junos

go 1.15

replace github.com/exsver/netconf/netconf => ../netconf

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/exsver/netconf/netconf v0.0.0-20220110193026-2de717e22109
	golang.org/x/crypto v0.0.0-20220507011949-2cf3adece122
	golang.org/x/sys v0.0.0-20220503163025-988cb79eb6c6 // indirect
)
