#include "property.h"

// 从json更新属性
bool Property::updateFormJson(const char *payload) {
    JsonDocument jsonDoc;

    // 解析JSON
    DeserializationError error = deserializeJson(jsonDoc, payload);
    if (error) return false; // 检查JSON解析是否成功

    JsonArray data = jsonDoc["data"];
    if (!data) return false;

    for (JsonVariant item : data) {
//        Serial.printf("name: %s, value: %d\n", item["name"].as<String>().c_str(), item["value"].as<bool>());
        const char* itemName = item["name"];
        if (!itemName) continue; // 确保"name"键存在
        auto switchUpdaterIt = switchUpdaters.find(itemName);
        if (switchUpdaterIt != switchUpdaters.end()) {
            // 如果找到了对应的更新函数，执行它
            JsonVariant value = item["value"];
            switchUpdaterIt->second(value);
        }
    }

    return true;
}

String Property::toJson() {
    JsonDocument jsonDoc;
    jsonDoc["data"] = JsonArray();

    for (auto& item : switchUpdaters) {
        JsonObject obj = jsonDoc["data"].add<JsonObject>();
        obj["name"] = item.first;
        obj["value"] = item.second(JsonVariant());
    }

    String jsonStr;
    serializeJson(jsonDoc, jsonStr);
    return jsonStr;
}
