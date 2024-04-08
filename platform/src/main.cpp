#include <Arduino.h>
#include "config/config.h"
#include "mqtt/mqtt.h"

void setup() {

    Serial.begin(115200);
    delay(500);
    Serial.println("Device Starting...");

    // 初始化WIFI
    initWifi();

    // 初始化MQTT
    initMqtt();

    Serial.println("Initialization complete!");
    Serial.println("===================================");
}

void loop() {
    mqttLoop();
}
