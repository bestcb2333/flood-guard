import sys
import json
import os
import requests
from humidity_temperature import get_humidity_temperature
from rain_fall import get_rain_fall
from water_level import get_water_level
from velocity import get_velocity

if len(sys.argv) != 2:
    print("用法: python history_data.py /path/to/config.json")
    sys.exit(1)

with open(sys.argv[1], 'r', encoding='utf-8') as file:
    config = json.load(file)

url = f"https://{config['host']}:{config['port']}/api/edit/historydata"

headers = {
    'Content-Type': 'application/json',
    'X-Password': config['password']
}

data = {}

if config['enable_temperature_humidity']:
    data['humidity'], data['temperature'] = get_humidity_temperature(
        config['DHT22_pin'],
    )

if config['enable_rain_fall']:
    data['rain_fall'] = get_rain_fall(
        config['rain_fall_pin'],
    )

if config['enable_velocity']:
    data['velocity'] = get_velocity(
        config['velocity_trig_pin'], 
        config['velocity_echo_pin'],
        config['velocity_sensor_distance'],
    )

if config['enable_water_level']:
    data['water_level'] = get_water_level(
        config['water_trig_pin'],
        config['water_echo_pin'],
    )

response = requests.post(url, json=data, headers=headers)
message = response.json().get('msg', '无消息')

with open(config['log_path'], 'a', encoding='utf-8'):
    file.write(f"HISTORY_DATA {response.status_code} {message}")
