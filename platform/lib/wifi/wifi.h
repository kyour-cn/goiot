
#ifndef GOIOT_LITE_WIFI_H
#define GOIOT_LITE_WIFI_H

//32: https://github.com/Edzelf/ESP32-Radio
//8266 : https://github.com/earlephilhower/ESP8266Audio
#ifdef ESP32
#include <WiFi.h>
#else

#include <ESP8266WiFi.h>

#endif

class WifiTool {
private:
public:
    static bool connect(const char *ssid, const char *password, int timeout);
    static bool connect(const char *ssid, const char *password);
    static bool smartConfig(int timeout);
};

void initWifi();

#endif //GOIOT_LITE_WIFI_H
