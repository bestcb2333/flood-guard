import Adafruit_DHT

def get_humidity_temperature(pin):
    return Adafruit_DHT.read_retry(Adafruit_DHT.DHT22, pin)
