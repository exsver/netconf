module github.com/exsver/netconf/rawxml

go 1.13

replace github.com/exsver/netconf/netconf => ../netconf

require (
	github.com/exsver/netconf/netconf v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20200208060501-ecb85df21340
)
