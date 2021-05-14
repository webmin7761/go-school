package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/webmin7761/go-school/homework/engineering/internal/biz"
	"github.com/webmin7761/go-school/homework/engineering/internal/conf"
	"github.com/webmin7761/go-school/homework/engineering/internal/data"
	"github.com/webmin7761/go-school/homework/engineering/internal/server"
	"github.com/webmin7761/go-school/homework/engineering/internal/service"
	"go.opentelemetry.io/otel/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, trace.TracerProvider, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
