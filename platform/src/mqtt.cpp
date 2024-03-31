#include "mqtt.h"

ESP32MQTTClient mqttClient;

/**
 * 初始化MQTT
 */
void initMqtt() {
    Serial.println("Connecting to MQTT...");

    mqttClient.enableDebuggingMessages();

    mqttClient.setURI(GOIOT_MQTT_SERVER, GOIOT_MQTT_USER, GOIOT_MQTT_PASS);
    mqttClient.setMqttClientName(GOIOT_DEVICE_KEY);
    mqttClient.enableLastWillMessage((std::string("device/basic/offline/") + GOIOT_DEVICE_KEY).c_str(), "offline");
    mqttClient.setKeepAlive(30);
    mqttClient.loopStart();

}

// 上次心跳发送时间
unsigned long lastHeartbeat = 0;

void mqttLoop()
{
    // 发送心跳
    if (millis() - lastHeartbeat > GOIOT_MQTT_HEARTBEAT_INTERVAL) {
        lastHeartbeat = millis();
        const std::string topic = std::string("device/basic/ping/") + GOIOT_DEVICE_KEY;
        mqttClient.publish(topic.c_str(), std::to_string(lastHeartbeat).c_str(), 0, false);
    }
}

/**
 * Callback when connection is established
 * @param client
 */
void onConnectionEstablishedCallback(esp_mqtt_client_handle_t client)
{
    Serial.println("MQTT connected");
    // 订阅主题
    const std::string subtopic = std::string("device/+/down/") + GOIOT_DEVICE_KEY;
    if (mqttClient.isMyTurn(client))
    {
        mqttClient.subscribe(subtopic.c_str(), [](const String &topic, const String &payload)
        {
            Serial.printf("%s: %s", topic.c_str(), payload.c_str());
        });
    }
}

esp_err_t handleMQTT(esp_mqtt_event_handle_t event)
{
    mqttClient.onEventCallback(event);
    return ESP_OK;
}