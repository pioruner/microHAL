package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"microHAL/devices"
)

func StartSubscriber(ctx context.Context, rdb *redis.Client, devs map[string]devices.Device) {
	pubsub := rdb.PSubscribe(ctx, "device:commands:*")
	ch := pubsub.Channel()
	for msg := range ch {
		log.Println("Received redis msg:", msg.Channel, msg.Payload)
		id := msg.Channel[len("device:commands:"):]
		if dev, ok := devs[id]; ok {
			err, _ := dev.Write([]byte(msg.Payload))
			if err != nil {
				log.Printf("Write to %s failed: %v", id, err)
			}
		}
	}
}
