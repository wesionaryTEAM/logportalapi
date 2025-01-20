package logw

func (f *SubFunc) GetPortalID() string {
	return f.portal.ID
}

func (f *SubFunc) IsActive() bool {
	return f.Active
}

func (f *SubFunc) SetActive(status bool) {
	f.Active = status
}

func (f *SubFunc) Write(p *Payload) {
	if f.Call == nil || !f.IsActive() {
		return
	}

	f.Call(p)
}
