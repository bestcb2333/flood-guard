这是我的毕业设计项目的代码，还在写。

That's the code of my graduation project, unfinished yet.

# 开发文档
## 布局相关
* 页面竖着割开，分为左右两栏，左侧为侧边栏导航菜单，余下右侧为各页面的内容。界面风格随意。
## 页面相关
* 主页：包括公告列表和一些装饰性内容（比如Logo和介绍）。
* 数据可视化页面：展示各种数据图表。
* 数据查询页面：可以查询数据，比如历史降水记录，内涝记录等。
* 数据修改页面：可以上传数据，修改或删除数据。
* 智能分析：调用ChatGPT API，传入用户问题和相关数据得到智能分析建议。
* 接口文档，提供关于如何调用api上传数据的说明。
* 个人页面：如果登录，则与个人信息相关；如果未登录，则要求登录。
## 逻辑相关
* 查数据不要求登录；增、删、改数据要求登录且具有管理员权限。因此不必登录即可进入主页面，未登录或没有管理员权限不允许访问数据修改页面。
## 数据库表相关
* 下面的每个字段的描述都是“含义 命名 数据类型”格式，比如“用户名 Username string”。
* int是整数，string是字符串，boolean是布尔值，timestring是有时间格式的字符串。float是小数
* 用户表，包括字段：ID ID int，用户名 Username string，密码 Password string，邮箱 Email string，个人介绍 Profile string，是否管理员Admin boolean。
* 地区表，包括字段：ID ID int，地区名 Name string，描述 Description string，坐标范围 Scope string。
* 内涝事件，包括字段：ID ID int，区域ID RegionID int，开始时间 StartTime timestring，结束时间 EndTime timestring，上传者用户ID Uploader int，严重性 Severity string，具体位置 Positon string，描述 Description string。
* 历史数据，包括字段：ID ID int，记录时间 RecordTime timestring，地区ID RegionID int，降水量 RainFall float，水位 WaterLevel float，流速 Velocity float，气温 Temperature float，湿度 Humidity float，数据源 DataSource string。
* 通知公告，包括字段：ID ID int，作者ID Author int，标题 Title string，内容 Content string，重要性 Importrance int。
* 用户评论，包括字段：ID ID int，作者ID Author int，内容 Content string，关联的表（公告或事件） Related string。
* 传感器列表，包括字段：ID ID int，名称 Name string，横坐标 Abscissa float, 纵坐标 Ordinate float，描述 Description string，区域ID RegionID int。
* 传感器状态，包括字段：ID ID int, 报告时间 Time timestring, 传感器ID SensorID int, 状态 Status string, 描述 Description string
## 接口相关
下面的响应体均为{"msg": "消息", "data": 内容}里的内容部分。
所有的后端路径都要以/api/开头，以和前端区分。例如人机验证码是/api/captcha
### 测试服务器地址
* 前端 https://flood.mcax.cn/
* 后端 https://flood.mcax.cn/api/
* 如果 mcax.cn 不可用，请替换成另一个保密域名和520端口
### POST /login 登录
* 请求体 application/json
  * CaptchaId string
  * CaptchaValue string
  * Username string
  * Password string
* 响应体 application/json
  * token string
  * username string
  * admin bool
  * email string
### GET /captcha 获取人机验证码
* 响应体是验证码图片
* 响应头
  * captchaID string
### POST /signup 注册
* 请求体 application/json
  * Username string
  * Password string
  * Email string
  * Authcode string 邮箱验证码
* 响应体 application/json
  * token string
  * username string
  * admin bool
  * email string
### GET /email 获取注册的邮件验证码
* 查询字符串参数
  * email string 收件人

### 关于增删改查
* 路径一般是“/行为类型/数据类型”的格式
* 带上“/api/”后即为“/api/行为类型/数据类型”格式
* 行为类型有三种，分别是edit(增或改), get(查), delete(删)
* 数据类型有八种，分别是
* user(用户列表), region(区域列表), floodevent(内涝事件列表)
* historydata(区域历史数据), notice(公告列表), comment(评论列表)
* sensor(传感器列表), sensorstatus(传感器历史状态记录)
* 例如, /api/get/region即为获取区域列表的路径
* 例如, /api/delete/user即为删除用户
* 例如, /api/edit/notice即为新增或编辑公告
* get是GET请求，edit和delete是POST请求

#### 对于 /api/get/???
* 对于get接口，可以在URL的查询参数里使用数据库字段来筛选。例如
* 如果数据库里面有时间，可以用URL参数来根据时间筛选，键`start_time`指定起始时间，`end_time`指定结束时间，值格式为`YYYY-MM-DD hh-mm-ss`，对于其他的筛选条件，将数据库里的字段名作为URL参数的键，筛选的目标值作为值即可，URL参数的键名用蛇形命名法，不知道准确的键名可以查数据库。

#### 对于 /api/edit/????
* 对于edit, 提供了id即为修改那个id的数据，没提供id则为新增数据
* edit一次只能修改/添加一条数据，请求体为内容对应数据库表的对象

#### 对于 /api/delete/????
* delete的请求体为application/json类型，内容为整数数组
* 代表被删除的记录的id，例如[2, 3, 4]为删除id为2 3 4的记录

### GET /get/user 获取用户列表
* 查询字符串参数
  * admin bool 仅筛选是否管理员（可选项）
* 响应体
  * 用户表的对象数组，对象的格式即为数据库表里的对应格式。
### GET /get/region 获取地区列表
* 响应体
  * 地区表的对象数组，对象的格式即为数据库表里的对应格式。
### GET /get/floodevent 获取内涝事件列表
* 查询字符串参数
  * region int 筛选特定地区，提供地区ID（可选）
  * from string 筛选起始时间，格式2006-01-02 15:04:05（可选）
  * to string 筛选结束时间，格式2006-01-02 15:04:05（可选）
  * uploader int 筛选上传者，提供用户ID（可选）
  * severity string 筛选严重等级（可选）
  * page int 页码（可选）
  * limit int 每页数量（可选）
* 响应体
  * 内涝事件表的对象数组，对象的格式即为数据库表里的对应格式。
### GET /get/historydata 获取历史记录列表
* 查询字符串参数
  * region int 筛选特定地区，提供地区ID（可选）
  * from string 筛选起始时间，格式2006-01-02 15:04:05（可选）
  * to string 筛选结束时间，格式2006-01-02 15:04:05（可选）
  * page int 页码（可选）
  * limit int 每页数量（可选）
* 响应体
  * 历史记录表的对象数组，对象的格式即为数据库表里的对应格式。
### GET /get/notice 获取公告列表
* 查询字符串参数
  * title string 标题（可选）
  * author int 作者ID（可选）
  * page int 页码（可选）
  * limit int 每页数量（可选）
* 响应体
  * 公告记录表的对象数组，对象的格式即为数据库表里的对应格式。

