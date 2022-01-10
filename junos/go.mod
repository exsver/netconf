module github.com/exsver/netconf/junos

go 1.15

replace github.com/exsver/netconf/netconf => ../netconf

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/exsver/netconf/netconf v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3
)
