package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"robot"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var robotType int // 0普通机器人模式 1 压力测试机器人模式
var conf *robot.Config
var roomID int
var roomType int
var roomCount int
var robotNum int
var add string

func readStartPram() {
	log.Println(`启动参数提示,
		启动格式1(1个参数，默认启动20个机器人): " ./robot 房间号" 
			示例: "./robot 10000"
		启动格式2(2个参数): "./robot 房间号 机器人数"
			示例:"./robot 10000 18"
		启动格式3(3参数): " ./robot add 房间类型 房间数 房间机器人数目" 
			示例: ”./robot 127.0.0.1:11111 1 10000 20“
		`)
	lenth := len(os.Args)
	var err error
	if lenth == 3 {
		param2 := os.Args[1]
		roomID, err = strconv.Atoi(param2)
		if err != nil {
			log.Println("roomID 解析失败 ", err)
		}

		param1 := os.Args[2]
		robotNum, err = strconv.Atoi(param1)
		if err != nil {
			log.Println("robotNum 解析失败 ", err)
		}
		robotType = 0
		log.Println("进入房间", roomID, robotNum)
		log.Println("进入房间", "roomID", roomID, "robotNum", robotNum)
	} else if lenth == 2 {
		param2 := os.Args[1]
		roomID, err = strconv.Atoi(param2)
		if err != nil {
			log.Println("roomId 解析失败 ", err)
		}
		robotNum = 20
		robotType = 0
		log.Println("进入房间", "roomID", roomID, "robotNum", robotNum)
	} else if lenth == 5 {
		add = os.Args[1]

		param2 := os.Args[2]
		roomType, err = strconv.Atoi(param2)
		if err != nil {
			log.Println("roomType 解析失败 ", err)
		}

		param3 := os.Args[3]
		roomCount, err = strconv.Atoi(param3)
		if err != nil {
			log.Println("roomCount 解析失败 ", err)
		}
		if roomCount <= 0 {
			roomCount = 1
			log.Println("roomCount 不可以比开房间需要的人数少 ", err)
		}

		param4 := os.Args[4]
		robotNum, err = strconv.Atoi(param4)
		if err != nil {
			log.Println("robotNum 解析失败 ", err)
		}
		robotType = 1
		log.Println("压力测试", "add:", add, "roomCount", roomCount, "robotNum", robotNum)
	} else {
		log.Fatalln(`启动参数错误 请检查`)
	}
}

func getCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func readConfig() {
	configPath := getCurrentPath() + "/Config.json"
	log.Println("读取配置文件路径为: ", configPath)
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalln("读取配置文件失败: ", err)
	} else {
		log.Println("读取文件成功。")
	}

	conf = &robot.Config{}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalln("配置文件解析失败: ", err)
	} else {
		log.Println("配置文件内容为：", conf)
	}

	add = conf.Adds[conf.AddIdx]
}

func unlimitSocket() {
	var rlim syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlim)
	if err != nil {
		fmt.Println("get rlimit error: " + err.Error())
		os.Exit(1)
	}
	rlim.Cur = 65535
	rlim.Max = 65535
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlim)
	if err != nil {
		fmt.Println("set rlimit error: " + err.Error())
		os.Exit(1)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	readConfig()
	readStartPram()
	unlimitSocket()
	if robotType == 0 {
		log.Println("本次启动为普通进入房间 开始创建机器人 ")
	} else {
		log.Println("本次启动为压力测试 开始创建机器人 ")
	}
	time.Sleep(1 * time.Second)
	if robotType == 0 {
		for i := 0; i < robotNum; i++ {
			log.Println("创建机器人 ", i+1)
			_ = robot.CreateRobot(roomID, add)
		}
	} else if robotType == 1 {
		for i := 0; i < roomCount; i++ {
			log.Println("创建房间 ", i+1)
			_ = robot.CreateRobotStressTest(roomType, add, robotNum)
		}
	}
	robot.RunRobotsLogic()
	log.Println(`创建结束------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------`)

	var comm string
	for {
		fmt.Scan(&comm)
		fmt.Println("get a comm :", &comm)
		if comm == "q" {
			break
		}
	}

	log.Println("关闭完成")
}
