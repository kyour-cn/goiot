#include <Arduino.h>
#include "property/property.h"

void setup() {
    Serial.begin(115200);
    Serial.println("===================================");

    Property test;

    Serial.printf("test: %d\n", test.size);

    bool res = test.updateFormJson(R"({"data":[{"name":"switch1","value":true},{"name":"switch2","value":true},{"name":"size","value":3}]})");

    Serial.printf("res: %d\n", res);

    Serial.printf("test: %d\n", test.size);
    Serial.println("===================================");

    Serial.println(test.toJson());

}

void loop() {
}
