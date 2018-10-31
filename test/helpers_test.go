package test

import (
	"GolangOrdering/helpers"
	"testing"
)

func TestCalcDistance_1(t *testing.T) {
	point := helpers.Distance{
		Src: "[40.6905615,-73.9976592]",
		Dst: "[40.6655101,-73.89188969999998]",
	}

	_, err := point.CalcDistance()
	ok(t, err)
}

func TestCalcDistance_2(t *testing.T) {
	point := helpers.Distance{
		Src: "[40.6905615,-73.9976592]",
		Dst: "[40.6655101,-73.89188969999998]",
	}

	rsp, _ := point.CalcDistance()
	equals(t, 36, rsp)
}
