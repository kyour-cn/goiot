//
// Created by kyour on 2024/2/6.
//

#include "wifi.h"

bool WifiTool::connect(const char *ssid, const char *password)
{
    return connect(ssid, password, 20);
}

bool WifiTool::connect(const char *ssid, const char *password, int timeout)
{
    Serial.println("Connecting to WiFi");

    WiFi.disconnect();
    WiFi.softAPdisconnect(true);
    WiFi.mode(WIFI_STA);

    WiFi.begin(ssid, password);

    int count = 0;

    while (WiFi.status() != WL_CONNECTED) {
        Serial.print(".");
        delay(1000);
        count++;
        if (count > timeout) {
            WiFi.disconnect();
            return false;
        }
    }

    return true;
}

/**
 * Airkiss一键配网
 */
bool WifiTool::smartConfig(int timeout)
{
    WiFi.mode(WIFI_STA);
    Serial.println("Wait for SmartConfig...");
    WiFi.beginSmartConfig();

    int count = 0;

    while (WiFi.status() != WL_CONNECTED) {

        Serial.print(".");
        delay(1000);
        // 超时时间
        if(timeout > 0) {
            count++;
            if (count > timeout) {
                WiFi.disconnect();
                return false;
            }
        }

        if (WiFi.smartConfigDone()) {
            Serial.println("SmartConfig Success");
            Serial.printf("SSID:%s\r\n", WiFi.SSID().c_str());
            Serial.printf("PSW:%s\r\n", WiFi.psk().c_str());
//            Serial.print("IP Address: ");
//            Serial.println(WiFi.localIP());

            return WifiTool::connect(WiFi.SSID().c_str(), WiFi.psk().c_str());

//            return true;
        }
    }
    return false;
}