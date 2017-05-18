package robot

import (
	"encoding/json"
)

//GameServer .
type GameServer struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	ID   string `json:"id"`
}

//Profile .
type Profile struct {
	PlayerID      int     `json:"playerId"`
	Level         int     `json:"level"`
	Gold          int     `json:"gold"`
	HappyRound    int     `json:"happyRound"`
	StandardRound int     `json:"standardRound"`
	WinRate       int     `json:"winRate"`
	Experience    int     `json:"experience"`
	EscapeCount   int     `json:"escapeCount"`
	WinCount      int     `json:"winCount"`
	LoseCount     int     `json:"loseCount"`
	TotalCount    int     `json:"totalCount"`
	Sex           string  `json:"sex"`
	Age           int     `json:"age"`
	WorthIndex    int     `json:"worthIndex"`
	PopularityVal int     `json:"popularityVal"`
	FreshRouds    int     `json:"freshRouds"`
	FreshWinRate  int     `json:"freshWinRate"`
	HappyWinRate  int     `json:"happyWinRate"`
	StandWinRate  int     `json:"standWinRate"`
	Hometown      string  `json:"hometown"`
	VipLevel      int     `json:"vipLevel"`
	IsVipYear     bool    `json:"isVipYear"`
	GiftNumber    int     `json:"giftNumber"`
	GiftList      []int   `json:"giftList"`
	Longitude     float32 `json:"longitude"`
	Latitude      float32 `json:"latitude"`
	Balance       int     `json:"balance"`
	GiftGoto      string  `json:"giftGoto"`
}

//QuickLoginMsgBody .
type QuickLoginMsgBody struct {
	GameServer `json:"gameserver"`
	Profile    `json:"profile"`
	RoomID     int `json:"roomId"`
}

//QuickLoginRetModule .
type QuickLoginRetModule struct {
	QuickLoginMsgBody `json:"data"`
	Em                string `json:"em"`
	Timesec           int64  `json:"timesec"`
	Ec                int    `json:"ec"`
}

//Serialize 序列化
func (m *QuickLoginRetModule) Serialize() []byte {
	b, _ := json.Marshal(m)
	return b
}
