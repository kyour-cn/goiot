#pragma once
#include <Arduino.h>
#include <PicoMQTT.h>
#include "config/config.h"

void initMqtt();
void mqttLoop();
