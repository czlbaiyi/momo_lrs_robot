package robot

import "encoding/json"

//HeartModule 心跳
type startGameModule struct {
}

//Serialize 序列化
func (m *startGameModule) Serialize() []byte {
	b, _ := json.Marshal(m)
	return b
}
