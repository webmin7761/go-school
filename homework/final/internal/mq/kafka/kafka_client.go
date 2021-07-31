package kafka

import (
	"context"
	"log"

	"github.com/webmin7761/go-school/homework/final/internal/conf"
	"golang.org/x/sync/errgroup"
)

type KafkaClient struct {
	address string
	topic   string
	queue   chan string
}

func NewKafkaClient(conf *conf.MessageQueue) *KafkaClient {
	//使用channel来mock kafka
	return &KafkaClient{
		address: conf.Connect.Source,
		topic:   conf.Connect.Topic,
		queue:   make(chan string, 100),
	}
}

func (k *KafkaClient) Produce(msg string) error {
	k.queue <- msg
	return nil
}

func (k *KafkaClient) Consume(consume func(ctx context.Context, msg string, err error) error) error {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		for {
			select {
			case msg := <-k.queue:
				err := consume(ctx, msg, nil)
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
