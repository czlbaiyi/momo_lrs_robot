package robot

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type loginModule struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
	MomoID    string  `json:"momoId"`
	Name      string  `json:"name"`
	HeadIcon  string  `json:"headIcon"`
	AppID     string  `json:"appId"`
	AuthTime  string  `json:"authTime"`
	Encrypted string  `json:"encrypted"`
	OS        string  `json:"os"`
	Version   string  `json:"version"`
	Sex       string  `json:"sex"`
	Hometown  string  `json:"hometown"`
}

//Serialize 序列化
func (m *loginModule) Serialize() []byte {
	b, _ := json.Marshal(m)
	return b
}

var currentIdx = 0

// GenerationMomoID 生成陌陌Id
func GenerationMomoID() string {
	currentIdx++
	timeN := time.Now().UnixNano()
	r := rand.New(rand.NewSource(timeN))
	var momoID = strconv.Itoa(r.Intn(100))
	for i := 0; i < 10; i++ {
		momoID += strconv.Itoa(int(r.Intn(10)))
	}
	momoID += strconv.Itoa(currentIdx)
	return momoID
}

func robotLoginPacketGeneration() *loginModule {
	momoID := GenerationMomoID()
	var nameValue = nameList[currentIdx%len(nameList)]

	var headIconValue = headIconList[currentIdx%len(headIconList)]
	var encrypedstr = fmt.Sprintf("appId=mm_sdk_test_3kKsqwvk&authTime=1491369793557&momoId=%s&auth_Login_Secret=immomo", momoID)
	encrypedByte, _ := RsaEncrypt([]byte(encrypedstr))
	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	var encryptedValue = b64.EncodeToString(encrypedByte)

	loginModule := &loginModule{
		Longitude: 0,
		Latitude:  0,
		MomoID:    momoID,
		Name:      nameValue,
		HeadIcon:  headIconValue,
		AppID:     "mm_sdk_test_3kKsqwvk",
		AuthTime:  "1491369793557",
		Encrypted: encryptedValue,
		OS:        "IOS",
		Version:   "1.0.0",
		Sex:       "M",
		Hometown:  "北京",
	}

	return loginModule
}
