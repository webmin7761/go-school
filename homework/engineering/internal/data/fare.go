package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"github.com/webmin7761/go-school/homework/engineering/internal/biz"
	"github.com/webmin7761/go-school/homework/engineering/internal/data/ent"
	"github.com/webmin7761/go-school/homework/engineering/internal/data/ent/fare"
)

type fareRepo struct {
	data *Data
	log  *log.Helper
}

// NewfareRepo .
func NewFareRepo(data *Data, logger log.Logger) biz.FareRepo {
	return &fareRepo{
		data: data,
		log:  log.NewHelper("fare_repo", logger),
	}
}

func (ar *fareRepo) Pricing(ctx context.Context, bizFare *biz.Fare) (*biz.Fare, error) {
	p, err := ar.data.db.Fare.
		Query().
		Where(
			fare.And(
				fare.FirstTravelDateLTE(bizFare.FirstTravelDate),
				fare.LastTravelDateGTE(bizFare.LastTravelDate),
				fare.OrgAirportEQ(bizFare.OrgAirport),
				fare.ArrAirportEQ(bizFare.ArrAirport),
				fare.PassageTypeEQ(bizFare.PassageType))).
		First(ctx)

	if err != nil {
		return nil, err
	}
	return &biz.Fare{
		Id:              p.ID,
		OrgAirport:      p.OrgAirport,
		ArrAirport:      p.ArrAirport,
		PassageType:     p.PassageType,
		FirstTravelDate: p.FirstTravelDate,
		LastTravelDate:  p.LastTravelDate,
		Amount:          p.Amount,
		CreatedAt:       p.CreatedAt,
		UpdatedAt:       p.UpdatedAt,
	}, nil
}

func (ar *fareRepo) CreateFare(ctx context.Context, fare *biz.Fare) error {

	p, err := ar.Pricing(ctx, fare)
	if err != nil && !ent.IsNotFound(err) {
		return err
	}

	if p != nil {
		return errors.New("data: create to fare error")
	}

	_, err = ar.data.db.Fare.
		Create().
		SetOrgAirport(fare.OrgAirport).
		SetArrAirport(fare.ArrAirport).
		SetPassageType(fare.PassageType).
		SetFirstTravelDate(fare.FirstTravelDate).
		SetLastTravelDate(fare.LastTravelDate).
		SetAmount(fare.Amount).
		Save(ctx)
	return err
}

func (ar *fareRepo) UpdateFare(ctx context.Context, id int64, fare *biz.Fare) error {
	p, err := ar.data.db.Fare.Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = p.Update().
		SetOrgAirport(fare.OrgAirport).
		SetArrAirport(fare.ArrAirport).
		SetPassageType(fare.PassageType).
		SetFirstTravelDate(fare.FirstTravelDate).
		SetLastTravelDate(fare.LastTravelDate).
		SetAmount(fare.Amount).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	return err
}

func (ar *fareRepo) DeleteFare(ctx context.Context, id int64) error {
	return ar.data.db.Fare.DeleteOneID(id).Exec(ctx)
}
