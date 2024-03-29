
#ifndef PLATFORM_CONFIG_H
#define PLATFORM_CONFIG_H

// 服务器信息和其他常量声明为 extern
extern const char *GOIOT_MQTT_SERVER;
extern const char *GOIOT_MQTT_USER;
extern const char *GOIOT_MQTT_PASS;

// 设备和产品相关 KEY 声明为 extern
extern const char *GOIOT_DEVICE_KEY;
//extern const char *GOIOT_DEVICE_SECRET; // 如果需要在其他地方定义则保留此行
extern const char *GOIOT_PRODUCT_KEY;

#endif //PLATFORM_CONFIG_H