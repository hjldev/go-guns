##### 基于Gin + Vue + Element UI的前后端分离权限管理系统 

系统初始化极度简单，只需要配置文件中，修改数据库连接，系统启动后会自动初始化数据库信息以及必须的基础数据


## ✨ 特性

- 遵循 RESTful API 设计规范

- 基于 GIN WEB API 框架，提供了丰富的中间件支持（用户认证、跨域、访问日志、追踪ID等）

- 基于Casbin的 RBAC 访问控制模型

- JWT 认证

- 基于 GORM 的数据库存储，目前仅支持mysql数据库

- 配置文件简单的模型映射，快速能够得到想要的配置

- 多命令模式

## 准备工作

你需要在本地安装 [go] [gin] [node](http://nodejs.org/) 和 [git](https://git-scm.com/) 


## 📦 本地开发

### 启动说明

#### 服务端启动说明

```bash
# 进入 go-guns 后端项目
cd ./go-guns

# 编译项目
go build

# 修改配置 
# 文件路径  go-guns/config/admin-dev.yml
vi ./config/admin-dev.yml 

# 1. 配置文件中修改数据库信息 
# 注意: database 下对应的配置数据
# 2. 确认log路径
```

#### 服务启动
```

# 启动项目，也可以用IDE进行调试
go run main.go admin -e dev

```


