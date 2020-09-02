/**
 * Author: richen
 * Date: 2020-09-02 09:31:38
 * LastEditTime: 2020-09-02 16:22:12
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */

/**
NSQ
---------------------------
* nsqd,负责消息接收、保存以及发送消息给消费者的进程
* nsqlookupd， 负责维护所有nsqd的状态，提供服务发现的过程
* nsqadmin web管理平台，实时监控集群以及执行管理任务
* 消息默认不持久化，可以配置持久化
* 每条消息至少传递一次
* 消息不保证有序

---------------------------
topic 事件队列
channel  每个消费者对应一个channel，实现消息可重复消费
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

var (
	producer *nsq.Producer
)

type Consumer struct {
}

func (*Consumer) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message", string(msg.Body))
	return nil
}

func initProducer(dsn string) error {
	var err error
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(dsn, config)
	if err != nil {
		return err
	}
	return nil
}

func Produce() {
	nsqAddress := "9.135.12.51:4150"
	err := initProducer(nsqAddress)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	//读取控制台输入
	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read string failed error: %v\n", err)
		}

		data = strings.TrimSpace(data)
		if data == "stop" {
			break
		}

		err = producer.Publish("order_queue", []byte(data))
		if err != nil {
			fmt.Printf("publish message failed, error: %v\n", err)
		}
		fmt.Printf("publish data: %s success", data)
	}
}

func initConsumer(topic string, channel string, address string) error {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = 15 * time.Second

	c, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		return err
	}
	consumer := &Consumer{}
	c.AddHandler(consumer)

	//建立NSQLookupd连接
	if err := c.ConnectToNSQLookupd(address); err != nil {
		return err
	}

	return nil
}

func Consume() {
	err := initConsumer("order_queue", "first", "9.135.12.51:4161")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	<-c
}

func main() {
	Consume()
	// Produce()
}
