#include "store.h"

Store::Store(const char* name) {
    this->name = name;
}

void Store::initStore() {
    if (this->isInit) return;
    this->store = Preferences();
    Serial.print("Store::initStore: ");
    Serial.println(this->name);

    this->store.begin(this->name);
    this->isInit = true;
}

//检查Key是否存在
bool Store::check(const char *key) {
    this->initStore();
    bool ret = this->store.isKey(key);
    return ret;
}

//获取缓存
String Store::get(const char *key, const String &defaultVal) {
    this->initStore();
    String data = this->store.getString(key, defaultVal);
    return data;
}

//设置缓存
void Store::set(const char *key, const String &data) {
    this->initStore();
    this->store.putString(key, data);
}

//移除缓存
void Store::remove(const char *key) {
    initStore();
    this->store.remove(key);
}

void Store::clear() {
    this->store.clear();
}
