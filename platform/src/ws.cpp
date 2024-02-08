//
// Created by kyour on 2024/2/6.
//

#include "ws.h"
#include "../lib/gimp/gimp.h"

WebSocketsClient webSocket;

void initWs()
{
    Serial.println("Connecting to Ws");

    // server address, port and URL
    webSocket.begin("192.168.1.100", 8888, "/ws/dev");

    // event handler
    webSocket.onEvent(webSocketEvent);

    // use HTTP Basic Authorization this is optional remove if not needed
//    webSocket.setAuthorization("user", "Password");

    // try ever 5000 again if connection has failed
    webSocket.setReconnectInterval(5000);

    Serial.println("Ws Connected");
}

void onWsConnect()
{
    // 初始化Gimp对象
    Gimp data = Gimp();

    // 发送心跳
    data.setCmd("ONLINE");
    data.setHeader("id", "4545151fdf15ddsd");
    data.setHeader("mac", "AB1234567890");
    data.setBody("1234567890");

    webSocket.sendTXT(data.encode().c_str());
}

/**
 * ws事件监听
 * @param type
 * @param payload
 * @param length
 */
void webSocketEvent(WStype_t type, uint8_t * payload, size_t length) {

    switch(type) {
        case WStype_DISCONNECTED:
            Serial.printf("[WSc] Disconnected!\n");
            break;
        case WStype_CONNECTED:
            Serial.printf("[WSc] Connected to url: %s\n", payload);
            onWsConnect();
            break;
        case WStype_TEXT:

            if(strcmp((char *)payload, (char *)"push") == 0){
//                Serial.println("[WSc]  请求推流");
//                captureVideo();
            }

            // send message to server
            // webSocket.sendTXT("message here");
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

void sendBin(uint8_t * payload, size_t length)
{
    webSocket.sendBIN(payload, length);
}

bool wsIsConnected()
{
    return webSocket.isConnected();
}

unsigned long lastTime = millis();

void wsLoop()
{
    unsigned long time = millis();
    //每10秒发送心跳
    if(time - 10000 > lastTime){
        lastTime = time;
        if(!wsIsConnected()) return;

        // 发送心跳
        Gimp ping = Gimp();
        ping.setCmd("ping");

        webSocket.sendTXT(ping.encode().c_str());
    }

    webSocket.loop();
}