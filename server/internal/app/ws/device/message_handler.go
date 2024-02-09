package device

import (
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
		nodeId := data.GetHeader("node_id")
		productId, _ := strconv.Atoi(data.GetHeader("product_id"))
		secret := data.GetHeader("secret")
		mac := data.GetHeader("mac")
		if nodeId == "" || productId == 0 {
			sendError(conn, "设备信息不完整")
			return
		}

		//查询数据库中设备信息并校验
		d := query.Device
		deviceInfo, err := d.Where(d.NodeID.Eq(nodeId)).
			Select(d.ID, d.NodeID, d.Secret, d.ProductID, d.Mac).
			First()
		if err != nil ||
			deviceInfo.ProductID != int32(productId) ||
			deviceInfo.Secret != secret {
			// 未查询到设备信息
			sendError(conn, "设备信息不匹配")
			_ = conn.Close()
			return
		}

		updateData := &model.Device{}

		// 更新mac
		nowTime := time.Now().Unix()
		updateData.Mac = mac
		updateData.OnlineTime = int32(nowTime)

		_ = d.Save(deviceInfo)

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
