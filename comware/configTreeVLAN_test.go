package comware

import "testing"

func TestIfLinkType_String(t *testing.T) {
	tests := []struct {
		name     string
		linkType InterfaceLinkType
		want     string
	}{
		{
			name:     "Access",
			linkType: InterfaceLinkTypeAccess,
			want:     InterfaceLinkTypeAccessString,
		},
		{
			name:     "Trunk",
			linkType: InterfaceLinkTypeTrunk,
			want:     InterfaceLinkTypeTrunkString,
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
