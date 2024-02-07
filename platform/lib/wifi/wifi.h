
#ifndef GOIOT_LITE_WIFI_H
#define GOIOT_LITE_WIFI_H

//32: https://github.com/Edzelf/ESP32-Radio
//8266 : https://github.com/earlephilhower/ESP8266Audio
#ifdef ESP32
#include <WiFi.h>
#else

#include <ESP8266WiFi.h>
#include <EEPROM.h>         //加载EEPROM的库

#endif

struct WifiConfig {
    char ssid[32];
    char pwd[64];
};

class WifiTool {
public:
    static bool connect(const char *ssid, const char *password, int timeout);

    static bool connect(const char *ssid, const char *password);

    static bool smartConfig(int timeout);

    // wifi连接配置
    static void saveConfig(WifiConfig *config);

    static WifiConfig *loadConfig();

    static void clearConfig();
};

#endif //GOIOT_LITE_WIFI_H
