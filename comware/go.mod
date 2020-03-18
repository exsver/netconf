module github.com/exsver/netconf/comware

go 1.13

replace github.com/exsver/netconf/netconf => ../netconf

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/exsver/netconf/netconf v0.0.0-20200318094214-dc404f622392
	golang.org/x/crypto v0.0.0-20200317142112-1b76d66859c6
	golang.org/x/sys v0.0.0-20200317113312-5766fd39f98d // indirect
)
