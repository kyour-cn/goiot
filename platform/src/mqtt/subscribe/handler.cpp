#include "handler.h"

namespace Subscribe {

    /**
     * 将字符串按照指定的分隔符分割成字符串数组
     * @param data
     * @param separator
     * @param result
     * @param maxSize
     */
    void splitString(const String &data, char separator, String result[], int maxSize) {
        int found = 0;
        int strIndex[] = {0, -1};
        int maxIndex = data.length() - 1;
        for (int i = 0; i <= maxIndex && found < maxSize; i++) {
            if (data.charAt(i) == separator || i == maxIndex) {
                found++;
                strIndex[0] = strIndex[1] + 1;
                strIndex[1] = (i == maxIndex) ? i + 1 : i;
                result[found - 1] = data.substring(strIndex[0], strIndex[1]);
            }
        }
    }

    /**
     * 处理MQTT订阅服务端下发的消息
     * @param topic
     * @param payload
     */
    void Handler::down(const char *topic, const char *payload) {

        // 使用函数将MQTT topic分割成字符串数组
        String topicParts[10]; // 假设最多有10个部分
        splitString(String(topic), '/', topicParts, 10);

        if (topicParts[1] == "pong") {
            // 由设备发送的ping后，服务端回复的pong
            return;
        } else if (topicParts[1] == "ping") {
            // 由服务端发送的ping消息，需要回复pong
            mqttPublish(topic, "pong", 0, false);
            return;
        } else if (topicParts[1] == "property") {
            // 服务端下发属性

        } else {
            Serial.printf("%s: %s\n", topic, payload);
        }

    }

}