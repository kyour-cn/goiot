#pragma once
#include <Arduino.h>
#include "ESP32MQTTClient.h"
#include "config/config.h"

void initMqtt();
void mqttLoop();
bool mqttPublish(const String &topic, const String &payload, int qos, bool retain);
bool mqttPublishBinary(const char* topic, const uint8_t* payload, size_t payload_len, int qos, bool retain);