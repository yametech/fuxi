package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	giter "github.com/yametech/fuxi/srv/giter/proto/giter"
)

type Giter struct{}

// Call is a single request handler called via clientv2.Call or the generated clientv2 code
func (e *Giter) Call(ctx context.Context, req *giter.Request, rsp *giter.Response) error {
	log.Log("Received Giter.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via clientv2.Stream or the generated clientv2 code
func (e *Giter) Stream(ctx context.Context, req *giter.StreamingRequest, stream giter.Giter_StreamStream) error {
	log.Logf("Received Giter.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&giter.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via clientv2.Stream or the generated clientv2 code
func (e *Giter) PingPong(ctx context.Context, stream giter.Giter_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&giter.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
