package dev_service

import (
	"github.com/gorilla/websocket"
	"sync"
)

type DeviceFd int32

var clients map[DeviceFd]*websocket.Conn
var mutex sync.Mutex
var connID DeviceFd

func checkFdExist(fd DeviceFd) bool {
	_, ok := clients[fd]
	return ok
}

func PushConn(conn *websocket.Conn) DeviceFd {
	if clients == nil {
		clients = make(map[DeviceFd]*websocket.Conn)
	}

	mutex.Lock()
	defer mutex.Unlock()

	// fd最大值为1亿，将达到1亿则清零
	if connID >= 10 { //100000000
		connID = 0
	}

	connID++
	// 判断fd是否存在，如果存在则加一，直到不存在
	for checkFdExist(connID) {
		connID++
	}

	clients[connID] = conn
	//fmt.Println(">>>", clients, connID)

	return connID
}

func CloseConn(fd DeviceFd) {
	delete(clients, fd)
	//fmt.Println("<<<", clients, fd)
}
