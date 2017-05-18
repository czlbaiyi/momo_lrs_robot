package robot

import "encoding/json"

type joinRoomModule struct {
	RoomID int `json:"roomId"`
}

func (m *joinRoomModule) Serialize() []byte {
	b, _ := json.Marshal(m)
	return b
}
