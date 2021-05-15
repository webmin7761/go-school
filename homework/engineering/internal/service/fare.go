package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/webmin7761/go-school/homework/engineering/api/fare/v1"
	"github.com/webmin7761/go-school/homework/engineering/internal/biz"
	"go.opentelemetry.io/otel"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func NewFareService(fare *biz.FareUsecase, logger log.Logger) *FareService {
	return &FareService{
		fare: fare,
		log:  log.NewHelper("fare", logger),
	}
}

func (s *FareService) CreateFare(ctx context.Context, req *pb.CreateFareRequest) (*pb.CreateFareReply, error) {
	s.log.Infof("input data %v", req)
	err := s.fare.Create(ctx, &biz.Fare{
		OrgAirport:      req.Fare.OrgAirport,
		ArrAirport:      req.Fare.ArrAirport,
		PassageType:     req.Fare.PassageType.String(),
		FirstTravelDate: req.Fare.FirstTravelDate.AsTime(),
		LastTravelDate:  req.Fare.LastTravelDate.AsTime(),
		Amount:          req.Fare.Amount.Value,
	})
	return &pb.CreateFareReply{}, err
}

func (s *FareService) UpdateFare(ctx context.Context, req *pb.UpdateFareRequest) (*pb.UpdateFareReply, error) {
	s.log.Infof("input data %v", req)
	err := s.fare.Update(ctx, req.Fare.Id, &biz.Fare{
		OrgAirport:      req.Fare.OrgAirport,
		ArrAirport:      req.Fare.ArrAirport,
		PassageType:     req.Fare.PassageType.String(),
		FirstTravelDate: req.Fare.FirstTravelDate.AsTime(),
		LastTravelDate:  req.Fare.LastTravelDate.AsTime(),
		Amount:          req.Fare.Amount.Value,
	})
	return &pb.UpdateFareReply{}, err
}

func (s *FareService) DeleteFare(ctx context.Context, req *pb.DeleteFareRequest) (*pb.DeleteFareReply, error) {
	s.log.Infof("input data %v", req)
	err := s.fare.Delete(ctx, req.Id)
	return &pb.DeleteFareReply{}, err
}

func (s *FareService) Pricing(ctx context.Context, req *pb.PricingRequest) (*pb.PricingResponse, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "PricingFare")
	defer span.End()
	p, err := s.fare.Pricing(ctx, &biz.Fare{
		OrgAirport:      req.OrgAirport,
		ArrAirport:      req.ArrAirport,
		PassageType:     req.PassageType.String(),
		FirstTravelDate: req.FlightDatetime.AsTime(),
		LastTravelDate:  req.FlightDatetime.AsTime(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.PricingResponse{
		Result:         &pb.Result{Code: "0"},
		OrgAirport:     req.OrgAirport,
		ArrAirport:     req.ArrAirport,
		FlightDatetime: req.FlightDatetime,
		PassageType:    req.PassageType,
		Amount:         wrapperspb.Double(p.Amount),
	}, nil
}
