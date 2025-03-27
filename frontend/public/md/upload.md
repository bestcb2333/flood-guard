## 上传历史数据和传感器信息

* 如果你要自行设计传感器，也可以通过应用的HTTP接口上传数据。

### 上传历史数据

* 请求地址：`https://your_domain/api/edit/historydata`
* 请求方法：`POST`
* 请求头：
* `Content-Type: application/json`
* `X-Password: your_password`
* 请求体格式：
```json
{
    "temperature": 10.00,
    "humidity": 10.00,
    "water_level": 10.00,
    "velocity": 10.00,
    "rain_fall": 10.00
}
```

### 上传传感器信息

* 请求地址：`https://your_domain/api/edit/sensorstatus`
* 请求方法：`POST`
* 请求头：
* `Content-Type: application/json`
* `X-Password: your_password`
* 请求体格式：
```json
{
    "sensor_id": 1,
    "status": "good",
    "description": "This sensor works well."
}
```
