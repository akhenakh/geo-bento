package s2

import (
	"github.com/golang/geo/s2"
	"github.com/warpstreamlabs/bento/public/bloblang"
)

func init() {
	s2Spec := bloblang.NewPluginSpec().
		Param(bloblang.NewFloat64Param("lat")).
		Param(bloblang.NewFloat64Param("lng")).
		Param(bloblang.NewInt64Param("resolution"))

	err := bloblang.RegisterFunctionV2(
		"geos2", s2Spec, func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			lat, err := args.GetFloat64("lat")
			if err != nil {
				return nil, err
			}

			lng, err := args.GetFloat64("lng")
			if err != nil {
				return nil, err
			}

			resolution, err := args.GetInt64("resolution")
			if err != nil {
				return nil, err
			}

			return func() (interface{}, error) {
				latLng := s2.LatLngFromDegrees(lat, lng)

				c := s2.CellIDFromLatLng(latLng)
				c = c.Parent(int(resolution))

				return c.String(), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}
}
