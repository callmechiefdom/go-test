﻿/*
 * @Author: Chiefdom
 * @Date: 2023-12-06 16:15:12
 * Copyright (c) 2023 by https://github.com/callmechiefdom, All Rights Reserved.
 */
package main

import (
	"time"

	"github.com/callmechiefdom/go-test/asynq_t/test_delivery/client"
)

func main() {
	for i := 0; i < 3; i++ {
		client.EmailDeliveryTaskAdd(i)
		time.Sleep(time.Second * 3)
	}
}
