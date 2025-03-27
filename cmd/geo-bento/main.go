package main

import (
	"context"

	"github.com/warpstreamlabs/bento/public/service"

	// Import all standard Bento components
	_ "github.com/warpstreamlabs/bento/public/components/all"

	// Add the needed needed geo plugins
	_ "github.com/akhenakh/geo-bento/country"
	_ "github.com/akhenakh/geo-bento/h3"
	_ "github.com/akhenakh/geo-bento/randpos"
	_ "github.com/akhenakh/geo-bento/s2"
	_ "github.com/akhenakh/geo-bento/tz"
)

func main() {
	service.RunCLI(context.Background())
}
