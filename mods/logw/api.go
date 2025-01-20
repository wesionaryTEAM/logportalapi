package logw

func (p *Config) Subscribe(s Subscriber) {
	s.SetActive(true)

	defer func() {
		if sse, ok := s.(*SubSSE); ok {
			sse.portal = p
			sse.run()
		}

		if fnc, ok := s.(*SubFunc); ok {
			fnc.portal = p
		}
	}()

	p.mutex.Lock()

	defer p.mutex.Unlock()

	p.subscribers = append(p.subscribers, s)
}

func (p *Config) Log(entry interface{}) {
	p.history = append(p.history, entry)

	if p.HistoryLen > -1 && len(p.history) > p.HistoryLen {
		p.history = p.history[len(p.history)-p.HistoryLen:]
	}

	payload := &Payload{
		PortalID: p.ID,
		Type:     PAYLOAD_TYPE_SINGLE,
		Entry:    entry,
	}

	for _, s := range p.subscribers {
		go s.Write(payload)
	}
}

func (p *Config) Type() string {
	return "logw"
}
