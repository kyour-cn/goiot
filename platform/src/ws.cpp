//
// Created by kyour on 2024/2/6.
//

#include "ws.h"
#include "../lib/gimp/gimp.h"

WebSocketsClient webSocket;

void initWs() {
    Serial.println("Connecting to Ws");

    // server address, port and URL
    webSocket.begin("192.168.1.100", 8888, "/ws/dev");

    // event handler
    webSocket.onEvent(webSocketEvent);

    // use HTTP Basic Authorization this is optional remove if not needed
//    webSocket.setAuthorization("user", "Password");

    // try ever 5000 again if connection has failed
    webSocket.setReconnectInterval(5000);

}

/**
 * ws连接成功后回调
 */
void onWsConnect() {
    // 初始化Gimp对象
    Gimp data = Gimp();

    // 发送心跳
    data.setCmd("ONLINE");
    data.setHeader("id", "4545151fdf15ddsd");
    data.setHeader("mac", "AB1234567890");

    char bin1[] = {0x1A, 0x23, 0x45, 0x67, 0x68, 0x0A, 0x0A, 0x0A};

    // 将bin1转为字符串
    std::string hex2bin = (std::string) bin1;

    data.setBody("12测试中文34567890" + hex2bin + "hhh=111");

    webSocket.sendTXT(data.encode().c_str());

//    auto bin = data.encode();
//    webSocket.sendBIN((uint8_t * )bin.c_str(), bin.length());
}

void onWsMessage(char *message) {
    Serial.println(message);
}

/**
 * ws事件监听
 * @param type
 * @param payload
 * @param length
 */
void webSocketEvent(WStype_t type, uint8_t *payload, size_t length) {

    switch (type) {
        case WStype_DISCONNECTED:
            Serial.printf("[WSc] Disconnected!\n");
            break;
        case WStype_CONNECTED:
            Serial.printf("[WSc] Connected to url: %s\n", payload);
            onWsConnect();
            break;
        case WStype_TEXT:

            onWsMessage((char *) payload);

            Serial.printf("[WSc] get text: %s\n", payload);

            break;
        case WStype_BIN:
            Serial.printf("[WSc] get binary length: %u\n", length);
//            hexdump(payload, length);

            // send data to server
            // webSocket.sendBIN(payload, length);
            break;
        case WStype_ERROR:
        case WStype_FRAGMENT_TEXT_START:
        case WStype_FRAGMENT_BIN_START:
        case WStype_FRAGMENT:
        case WStype_FRAGMENT_FIN:
            break;
        case WStype_PING:
            break;
        case WStype_PONG:
            break;
    }

}

void wsSendBin(uint8_t *payload, size_t length) {
    webSocket.sendBIN(payload, length);
}

bool wsIsConnected() {
    return webSocket.isConnected();
}

unsigned long lastTime = millis();

void wsLoop() {
    unsigned long time = millis();
    //每10秒发送心跳
    if (time - 10000 > lastTime) {
        lastTime = time;
        if (!wsIsConnected()) return;

        // 发送心跳
        Gimp ping = Gimp();
        ping.setCmd("ping");

        webSocket.sendTXT(ping.encode().c_str());
    }

    webSocket.loop();
}