package robot

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"

	"strconv"

	"github.com/robfig/cron"
)

// Robot has
type Robot struct {
	CloseSign          chan bool
	isConnSucess       bool
	isConnSucessRWLock *sync.RWMutex
	conn               net.Conn
	msgbuf             []byte
	roomType           int
	roomID             int
	add                string
	tickIdx            int
	roomRobotCount     int
	reSendTimes        int
	isRoomOwnner       bool
	hasStartGame       bool
	roomPlayerCount    int
}

//LogType 0全日志 //1关键日志
var LogType = 0
var robots = [](*Robot){}
var timer *cron.Cron
var robotAddLock = new(sync.RWMutex)

func (r *Robot) init() {
	r.isConnSucessRWLock = new(sync.RWMutex)
	r.isConnSucess = false
}

func (r *Robot) stickyPacket() {

}

func (r *Robot) readMsg() {
	buf := make([]byte, 512)
	for {
		r.isConnSucessRWLock.RLock()
		var isConnSucess = r.isConnSucess
		r.isConnSucessRWLock.RUnlock()
		if !isConnSucess {
			continue
		}
		select {
		case <-r.CloseSign:
			{
				log.Println("close readMsg")
				return
			}
		default:
			if LogType == 0 {
				fmt.Print(".")
			}
		}
		lenght, err := r.conn.Read(buf)
		if err != nil {
			log.Println("conn read error  " + err.Error())
			return
		}
		if lenght > 0 {
			r.msgbuf = append(r.msgbuf, buf[0:lenght]...)
			if len(r.msgbuf) >= 12 {
				msgLenth := binary.BigEndian.Uint16(r.msgbuf[2:4])
				if len(r.msgbuf) >= (int(msgLenth) + 4) {
					r.unpacket(r.msgbuf[0 : int(msgLenth)+4])
					r.msgbuf = r.msgbuf[int(msgLenth)+4:]
				}
			}
		}

	}
}

func (r *Robot) sendMsg(buf []byte) (err error) {
	r.isConnSucessRWLock.RLock()
	var isConnSucess = r.isConnSucess
	r.isConnSucessRWLock.RUnlock()
	if !isConnSucess {
		log.Println("send msg err,robot connect has closed:", err.Error())
		return
	}
	_, err = r.conn.Write(buf)
	if err != nil {
		log.Println("send msg error:", err.Error())
		// I thought, listener can continue work another cases
		neterr, ok := err.(net.Error)
		if ok && neterr.Timeout() {
			if r.reSendTimes < 3 {
				r.reSendTimes++
				r.conn.Write(buf)
			} else {
				r.isConnSucessRWLock.Lock()
				r.isConnSucess = false
				r.isConnSucessRWLock.Unlock()
				r.conn.Close()
			}
		} else if err == io.EOF {
			r.isConnSucessRWLock.Lock()
			r.isConnSucess = false
			r.isConnSucessRWLock.Unlock()
			r.conn.Close()
		}
	}
	r.reSendTimes = 0
	return
}

func (r *Robot) auth() {
	loginModule := robotLoginPacketGeneration()
	head := &Header{(uint16)(0x0000), (uint16)(0), (uint8)(0), (uint8)(0), (uint16)(0x0000), (uint16)(currentIdx), (int8)(2), (int8)(1)}
	packet := &Packet{head, loginModule}
	r.sendMsg(packet.Serialize())
}

func (r *Robot) createRoom(roomType int) {
	createRoomModule := &createRoomModule{roomType}
	head := &Header{(uint16)(0x0000), (uint16)(0), (uint8)(0), (uint8)(0), (uint16)(0x0000), (uint16)(currentIdx), (int8)(3), (int8)(1)}
	packet := &Packet{head, createRoomModule}
	r.sendMsg(packet.Serialize())
}

func (r *Robot) joinRoom(roomID int) {
	joinRoomModule := &joinRoomModule{roomID}
	head := &Header{(uint16)(0x0000), (uint16)(0), (uint8)(0), (uint8)(0), (uint16)(0x0000), (uint16)(currentIdx), (int8)(3), (int8)(2)}
	packet := &Packet{head, joinRoomModule}
	r.sendMsg(packet.Serialize())
}

func (r *Robot) heart() {
	heartModule := &HeartModule{}
	head := &Header{(uint16)(0x0000), (uint16)(0), (uint8)(0), (uint8)(0), (uint16)(0x0000), (uint16)(currentIdx), (int8)(1), (int8)(1)}
	packet := &Packet{head, heartModule}
	r.sendMsg(packet.Serialize())
}

func (r *Robot) chat(text string) {
	chatModule := &ChatModule{text}
	head := &Header{(uint16)(0x0000), (uint16)(0), (uint8)(0), (uint8)(0), (uint16)(0x0000), (uint16)(currentIdx), (int8)(5), (int8)(1)}
	packet := &Packet{head, chatModule}
	r.sendMsg(packet.Serialize())
}

func (r *Robot) startGame() {
	startGameModule := &startGameModule{}
	head := &Header{(uint16)(0x0000), (uint16)(0), (uint8)(0), (uint8)(0), (uint16)(0x0000), (uint16)(currentIdx), (int8)(4), (int8)(5)}
	packet := &Packet{head, startGameModule}
	r.sendMsg(packet.Serialize())
}

func (r *Robot) authAndJoinRoom(roomID int) {
	r.auth()
	r.joinRoom(roomID)
}

func (r *Robot) authAndCreateRoom(roomType int) {
	r.auth()
	r.createRoom(roomType)
}

//Stop 停止运行机器人
func (r *Robot) Stop() {
	r.CloseSign <- false
	r.conn.Close()
}

func (r *Robot) test() {
	fmt.Println("sfsfsdfsfa")
}

func createNewRobots() *Robot {
	robotAddLock.Lock()
	r := new(Robot)
	r.init()
	robots = append(robots, r)
	robotAddLock.Unlock()
	return r
}

//CreateRobotFollowRoom 跟随房间 创建机器人
func CreateRobotFollowRoom(newRoomID int, newAdd string) {
	r := createNewRobots()
	r.add = newAdd
	r.roomID = newRoomID
	var err error
	r.conn, err = net.Dial("tcp", r.add)
	if err != nil {
		panic(err)
	}
	if LogType == 0 {
		log.Println("连接服务器地址：", r.add)
	}
	r.isConnSucess = true
	go r.readMsg()
	go r.authAndJoinRoom(newRoomID)
	return
}

//CreateRobotStressTest 创建房间 然后根据参数创建机器人
func CreateRobotStressTest(romeType int, newAdd string, roomRobotCount int) {
	r := createNewRobots()
	r.add = newAdd
	r.roomRobotCount = roomRobotCount
	var err error
	r.conn, err = net.Dial("tcp", r.add)
	if err != nil {
		panic(err)
	}
	if LogType == 0 {
		log.Println("连接服务器地址：", r.add)
	}
	r.isConnSucess = true
	go r.readMsg()
	go r.authAndCreateRoom(romeType)
	return
}

// QuickLoginTest 快速登录
func QuickLoginTest(roomType int, quickHTTPAdd string) {
	reqFullPath := quickHTTPAdd
	reqFullPath += `?lg=0&la=0&roomType=1&sessionId=`
	reqFullPath += GenerationMomoID()
	log.Println(reqFullPath)
	resp, err := http.Get(reqFullPath)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("请求错误，reqFullPath", reqFullPath)
	}
	if LogType == 0 {
		log.Println(string(body))
	}
	ret := &QuickLoginRetModule{}
	err1 := json.Unmarshal(body, &ret)
	if err1 != nil {
		log.Println("消息解析失败", err1)
	} else {
		tcpAdd := ret.Host
		tcpAdd += ":"
		tcpAdd += strconv.Itoa(ret.Port)
		if ret.RoomID > 0 {
			//跟随房间
			log.Println("跟随房间")
			CreateRobotFollowRoom(ret.RoomID, tcpAdd)
		} else {
			//创建房间
			log.Println("创建房间")
			CreateRobotStressTest(roomType, tcpAdd, 1)
		}
	}
}

//RunRobotsLogic 游戏逻辑开启
func RunRobotsLogic() {
	timer = cron.New()
	timer.AddFunc("*/5 * * * * ?", func() {
		log.Println("5s定时器执行-------------------------")
		robotAddLock.RLock()
		for _, r := range robots {
			r.tickIdx += 5
			r.chat("test chart xxxxxxxxx")
			if r.tickIdx%25 == 0 {
				r.heart()
			}
		}
		robotAddLock.RUnlock()
	})
	timer.Start()
}
