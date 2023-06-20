package comware

func (base *Base) ConvertToTop() *Top {
	return &Top{
		Device: &Device{
			Base: base,
		},
	}
}
