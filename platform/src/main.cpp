#include <Arduino.h>
#include "../lib/wifi/wifi.h"
#include "ws.h"
#include "../lib/store/store.h"

void setup() {

    const char* ssid = "70701";

    Serial.begin(115200);
    delay(500);
    Serial.println("Device Starting...");

    Store cache = Store("test");

    cache.set("test", "123456");

    String val = cache.get("test", "null");
    Serial.println("Cache test success!"+ val);

    bool ret = WifiTool::connect(ssid, "2653907035");
    if (!ret) {
        Serial.println("Connect failed!");
        WifiTool::smartConfig(0);
    }

    Serial.println("Initialization complete!");
    Serial.println("===================================");

    initWs();

}

void loop() {
    wsLoop();
//    delay(1000);                  // 留时间给蓝牙缓冲
//    Serial.println("loop!");
}