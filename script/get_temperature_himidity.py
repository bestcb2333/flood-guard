import Adafruit_DHT

def get_humi_temp(pin):
    return Adafruit_DHT.read_retry(Adafruit_DHT.DHT22, pin)
