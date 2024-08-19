import sys
import json
import os
import requests
from get_temperature_himidity import get_humi_temp
from get_rain_fall import get_rain_fall
from get_water_level import get_water_level
from get_velocity import get_velocity

if len(sys.argv) != 6:
    print("用法: python upload_history_data.py 配置文件路径(str) 降水量(float) 水位(float) 流速(float) 温度(float) 适度(float)")
    sys.exit(1)

with open(sys.argv[1], 'r', encoding='utf-8') as file:
    config = json.load(file)

url = f"https://{config['host']}:{config['port']}/api/edit/historydata"

headers = {
    'Content-Type': 'application/json',
    'X-Password': config['password']
}

data = {
    'rain_fall': sys.argv[2],
    'water_level': sys.argv[3],
    'velocity': sys.argv[4],
}

if config['enable_temperature_humidity']:
    data['humidity'], data['temperature'] = get_humi_temp(config['DHT22_bin'])

if config['enable_rain_fall']:
response = requests.post(url, json=data, headers=headers)
message = response.json().get('msg', '无消息')

with open(config['log_path'], 'a', encoding='utf-8'):
    file.write(f"HISTORY_DATA {response.status_code} {message}")
