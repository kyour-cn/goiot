package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"gourd/internal/app/ws/dev_service"
	"log"
	"net/http"
	"strconv"
)

// 定义一个 websocket.Upgrader，用于升级 HTTP 连接为 websocket 连接
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允许跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// HandlerDev 定义一个 websocket 处理器函数
func HandlerDev(w http.ResponseWriter, r *http.Request) {
	// 升级 HTTP 连接为 websocket 连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// 关闭连接时执行
	defer conn.Close()

	fd := dev_service.PushConn(conn)

	fmt.Println("设备连接:" + strconv.FormatInt(int64(fd), 10))

	// 循环读取和写入数据
	for {
		// 读取数据
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// 将收到的数据交给业务处理函数处理
		dev_service.MsgHandler(string(message), conn)

		fmt.Println("设备收到:" + string(message))

		// 写入数据
		//err = conn.WriteMessage(messageType, message)
		//if err != nil {
		//	//log.Println(err)
		//	break
		//}
	}

	dev_service.CloseConn(fd)
	fmt.Println("设备断开:" + strconv.FormatInt(int64(fd), 10))
}
