import RPi.GPIO as GPIO
import time

# 设置 GPIO 引脚编号
RAIN_SENSOR_PIN = 18  # 假设信号连接到 GPIO18

def get_rain_fall(RAIN_SENSOR_PIN):
    return 1
# 初始化 GPIO
GPIO.setmode(GPIO.BCM)
GPIO.setup(RAIN_SENSOR_PIN, GPIO.IN, pull_up_down=GPIO.PUD_UP)

# 初始化计数器
rain_count = 0

# 定义雨量计算公式
def calculate_rainfall(count):
    # 每次翻转对应的降水量，单位：毫米
    mm_per_tip = 0.2  # 根据你的设备选择正确的值
    return count * mm_per_tip

# 定义回调函数，翻斗翻转时触发
def rain_detected(channel):
    global rain_count
    rain_count += 1
    print(f"Rain detected! Count: {rain_count}")
    print(f"Total Rainfall: {calculate_rainfall(rain_count)} mm")

# 设置事件检测，检测翻斗翻转
GPIO.add_event_detect(RAIN_SENSOR_PIN, GPIO.FALLING, callback=rain_detected, bouncetime=300)

try:
    while True:
        # 保持程序运行，等待雨量计信号
        time.sleep(1)

except KeyboardInterrupt:
    # 清理 GPIO 设置
    GPIO.cleanup()
