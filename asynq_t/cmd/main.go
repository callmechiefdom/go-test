/*
 * @Author: Chiefdom
 * @Date: 2023-12-06 16:12:30
 */
package main

import (
	"log"

	"github.com/callmechiefdom/go-test/asynq_t/test_delivery"

	"github.com/hibiken/asynq"
)

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     "10.45.11.94:31599",
			Password: "",
			DB:       2,
		},
		asynq.Config{
			// 每个进程并发执行的worker数量
			Concurrency: 5,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(test_delivery.TypeEmailDelivery, test_delivery.HandleEmailDeliveryTask)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
