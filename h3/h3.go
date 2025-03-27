package h3

import (
	h3 "github.com/akhenakh/goh3"
	"github.com/warpstreamlabs/bento/public/bloblang"
	"modernc.org/libc"
)

func init() {
	h3Spec := bloblang.NewPluginSpec().
		Param(bloblang.NewFloat64Param("lat")).
		Param(bloblang.NewFloat64Param("lng")).
		Param(bloblang.NewInt64Param("resolution"))

	tls := libc.NewTLS()

	err := bloblang.RegisterFunctionV2(
		"h3", h3Spec, func(args *bloblang.ParsedParams) (bloblang.Function, error) {
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
				latLng := h3.NewLatLng(lat, lng)

				c := h3.LatLngToCell(latLng, int(resolution), h3.WithTLS(tls))

				return c.String(), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}
}
