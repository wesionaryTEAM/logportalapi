package logw

import "encoding/json"

func (p *Payload) Json() []byte {
	if b, ok := p.Entry.([]byte); ok {
		p.Entry = string(b)
	}

	for i, entry := range p.Entries {
		if b, ok := entry.([]byte); ok {
			p.Entries[i] = string(b)
		}
	}

	b, _ := json.Marshal(p)

	return b
}
