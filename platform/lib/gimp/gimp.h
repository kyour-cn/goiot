//
// Created by kyour on 2024/2/8.
//

#ifndef PLATFORM_GIMP_H
#define PLATFORM_GIMP_H

#include <map>
#include <string>

class Gimp {
public:
    std::string cmd;
    std::map<std::string, std::string> headers;
    std::string body;

    void setCmd(std::string cmd);

    std::string getCmd();

    void setHeader(std::string key, std::string value);

    std::string getHeader(std::string key);

    void clearHeader(std::string key);

    void clearHeader();

    void setBody(std::string body);

    std::string getBody();

    std::string toString();

};

#endif //PLATFORM_GIMP_H
