package robot

import "encoding/json"

//ChatModule 心跳
type ChatModule struct {
	Text string `json:"text"`
}

//Serialize 序列化
func (m *ChatModule) Serialize() []byte {
	b, _ := json.Marshal(m)
	return b
}
