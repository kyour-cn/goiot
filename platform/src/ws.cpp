#include "ws.h"
#include "../lib/gimp/gimp.h"
#include "config.h"
#include "wifi.h"

WebSocketsClient webSocket;

void initWs() {
    Serial.println("Connecting to Ws");

    // server address, port and URL
    webSocket.begin(GOIOT_WS_SERVER, GOIOT_WS_PORT, GOIOT_WS_URL);

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

    String mac = WifiTool::getMacAddress();

    // 发送心跳
    data.setCmd("online");
    data.setHeader("device_key", GOIOT_DEVICE_KEY);
    data.setHeader("secret", GOIOT_DEVICE_SECRET);
    data.setHeader("product_key", GOIOT_PRODUCT_KEY);
    data.setHeader("mac", mac.c_str());

    webSocket.sendTXT(data.encode().c_str());

//    auto binData = data.encode();
//    webSocket.sendBIN((uint8_t * )binData.c_str(), binData.length());
}

/**
 * ws收到消息回调
 * @param message
 */
void onWsMessage(const char *message) {
//    Serial.println(message);
    auto data = Gimp();
    data.decode(message);

    if(data.cmd == "pong") {
        return;
    }else if (data.cmd == "online") {
    Serial.println("Device online successful!");
//        data.setCmd("pong");
//        webSocket.sendTXT(data.encode().c_str());
    }
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

//            Serial.printf("[WSc] get text: %s\n", payload);

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