//
// Created by kyour on 2024/2/8.
//
#include "../lib/gimp/gimp.h"
#include <Arduino.h>

void test() {

    Gimp data = Gimp();

    data.setCmd("test");
    data.setBody("testbody");

    Serial.println("===================================");
    Serial.println(data.encode().c_str());
    Serial.println("===================================");

    Gimp demo = Gimp();

    Serial.println("===================================");
    demo.parse("cmd\ntest:1234567890\nhh:666\n\nbody:1234567890");
    Serial.println("cmd\ntest:1234567890\nhh:666\n\nbody:1234567890");
    Serial.println(demo.getCmd().c_str());
    Serial.println(demo.getHeader("test").c_str());
    Serial.println(demo.getHeader("hh").c_str());
    Serial.println(demo.getBody().c_str());
    Serial.println("===================================");

    Serial.println("===================================");
    demo.parse("cmd2\n\nbodyaaa12\n34567890\naaa");
    Serial.println("cmd2\n\nbodyaaa12\n34567890\naaa");
    Serial.println(demo.getCmd().c_str());
    Serial.println(demo.getHeader("test2").c_str());
    Serial.println(demo.getBody().c_str());
    Serial.println("===================================");

}