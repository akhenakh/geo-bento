package tz

import (
	"github.com/ringsaturn/tzf"
	"github.com/warpstreamlabs/bento/public/bloblang"
)

func init() {
	tzSpec := bloblang.NewPluginSpec().
		Param(bloblang.NewFloat64Param("lat")).
		Param(bloblang.NewFloat64Param("lng"))

	finder, err := tzf.NewDefaultFinder()
	if err != nil {
		panic(err)
	}

	err = bloblang.RegisterFunctionV2(
		"tz", tzSpec, func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			lat, err := args.GetFloat64("lat")
			if err != nil {
				return nil, err
			}

			lng, err := args.GetFloat64("lng")
			if err != nil {
				return nil, err
			}

			return func() (interface{}, error) {
				return finder.GetTimezoneName(lng, lat), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}
}
