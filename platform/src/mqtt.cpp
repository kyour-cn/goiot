//
// Created by kyour on 2024/3/29.
//

#include "mqtt.h"

ESP32MQTTClient mqttClient; // all params are set later

void initMqtt() {
    Serial.println("Connecting to server...");

    mqttClient.enableDebuggingMessages();

    mqttClient.setURI(GOIOT_MQTT_SERVER, GOIOT_MQTT_USER, GOIOT_MQTT_PASS);
    mqttClient.enableLastWillMessage((std::string("device/office/") + GOIOT_DEVICE_KEY).c_str(), "I am going offline");
    mqttClient.setKeepAlive(30);
    mqttClient.loopStart();

}

//int pubCount = 0;

void mqttLoop()
{
//    String msg = "Hello: " + String(pubCount++);
//
//    const std::string topic = std::string("device/ping/") + GOIOT_DEVICE_KEY + "/online";
//    mqttClient.publish(topic.c_str(), msg, 0, false);
//
//    delay(2000);
}

/**
 * Callback when connection is established
 * @param client
 */
void onConnectionEstablishedCallback(esp_mqtt_client_handle_t client)
{
    const std::string subtopic = std::string("device/message/") + GOIOT_DEVICE_KEY + "/#";
    if (mqttClient.isMyTurn(client)) // can be omitted if only one client
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