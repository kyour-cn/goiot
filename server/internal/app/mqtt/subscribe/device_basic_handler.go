package subscribe

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gourd/internal/app/mqtt/service"
	"strings"
)

func DeviceBasicHandler(client mqtt.Client, msg mqtt.Message) {

	topic := msg.Topic()

	// 取出topic第三个字段
	topicArr := strings.Split(topic, "/")
	commend := topicArr[2]
	deviceKey := topicArr[3]

	if commend == "ping" { // ping
		s := service.DeviceService{}
		_ = s.Ping(deviceKey)
		client.Publish("device/pong/down/"+topicArr[3], 0, false, "pong")
	} else if commend == "online" { // connect
		data := string(msg.Payload())
		var deviceOnlineMsg service.DeviceOnlineMsg
		err := json.Unmarshal([]byte(data), &deviceOnlineMsg)
		if err != nil {
			return
		}

		s := service.DeviceService{}
		_, err = s.Online(deviceOnlineMsg)
		if err != nil {
			return
		}

		client.Publish("device/online/down/"+deviceKey, 0, false, "pong")
	} else {
		fmt.Printf("Device basic message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	}

}
