package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "github.com/webmin7761/go-school/homework/final/api/fare/v1"
	"github.com/webmin7761/go-school/homework/final/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewFareService)

type FareService struct {
	pb.UnimplementedFareServiceServer

	log *log.Helper

	fare *biz.FareUsecase
}
