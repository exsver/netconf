package comware

func (lagg *LAGG) ConvertToTop() *Top {
	return &Top{
		LAGG: lagg,
	}
}

func (base *LAGGBase) ConvertToTop() *Top {
	return &Top{
		LAGG: &LAGG{
			Base: base,
		},
	}
}

func (group *LAGGGroup) ConvertToTop() *Top {
	return &Top{
		LAGG: &LAGG{
			LAGGGroups: &LAGGGroups{
				Groups: []LAGGGroup{*group},
			},
		},
	}
}

func (member *LAGGMember) ConvertToTop() *Top {
	return &Top{
		LAGG: &LAGG{
			LAGGMembers: &LAGGMembers{
				[]LAGGMember{*member},
			},
		},
	}
}
