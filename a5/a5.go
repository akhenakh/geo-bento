package a5

import (
	"github.com/akhenakh/a5-go"
	"github.com/warpstreamlabs/bento/public/bloblang"
)

func init() {
	a5Spec := bloblang.NewPluginSpec().
		Param(bloblang.NewFloat64Param("lat")).
		Param(bloblang.NewFloat64Param("lng")).
		Param(bloblang.NewInt64Param("resolution"))

	err := bloblang.RegisterFunctionV2(
		"a5", a5Spec, func(args *bloblang.ParsedParams) (bloblang.Function, error) {
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
				cell, err := a5.FromLatLng(lat, lng, int(resolution))
				if err != nil {
					return nil, err
				}

				return cell.String(), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}
}
