package robot

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"unsafe"
)

var msgSequenceNum uint16

//Header ex
//"消息头开始标示:2个字节0x0000"
//"消息体长度:2个字节,不包含消息体长度本身的长度"
//"消息协议版本号:1个字节，默认从1开始映射到实际的字符串版本号"
//"消息格式(消息body序列化反序列化格式):1个字节,0-json,1-pb,现在默认0(json)"
//"消息状态:2个字节,0x0000－客服端正常请求,0x0001-服务器响应客服端请求,1<<1服务器主动推送,1<<2压缩"
//"消息序列号:2个无符号字节,从1开始自增到最大又从1开始"
//"消息模块号:1个无符号字节"
//"消息功能号:1个无符号字节"
//服务器和客服端消息格式全部使用大端模式
type Header struct {
	MsgBegin       uint16 `json:"msgBegin"`
	MsgLong        uint16 `json:"msgLong"`
	MsgVersionCode uint8  `json:"msgVersionCode"`
	MsgFormatType  uint8  `json:"msgFormatType"`
	MsgState       uint16 `json:"msgState"`
	MsgSequenceNum uint16 `json:"msgSequenceNum"`
	MsgModuleID    int8   `json:"msgModuleId"`
	MsgFunctionID  int8   `json:"msgFunctionId"`
}

//Packet use by net
type Packet struct {
	*Header
	IPacketModle
}

func (p *Packet) setMFD(msgModuleID int8, msgFunctionID int8, ip IPacketModle) {
	p.Header.MsgModuleID = msgModuleID
	p.Header.MsgFunctionID = msgFunctionID
	p.IPacketModle = ip
}

//Serialize 序列化包
func (p *Packet) Serialize() []byte {
	msgSequenceNum++
	p.Header.MsgSequenceNum = msgSequenceNum

	p.Header.MsgLong += uint16(unsafe.Sizeof(p.Header.MsgVersionCode))
	p.Header.MsgLong += uint16(unsafe.Sizeof(p.Header.MsgFormatType))
	p.Header.MsgLong += uint16(unsafe.Sizeof(p.Header.MsgState))
	p.Header.MsgLong += uint16(unsafe.Sizeof(p.Header.MsgSequenceNum))
	p.Header.MsgLong += uint16(unsafe.Sizeof(p.Header.MsgModuleID))
	p.Header.MsgLong += uint16(unsafe.Sizeof(p.Header.MsgFunctionID))
	body := p.IPacketModle.Serialize()
	p.Header.MsgLong += uint16(len(body))

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, p.Header)
	binary.Write(buf, binary.BigEndian, body)
	if LogType == 0 {
		fmt.Println("Send MsgModuleID", p.Header.MsgModuleID, "Send MsgFunctionID", p.Header.MsgFunctionID, "body", p.IPacketModle)
	}
	return buf.Bytes()
}

//Unpacket 解包
func (r *Robot) unpacket(buf []byte) {
	header := new(Header)
	bufIO := bytes.NewReader(buf[0:12])
	binary.Read(bufIO, binary.BigEndian, header)

	serverError := header.MsgState&uint16(1<<7) == uint16(1<<7)
	if serverError {
		log.Println("服务器内部错误")
	} else {
		if LogType == 0 {
			if header.MsgModuleID == 5 && header.MsgFunctionID == 1 {
				log.Print(".")
			} else {
				log.Println("收到服务器数据消息头：", header)
			}
		}

		if header.MsgModuleID == 3 && header.MsgFunctionID == 1 {

			ret := &CreateRoomRetModule{}
			err := json.Unmarshal(buf[12:], &ret)
			if err != nil {
				log.Println("消息解析失败")
			} else {
				if LogType == 0 {
					log.Println("消息体内容为:", ret)
				}
				r.roomID = ret.RoomID
				for i := 0; i < r.roomRobotCount-1; i++ {
					CreateRobotFollowRoom(r.roomID, r.add)
				}
			}
			r.isRoomOwnner = true
		} else if r.isRoomOwnner && header.MsgModuleID == 4 && header.MsgFunctionID == -2 {

			ret := &PosChanPushModule{}
			err := json.Unmarshal(buf[12:], &ret)
			if err != nil {
				log.Println("消息解析失败")
			} else {
				if ret.Op == 0 {
					r.roomPlayerCount++
				} else if ret.Op == 1 {
					r.roomPlayerCount--
				}
				if LogType == 0 {
					log.Println("xxxxxxxxxx roomId:", r.roomID, "count:", r.roomPlayerCount)
				}

				if !r.hasStartGame && r.roomPlayerCount == 9 {
					log.Println("发送开始游戏 roomId = ", r.roomID)
					r.hasStartGame = true
					r.startGame()
				}
			}
		}
	}
}
