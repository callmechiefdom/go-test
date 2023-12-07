+ 启动redis
  + docker-compose up -d
+ 创建项目asynq_task

### asynq_task 目录

```bash
.
|-- README.md
|-- cmd
|   `-- main.go  # 启动消费者监听
|-- go.mod
|-- go.sum
|-- test.go # 生产者 发送测试数据
`-- test_delivery
    |-- client 
    |   `-- client.go # 生产者 具体发送测试数据的逻辑
    `-- test_delivery.go  # 消费者，执行任务具体处理逻辑
├── conf
│   └── redis.conf
└── docker-compose.yml


```

#### redis连接

```go
// Asynq 使用 Redis 作为消息代理。
// client.go 和 main.go 都需要连接到 Redis 进行写入和读取。
// 我们将使用 asynq.RedisClientOpt 指定如何连接到本地 Redis 实例。
asynq.RedisClientOpt{
	Addr:     "10.45.11.94:31599",
	Password: "",
	DB:       2,
}



```

### task

```go
type Task struct {
	// 一个简单的字符串值，表示要执行的任务的类型.
	typename string

	// 有效载荷保存执行任务所需的数据，有效负载值必须是可序列化的.
	payload []byte

	// 保存任务的选项.
	opts []Option

	// 任务的结果编写器.
	w *ResultWriter
}

```

+ go mod init asynq_t
  + go.mod 文件中一直是 module github.com/callmechiefdom/go-test
  + 实际需要初始化的路径是 github.com/callmechiefdom/go-test/asynq_t， 修改继续执行 go get
+ go mod tidy
+ go get -u github.com/hibiken/asynq
+ go doc -u asynq.Config
  + 输出对应的结构体doc

### deploy

```bash
docker run -it --privileged=true --name go-test -p 8085:8080 -v /data/app/aigc/go-test:/data golang:latest bash
cd /data
go mod tidy
go run cmd/main.go

```


### asynqmon

+ docker-compose.yaml
+ docker-compose up
