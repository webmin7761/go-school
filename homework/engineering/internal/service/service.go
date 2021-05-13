package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "github.com/webmin7761/go-school/homework/engineering/api/price/v1"
	"github.com/webmin7761/go-school/homework/engineering/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewPriceService)

type PriceService struct {
	pb.UnimplementedPriceServiceServer

	log *log.Helper

	fare *biz.FareUsecase
}
