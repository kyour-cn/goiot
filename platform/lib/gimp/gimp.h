//
// Created by kyour on 2024/2/8.
//

#ifndef GOIOT_GIMP_H
#define GOIOT_GIMP_H

#include <map>
#include <string>
#include <sstream>

class Gimp {
public:
    std::string cmd;
    std::map<std::string, std::string> headers;
    std::string body;

    Gimp() = default;

    void setCmd(std::string cmd);

    std::string getCmd();

    void setHeader(std::string key, std::string value);

    std::string getHeader(std::string key);

    void clearHeader(std::string key);

    void clearHeader();

    void setBody(std::string body);

    std::string getBody();

    // 封装消息数据用于发送
    std::string encode();

    // 解析消息数据
    bool parse(std::string data);

};

#endif //GOIOT_GIMP_H
