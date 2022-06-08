package comware

const (
	// Interface administration status
	InterfaceAdminStatusUP   = 1 // Admin Up
	InterfaceAdminStatusDown = 2 // Admin Down

	// Interface Speed
	InterfaceSpeedAuto = 1
	InterfaceSpeed10M  = 2
	InterfaceSpeed100M = 4
	InterfaceSpeed1G   = 32
	InterfaceSpeed10G  = 1024
	InterfaceSpeed40G  = 8192
	InterfaceSpeed100G = 16384

	// Interface Duplex
	InterfaceDuplexFull = 1 // Full Duplex
	InterfaceDuplexHalf = 2 // Half Duplex
	InterfaceDuplexAuto = 3 // Duplex Auto

	// Interface Link Type (VLAN type of an interface)
	InterfaceLinkTypeAccess = 1
	InterfaceLinkTypeTrunk  = 2
	InterfaceLinkTypeHybrid = 3

	// Suppression units
	SuppressionUnitRatio = 1
	SuppressionUnitPps   = 2
	SuppressionUnitKbps  = 3

	// Protocols
	ProtocolICMP = 1
	ProtocolTCP  = 6
	ProtocolUDP  = 17
	ProtocolAny  = 256

	// Operations
	OperationLess     = 1
	OperationEqual    = 2
	OperationGreater  = 3
	OperationNotEqual = 4
	OperationRange    = 5

	// ACL apply directions
	PFilterApplyDirectionInbound  = 1
	PFilterApplyDirectionOutbound = 2

	// ACL apply object types
	PFilterAppObjTypeInterface = 1
	PFilterAppObjTypeVlan      = 2
	PFilterAppObjTypeGlobal    = 3

	// ACL group types
	ACLGroupTypeIPv4        = 1
	ACLGroupTypeIPv6        = 2
	ACLGroupTypeMAC         = 3
	ACLGroupTypeUserDefined = 4
	ACLGroupTypeDefault     = 5

	// ACL group Category
	ACLGroupCategoryNone     = 0 // for ACLGroupTypeMAC (3) or ACLGroupTypeUserDefined (4) only
	ACLGroupCategoryBasic    = 1
	ACLGroupCategoryAdvanced = 2

	// If you set column RuleID to ACLRuleIDAuto, the system automatically assigns a new rule ID.
	ACLRuleIDAuto = 65535

	// ACL rule Actions
	ACLRuleActionDeny   = 1
	ACLRuleActionPermit = 2

	// STP modes
	STPModeSTP  = 0
	STPModeRSTP = 2
	STPModeMSTP = 3
	STPModePVST = 4
)
