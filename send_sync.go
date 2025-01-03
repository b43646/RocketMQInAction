package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
)

// implements a simple producer to send message.
func main() {
	p, _ := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"129.226.169.221:9876"})),
		producer.WithRetry(2),
	)
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}
	msg := &primitive.Message{
		Topic: "demo",
		Body:  []byte("Hello RocketMQ Go Client!"),
	}
	msg.WithTag("TagA")
	msg.WithKeys([]string{"KeyA"})
	res, err := p.SendSync(context.Background(), msg)

	if err != nil {
		fmt.Printf("send message error: %s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}
