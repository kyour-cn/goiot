#pragma once
#include <Arduino.h>
#include <PicoMQTT.h>
#include "config/config.h"

void initMqtt();
void mqttLoop();
bool mqttPublish(const String &topic, const String &payload, int qos, bool retain);
bool mqttPublishBinary(const char* topic, const char *payload, size_t payload_len, int qos, bool retain);