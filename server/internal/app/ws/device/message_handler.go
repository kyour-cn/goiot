package device

import (
	"fmt"
	"github.com/gorilla/websocket"
	"gourd/internal/orm/model"
	"gourd/internal/orm/query"
	"gourd/pkg/gimp"
	"strconv"
	"time"
)

// MessageHandler 消息处理
func MessageHandler(msg string, conn *websocket.Conn) {

	// 解析GIMP协议
	data := gimp.Gimp{}
	err := data.Decode(msg)
	if err != nil {
		// 非法消息处理
		return
	}

	switch data.GetCmd() {
	case "ping": // 设备心跳
		pong := gimp.Gimp{}
		pong.SetCmd("pong")

		// 发送消息
		_ = conn.WriteMessage(websocket.TextMessage, []byte(pong.Encode()))
		break
	case "online": //设备上线

		// 校验设备信息
		deviceKey := data.GetHeader("device_key")
		productKey := data.GetHeader("product_key")
		secret := data.GetHeader("secret")
		mac := data.GetHeader("mac")
		if deviceKey == "" || productKey == "" {
			sendError(conn, "Request data incomplete")
			_ = conn.Close()
			return
		}

		p := query.Product

		fmt.Println("productKey:", productKey)

		productInfo, err := p.Where(p.Key.Eq(productKey)).
			Select(p.ID, p.Key).
			First()
		if err != nil {
			// 未查询到产品信息
			sendError(conn, "Incomplete product information")
			_ = conn.Close()
			return
		}

		//查询数据库中设备信息并校验
		d := query.Device
		deviceInfo, err := d.Where(d.DeviceKey.Eq(deviceKey)).
			Select(d.ID, d.DeviceKey, d.Secret, d.ProductID, d.Mac).
			First()
		if err != nil ||
			productInfo.Key != productKey ||
			deviceInfo.Secret != secret {
			// 未查询到设备信息
			sendError(conn, "Incomplete device information")
			_ = conn.Close()
			return
		}

		// 更新mac
		nowTime := time.Now().Unix()
		_, _ = d.Where(d.ID.Eq(deviceInfo.ID)).Updates(&model.Device{
			Mac:        mac,
			OnlineTime: int32(nowTime),
		})

		res := gimp.Gimp{}
		res.SetCmd("online")
		res.SetHeader("time", strconv.Itoa(int(nowTime)))
		_ = conn.WriteMessage(websocket.TextMessage, []byte(res.Encode()))
	}

}

func sendError(conn *websocket.Conn, msg string) {
	res := gimp.Gimp{}
	res.SetCmd("error")
	res.SetBody(msg)
	_ = conn.WriteMessage(websocket.TextMessage, []byte(res.Encode()))
}
