package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/webmin7761/go-school/homework/engineering/api/price/v1"
	"github.com/webmin7761/go-school/homework/engineering/internal/conf"
	"github.com/webmin7761/go-school/homework/engineering/internal/service"
	"go.opentelemetry.io/otel/trace"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, tracer trace.TracerProvider, price *service.PriceService) *http.Server {
	var opts []http.ServerOption
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	m := http.Middleware(
		middleware.Chain(
			tracing.Server(tracing.WithTracerProvider(tracer)),
			logging.Server(log.DefaultLogger),
			recovery.Recovery(),
		),
	)
	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", v1.NewPriceServiceHandler(price, m))
	return srv
}
