package comware

import "testing"

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
