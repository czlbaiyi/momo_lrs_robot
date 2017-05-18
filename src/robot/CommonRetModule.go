package robot

import "encoding/json"

//CommonRetModule .
type CommonRetModule struct {
	Code    int    `json:"code"`
	Level   int    `json:"level"`
	Content []byte `json:"content"`
}

//Serialize 序列化
func (m *CommonRetModule) Serialize() []byte {
	b, _ := json.Marshal(m)
	return b
}
