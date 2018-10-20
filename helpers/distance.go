//Author Mohammad Naser Abbasanadi
//Creating Date 2018-11-20
// distance.go is to connect to google api and retrieve needed information through apis

package helpers

import (
	"GolangOrdering/config"
	"GolangOrdering/logger"

	"strings"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

//CalcDistance is responsable for providing Distance betwwen two pont based on google map matrix
func CalcDistance(src, dst string) (int, error) {
	cnf := config.GetConfigInstance()
	replacer := strings.NewReplacer("]", "", "[", "", `"`, "")

	src = replacer.Replace(src)
	dst = replacer.Replace(dst)

	logger.Log.Printf("origin: %v and dist : %v", src, dst)

	c, err := maps.NewClient(maps.WithAPIKey(cnf.APIKEY))
	if err != nil {
		logger.Log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DistanceMatrixRequest{
		Origins:      []string{src},
		Destinations: []string{dst},
		Units:        maps.UnitsMetric,
		Language:     "en",
		Mode:         maps.TravelModeDriving,
	}
	route, err := c.DistanceMatrix(context.Background(), r)

	if err != nil {
		logger.Log.Fatalf("fatal error: %s", err)
	}
	return int(route.Rows[0].Elements[0].Duration.Minutes()), nil
}
