package dev_service

import (
	"github.com/gorilla/websocket"
	"gourd/pkg/gimp"
)

func MsgHandler(msg string, conn *websocket.Conn) {

	// 解析GIMP协议
	data := gimp.Gimp{}
	err := data.Decode(msg)
	if err != nil {
		return
	}

	switch data.GetCmd() {
	case "ping":
		pong := gimp.Gimp{}
		pong.SetCmd("pong")

		// 发送消息
		err := conn.WriteMessage(websocket.TextMessage, []byte(pong.Encode()))
		if err != nil {
			return
		}
	}

}
