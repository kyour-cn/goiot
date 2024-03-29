#include <Arduino.h>
#include "wifi_tool.h"
#include "mqtt.h"

void setup() {

    Serial.begin(115200);
    delay(500);
    Serial.println("Device Starting...");

//    auto mac = WifiTool::getMacAddress();
//    Serial.println("MAC:" + mac);

    // 先加载上一次配网账号密码
    WifiConfig *readConfigs = WifiTool::loadConfig();
    if (strlen(readConfigs->ssid) > 0) {
        Serial.println("ssid:" + String(readConfigs->ssid));
        Serial.println("pwd:" + String(readConfigs->pwd));

        bool ret = WifiTool::connect(readConfigs->ssid, readConfigs->pwd);
        if (!ret) {
            // 连接失败, 重新配网
            Serial.println("Connect failed!");
            WifiTool::smartConfig(0);
        }
    } else {
        // 重新配网
        WifiTool::smartConfig(0);
    }

    Serial.println("Initialization complete!");
    Serial.println("===================================");

    initMqtt();

}

void loop() {
//    delay(1000);                  // 留时间给蓝牙缓冲
//    Serial.println("loop!");
}