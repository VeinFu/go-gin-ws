## 写在前面

该项目就是一个基于Go Gin框架的Demo，并且添加了Swagger，仅供调研记录和参考之用。

## 实现什么
该Demo基于Gin简单实现了四个api：

```bash
# 获取指定user的信息
GET /api/v1/users/<user_name>

# 创建一个user
POST /api/v1/users

# Gin query string api test
GET /api/test?name=<name>&api_feature=<api_feature>

# swagger ui
GET /api/v1/swaggerui/dist
```

运行起Web Server之后直接访问Swagger UI便可对其接口文档进行查看和接口测试。

## 参考资料
[Go Gin文档](https://github.com/gin-gonic/gin)

[Swagger对接Gin参考资料1](https://www.ribice.ba/swagger-golang/)

[Swagger对接Gin参考资料2](https://medium.com/@ribice/serve-swaggerui-within-your-golang-application-5486748a5ed4)
