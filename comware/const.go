package comware

const (
	// Common constants
	UnknownString = "Unknown"

	// Interface administration status
	InterfaceAdminStatusUP         = 1         // Admin Up
	InterfaceAdminStatusUPString   = "AdmUp"   // Admin Up
	InterfaceAdminStatusDown       = 2         // Admin Down
	InterfaceAdminStatusDownString = "AdmDown" // Admin Down

	// Interface Speed
	InterfaceSpeedAuto     = 1
	InterfaceSpeed10M      = 2
	InterfaceSpeed100M     = 4
	InterfaceSpeed1G       = 32
	InterfaceSpeed2point5G = 128
	InterfaceSpeed10G      = 1024
	InterfaceSpeed40G      = 8192
	InterfaceSpeed100G     = 16384
	InterfaceSpeed5G       = 32768

	// Interface Duplex
	InterfaceDuplexFull       = 1      // Full Duplex
	InterfaceDuplexFullString = "Full" // Full Duplex
	InterfaceDuplexHalf       = 2      // Half Duplex
	InterfaceDuplexHalfString = "Half" // Half Duplex
	InterfaceDuplexAuto       = 3      // Duplex Auto
	InterfaceDuplexAutoString = "Auto" // Duplex Auto

	// Interface Type (IANAifType)
	InterfaceTypeOther             = 1 // NULL interfaces
	InterfaceTypeEthernet          = 6
	InterfaceTypeLoopback          = 24
	InterfaceTypeVlanInterface     = 136
	InterfaceTypeBridgeAggregation = 161

	// Interface status
	InterfaceStatusUp                   = 1
	InterfaceStatusUpString             = "Up"
	InterfaceStatusDown                 = 2
	InterfaceStatusDownString           = "Down"
	InterfaceStatusTesting              = 3
	InterfaceStatusTestingString        = "Testing"
	InterfaceStatusUnknown              = 4
	InterfaceStatusUnknownString        = "Unknown"
	InterfaceStatusDormant              = 5
	InterfaceStatusDormantString        = "Dormant"
	InterfaceStatusNotPresent           = 6
	InterfaceStatusNotPresentString     = "NotPresent"
	InterfaceStatusLowerLayerDown       = 7
	InterfaceStatusLowerLayerDownString = "LowerLayerDown"

	// Interface Link Type (VLAN type of an interface)
	InterfaceLinkTypeAccess       = 1
	InterfaceLinkTypeAccessString = "Access"
	InterfaceLinkTypeTrunk        = 2
	InterfaceLinkTypeTrunkString  = "Trunk"
	InterfaceLinkTypeHybrid       = 3
	InterfaceLinkTypeHybridString = "Hybrid"

	// Suppression units
	SuppressionUnitRatio       = 1
	SuppressionUnitRatioString = "Ratio"
	SuppressionUnitPps         = 2
	SuppressionUnitPpsString   = "Pps"
	SuppressionUnitKbps        = 3
	SuppressionUnitKbpsString  = "Kbps"

	// Protocols
	ProtocolICMP    = 1
	ProtocolIGMP    = 2
	ProtocolIP      = 4
	ProtocolTCP     = 6
	ProtocolUDP     = 17
	ProtocolGRE     = 47
	ProtocolIPv6esp = 50
	ProtocolIPv6ah  = 51
	ProtocolICMPv6  = 58
	ProtocolOSPF    = 89
	ProtocolAny     = 256

	// Operations
	OperationLess           = 1
	OperationLessString     = "Less"
	OperationEqual          = 2
	OperationEqualString    = "Equal"
	OperationGreater        = 3
	OperationGreaterString  = "Greater"
	OperationNotEqual       = 4
	OperationNotEqualString = "NotEqual"
	OperationRange          = 5
	OperationRangeString    = "Range"

	// ACL apply directions
	PFilterApplyDirectionInbound        = 1
	PfilterApplyDirectionInboundString  = "Inbound"
	PFilterApplyDirectionOutbound       = 2
	PfilterApplyDirectionOutboundString = "Outbound"

	// ACL apply object types
	PFilterAppObjTypeInterface = 1
	PFilterAppObjTypeVlan      = 2
	PFilterAppObjTypeGlobal    = 3

	// ACL group types
	ACLGroupTypeIPv4              = 1
	ACLGroupTypeIPv4String        = "IPv4"
	ACLGroupTypeIPv6              = 2
	ACLGroupTypeIPv6String        = "IPv6"
	ACLGroupTypeMAC               = 3
	ACLGroupTypeMACString         = "MAC"
	ACLGroupTypeUserDefined       = 4
	ACLGroupTypeUserDefinedString = "User-defined"
	ACLGroupTypeDefault           = 5
	ACLGroupTypeDefaultString     = "Default"

	// ACL group Category
	ACLGroupCategoryNone           = 0 // for ACLGroupTypeMAC (3) or ACLGroupTypeUserDefined (4) only
	ACLGroupCategoryNoneString     = "None"
	ACLGroupCategoryBasic          = 1
	ACLGroupCategoryBasicString    = "Basic"
	ACLGroupCategoryAdvanced       = 2
	ACLGroupCategoryAdvancedString = "Advanced"

	// If you set column RuleID to ACLRuleIDAuto, the system automatically assigns a new rule ID.
	ACLRuleIDAuto = 65535

	// ACL rule Actions
	ACLRuleActionDeny         = 1
	ACLRuleActionDenyString   = "Deny"
	ACLRuleActionPermit       = 2
	ACLRuleActionPermitString = "Permit"

	// ACL rule status
	ACLRuleStatusActive         = 1
	AclRuleStatusActiveString   = "Active"
	ACLRuleStatusInactive       = 2
	ACLRuleStatusInactiveString = "Inactive"

	// LAGG modes
	LAGGLinkModeStatic        = 1
	LAGGLinkModeStaticString  = "Static"
	LAGGLinkModeDynamic       = 2
	LAGGLinkModeDynamicString = "Dynamic"

	// LAGG member selected status
	LAGGMemberSelectedStatusSelected         = 1
	LAGGMemberSelectedStatusSelectedString   = "Selected"
	LAGGMemberSelectedStatusUnselected       = 2
	LAGGMemberSelectedStatusUnselectedString = "Unselected"
	LAGGMemberSelectedStatusIndividual       = 3
	LAGGMemberSelectedStatusIndividualString = "Individual"

	// STP modes
	STPModeSTP  = 0
	STPModeRSTP = 2
	STPModeMSTP = 3
	STPModePVST = 4

	// Port filter actions
	PfilterActionPermit       = 1
	PfilterActionPermitString = "Permit"
	PfilterActionDeny         = 2
	PfilterActionDenyString   = "Deny"
)
