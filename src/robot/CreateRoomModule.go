package robot

import "encoding/json"

type createRoomModule struct {
	RoomType int `json:"roomType"`
}

//Serialize 序列化
func (m *createRoomModule) Serialize() []byte {
	b, _ := json.Marshal(m)
	return b
}
