package comware

func (route *IPv4StaticRoute) ConvertToTop() *Top {
	return &Top{
		StaticRoute: &StaticRoute{
			Ipv4StaticRouteConfigurations: &Ipv4StaticRouteConfigurations{
				[]IPv4StaticRoute{*route},
			},
		},
	}
}

func (route *IPv6StaticRoute) ConvertToTop() *Top {
	return &Top{
		StaticRoute: &StaticRoute{
			Ipv6StaticRouteConfigurations: &Ipv6StaticRouteConfigurations{
				[]IPv6StaticRoute{*route},
			},
		},
	}
}
