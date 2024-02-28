package pkg

import (
	"context"
	"grpc_course/internal/domain/model"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Shutdown(ctx context.Context, timeout time.Duration, opeations model.Operations) chan struct{} {
	waiters := make(chan struct{})
	go func() {
		s := make(chan os.Signal)
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Println("Shutting down")

		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		var wg sync.WaitGroup

		for key, operation := range opeations {
			wg.Add(1)
			innerKey := key
			innerOp := operation
			go func() {
				defer wg.Done()
				log.Printf("Cleaning Up %s\n", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Printf("Cleaning Failed %s \n", err.Error())
					return
				}
				log.Printf("%s was shutdown gracefully", innerKey)
			}()
			wg.Wait()
			close(waiters)
		}
	}()
	return waiters
}
