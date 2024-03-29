package comware

func (acl *Group) ConvertToTop() *Top {
	return &Top{
		ACL: &ACL{
			Groups: &Groups{
				Groups: []Group{*acl},
			},
		},
	}
}

func (namedGroup *NamedGroup) ConvertToTop() *Top {
	return &Top{
		ACL: &ACL{
			NamedGroups: &NamedGroups{
				Groups: []NamedGroup{*namedGroup},
			},
		},
	}
}

func (rules *IPv4NamedAdvanceRules) ConvertToTop() *Top {
	return &Top{
		ACL: &ACL{
			IPv4NamedAdvanceRules: rules,
		},
	}
}

func (rule *IPv4NamedAdvanceRule) ConvertToTop() *Top {
	return &Top{
		ACL: &ACL{
			IPv4NamedAdvanceRules: &IPv4NamedAdvanceRules{
				IPv4NamedAdvanceRules: []IPv4NamedAdvanceRule{*rule},
			},
		},
	}
}

func (rule *IPv4NamedBasicRule) ConvertToTop() *Top {
	return &Top{
		ACL: &ACL{
			IPv4NamedBasicRules: &IPv4NamedBasicRules{
				IPv4NamedBasicRules: []IPv4NamedBasicRule{*rule},
			},
		},
	}
}

func (rules *IPv6NamedAdvanceRules) ConvertToTop() *Top {
	return &Top{
		ACL: &ACL{
			IPv6NamedAdvanceRules: rules,
		},
	}
}

func (rule *IPv6NamedAdvanceRule) ConvertToTop() *Top {
	return &Top{
		ACL: &ACL{
			IPv6NamedAdvanceRules: &IPv6NamedAdvanceRules{
				IPv6NamedAdvanceRules: []IPv6NamedAdvanceRule{*rule},
			},
		},
	}
}

func (rule *IPv6NamedBasicRule) ConvertToTop() *Top {
	return &Top{
		ACL: &ACL{
			IPv6NamedBasicRules: &IPv6NamedBasicRules{
				IPv6NamedBasicRules: []IPv6NamedBasicRule{*rule},
			},
		},
	}
}

func (pfilter *Pfilter) ConvertToTop() *Top {
	return &Top{
		ACL: &ACL{
			PfilterApply: &PfilterApply{
				Pfilters: []Pfilter{*pfilter},
			},
		},
	}
}
