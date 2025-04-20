#include <WiFi.h>
#include <ArduinoJson.h>
#include <HTTPClient.h>
#define ANALOG_SENSOR 32
#define BOMB_OUTPUT 33
#define dry 3550
#define wet 1059

const char* ssid = "Ceetros";
const char* password = "jR110501";
const char* serverName = "http://51.38.188.158:420/api/v1/sesor/update";

int wifiConnect (){
  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);
  Serial.println("Conectando");

  unsigned long startAttemptTime = millis();
  const unsigned long timeout = 10000;

  while (WiFi.status() != WL_CONNECTED && millis() - startAttemptTime < timeout) {
    delay(500);
    Serial.print(".");
  }

  if(WiFi.status() != WL_CONNECTED){
    Serial.println("\nFalha na conexão.");
    return WiFi.status();
  } else {
    Serial.print("\nConectado na rede WiFi com IP ");
    Serial.println(WiFi.localIP());
    return WiFi.status();
  }
}


int postData(int sensorData){
  WiFiClient client;
  HTTPClient http;
  StaticJsonDocument<200> doc;
  doc["sensor"] = "wd001";
  doc["value"] = sensorData;

  http.begin(client, serverName);
  http.addHeader("Content-Type", "application/json");
  String json;
  serializeJson(doc, json);
  int httpResponse = http.POST(json);

  Serial.println(httpResponse);

  http.end();

  return httpResponse;
}

void setup() {
  Serial.begin(115200);
  delay(200);
  pinMode(BOMB_OUTPUT, OUTPUT);

  int wifiReturn = wifiConnect();
  while(wifiReturn != 3){
    Serial.println("Tentando reconexão.");
    wifiReturn = wifiConnect();
  }
}

void loop() {
  if(WiFi.status() != WL_CONNECTED){
    int wifiReturn = wifiConnect();
    while(wifiReturn != 3){
      Serial.println("Tentando reconexão.");
      wifiReturn = wifiConnect();
    }
  }

  int humidity = analogRead(ANALOG_SENSOR);
  Serial.println(humidity);
  humidity = map(humidity, dry, wet, 0, 100);


  if(humidity >= 70)
  {
    digitalWrite(BOMB_OUTPUT, LOW);
    Serial.println("TUA DJAMBA TA HUMILDEEEE ("+String(humidity)+"%)");

  }
  else
  {
    Serial.println("TA SECO, PORCENTAGEM:" + String(humidity) + "%");
    Serial.println("LIGANGO BOMBA");
    digitalWrite(BOMB_OUTPUT, HIGH);

  }

  postData(humidity);
  delay(1000 * 5 * 60);
}
