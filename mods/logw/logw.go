package logw

var portals = map[string]*Config{
	// ...
}

func New(p *Config) *Config {
	if p.ID != "" {
		portals[p.ID] = p
	}

	return p
}

func Use(name string) *Config {
	return portals[name]
}
