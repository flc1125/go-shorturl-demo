# 基于 Golang &amp; Echo 框架实现的短网址 Demo

## 安装

```bash
go mod vendor
```

## 使用

```bash
go run main.go
```

## 路由

- http://example.com/create 生成短网址
- http://example.com/{{url}} 短网址重定向

## TODO

- 封装逻辑层
- DB 接入
- 代码优化