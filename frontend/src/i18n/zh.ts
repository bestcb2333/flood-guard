export default {
  global: {
    yes: '是',
    no: '否',
  },
  navbar: {
    monitor: {
      title: '监控与分析系统',
      dashboard: '仪表盘',
      history: '历史水位数据',
      analyze: '智能分析',
    },
    resource: {
      title: '资源管理系统',
      sensor: '传感器管理',
      resource: '救援资源管理',
      docs: '传感器部署文档',
    },
    event: {
      title: '历史事件评估',
      history: '历史事件查看',
      optimize: '分析与优化',
    },
    public: {
      title: '公众参与系统',
      report: '内涝事件报告',
      volunteer: '志愿者报名',
    },
    user: '个人中心',
    setting: '系统设置',
    hide: '折叠菜单栏',
  },
  monitor: {
    dash: {
      notice: {
        title: '公告列表',
      },
      dashboard: {
        title: '仪表盘',
        sensor: '运行的传感器',
      },
      map: {
        title: '实时地图',
      },
    },
  },
  resource: {
    sensor: {
      dashTitle: '传感器状态总览',
      mapTitle: '传感器分布地图',
      tableTitle: '传感器列表',
    },
    resource: {
      radio: {
        all: '所有资源',
        staff: '人员管理',
        sandbox: '沙袋管理',
        vehicle: '车辆管理',
        pump: '泵管理',
      },
      map: '资源地图',
      table: '资源图表',
    },
    docs: {
      radio: {
        sensor: '传感器部署文档',
        upload: '数据上传文档',
      },
    },
  },
  event: {
    history: {
      statusTitle: '事件信息',
      mapTitle: '事件地图',
      tableTitle: '事件列表',
    },
  },
  user: {
    title: '欢迎使用Flood Guard！',
    login: {
      title: '登录',
      username: '用户名',
      password: '密码',
    },
    signup: {
      title: '注册',
      username: '用户名',
      email: '邮箱',
      authcode: '邮箱验证码',
      password: '密码',
      rePassword: '确认密码',
      captcha: '验证码',
      send: '发送邮件',
    },
    retrieve: {
      title: '找回密码',
      email: '邮箱',
      authcode: '邮箱验证码',
      password: '新密码',
      rePassword: '确认密码',
      send: '发送邮件',
    },
    myinfo: '我的个人信息',
    username: '用户名',
    userId: 'UID',
    email: '邮箱',
    createdAt: '账号创建时间',
    updatedAt: '信息修改时间',
    admin: '管理员',
    regionName: '志愿服务的区域',
    profile: '个人简介',
  },
  setting: {
    title: '系统设置',
    client: {
      title: '客户端设置',
      backendAddr: {
        title: '后端服务器地址',
        default: '恢复默认',
        test: '测试连接',
      },
      dark: {
        title: '启用暗色模式',
        false: '亮色',
        true: '暗色',
      },
      theme: '选取主题色',
      fontSize: '设置字体大小',
      language: '语言',
    },
    server: {
      title: '服务器设置',
      upload: '上传区域',
    },
  },
  table: {
    true: '是',
    false: '否',
  },
}
