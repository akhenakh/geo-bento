package country

import (
	"github.com/akhenakh/coord2country"
	"github.com/warpstreamlabs/bento/public/bloblang"
)

func init() {
	countrySpec := bloblang.NewPluginSpec().
		Param(bloblang.NewFloat64Param("lat")).
		Param(bloblang.NewFloat64Param("lng"))

	idx, err := coord2country.OpenIndex()
	if err != nil {
		panic(err)
	}

	err = bloblang.RegisterFunctionV2(
		"country", countrySpec, func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			lat, err := args.GetFloat64("lat")
			if err != nil {
				return nil, err
			}

			lng, err := args.GetFloat64("lng")
			if err != nil {
				return nil, err
			}

			return func() (interface{}, error) {
				resp := idx.Query(lat, lng)

				countries := make([]string, len(resp))
				for i, l := range resp {
					countries[i] = l.Name
				}

				return countries, nil
			}, nil
		})
	if err != nil {
		panic(err)
	}
}
