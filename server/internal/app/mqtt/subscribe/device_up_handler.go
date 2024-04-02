package subscribe

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-gourd/gourd/log"
	"os"
	"strconv"
	"strings"
)

func DeviceUpHandler(client mqtt.Client, msg mqtt.Message) {

	topicInfo := strings.Split(msg.Topic(), "/")

	// 取出上报类型
	dataType := topicInfo[2]

	switch dataType {
	case "image":
		// 将图片数据写入到文件 ./test.jpg

		fmt.Println("image data write to file" + strconv.Itoa(len(msg.Payload())))

		err := os.WriteFile("./test.png", msg.Payload(), 0644)
		if err != nil {
			log.Error("write file error:" + err.Error())
			return
		}

	case "property":
		DeviceUpHandler(client, msg)
	default:
		log.Error("unknown data type:" + dataType)
	}
}
