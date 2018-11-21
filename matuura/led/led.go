package led

import (
	"time"

	"github.com/tarm/serial"
)

// 信号線は13番
/* led.ino
```led.ino
#include <Adafruit_NeoPixel.h>

#define LED_NUM       ( 500 )
#define DIN_PIN       (  13 )

Adafruit_NeoPixel* pixels;

int i = 0;
char cl[3]={0,0,0};

void setup() {
 Serial.begin(9600);
 pixels = new Adafruit_NeoPixel( LED_NUM, DIN_PIN, NEO_RGB + NEO_KHZ800);
 pixels->begin();
 led_on();
}

void loop() {
 if(Serial.available()>0){
   cl[0] = Serial.read();
   cl[1] = Serial.read();
   cl[2] = Serial.read();
   led_on();
 }
 delay(500);
}

void led_on(){
 pixels->setBrightness(255);
 for( int i = 0; i < pixels->numPixels(); i++ ){
   pixels->setPixelColor(i,Adafruit_NeoPixel::Color( cl[0], cl[1], cl[2]) );
 }
 pixels->show();
}
```

*/

type LED struct {
	p *serial.Port
}

func NewLED(port string) (*LED, error) {

	c := &serial.Config{Name: port, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}
	time.Sleep(2 * time.Second)
	return &LED{p: s}, nil
}

func (l *LED) Send(R, G, B uint8) error {
	var res []byte
	res = append(res, ToByte(G))
	res = append(res, ToByte(R))
	res = append(res, ToByte(B))
	_, err := l.p.Write(res)
	if err != nil {
		return err
	}
	err = l.p.Flush()
	return err
}

func ToByte(i uint8) byte {
	return byte(i)
}
