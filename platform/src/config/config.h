#pragma once

// 服务器信息和其他常量声
extern const char *GOIOT_MQTT_SERVER;
extern const int GOIOT_MQTT_HOST;
extern const char *GOIOT_MQTT_USER;
extern const char *GOIOT_MQTT_PASS;
extern const unsigned long GOIOT_MQTT_HEARTBEAT_INTERVAL;

// 设备和产品相关 KEY
extern const char *GOIOT_PRODUCT_KEY;
extern const char *GOIOT_DEVICE_KEY;

void initWifi();