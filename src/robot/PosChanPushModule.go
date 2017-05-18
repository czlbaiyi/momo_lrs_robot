package robot

import "encoding/json"

//PosChanPushModule ...
type PosChanPushModule struct {
	Solt     int    `json:"solt"`
	Op       int    `json:"op"`
	MomoID   string `json:"momoId"`
	Name     string `json:"name"`
	HeadIcon string `json:"headIcon"`
}

//Serialize 序列化
func (m *PosChanPushModule) Serialize() []byte {
	b, _ := json.Marshal(m)
	return b
}
