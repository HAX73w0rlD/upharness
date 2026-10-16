package catalog

// Provider lists only publishable providers. No keys here.
type Provider struct {
	ID string
	Free bool
}

func Providers() []Provider {
	return []Provider{{ID: "nvidia", Free: true}, {ID: "opencode", Free: true}}
}
