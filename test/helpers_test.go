package test

import (
	"GolangOrdering/helpers"
	"log"
	"testing"
)

func TestCalcDistance_1(t *testing.T) {
	src := "[40.6905615,-73.9976592]"
	dst := "[40.6655101,-73.89188969999998]"
	_, err := helpers.CalcDistance(src, dst)
	ok(t, err)
}

func TestCalcDistance_2(t *testing.T) {
	src := "0.6905615,-73.9976592"
	dst := "-40.6655101,-73.89188969999998"
	dist, _ := helpers.CalcDistance(src, dst)
	log.Println(dist)
	equals(t, dist, 0)
}
