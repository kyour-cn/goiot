#pragma once
#include <Arduino.h>
#include "mqtt/mqtt.h"

namespace Subscribe {

    class Handler {
    public:
        static void down(const char *topic, const char *payload);
    };
}