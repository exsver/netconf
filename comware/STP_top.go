package comware

func (base *STPBase) ConvertToTop() *Top {
	return &Top{
		STP: &STP{
			Base: base,
		},
	}
}
