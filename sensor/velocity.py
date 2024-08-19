import RPi.GPIO as GPIO
import time

def measure_time(TRIG_PIN, ECHO_PIN):
    GPIO.output(TRIG_PIN, True)
    time.sleep(0.00001)
    GPIO.output(TRIG_PIN, False)
    start_time = time.time()
    stop_time = time.time()
    while GPIO.input(ECHO_PIN) == 0:
        start_time = time.time()
    while GPIO.input(ECHO_PIN) == 1:
        stop_time = time.time()
    time_elapsed = stop_time - start_time
    return time_elapsed

def get_velocity(TRIG_PIN, ECHO_PIN, DISTANCE):
    GPIO.setmode(GPIO.BCM)
    GPIO.setup(TRIG_PIN, GPIO.OUT)
    GPIO.setup(ECHO_PIN, GPIO.IN)
    while True:
        time_downstream = measure_time(TRIG_PIN, ECHO_PIN)
        time.sleep(1)

        time_upstream = measure_time(TRIG_PIN, ECHO_PIN)

        return (DISTANCE * (time_upstream - time_downstream)) / (2 * time_downstream * time_upstream)

