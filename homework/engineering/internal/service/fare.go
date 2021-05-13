package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/webmin7761/go-school/homework/engineering/internal/biz"
	"go.opentelemetry.io/otel"
)

func NewPriceService(fare *biz.FareUsecase, logger log.Logger) *PriceService {
	return &PriceService{
		fare: fare,
		log:  log.NewHelper("fare", logger),
	}
}

func (s *PriceService) CreateFare(ctx context.Context, req *pb.CreateFareRequest) (*pb.CreateFareReply, error) {
	s.log.Infof("input data %v", req)
	err := s.fare.Create(ctx, &biz.Fare{
		Title:   req.Title,
		Content: req.Content,
	})
	return &pb.CreateFareReply{}, err
}

func (s *PriceService) UpdateFare(ctx context.Context, req *pb.UpdateFareRequest) (*pb.UpdateFareReply, error) {
	s.log.Infof("input data %v", req)
	err := s.fare.Update(ctx, req.Id, &biz.Fare{
		Title:   req.Title,
		Content: req.Content,
	})
	return &pb.UpdateFareReply{}, err
}

func (s *PriceService) DeleteFare(ctx context.Context, req *pb.DeleteFareRequest) (*pb.DeleteFareReply, error) {
	s.log.Infof("input data %v", req)
	err := s.fare.Delete(ctx, req.Id)
	return &pb.DeleteFareReply{}, err
}

func (s *BlogService) GetFare(ctx context.Context, req *pb.GetFareRequest) (*pb.GetFareReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetFare")
	defer span.End()
	p, err := s.fare.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetFareReply{Fare: &pb.Fare{Id: p.Id, Title: p.Title, Content: p.Content, Like: p.Like}}, nil
}
