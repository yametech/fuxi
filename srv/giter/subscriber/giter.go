package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	giter "github.com/yametech/fuxi/srv/giter/proto/giter"
)

type Giter struct{}

func (e *Giter) Handle(ctx context.Context, msg *giter.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *giter.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
