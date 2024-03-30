
#ifndef PLATFORM_MQTT_H
#define PLATFORM_MQTT_H

#include <Arduino.h>
#include "ESP32MQTTClient.h"
#include "config.h"

void initMqtt();
void mqttLoop();
void onConnectionEstablishedCallback(esp_mqtt_client_handle_t client);
esp_err_t handleMQTT(esp_mqtt_event_handle_t event);

#endif //PLATFORM_MQTT_H
