package comware

import "testing"

func TestSrcIPv4_String(t *testing.T) {
	type fields struct {
		SrcIPv4Addr     string
		SrcIPv4Wildcard string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "1",
			fields: fields{
				SrcIPv4Addr:     "10.10.10.0",
				SrcIPv4Wildcard: "0.0.0.255",
			},
			want: "10.10.10.0/24",
		},
		{
			name: "2",
			fields: fields{
				SrcIPv4Addr:     "10.10.0.0",
				SrcIPv4Wildcard: "0.0.127.255",
			},
			want: "10.10.0.0/17",
		},
		{
			name: "3",
			fields: fields{
				SrcIPv4Addr:     "10.0.0.0",
				SrcIPv4Wildcard: "0.255.255.255",
			},
			want: "10.0.0.0/8",
		},
		{
			name: "4",
			fields: fields{
				SrcIPv4Addr:     "0.0.0.0",
				SrcIPv4Wildcard: "255.255.255.255",
			},
			want: "0.0.0.0/0",
		},
		{
			name: "5",
			fields: fields{
				SrcIPv4Addr:     "10.10.10.10",
				SrcIPv4Wildcard: "0.0.0.0",
			},
			want: "10.10.10.10/32",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip := &SrcIPv4{
				SrcIPv4Addr:     tt.fields.SrcIPv4Addr,
				SrcIPv4Wildcard: tt.fields.SrcIPv4Wildcard,
			}

			if got := ip.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wildcardToPrefix(t *testing.T) {
	tests := []struct {
		name            string
		wildcardAddress string
		want            int
		wantErr         bool
	}{
		{
			name:            "0",
			wildcardAddress: "0.0.0.0",
			want:            32,
			wantErr:         false,
		},
		{
			name:            "1",
			wildcardAddress: "0.0.0.1",
			want:            31,
			wantErr:         false,
		},
		{
			name:            "2",
			wildcardAddress: "0.0.0.2",
			want:            -1,
			wantErr:         true,
		},
		{
			name:            "3",
			wildcardAddress: "0.0.0.3",
			want:            30,
			wantErr:         false,
		},
		{
			name:            "4",
			wildcardAddress: "0.0.0.4",
			want:            -1,
			wantErr:         true,
		},
		{
			name:            "5",
			wildcardAddress: "0.0.0.5",
			want:            -1,
			wantErr:         true,
		},
		{
			name:            "6",
			wildcardAddress: "0.0.0.6",
			want:            -1,
			wantErr:         true,
		},
		{
			name:            "7",
			wildcardAddress: "0.0.0.7",
			want:            29,
			wantErr:         false,
		},
		{
			name:            "8",
			wildcardAddress: "0.0.0.255",
			want:            24,
			wantErr:         false,
		},
		{
			name:            "9",
			wildcardAddress: "0.0.1.255",
			want:            23,
			wantErr:         false,
		},
		{
			name:            "10",
			wildcardAddress: "0.0.1.0",
			want:            -1,
			wantErr:         true,
		},
		{
			name:            "11",
			wildcardAddress: "0.255.255.255",
			want:            8,
			wantErr:         false,
		},
		{
			name:            "12",
			wildcardAddress: "255.255.255.255",
			want:            0,
			wantErr:         false,
		},
		{
			name:            "13",
			wildcardAddress: "0.0.0.0.0",
			want:            -1,
			wantErr:         true,
		},
		{
			name:            "14",
			wildcardAddress: "0.0.0",
			want:            -1,
			wantErr:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := wildcardToPrefix(tt.wildcardAddress)
			if (err != nil) != tt.wantErr {
				t.Errorf("wildcardToPrefix() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if got != tt.want {
				t.Errorf("wildcardToPrefix() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestACLRuleStatus_String(t *testing.T) {
	tests := []struct {
		name   string
		status ACLRuleStatus
		want   string
	}{
		{
			name:   "Active",
			status: ACLRuleStatusActive,
			want:   AclRuleStatusActiveString,
		},
		{
			name:   "Inactive",
			status: ACLRuleStatusInactive,
			want:   ACLRuleStatusInactiveString,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestACLRuleAction_String(t *testing.T) {
	tests := []struct {
		name   string
		action ACLRuleAction
		want   string
	}{
		{
			name:   "Permit",
			action: ACLRuleActionPermit,
			want:   ACLRuleActionPermitString,
		},
		{
			name:   "Deny",
			action: ACLRuleActionDeny,
			want:   ACLRuleActionDenyString,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.action.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestACLGroupType_String(t *testing.T) {
	tests := []struct {
		name  string
		gType ACLGroupType
		want  string
	}{
		{
			name:  "IPv4",
			gType: ACLGroupTypeIPv4,
			want:  ACLGroupTypeIPv4String,
		},
		{
			name:  "IPv6",
			gType: ACLGroupTypeIPv6,
			want:  ACLGroupTypeIPv6String,
		},
		{
			name:  "Other",
			gType: 100,
			want:  UnknownString,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gType.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestACLGroupCategory_String(t *testing.T) {
	tests := []struct {
		name     string
		category ACLGroupCategory
		want     string
	}{
		{
			name:     "Basic",
			category: ACLGroupCategoryBasic,
			want:     ACLGroupCategoryBasicString,
		},
		{
			name:     "Advanced",
			category: ACLGroupCategoryAdvanced,
			want:     ACLGroupCategoryAdvancedString,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.category.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceOperStatus_String(t *testing.T) {
	tests := []struct {
		name   string
		status InterfaceOperStatus
		want   string
	}{
		{
			name:   "Up",
			status: InterfaceStatusUp,
			want:   InterfaceStatusUpString,
		},
		{
			name:   "Down",
			status: InterfaceStatusDown,
			want:   InterfaceStatusDownString,
		},
		{
			name:   "Testing",
			status: InterfaceStatusTesting,
			want:   InterfaceStatusTestingString,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceLinkType_String(t *testing.T) {
	tests := []struct {
		name     string
		linkType InterfaceLinkType
		want     string
	}{
		{
			name:     "Access",
			linkType: IfLinkTypeAccess,
			want:     IfLinkTypeAccessString,
		},
		{
			name:     "Trunk",
			linkType: IfLinkTypeTrunk,
			want:     IfLinkTypeTrunkString,
		},
		{
			name:     "Other",
			linkType: 100,
			want:     UnknownString,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.linkType.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceDuplex_String(t *testing.T) {
	tests := []struct {
		name   string
		duplex InterfaceDuplex
		want   string
	}{
		{
			name:   "Full",
			duplex: InterfaceDuplexFull,
			want:   InterfaceDuplexFullString,
		},
		{
			name:   "Half",
			duplex: InterfaceDuplexHalf,
			want:   InterfaceDuplexHalfString,
		},
		{
			name:   "Other",
			duplex: 100,
			want:   UnknownString,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.duplex.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceAdminStatus_String(t *testing.T) {
	tests := []struct {
		name   string
		status InterfaceAdminStatus
		want   string
	}{
		{
			name:   "Up",
			status: InterfaceAdminStatusUP,
			want:   InterfaceAdminStatusUPString,
		},
		{
			name:   "Down",
			status: InterfaceAdminStatusDown,
			want:   InterfaceAdminStatusDownString,
		},
		{
			name:   "Testing",
			status: 300,
			want:   UnknownString,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestACLRulePortOperation_String(t *testing.T) {
	tests := []struct {
		name      string
		operation ACLRulePortOperation
		want      string
	}{
		{
			name:      "Less",
			operation: OperationLess,
			want:      OperationLessString,
		},
		{
			name:      "Equal",
			operation: OperationEqual,
			want:      OperationEqualString,
		},
		{
			name:      "Greater",
			operation: OperationGreater,
			want:      OperationGreaterString,
		},
		{
			name:      "Range",
			operation: 2114,
			want:      UnknownString,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.operation.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
