package robot

import "encoding/json"

//CreateRoomRetBodyModule ...
type CreateRoomRetBodyModule struct {
	RoomID     int   `json:"roomId"`
	RoomType   int   `json:"roomType"`
	Solt       int   `json:"solt"`
	CloseSeats []int `json:"closeSeats"`
}

//CreateRoomRetModule ...
type CreateRoomRetModule struct {
	Code                    int `json:"code"`
	Level                   int `json:"level"`
	CreateRoomRetBodyModule `json:"content"`
}

//Serialize 序列化
func (m *CreateRoomRetModule) Serialize() []byte {
	b, _ := json.Marshal(m)
	return b
}
