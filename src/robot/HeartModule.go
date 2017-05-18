package robot

import "encoding/json"

//HeartModule 心跳
type HeartModule struct {
}

//Serialize 序列化
func (m *HeartModule) Serialize() []byte {
	b, _ := json.Marshal(m)
	return b
}
