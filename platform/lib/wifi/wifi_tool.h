#pragma once

//32: https://github.com/Edzelf/ESP32-Radio
//8266 : https://github.com/earlephilhower/ESP8266Audio
#ifdef ESP32
#include <WiFi.h>
#else
#include <ESP8266WiFi.h>
#endif

#include <EEPROM.h>         //加载EEPROM的库

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

    static String getMacAddress();
};
