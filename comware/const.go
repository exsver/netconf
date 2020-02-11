package comware

const (
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
)
