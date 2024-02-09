//
// Created by kyour on 2024/2/6.
//

#ifndef GOIOT_LITE_WS_H
#define GOIOT_LITE_WS_H

#include <Arduino.h>
#include "WebSocketsClient.h"

void initWs();
void wsLoop();
void webSocketEvent(WStype_t type, uint8_t * payload, size_t length);
void wsSendBin(uint8_t * payload, size_t length);
bool wsIsConnected();

#endif //GOIOT_LITE_WS_H
