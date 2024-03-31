#pragma once
#include <Arduino.h>
#include "ESP32MQTTClient.h"
#include "config/config.h"

void initMqtt();
void mqttLoop();
