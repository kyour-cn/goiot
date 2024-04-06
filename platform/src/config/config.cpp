#include "config.h"

// MQTT相关配置
const char *GOIOT_MQTT_SERVER = "192.168.1.123";
const int GOIOT_MQTT_HOST = 1883;
const char *GOIOT_MQTT_USER = "test";
const char *GOIOT_MQTT_PASS = "test123";
// MQTT心跳发送间隔
const unsigned long GOIOT_MQTT_HEARTBEAT_INTERVAL = 20000;

// 设备相关配置
const char *GOIOT_DEVICE_KEY = "test-8266";
const char *GOIOT_DEVICE_SECRET = "test123";

// 配网模式 _SC：SmartConfig模式 _AP：AP热点默认
#define WIFI_CONFIG_MODE_AP

#ifdef WIFI_CONFIG_MODE_AP
#include <WiFiManager.h>
#else
#include "wifi_tool.h"
#endif

// 初始化WIFI
void initWifi() {

#ifdef WIFI_CONFIG_MODE_AP
    WiFiManager wifiManager;
    wifiManager.autoConnect("GOIOT-DEV");
#else
    // 先加载上一次配网账号密码
    WifiConfig *readConfigs = WifiTool::loadConfig();
    if (strlen(readConfigs->ssid) > 0) {
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
#endif
    delay(500);
}