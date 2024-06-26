#include "wifi_tool.h"

#define EEPROM_START 1024

/*
  保存参数到eeprom
*/
void WifiTool::saveConfig(WifiConfig *config) {
    EEPROM.begin(EEPROM_START);
    auto *p = (uint8_t *) (config);
    for (unsigned int i = 0; i < sizeof(*config); i++) {
        EEPROM.write(i, *(p + i));
    }
    EEPROM.commit();
    // 释放内存
    delete config;
}

/*
   获取wifi账号密码信息
*/
WifiConfig *WifiTool::loadConfig() {
    // 为变量请求内存
    auto *pvalue = new WifiConfig;
    EEPROM.begin(EEPROM_START);
    auto *p = (uint8_t *) (pvalue);
    for (unsigned int i = 0; i < sizeof(*pvalue); i++) {
        *(p + i) = EEPROM.read(i);
    }
    EEPROM.commit();
    return pvalue;
}

/**
   清空wifi账号和密码
*/
void WifiTool::clearConfig() {
    EEPROM.begin(EEPROM_START);
    // 这里为啥是96 ，因为在结构体声明的长度之和就是96
    for (int i = 0; i < 96; i++) {
        EEPROM.write(i, 0);
    }
    EEPROM.commit();
}

bool WifiTool::connect(const char *ssid, const char *password) {
    return WifiTool::connect(ssid, password, 30);
}

bool WifiTool::connect(const char *ssid, const char *password, int timeout) {
    Serial.println("Connecting to WiFi");

//    WiFi.disconnect();
    WiFi.softAPdisconnect(true);
    WiFi.mode(WIFI_STA);

    WiFi.begin(ssid, password);

    int count = 0;

    while (WiFi.status() != WL_CONNECTED) {
        Serial.print(".");
        delay(1000);
        count++;
        if (count > timeout) {
            return false;
        }
    }

    Serial.print("\nIP Address: ");
    Serial.println(WiFi.localIP());

    return true;
}

/**
 * Airkiss一键配网
 */
bool WifiTool::smartConfig(int timeout) {
    WiFi.mode(WIFI_STA);
    Serial.println("Wait for SmartConfig...");
    WiFi.beginSmartConfig(SC_TYPE_ESPTOUCH_AIRKISS);

    int count = 0;

    while (!WiFi.smartConfigDone()) {
        Serial.print(".");

        delay(1000);
        // 超时时间
        if (timeout > 0) {
            count++;
            if (count > timeout) {
                break;
            }
        }
    }

    Serial.println("\nSmartConfig Success");
    Serial.printf("SSID:%s\n", WiFi.SSID().c_str());
    Serial.printf("PSW:%s\n", WiFi.psk().c_str());

    // 保存账号密码
    auto *config = new WifiConfig;
    strcpy(config->ssid, WiFi.SSID().c_str());
    strcpy(config->pwd, WiFi.psk().c_str());

    WifiTool::saveConfig(config);

    while (WiFi.status() != WL_CONNECTED) {

        Serial.print(".");

        delay(1000);
        // 超时时间
        if (timeout > 0) {
            count++;
            if (count > timeout) {
                return false;
            }
        }
    }
    return true;
}

String WifiTool::getMacAddress() {
    return WiFi.macAddress();
}