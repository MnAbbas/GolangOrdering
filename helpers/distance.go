//Author Mohammad Naser Abbasanadi
//Creating Date 2018-10-20
// distance.go is to connect to google api and retrieve needed information through apis

package helpers

import (
	"GolangOrdering/config"
	"GolangOrdering/logger"
	"fmt"
	"sync"

	"strings"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

type IGoogle interface {
	initClient() *maps.Client
}
type Distance struct {
	Src, Dst string
}

func (dist *Distance) initClient() *maps.Client {
	cnf := config.GetConfigInstance()

	c, err := maps.NewClient(maps.WithAPIKey(cnf.APIKEY))
	if err != nil {
		logger.Log.Fatalf("fatal error: %s", err)
		return nil
	}
	return c
}

var (
	instace      *Distance
	distanceOnce sync.Once
)

func mapClient() IGoogle {
	if instace == nil {
		distanceOnce.Do(func() {
			instace = &Distance{}
		})
	}
	return instace
}

//CalcDistance is responsable for providing Distance betwwen two pont based on google map matrix
func (dist *Distance) CalcDistance() (int, error) {
	replacer := strings.NewReplacer("]", "", "[", "", `"`, "")

	src := replacer.Replace(dist.Src)
	dst := replacer.Replace(dist.Dst)

	// logger.Log.Printf("origin: %v and dist : %v", src, dst)
	c := dist.initClient()
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
	fmt.Println("distance", route)
	return int(route.Rows[0].Elements[0].Duration.Minutes()), nil
}
