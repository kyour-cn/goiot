package dev_service

import "github.com/gorilla/websocket"

func MsgHandler(msg string, conn *websocket.Conn) {

	switch msg {
	case "ping":
		// 发送消息
		err := conn.WriteMessage(websocket.TextMessage, []byte("pong"))
		if err != nil {
			return
		}
	}

}
