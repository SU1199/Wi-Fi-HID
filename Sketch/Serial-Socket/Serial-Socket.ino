#include "ESP8266WiFi.h"
 
const char* ssid = "";
const char* password =  "";

IPAddress local_IP(192, 168, 1, 69);
IPAddress gateway(192, 168, 1, 254);
IPAddress subnet(255, 255, 255, 0);

WiFiServer wifiServer(80);
 
void setup() {
 
  Serial.begin(115200);

  if (!WiFi.config(local_IP, gateway, subnet)) {
    Serial.println("STA Failed to configure");
  }
 
  delay(1000);
 
  WiFi.begin(ssid, password);
  
  while (WiFi.status() != WL_CONNECTED) {
    delay(1000);
    Serial.println("Connecting..");
  }
 
  Serial.print("Connected to WiFi. IP:");
  Serial.println(WiFi.localIP());
 
  wifiServer.begin();
}
 
void loop() {
  WiFiClient client = wifiServer.available();
  if (client) {
    Serial.println("Client Connected");
    while (client.connected()) {
      while (client.available() > 0) {
        char c = client.read();
        Serial.write(c);
        delay(10);
      }
        if (Serial.available()) {
          size_t len = Serial.available();
          uint8_t sbuf[len];
          Serial.readBytes(sbuf, len);
          client.write(sbuf, len);
          delay(1);
        }
    }
    client.stop();
    Serial.println("Client disconnected");
  }
}
