
#ifndef PLATFORM_CONFIG_H
#define PLATFORM_CONFIG_H

// WS服务器地址
const char *GOIOT_WS_SERVER = "192.168.1.26";
const int GOIOT_WS_PORT = 8888;
const char *GOIOT_WS_URL = "/ws/dev";

// 设备KEY（设备在平台的唯一码）
const char *GOIOT_DEVICE_KEY = "test_device_01";
// 设备Secret（接入密钥）
const char *GOIOT_DEVICE_SECRET = "goiot_device";
// 产品KEY（用于区分设备类型）
const char *GOIOT_PRODUCT_KEY = "test";

#endif //PLATFORM_CONFIG_H
