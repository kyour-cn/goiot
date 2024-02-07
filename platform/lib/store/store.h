#include "Preferences.h"

class Store {
private:
    Preferences store;
    void initStore();
    bool isInit = false;
    const char *name = "";
public:
    explicit Store(const char* name);
    bool check(const char* key);
    String get(const char* key, const String &defaultVal);
    void set(const char* key, const String& data);
    void remove(const char* key);
    void clear();
};

//bool checkCache(const char* key);
//
//String getCache(const char* key);
//
//void setCache(const char* key, const String& data);
