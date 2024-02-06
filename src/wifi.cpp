//
// Created by kyour on 2024/2/6.
//

#include "wifi.h"

//32: https://github.com/Edzelf/ESP32-Radio
//8266 : https://github.com/earlephilhower/ESP8266Audio
#ifdef ESP32
#include <WiFi.h>
#else
#include <ESP8266WiFi.h>
#endif

// Enter your Wi-Fi setup here:
const char *WIFI_SSID = "70701";
const char *WIFI_PASSWORD = "2653907035";

/**
 * 初始化WIFI连接
 */
void initWifi() {
    Serial.println("Connecting to WiFi");

    WiFi.disconnect();
    WiFi.softAPdisconnect(true);
    WiFi.mode(WIFI_STA);

    WiFi.begin(WIFI_SSID, WIFI_PASSWORD);

    // Try forever
    while (WiFi.status() != WL_CONNECTED) {
        Serial.println("...Connecting to WiFi");
        delay(1000);
    }
    Serial.println("Wifi Connected");
    Serial.print("Wifi Connected, IP: ");
    Serial.println(WiFi.localIP());

}