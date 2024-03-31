#include "cam.h"
#include "mqtt/mqtt.h"

void initCam() {
    // 初始化摄像头
    cameraInit();
}

// 上次发送时间
unsigned long lastSendTime = 0;

void camLoop() {
    unsigned long millisVal = millis();
    // 每500ms发送一次
    if (millisVal - lastSendTime > 2000 || lastSendTime / 2 > millisVal) {
        lastSendTime = millisVal;

        sendCamera();
    }
}

// 发送摄像头图像
void sendCamera() {

    // 发送设备图像
    camera_fb_t * fb = esp_camera_fb_get();
    if (fb) {

        String imageData((const char*)fb->buf, fb->len);

        const String topic = String("device/image/up/") + GOIOT_DEVICE_KEY;
        mqttPublish(topic, imageData, 0, false);
    }
    esp_camera_fb_return(fb);
}
