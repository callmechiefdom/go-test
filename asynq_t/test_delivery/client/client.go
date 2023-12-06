/*
 * @Author: Chiefdom
 * @Date: 2023-12-06 16:15:53
 * Copyright (c) 2023 by https://github.com/callmechiefdom, All Rights Reserved.
 */
package client

import (
	"fmt"
	"log"

	"github.com/callmechiefdom/go-test/asynq_t/test_delivery"

	"github.com/hibiken/asynq"
)

func EmailDeliveryTaskAdd(i int) {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     "10.45.11.94:31599",
		Password: "",
		DB:       2,
	})
	defer client.Close()

	// 初使货需要传递的数据
	task, err := test_delivery.NewEmailDeliveryTask(42, fmt.Sprintf("some:template:id:%d", i), `{"name":"lisi"}`)
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}
	// 任务入队
	info, err := client.Enqueue(task)

	//info, err := client.Enqueue(task, time.Now())
	// 延迟执行
	// info, err := client.Enqueue(task, asynq.ProcessIn(3*time.Second))
	// MaxRetry 重度次数 Timeout超时时间
	//info, err = client.Enqueue(task, asynq.MaxRetry(10), asynq.Timeout(3*time.Second))
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}
