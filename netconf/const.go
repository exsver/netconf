package netconf

const (
	messageSeparator = `]]>]]>`
	BaseURI          = "urn:ietf:params:xml:ns:netconf:base:1.0"
	XMLHeader        = `<?xml version="1.0" encoding="UTF-8"?>`
	XMLHello         = `<hello xmlns="urn:ietf:params:xml:ns:netconf:base:1.0"><capabilities><capability>urn:ietf:params:netconf:base:1.0</capability></capabilities></hello>`
	XMLClose         = `<rpc xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="close-uuid"><close-session/></rpc>`
)
