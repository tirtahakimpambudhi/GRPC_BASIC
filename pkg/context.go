package pkg

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func CtxError(ctx context.Context) error {
	switch ctx.Err() {
	case context.DeadlineExceeded:
		log.Println("deadline is exceeded")
		return status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	case context.Canceled:
		log.Println("request canceled")
		return status.Error(codes.Canceled, "request canceled")
	}
	return nil
}
