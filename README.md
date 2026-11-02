# r2id — 短码、订单号与排队号生成

> 仓库：https://github.com/open4go/r2id ｜ 模块：`github.com/open4go/r2id` ｜ Go 1.25.13 ｜ 最近提交：2026-08-21

## 项目简介
基于 snowflake 截取的短数字码（6 位、3 位）、管理员账号码、随机字符串、商户订单号（`GenerateOutTradeNo`），以及按天自增的 Redis 排队号。

## 目录结构
```
r2id/
├── code.go    InitNode / Cutter / D6Code / D3Code / S6Code / S3Code / X5Code（random.r2id.node 配置节点号）
├── pay.go     GenerateRandomString / GenerateOutTradeNo(prefix...)
├── ticket.go  QueueTicketNumber(ctx, ticketPrefix, namespace)（Redis INCR，48h 过期）
└── db.go      GetRedisMiddleHandler（open4go/db "cache"）
```

## 使用示例
```go
code := r2id.S6Code()                      // 例如 "6123456"
no := r2id.GenerateOutTradeNo("P")         // P + 上海时区 yyyyMMddHHmmss + 随机大写字母
ticket := r2id.QueueTicketNumber(ctx, "A", merchantID) // A0001
```

## 依赖
`bwmarrin/snowflake`、`go-redis/v9 v9.0.3`、`viper`、`open4go/db v0.0.6`、`open4go/log v0.0.12`

## 构建
**`go build ./...` 失败**：依赖的 `open4go/db v0.0.6` 调用的是旧版签名 `log.Log()`，与 go.mod 里解析出的 `log v0.0.12`（签名是 `Log(ctx)`）不兼容：
```
github.com/open4go/db@v0.0.6/redis/handler.go:56:3: not enough arguments in call to log.Log
```
没有测试，没有 CI。

## 与其他仓库的关系
依赖 `db`、`log`；已停更的 `pay` 仓库依赖 r2id v0.0.4；业务侧可能通过 r2day 服务使用。
