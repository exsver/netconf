package comware

func (filterSource *FilterSource) ConvertToTop() *Top {
	return &Top{
		ARP: &ARP{
			ArpFilterSource: &ArpFilterSource{
				FilterSources: []FilterSource{*filterSource},
			},
		},
	}
}

func (arpRateLimitLog *ArpRateLimitLog) ConvertToTop() *Top {
	return &Top{
		ARP: &ARP{
			ArpRateLimitLog: arpRateLimitLog,
		},
	}
}

func (ifFilterBinding *ArpInterfaceFilter) ConvertToTop() *Top {
	return &Top{
		ARP: &ARP{
			ArpFilterBinding: &ArpFilterBinding{
				FilterBindings: []ArpInterfaceFilter{*ifFilterBinding},
			},
		},
	}
}

func (ifFilterSource *ArpFilterSource) ConvertToTop() *Top {
	return &Top{
		ARP: &ARP{
			ArpFilterSource: ifFilterSource,
		},
	}
}
