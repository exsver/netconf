module github.com/exsver/netconf/comware

go 1.13

replace github.com/exsver/netconf/netconf => ../netconf

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/exsver/netconf/netconf v0.0.0-20200311172139-4100e46a7c26
	golang.org/x/crypto v0.0.0-20200208060501-ecb85df21340
)
