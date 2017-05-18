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

var (
	roomID                  int
	roomType                int
	roomStartGameRobotCount int
	roomCount               int
	robotNum                int
	add                     string
	conf                    *robot.Config
)

func readStartPram() {
	lenth := len(os.Args)
	var err error
	if conf.RobotType == 0 {
		if lenth == 3 {
			param1 := os.Args[1]
			roomID, err = strconv.Atoi(param1)
			if err != nil {
				log.Fatalln("roomID 解析失败 ", err)
			}

			param2 := os.Args[2]
			robotNum, err = strconv.Atoi(param2)
			if err != nil {
				log.Fatalln("robotNum 解析失败 ", err)
			}
			log.Println("进入房间", "roomID", roomID, "robotNum", robotNum)
		} else if lenth == 2 {
			param2 := os.Args[1]
			roomID, err = strconv.Atoi(param2)
			if err != nil {
				log.Fatalln("roomId 解析失败 ", err)
			}
			robotNum = 20
			log.Println("进入房间", "roomID", roomID, "robotNum", robotNum)
		} else {
			log.Fatalln(`启动参数错误 请检查`)
		}
	} else if conf.RobotType == 1 {
		if lenth == 5 {
			param1 := os.Args[1]
			roomType, err = strconv.Atoi(param1)
			if err != nil {
				log.Fatalln("roomType 解析失败 ", err)
			}

			param2 := os.Args[2]
			roomStartGameRobotCount, err = strconv.Atoi(param2)
			if err != nil {
				log.Fatalln("roomStartGameRobotCount 解析失败 ", err)
			}

			param3 := os.Args[3]
			roomCount, err = strconv.Atoi(param3)
			if err != nil {
				log.Fatalln("roomCount 解析失败 ", err)
			}

			param4 := os.Args[4]
			robotNum, err = strconv.Atoi(param4)
			if err != nil {
				log.Fatalln("robotNum 解析失败 ", err)
			}

			if roomCount < roomStartGameRobotCount {
				log.Fatalln("房间人数不可以比开房间需要的人数少 ", err)
			}
			log.Println("压力测试", "add:", add, "roomCount", roomCount, "robotNum", robotNum)
		} else {
			log.Fatalln(`启动参数错误 请检查`)
		}
	} else if conf.RobotType == 2 {
		param1 := os.Args[1]
		roomType, err = strconv.Atoi(param1)
		if err != nil {
			log.Fatalln("roomType 解析失败 ", err)
		}

		param2 := os.Args[2]
		robotNum, err = strconv.Atoi(param2)
		if err != nil {
			log.Fatalln("robotNum 解析失败 ", err)
		}
		log.Println("压力测试", "add:", add, "roomType", roomCount, "robotNum", robotNum)
	} else {
		log.Fatalln(`配置表机器人启动类型参数错误 请检查`)
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
	if conf.RobotType == 0 {
		log.Println("本次启动为普通进入房间 开始创建机器人 ")
	} else if conf.RobotType == 1 {
		log.Println("本次启动为压力测试 开始创建机器人 ")
	} else if conf.RobotType == 2 {
		log.Println("本次启动快速匹配测试 开始创建机器人 ")
	}

	time.Sleep(1 * time.Second)
	if conf.RobotType == 0 {
		for i := 0; i < robotNum; i++ {
			log.Println("创建机器人 ", i+1)
			robot.CreateRobotFollowRoom(roomID, add)
		}
	} else if conf.RobotType == 1 {
		for i := 0; i < roomCount; i++ {
			log.Println("创建房间 ", i+1)
			robot.CreateRobotStressTest(roomType, add, robotNum)
		}
	} else if conf.RobotType == 2 {
		for i := 0; i < robotNum; i++ {
			log.Println("快速进入 ", i+1)
			robot.QuickLoginTest(roomType, add)
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
