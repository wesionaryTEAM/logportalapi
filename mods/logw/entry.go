package logw

import "encoding/json"

func (message *Entry) Json() []byte {
	b, _ := json.Marshal(message)
	return b
}
