# r2id 优化方案

| # | 优先级 | 问题 | 位置 | 建议 | 工作量 |
|---|---|---|---|---|---|
| 1 | 高 | **当前无法编译**：`open4go/db v0.0.6` 和当前解析出的 `log` 版本不兼容 | `go.mod` | `go get github.com/open4go/db@v0.0.15 github.com/open4go/log@v0.0.23 && go mod tidy`；补 CI 防止再次出现 | S |
| 2 | 高 | **订单号随机部分只有 2 位**：`for i := 0; i < OrderLength-len(result.String()); i++` 每次循环 i 加一、长度也加一，结果只追加 2 个随机字母（总长 17，不是注释说的 32，也不是 18）。同一秒内只有 26²=676 种可能，并发下很容易撞号 | `pay.go:48` | 先算 `n := OrderLength - len(prefix+date)` 再循环；把随机部分加长到至少 8 位，或者改用 snowflake 加时间；入库时加唯一索引 | S |
| 3 | 高 | `QueueTicketNumber` 在 Redis 出错时调用标准库 `log.Fatal`，**整个进程退出** | `ticket.go:24` | 返回 `(string, error)` | S |
| 4 | 中 | 短码不唯一：`Cutter` 只取 snowflake 的后 6 位或 3 位，同一毫秒的不同序列号、不同时间都可能重复；`X5Code` 实际是 “X5”+3 位 | `code.go` | 文档里写明“非唯一，只用于展示”；需要唯一时用 Redis INCR 或数据库唯一索引重试 | S |
| 5 | 中 | 排队号的 key 有双冒号（prefix 已经以 `:` 结尾，再拼 `%s:%s:%s`）；日期用的是服务器本地时区，而订单号用的是 Asia/Shanghai | `ticket.go:11,19` | 去掉多余的冒号；统一用 `r3time.LocTime()` | S |
| 6 | 低 | 每次 INCR 之后都执行一次 Expire，没有用 pipeline | `ticket.go` | 用 pipeline 或 Lua，只在第一次（值为 1 时）设置过期时间 | S |
| 7 | 低 | 用 `fmt.Println` 打错误 | `code.go`、`pay.go` | 改用 `open4go/log` | S |
| 8 | 低 | 和 `util9s/q.go`、`p7/oqueue` 的排队功能重复 | 跨仓库 | 统一保留一套 | M |
| 9 | 中 | 没有测试 | — | 对 `GenerateOutTradeNo` 的长度和唯一性写测试（能直接发现第 2 条的问题） | S |
