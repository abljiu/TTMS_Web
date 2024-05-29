# TTMS_Web

#项目结构
gin-mall
├── api             # 用于定义接口函数，也就是controller的作用、
├── cache           # Redis缓存
├── conf            # 配置文件
├── dao             # dao层，对db进行操作
├── middleware      # 中间件
├── model           # 定义mysql的模型
├── pkg
│  ├── e            # 错误码
│  └── util         # 工具函数  
├── routes          # 路由逻辑处理
├── serializer      # 将数据序列化为 json 的函数，便于返回给前端
├── service         # 接口函数的实现
└── static          # 存放静态文件
