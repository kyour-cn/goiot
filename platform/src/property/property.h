#pragma once

#include <Arduino.h>
#include <ArduinoJson.h>
#include <map>

class Property {

public:
    bool switch1 = false;
    bool switch2 = false;
    u_int size = 0;

    bool updateFormJson(const char *payload);

    String toJson();

private:

    std::map<std::string, std::function<bool(JsonVariant)>> switchUpdaters = {
            {"switch1", [this](JsonVariant value) { return this->switch1 = value.as<bool>(); }},
            {"switch2", [this](JsonVariant value) { return this->switch2 = value.as<bool>(); }},
            {"size",    [this](JsonVariant value) { return this->size = value.as<u_int>(); }}
    };
};
