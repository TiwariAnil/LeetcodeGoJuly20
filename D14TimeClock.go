package main

import (
	"fmt"
	"math"
)

func angleClock(hour int, minutes int) float64 {
	alpha := ((float64(hour) * 60  + float64(minutes)) / 2.0);
	intAlpha := int64(alpha)
	modDiff := float64(alpha) - float64(intAlpha)
	modAlpha := float64(intAlpha % 360);
	finalAlpha := modDiff + modAlpha
	beta := minutes * 6;
	angle := math.Abs(float64(beta) - float64(finalAlpha));
	if angle < (360-angle){
		return angle
	}
	return 360-angle
}

func main() {
	fmt.Println(angleClock(3,15))
}
