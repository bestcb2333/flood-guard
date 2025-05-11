export const rawWikiGroups: [
  path: string,
  title: string,
  docs: [
    path: string,
    title: string,
  ][],
][] = [
  ['guide', '系统使用指南', [
    ['login-permissions', '用户登录与权限管理说明'],
    ['home-navigation', '系统首页导航介绍'],
    ['query-export', '数据查询与报表导出流程'],
    ['hazard-reporting', '隐患点位上报操作手册'],
    ['mobile-app', '移动端APP操作指南'],
  ]],
  ['knowledge', '内涝隐患基础知识', [
    ['causes-types', '城市内涝成因及分类'],
    ['hazard-types', '常见城市内涝隐患点类型解析'],
    ['risk-assessment', '内涝风险评估方法简介'],
    ['infrastructure-impact', '与洪涝相关的城市基础设施影响分析'],
  ]],
  ['data-specs', '数据管理规范', [
    ['field-description', '内涝隐患点属性字段说明'],
    ['input-standards', '数据录入标准与示例'],
    ['gis-standards', 'GIS空间数据管理规范'],
    ['data-quality', '数据质量控制流程'],
  ]],
  ['standards', '标准与法规政策', [
    ['drainage-construction', '城市排水防涝设施建设标准'],
    ['monitoring-warning', '城市内涝监测与预警技术规范'],
    ['local-guidelines', '地方政府关于城市防涝工作的指导意见'],
    ['emergency-response', '国家应急响应等级划分与启动机制'],
  ]],
  ['emergency', '应急响应与处置流程', [
    ['warning-response', '内涝预警接收与响应机制'],
    ['multi-dept-flow', '多部门联动处置流程图'],
    ['pump-deployment', '应急排水设备调用流程'],
    ['info-reporting', '信息上报与发布机制'],
  ]],
]

export const wikiGroups = rawWikiGroups.map(([path, title, docs]) => ({
  path, title, docs: docs.map(([path, title]) => ({path, title})),
}))

