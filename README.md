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
* 接口文档，提供关于如何调用api上传数据的说明。
* 个人页面：如果登录，则与个人信息相关；如果未登录，则要求登录。
## 逻辑相关
* 查数据不要求登录；增、删、改数据要求登录且具有管理员权限。因此不必登录即可进入主页面，未登录或没有管理员权限不允许访问数据修改页面。
## 数据库表相关
* 用户表，包括字段：ID int，用户名string，密码string，邮箱string，个人介绍string，是否管理员boolean。
* 地区表，包括字段：ID int，地区名string，描述string。
* 内涝事件，包括字段：ID int，区域ID int，开始时间datetime，结束时间datetime，上传者用户ID int，严重性string，具体位置string，描述string。
* 历史数据，包括字段：ID int，记录时间datetime，地区ID int，降水量float，水位float，流速float，气温float，湿度float，数据源string。
* 通知公告，包括字段：ID int，作者ID int，标题string，内容string，重要性int。
* 用户评论，包括字段：ID int，作者ID int，内容string，关联的表（公告或事件）string，关联的数据ID。
## 接口相关
下面的响应体均为{"msg": "消息", "data": 内容}里的内容部分。
### POST /login 登录
* 请求体 application/json
  * CaptchaID string
  * CaptchaValue string
  * Username string
  * Password string
* 响应体 application/json
  * token string
  * username string
  * admin bool
  * email string
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

