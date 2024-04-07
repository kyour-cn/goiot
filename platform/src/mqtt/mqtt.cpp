#include "mqtt.h"

PicoMQTT::Client mqtt(
    GOIOT_MQTT_SERVER,
    GOIOT_MQTT_HOST,
    GOIOT_DEVICE_KEY,
    GOIOT_MQTT_USER,
    GOIOT_MQTT_PASS
);

/**
 * 初始化MQTT
 */
void initMqtt() {
    Serial.println("Connecting to MQTT...");

    // 订阅主题
    const String subtopic = String("device/+/down/") + GOIOT_DEVICE_KEY;
    mqtt.subscribe(subtopic, [](const char *topic, const char *payload) {
        Serial.printf("%s: %s", topic, payload);
    });

    mqtt.connected_callback = [] {
        Serial.println("MQTT connected");

        // 发送上线通知
        const String topic = String("device/basic/online/") + GOIOT_DEVICE_KEY;
        JsonDocument payload;
        payload["device_key"] = GOIOT_DEVICE_KEY;
        payload["device_name"] = GOIOT_DEVICE_SECRET;
        String payloadStr;
        serializeJson(payload, payloadStr);
        mqtt.publish(topic, payloadStr);
    };

    mqtt.begin();

}

// 上次心跳发送时间
unsigned long lastHeartbeat = 0;

void mqttLoop() {

    mqtt.loop();

    unsigned long millisVal = millis();
    // 发送心跳
    if (millisVal - lastHeartbeat > GOIOT_MQTT_HEARTBEAT_INTERVAL || lastHeartbeat / 2 > millisVal) {
        lastHeartbeat = millisVal;
        const String topic = String("device/basic/ping/") + GOIOT_DEVICE_KEY;
        mqtt.publish(topic, String(lastHeartbeat));
    }
}

// 发布消息
bool mqttPublish(const String &topic, const String &payload, int qos, bool retain) {
    return mqtt.publish(topic, payload, qos, retain);
}

// 发布消息二进制数据
bool mqttPublishBinary(const char *topic, const char *payload, size_t payload_len, int qos, bool retain) {
    return mqtt.publish_P(topic, payload, payload_len, qos, retain);
}
