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
std::string Gimp::encode() {
    std::string msg = this->cmd;

    if (this->headers.empty() && this->body.empty()) {
        return msg;
    }

    // 指令结束符
    msg += "\n";

    for (auto it = this->headers.begin(); it != this->headers.end(); ++it) {
        msg += it->first;
        msg += ":";
        msg += it->second;
        msg += "\n";
    }

    if (this->body.empty()) {
        return msg;
    }

    // 消息头结束符（形成空行），后面是消息体
    msg += "\n" + this->body;

    return msg;
}

// 解析请求信息
bool Gimp::decode(std::string data) {
    std::istringstream iss(data);
    std::string line;

    // 读取指令
    std::getline(iss, this->cmd);

    this->headers.clear();
    // 解析消息头(每行一个消息头，如果遇到空行，表示消息头结束)
    while (std::getline(iss, line) && !line.empty()) {
        size_t pos = line.find(':');
        if (pos != std::string::npos) {
            std::string key = line.substr(0, pos);
            std::string value = line.substr(pos + 1);
            this->headers[key] = value;
        }
    }

    this->body = "";
    // 后续的内容被视为消息体
    std::getline(iss, this->body, '\0');

    return true;
}