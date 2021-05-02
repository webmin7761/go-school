package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"

	_ "net/http/pprof"
)

func main() {

	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		return serverApp(ctx)
	})

	g.Go(func() error {
		return serverDebug(ctx)
	})

	g.Go(func() error {
		return NewSig()(ctx)
	})

	log.Println("Started")

	if err := g.Wait(); err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Println("Stopped")
}

func NewSig(sig ...os.Signal) func(context.Context) error {
	return func(ctx context.Context) error {

		if len(sig) == 0 {
			sig = append(sig, os.Interrupt)
		}

		done := make(chan os.Signal, len(sig))
		signal.Notify(done, sig...)

		var err error
		select {
		case <-ctx.Done():
		case s := <-done:
			err = errors.New("main: " + s.String())
		}

		signal.Stop(done)
		close(done)
		return err
	}
}

func serverApp(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, World!")
	})

	return server("0.0.0.0:8080", mux, ctx)
}

func serverDebug(ctx context.Context) error {
	return server("127.0.0.1:6060", nil, ctx)
}

func server(addr string, handler http.Handler, ctx context.Context) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		<-ctx.Done()
		log.Printf("%s stopping...\n", addr)
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}
