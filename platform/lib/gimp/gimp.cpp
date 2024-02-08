//
// Created by kyour on 2024/2/8.
//

#include "gimp.h"

void Gimp::setCmd(const std::string newCmd) {
    this->cmd = newCmd;
}

std::string Gimp::getCmd() {
    return this->cmd;
}

void Gimp::setHeader(std::string key, std::string value) {
    this->headers[key] = value;
}

std::string Gimp::getHeader(std::string key) {
    return this->headers[key];
}

void Gimp::clearHeader(std::string key) {
    this->headers.erase(key);
}

void Gimp::clearHeader() {
    this->headers.clear();
}

void Gimp::setBody(std::string newBody) {
    this->body = newBody;
}

std::string Gimp::getBody() {
    return this->body;
}

// 将请求信息转换为字符串
std::string Gimp::toString() {
    std::string msg = this->cmd;

    // 指令结束符
    msg += "\n";

    for (auto it = this->headers.begin(); it != this->headers.end(); ++it) {
        msg += it->first;
        msg += ":";
        msg += it->second;
        msg += "\n";
    }

    // 消息头结束符
    msg += "\n";

    msg.append(this->body);

    return msg;
}