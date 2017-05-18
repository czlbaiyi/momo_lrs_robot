package robot

import "encoding/json"

//QuikLoginModule ...
type QuikLoginModule struct {
	La        string `json:"la"`
	Lg        string `json:"lg"`
	RoomType  int    `json:"roomType"`
	SessionID string `json:"sessionId"`
}

//Serialize 序列化
func (m *QuikLoginModule) Serialize() []byte {
	b, _ := json.Marshal(m)
	return b
}
