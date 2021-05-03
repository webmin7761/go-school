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

type Server func(context.Context) error

func main() {
	run(context.Background())
}

func run(ctx context.Context) {
	g, ctx := errgroup.WithContext(ctx)

	servers := []Server{serverApp, serverDebug, newSig()}

	for _, f := range servers {
		svr := f
		g.Go(func() error {
			return svr(ctx)
		})
	}

	if err := g.Wait(); err != nil {
		log.Printf("error: %v\n", err)
	}

	log.Println("Stopped")
}

//处理signal信号
func newSig(sig ...os.Signal) func(context.Context) error {
	return func(ctx context.Context) error {

		if len(sig) == 0 {
			sig = append(sig, os.Interrupt)
		}

		done := make(chan os.Signal, len(sig))
		signal.Notify(done, sig...)

		log.Println("signal process started")

		var err error
		select {
		case <-ctx.Done():
		case s := <-done:
			err = errors.New("main: " + s.String())
		}

		signal.Stop(done)
		close(done)

		log.Printf("signal process stopped\n")

		return err
	}
}

//App
func serverApp(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, World!")
	})

	return server("0.0.0.0:8080", mux, ctx)
}

//pprof
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
		s.Shutdown(ctx)
	}()

	log.Printf("%s listen\n", addr)
	return s.ListenAndServe()
}
