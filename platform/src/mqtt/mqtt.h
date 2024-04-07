#pragma once
#include <Arduino.h>
#include <PicoMQTT.h>
#include "config/config.h"
#include <ArduinoJson.h>

void initMqtt();
void mqttLoop();
