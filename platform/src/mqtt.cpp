//
// Created by kyour on 2024/3/29.
//

#include "mqtt.h"

ESP32MQTTClient mqttClient; // all params are set later

void initMqtt() {
    Serial.println("Connecting to server...");

    mqttClient.enableDebuggingMessages();

    mqttClient.setURI(server);
    mqttClient.enableLastWillMessage("lwt", "I am going offline");
    mqttClient.setKeepAlive(30);
    WiFi.begin(ssid, pass);
    WiFi.setHostname("c3test");
    mqttClient.loopStart();


    // server address, port and URL
//    webSocket.begin(GOIOT_WS_SERVER, GOIOT_WS_PORT, GOIOT_WS_URL);
//
//    // event handler
//    webSocket.onEvent(webSocketEvent);
//
//    // use HTTP Basic Authorization this is optional remove if not needed
////    webSocket.setAuthorization("user", "Password");
//
//    // try ever 5000 again if connection has failed
//    webSocket.setReconnectInterval(5000);

}