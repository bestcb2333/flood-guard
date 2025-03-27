## 传感器相关

### 传感器是什么？

* 程序需要传感器来定点定时采集区域的水位、降水量、温度、湿度、流速，用于评估并预测风险。

### 如何制作传感器？

#### 材料准备

* 树莓派4B开发板
* SD卡（至少8G）
* 温湿度传感器DHT22
* 超声波测距传感器HC-SR04
* 气压传感器BMP280
* 水流速度传感器YF-S201

#### 制作过程

* 将SD卡与各种传感器组件接入开发板，组成传感器
* 给传感器安装Linux系统，例如Raspberry OS或Debian
* 在传感器上安装Python3和所需的模块

`sudo apt update`

`sudo apt install python3`

`pip install requests Adafruit_DHT RPi.GPIO`

* 下载相关的Python代码

`sudo apt install git`

`git clone https://github.com/bestcb2333/floodguard`

`mv floodguard/sensor/* ~`

* 修改配置文件和定时任务

`sudo apt install vim`

`vim ~/config.json`

`sudo crontab -e`

`0 * * * * python3 ~/history_data.py ~/config.json`

`0 * * * * python3 ~/sensor_status.py ~/config.json`

* 关机

`sudo shutdown -h now`


#### 部署过程

* 将传感器放置于需要测量的地方，并接入电源和网线，开机即可。
