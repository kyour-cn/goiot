#include <Arduino.h>
#include "wifi.h"
#include "ws.h"
#include "../lib/store/store.h"

//const char *ssid = WIFI_NAME;
//const char *password = WIFI_PASSWORD;

void setup() {

    Serial.begin(115200);
    delay(500);
    Serial.println("Device Starting...");

    Store cache = Store("test");

    cache.set("test", "123456");

    String val = cache.get("test", "null");
    Serial.println("Cache test success!"+ val);

    initWifi();

    Serial.println("Initialization complete!");
    Serial.println("===================================");

    initWs();

}

void loop() {
    wsLoop();
//    delay(1000);                  // 留时间给蓝牙缓冲
//    Serial.println("loop!");
}