package kafka

import (
	"context"
	"log"
	"time"

	"github.com/webmin7761/go-school/homework/final/internal/conf"
	"golang.org/x/sync/errgroup"
)

type KafkaClient struct {
	address string
	topic   string
}

func NewKafkaClient(conf *conf.MessageQueue) *KafkaClient {
	return &KafkaClient{
		address: conf.Connect.Source,
		topic:   conf.Connect.Topic,
	}
}

func (k *KafkaClient) Produce(msg string) error {
	return nil
}

func (k *KafkaClient) Consume(consume func(ctx context.Context, msg string, err error) error) error {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		chl := time.After(10 * time.Second)
		for {
			select {
			case <-chl:
				//mock 投消息
				err := consume(ctx, "mock msg", nil)
				if err != nil {
					return err
				}
			case <-ctx.Done():
				return nil
			}
		}
	})

	if err := g.Wait(); err != nil {
		log.Printf("error: %v\n", err)
		return err
	}

	return nil
}
