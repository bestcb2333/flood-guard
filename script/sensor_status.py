import sys
import json
import os
import requests

if len(sys.argv) != 4:
    print("用法: python upload_history_data.py 配置文件路径(str) 状态(str) 状态描述(str)")
    sys.exit(1)

with open(sys.argv[1], 'r', encoding='utf-8') as file:
    config = json.load(file)

url = f"https://{config['host']}:{config['port']}/api/edit/sensorstatus"

headers = {
    'Content-Type': 'application/json',
    'X-Password': config['password']
}

data = {
    'status': sys.argv[2],
    'description': sys.argv[3],
}

response = requests.post(url, json=data, headers=headers)
message = response.json().get('msg', '无消息')

with open(config['log_path'], 'a', encoding='utf-8'):
    file.write(f"SENSOR_STATUS {response.status_code} {message}")
