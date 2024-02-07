#include <ESP8266WiFi.h>

bool AutoConfig()//断电重连
{
    WiFi.begin();
    //如果觉得时间太长可改
    for (int i = 0; i < 20; i++)
    {
        int wstatus = WiFi.status();
        if (wstatus == WL_CONNECTED)
        {
            Serial.println("WIFI SmartConfig Success");
            Serial.printf("SSID:%s", WiFi.SSID().c_str());
            Serial.printf(", PSW:%s\r\n", WiFi.psk().c_str());
            Serial.print("LocalIP:");
            Serial.print(WiFi.localIP());
            Serial.print(" ,GateIP:");
            Serial.println(WiFi.gatewayIP());
            return true;
        }
        else
        {
            Serial.print("WIFI AutoConfig Waiting......");
            Serial.println(wstatus);//返回值一般为6说明未连接
            delay(1000);
        }
    }
    Serial.println("WIFI AutoConfig Faild!" );
    return false;
}

void SmartConfig()//一键配网
{
    WiFi.mode(WIFI_STA);
    Serial.println("\r\nWait for Smartconfig...");
    WiFi.beginSmartConfig();
    while (1)
    {
        Serial.print(".");
        delay(500);                   // wait for a second
        if (WiFi.smartConfigDone())
        {
            Serial.println("SmartConfig Success");
            Serial.printf("SSID:%s\r\n", WiFi.SSID().c_str());
            Serial.printf("PSW:%s\r\n", WiFi.psk().c_str());
            Serial.print("IP Address: ");
            Serial.println(WiFi.localIP());
            break;
        }
    }
}
void setup()
{
    Serial.begin(115200);
    delay(500);
    Serial.println("Device Starting...");

    if(!AutoConfig())
    {
        SmartConfig();
    }

}
void loop(){
}