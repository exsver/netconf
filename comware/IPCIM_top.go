package comware

func (binding *SourceBinding) ConvertToTop() *Top {
	return &Top{
		IPCIM: &IPCIM{
			IPSourceBindingInterface: &IPSourceBindingInterface{
				SourceBindings: []SourceBinding{*binding},
			},
		},
	}
}

func (verifySource *VerifySource) ConvertToTop() *Top {
	return &Top{
		IPCIM: &IPCIM{
			IPVerifySource: &IPVerifySource{
				[]VerifySource{*verifySource},
			},
		},
	}
}
