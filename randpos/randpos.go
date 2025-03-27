package randpos

import (
	"context"
	"encoding/json"
	"math/rand"

	"github.com/warpstreamlabs/bento/public/service"
)

var randPosConfigSpec = service.NewConfigSpec().
	Summary("Creates an input that generates random coordinates.").
	Field(service.NewFloatField("min_lat").Default(-90.0)).
	Field(service.NewFloatField("max_lat").Default(90.0)).
	Field(service.NewFloatField("min_lng").Default(-180.0)).
	Field(service.NewFloatField("max_lng").Default(180.0))

func newRandPosInput(conf *service.ParsedConfig) (service.Input, error) {
	minLat, err := conf.FieldFloat("min_lat")
	if err != nil {
		return nil, err
	}
	maxLat, err := conf.FieldFloat("max_lat")
	if err != nil {
		return nil, err
	}
	minLng, err := conf.FieldFloat("min_lng")
	if err != nil {
		return nil, err
	}
	maxLng, err := conf.FieldFloat("max_lng")
	if err != nil {
		return nil, err
	}

	return service.AutoRetryNacks(&randPosInput{
		minLat: minLat,
		maxLat: maxLat,
		minLng: minLng,
		maxLng: maxLng,
	}), nil
}

func init() {
	err := service.RegisterInput(
		"randpos", randPosConfigSpec,
		func(conf *service.ParsedConfig, mgr *service.Resources) (service.Input, error) {
			return newRandPosInput(conf)
		})
	if err != nil {
		panic(err)
	}
}

type randPosInput struct {
	minLat, maxLat, minLng, maxLng float64
}

func (i *randPosInput) Connect(ctx context.Context) error {
	return nil
}

func (i *randPosInput) Read(ctx context.Context) (*service.Message, service.AckFunc, error) {
	lat := i.minLat + rand.Float64()*(i.maxLat-i.minLat)
	lng := i.minLng + rand.Float64()*(i.maxLng-i.minLng)

	pos := struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}{
		Lat: lat,
		Lng: lng,
	}

	b, err := json.Marshal(pos)
	if err != nil {
		return nil, nil, err
	}

	return service.NewMessage(b), func(ctx context.Context, err error) error {
		// Nacks are retried automatically when we use service.AutoRetryNacks
		return nil
	}, nil
}

func (i *randPosInput) Close(ctx context.Context) error {
	return nil
}
